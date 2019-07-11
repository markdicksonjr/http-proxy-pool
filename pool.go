package pool

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"net/url"
)

var DefaultProxyUrls []string
var DefaultTlsConfig *tls.Config

type ProxyPool struct {
	urls      []string
	tlsConfig *tls.Config
}

func (p *ProxyPool) WithUrls(urls []string) *ProxyPool {
	p.urls = urls
	return p
}

func (p *ProxyPool) WithTlsConfig(tlsConfig *tls.Config) *ProxyPool {
	p.tlsConfig = tlsConfig
	return p
}

func (p *ProxyPool) GetClient() (*http.Client, error) {
	var proxy func(*http.Request) (*url.URL, error)

	// use the default proxy urls if none were given for this pool
	if len(p.urls) == 0 && len(DefaultProxyUrls) > 0 {
		p.urls = DefaultProxyUrls
	}

	// use the default TLS config, if none was given for this pool
	if p.tlsConfig == nil && DefaultTlsConfig != nil {
		p.tlsConfig = DefaultTlsConfig
	}

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
