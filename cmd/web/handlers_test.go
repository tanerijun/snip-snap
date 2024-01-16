package main

import (
	"net/http"
	"testing"

	"github.com/tanerijun/snip-snap/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/health")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, string(body), "OK")
}
