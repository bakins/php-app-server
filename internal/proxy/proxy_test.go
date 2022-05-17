package proxy_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"github.com/bakins/php-app-server/internal/proxy"
)

func TestServer(t *testing.T) {
	target := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				data := make([]byte, 16*1024)

				w.WriteHeader(http.StatusOK)

				_, err := w.Write(data)
				assert.NoError(t, err)
			},
		),
	)

	target.EnableHTTP2 = true
	target.StartTLS()

	defer target.Close()

	p, err := proxy.New(
		proxy.WithTransport(target.Client().Transport),
		proxy.WithLogger(zaptest.NewLogger(t)),
	)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	go func() {
		err = p.Run(ctx)
		assert.NoError(t, err)
	}()

	addr, err := p.WaitForAddress(ctx)
	require.NoError(t, err)

	client := &http.Client{
		Transport: newClientTransport("http://" + addr.String()),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target.URL+"/hello/world", nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, 16*1024, len(data))
}

type clientTransport struct {
	transport *http.Transport
}

func newClientTransport(proxy string) *clientTransport {
	u, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}

	transport := http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			return u, nil
		},
	}

	t := clientTransport{
		transport: &transport,
	}

	return &t
}

func (t *clientTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.URL.Scheme == "http" {
		return t.transport.RoundTrip(r)
	}

	// shallow clone
	r = r.WithContext(r.Context())
	r.Header.Set(proxy.SchemeHeader, r.URL.Scheme)
	r.URL.Scheme = "http"

	return t.transport.RoundTrip(r)
}
