// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkg

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/bborbe/errors"
	libkafka "github.com/bborbe/kafka"
	"github.com/golang/glog"
)

type DeleteSender interface {
	SendDelete(ctx context.Context, topic libkafka.Topic, key []byte) error
}

type deleteSenderFunc func(ctx context.Context, topic libkafka.Topic, key []byte) error

func (d deleteSenderFunc) SendDelete(ctx context.Context, topic libkafka.Topic, key []byte) error {
	return d(ctx, topic, key)
}

func NewDeleteSender(syncProducer libkafka.SyncProducer) DeleteSender {
	return deleteSenderFunc(func(ctx context.Context, topic libkafka.Topic, key []byte) error {
		partition, offset, err := syncProducer.SendMessage(ctx, &sarama.ProducerMessage{
			Topic: topic.String(),
			Key:   sarama.ByteEncoder(key),
		})
		if err != nil {
			return errors.Wrap(ctx, err, "send message failed")
		}
		glog.V(2).
			Infof("send message successful to %s with partition %d offset %d", topic, partition, offset)
		return nil
	})
}

func NewDeleteSenderDryRun() DeleteSender {
	return deleteSenderFunc(func(ctx context.Context, topic libkafka.Topic, key []byte) error {
		glog.V(2).Infof("would delete key %s", string(key))
		return nil
	})
}
