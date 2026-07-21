// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"context"
	"net/http"

	libkafka "github.com/bborbe/kafka"
	"github.com/bborbe/log"
	libsentry "github.com/bborbe/sentry"

	"github.com/bborbe/kafka-topic-purger/pkg"
	"github.com/bborbe/kafka-topic-purger/pkg/handler"
)

func CreateDeleter(dryRun bool, syncProducer libkafka.SyncProducer) pkg.DeleteSender {
	if dryRun {
		return pkg.NewDeleteSenderDryRun()
	}
	return pkg.NewDeleteSender(syncProducer)
}

func CreatePurgeHandler(
	ctx context.Context,
	sentryClient libsentry.Client,
	saramaClientProvider libkafka.SaramaClientProvider,
	deleteSender pkg.DeleteSender,
) http.Handler {
	return handler.NewPurgeHandler(
		ctx,
		pkg.NewPurger(
			sentryClient,
			saramaClientProvider,
			deleteSender,
			log.DefaultSamplerFactory,
		),
	)
}
