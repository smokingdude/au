package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

func getHttpClient() (*http.Client, error) {
	var proxyUrl *url.URL
	var err error
	if isProxySet() {
		proxyUrl, err = url.Parse(config.HttpClient.Proxy)
		if err != nil {
			errorText := fmt.Sprintf("ERROR PARSING PROXY: %s", err.Error())
			return nil, errors.New(errorText)
		}
	}
	proxy := http.ProxyURL(proxyUrl)
	tlsConfig := getTLSClientConfig()
	transport := &http.Transport{Proxy: proxy, TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}
	return client, nil
}

func getTLSClientConfig() *tls.Config {
	return &tls.Config{InsecureSkipVerify: config.HttpClient.Insecure}
}
