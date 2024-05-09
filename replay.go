/**
 * Package main
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/9
 */

package main

import (
	"bytes"
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"time"
)

func Replay(args Argument, httpMethod string, bodyBytes []byte, headers http.Header, targetURL url.URL) {
	for _, target := range args.Targets {
		targetURL.Host = target.Host
		targetURL.Path = target.Path
		targetURL.Scheme = target.Scheme
		req, err := http.NewRequest(httpMethod, targetURL.String(), bytes.NewReader(bodyBytes))
		if err != nil {
			log.Printf("Failed to create request to %s: %s\n", targetURL.String(), err)
			return
		}

		for key, values := range headers {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		client := &http.Client{
			Transport: tr,
			Timeout:   3 * time.Second, // TODO: extract the timeout to the command line arguments
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to send request to %s: %s\n", targetURL.String(), err)
			return
		}
		_ = resp.Body.Close()

		log.Printf("---> Send to %s, response status: %s\n", targetURL.String(), resp.Status)
	}
}

var (
	tr *http.Transport
)

func initTransport() {
	// TODO: extract the parameters to the command line arguments
	tr = &http.Transport{
		MaxIdleConns:        100,
		MaxConnsPerHost:     50,
		MaxIdleConnsPerHost: 50,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
}
