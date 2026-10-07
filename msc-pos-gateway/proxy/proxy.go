package proxy

import (
	"net/http/httputil"
	"net/url"
)

func CreateReverseProxy(
	target string,
) (*httputil.ReverseProxy, error) {

	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	return reverseProxy, nil
}
