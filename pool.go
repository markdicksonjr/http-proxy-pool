package pool

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"net/url"
)

type ProxyPool struct {
	urls      []string
	tlsConfig *tls.Config
}

// Init initializes the pool with proxies and options
func (p *ProxyPool) Init(urls []string, tlsConfig *tls.Config) error {
	p.urls = urls
	p.tlsConfig = tlsConfig
	return nil
}

func (p *ProxyPool) GetClient() (*http.Client, error) {
	var proxy func(*http.Request) (*url.URL, error)

	// IF there's an available proxy pool, assign a proxy from it
	if len(p.urls) > 0 {
		randIndex := randInt(0, len(p.urls))

		urlParsed, err := url.Parse(p.urls[randIndex])
		if err != nil {
			return nil, err
		}
		proxy = http.ProxyURL(urlParsed)
	}

	return &http.Client{Transport: &http.Transport{
		TLSClientConfig: p.tlsConfig,
		Proxy:           proxy,
	}}, nil
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
