package main

import (
	"net/url"
	"strings"

	"github.com/higress-group/proxy-wasm-go-sdk/proxywasm"
	"github.com/higress-group/wasm-go/pkg/wrapper"
)

// maybeTriggerBatch detects batch-type requests via header or query and emits a batch trigger metric.
// Detection rules:
// - Header: X-AI-Batch: true
// - Query:  ?batch=true
func maybeTriggerBatch(ctx wrapper.HttpContext) {
	val, _ := proxywasm.GetHttpRequestHeader("X-AI-Batch")
	if strings.EqualFold(val, "true") {
		metricsBatchTriggered()
		return
	}
	p, err := url.Parse(ctx.Path())
	if err != nil {
		return
	}
	q := p.Query().Get("batch")
	if strings.EqualFold(q, "true") {
		metricsBatchTriggered()
	}
}