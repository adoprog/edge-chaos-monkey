package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
)

func startProxyServer() {
	fmt.Println("Starting server on port 8080")

	demoURL, err := url.Parse("https://edge-platform.sitecorecloud.io")
	if err != nil {
		log.Fatal(err)
	}

	proxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		reqID := uuid.New().String()[:8]
		fmt.Printf("[%s] Request received: %s %s from %s\n", reqID, req.Method, req.URL.Path, req.RemoteAddr)

		mu.RLock()
		mode := currentMode
		mu.RUnlock()

		switch mode {
		case Mode2: // Logging mode
			bodyBytes, _ := io.ReadAll(req.Body)
			req.Body.Close() //  must close
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			reqBodyString := string(bodyBytes)
			fmt.Printf("[%s]Request body: [%s]\n", reqID, reqBodyString)
		case Mode3: // Throttling mode
			rw.WriteHeader(http.StatusTooManyRequests)
			fmt.Printf("[%s] 429 Too Many Requests\n", reqID)
			return
		case Mode4: // Slow mode
			fmt.Printf("[%s] Slow response\n", reqID)
			time.Sleep(1 * time.Second)
		case Mode5: // Internal Server Error
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("[%s] 500 Internal Server Error\n", reqID)
			return
		case Mode6:
			if rand.Intn(100) < 10 { // 10% chance
				rw.WriteHeader(http.StatusTooManyRequests)
				fmt.Printf("[%s] 429 Too Many Requests\n", reqID)
				return
			}

			if rand.Intn(100) < 10 { // 10% chance
				rw.WriteHeader(http.StatusInternalServerError)
				fmt.Printf("[%s] 500 Internal Server Error\n", reqID)
				return
			}

			if rand.Intn(100) < 10 { // 10% chance
				fmt.Printf("[%s] Slow response\n", reqID)
				time.Sleep(1 * time.Second)
			}
		}

		req.Host = demoURL.Host
		req.URL.Host = demoURL.Host
		req.URL.Scheme = demoURL.Scheme
		req.RequestURI = ""

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(rw, err)
			return
		}

		for key, values := range resp.Header {
			for _, value := range values {
				rw.Header().Add(key, value)
			}
		}

		rw.WriteHeader(resp.StatusCode)
		fmt.Printf("[%s] Response status code: [%d] (%s)\n", reqID, resp.StatusCode, resp.Header.Get("Content-Type"))

		switch mode {
		case Mode2: // Logging mode
			contentEncoding := resp.Header.Get("Content-Encoding")
			respBodyBytes, err := decompressIfNeeded(resp.Body, contentEncoding)
			if err != nil {
				fmt.Printf("[%s] Failed to decompress response body: %v\n", reqID, err)
				return
			}

			resp.Body.Close() //  must close
			resp.Body = io.NopCloser(bytes.NewBuffer(respBodyBytes))
			respBodyString := string(respBodyBytes)

			fmt.Printf("[%s] Response body: [%s]\n", reqID, respBodyString)
		}

		io.Copy(rw, resp.Body)
	})

	if err := http.ListenAndServe(":8080", proxy); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}

func decompressIfNeeded(body io.ReadCloser, encoding string) ([]byte, error) {
	var reader io.Reader = body

	if encoding == "gzip" {
		gzipReader, err := gzip.NewReader(body)
		if err != nil {
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	return io.ReadAll(reader)
}
