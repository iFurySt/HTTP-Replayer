/**
 * Package main
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/9
 */

package main

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
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

		// FIXME: use the http.Transport to reuse connections
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to send request to %s: %s\n", targetURL.String(), err)
			return
		}
		_ = resp.Body.Close()

		log.Printf("---> Send to %s, response status: %s\n", targetURL.String(), resp.Status)
	}
}
