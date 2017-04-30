// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package instrument

import (
	"time"

	"github.com/m3db/m3x/log"

	"github.com/uber-go/tally"
)

// Reporter reports information during runtime
type Reporter interface {
	// Report reports information during runtime
	Report()
}

// Options represents the options for instrumentation
type Options interface {
	// SetLogger sets the logger
	SetLogger(value xlog.Logger) Options

	// Logger returns the logger
	Logger() xlog.Logger

	// SetMetricsScope sets the metricsScope
	SetMetricsScope(value tally.Scope) Options

	// MetricsScope returns the metricsScope
	MetricsScope() tally.Scope

	// SetMetricsSamplingRate sets the metrics sampling rate
	SetMetricsSamplingRate(value float64) Options

	// SetMetricsSamplingRate returns the metrics sampling rate
	MetricsSamplingRate() float64

	// ReportInterval sets time between reporting many metrics within the system
	SetReportInterval(time.Duration) Options

	// GetReportInterval returns the time between reporting many metrics within the system
	ReportInterval() time.Duration
}