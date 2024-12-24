package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// scanPort sends an HTTP request through the proxy to test if a port is open
func scanPort(proxyURL *url.URL, target string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Configure the HTTP client to use the proxy
	proxyTransport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Transport: proxyTransport,
		Timeout:   5 * time.Second, // Set a timeout for the requests
	}

	// Attempt to make an HTTP request to the target host and port
	url := fmt.Sprintf("http://%s:%d", target, port)
	resp, err := client.Get(url)
	if err != nil {
		// Port is likely closed if there is an error connecting
		return
	}
	defer resp.Body.Close()

	// Check if we got a valid response indicating the port is open
	if resp.StatusCode == 200 || resp.StatusCode == 404 || resp.StatusCode == 401 {
		fmt.Printf("Port %d on %s seems OPEN\n", port, target)
	}
}

func main() {
	// Command-line arguments
	proxyAddr := flag.String("proxy", "", "Proxy server address (http://xxx:3128)")
	targetAddr := flag.String("target", "", "Target IP behind the proxy")
	startPort := flag.Int("start", 1, "Start of port range")
	endPort := flag.Int("end", 1024, "End of port range")
	flag.Parse()

	if *proxyAddr == "" || *targetAddr == "" {
		flag.Usage()
		return
	}

	if *startPort < 1 || *endPort > 65535 || *startPort > *endPort {
		fmt.Println("Invalid port range. Ensure the range is between 1 and 65535.")
		return
	}

	// Parse the proxy URL
	proxyURL, err := url.Parse(*proxyAddr)
	if err != nil {
		fmt.Printf("Failed to parse proxy address: %v\n", err)
		return
	}

	fmt.Printf("Using proxy address: %s\n", *proxyAddr)
	fmt.Printf("Scanning ports %d to %d on target %s\n", *startPort, *endPort, *targetAddr)

	var wg sync.WaitGroup
	guard := make(chan struct{}, 15) // Limits concurrency to 15 goroutines

	// Scan the range of ports
	for port := *startPort; port <= *endPort; port++ {
		guard <- struct{}{} // Block if max goroutines are running

		wg.Add(1)
		go func(p int) {
			defer func() { <-guard }() // Release guard when done
			scanPort(proxyURL, *targetAddr, p, &wg)
		}(port)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Port scanning complete.")
}
