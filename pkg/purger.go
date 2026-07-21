// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkg

import (
	"context"
	"sync/atomic"

	"github.com/IBM/sarama"
	"github.com/bborbe/errors"
	libkafka "github.com/bborbe/kafka"
	"github.com/bborbe/log"
	"github.com/bborbe/run"
	libsentry "github.com/bborbe/sentry"
	"github.com/golang/glog"
)

type Purger interface {
	Purge(ctx context.Context, topic libkafka.Topic) error
}

func NewPurger(
	sentryClient libsentry.Client,
	saramaClientProvider libkafka.SaramaClientProvider,
	deleteSender DeleteSender,
	logSamplerFactory log.SamplerFactory,
) Purger {
	return &purger{
		sentryClient:         sentryClient,
		saramaClientProvider: saramaClientProvider,
		deleteSender:         deleteSender,
		logSamplerFactory:    logSamplerFactory,
	}
}

type purger struct {
	sentryClient         libsentry.Client
	deleteSender         DeleteSender
	logSamplerFactory    log.SamplerFactory
	saramaClientProvider libkafka.SaramaClientProvider
}

func (p *purger) Purge(ctx context.Context, topic libkafka.Topic) error {
	saramaClient, err := p.saramaClientProvider.Client(ctx)
	if err != nil {
		return errors.Wrapf(ctx, err, "create sarama client failed")
	}

	highWaterMarks, err := libkafka.HighWaterMarks(ctx, saramaClient, topic)
	if err != nil {
		return errors.Wrapf(ctx, err, "get highwater marks failed")
	}

	trigger := run.NewTrigger()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var deleted, skip int64

	go func() {
		select {
		case <-ctx.Done():
		case <-trigger.Done():
			glog.V(2).Infof("purge finished with %d deleted and %d skipped records", deleted, skip)
			cancel()
		}
	}()

	return libkafka.NewSimpleConsumer(
		saramaClient,
		topic,
		libkafka.OffsetOldest,
		libkafka.MessageHanderList{
			libkafka.MessageHandlerFunc(
				func(ctx context.Context, msg *sarama.ConsumerMessage) error {
					if len(msg.Value) == 0 {
						atomic.AddInt64(&skip, 1)
						glog.V(2).Infof("key '%s' already deleted => skip", string(msg.Key))
						return nil
					}
					glog.V(2).Infof("key '%s' => delete", string(msg.Key))
					if err := p.deleteSender.SendDelete(ctx, libkafka.Topic(msg.Topic), msg.Key); err != nil {
						return errors.Wrap(ctx, err, "send delete failed")
					}
					atomic.AddInt64(&deleted, 1)
					return nil
				},
			),
			libkafka.NewOffsetTriggerMessageHandler(
				highWaterMarks,
				topic,
				trigger,
			),
		},
		p.logSamplerFactory,
	).Consume(ctx)
}
