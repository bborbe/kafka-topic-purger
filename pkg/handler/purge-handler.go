// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"context"
	"net/http"

	"github.com/bborbe/errors"
	libhttp "github.com/bborbe/http"
	libkafka "github.com/bborbe/kafka"
	"github.com/gorilla/mux"

	"github.com/bborbe/kafka-topic-purger/pkg"
)

func NewPurgeHandler(ctx context.Context, purger pkg.Purger) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		topic, ok := vars["topic"]
		if !ok {
			http.Error(resp, "parameter topic missing", http.StatusNotFound)
			return
		}
		libhttp.NewBackgroundRunHandler(ctx, func(ctx context.Context) error {
			err := purger.Purge(ctx, libkafka.Topic(topic))
			if err != nil {
				return errors.Wrapf(ctx, err, "purge topic %s failed", topic)
			}
			return nil
		}).ServeHTTP(resp, req)
	})
}
