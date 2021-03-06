package internal

import (
	"context"

	"go.opentelemetry.io/otel/sdk/export/trace"
)

var _ trace.SpanExporter = &Recorder{}

// Recorder records spans being synced through the SpanSyncer interface.
type Recorder struct {
	spans []*trace.SpanData
}

// ExportSpans records spans into the internal buffer
func (r *Recorder) ExportSpans(ctx context.Context, s []*trace.SpanData) error {
	r.spans = append(r.spans, s...)
	return nil
}

// Shutdown flushes the buffer
func (r *Recorder) Shutdown(_ context.Context) error {
	_ = r.Flush()
	return nil
}

// Flush returns the current recorded spans and reset the recordings
func (r *Recorder) Flush() []*trace.SpanData {
	spans := r.spans
	r.spans = nil
	return spans
}
