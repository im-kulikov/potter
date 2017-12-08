package config

import (
	"net/http"
	"net/url"
)

type Proxy string

func (p *Proxy) String() string {
	return string(*p)
}

func (p *Proxy) Apply() error {
	var (
		err      error
		proxyURL *url.URL
		strURL   = p.String()
	)

	// Not need if empty URL:
	if len(strURL) == 0 {
		return nil
	}

	// Try to parse url..
	if proxyURL, err = url.Parse(p.String()); err != nil {
		return err
	}

	// Set default transport:
	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	return nil
}
