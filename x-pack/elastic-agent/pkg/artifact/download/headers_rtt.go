// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package download

import (
	"fmt"
	"net/http"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/release"
)

// Headers is a collection of headers used in download client.
var Headers = map[string]string{
	"User-Agent": fmt.Sprintf("Beat elastic-agent v%s", release.Version()),
}

// WithHeaders injects specified headers into a download client.
func WithHeaders(rtt http.RoundTripper, headers map[string]string) http.RoundTripper {
	if rtt == nil {
		rtt = http.DefaultTransport
	}
	return &rttWithHeaders{target: rtt, headers: headers}
}

type rttWithHeaders struct {
	target  http.RoundTripper
	headers map[string]string
}

func (r *rttWithHeaders) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range r.headers {
		req.Header.Add(k, v)
	}
	return r.target.RoundTrip(req)
}
