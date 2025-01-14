// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package observabilityzap

import (
	"io"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// TracerProviderCloser is used to wrap a trace.TracerProvider with an io.Closer to use on shutdown.
type TracerProviderCloser interface {
	trace.TracerProvider
	io.Closer
}

// Start creates a Zap logging exporter for Opentelemetry traces and returns
// the exporter. The exporter implements io.Closer for clean-up.
func Start(logger *zap.Logger) TracerProviderCloser {
	exporter := newZapExporter(logger)
	tracerProviderOptions := []sdktrace.TracerProviderOption{
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporter)),
	}
	tracerProvider := newTracerProviderCloser(sdktrace.NewTracerProvider(tracerProviderOptions...))
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return tracerProvider
}
