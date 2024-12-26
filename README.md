# GoPose

A Go implementation of [Spose (Squid Pivoting Open Port Scanner)](https://github.com/aancw/spose), originally created by [Aan](https://github.com/aancw). This tool allows you to find open ports on target hosts behind Squid proxy servers.

## About

This is a high-performance Golang port of the original Python-based Spose tool, optimized for speed through aggressive concurrency and efficient connection handling. It scans ports through HTTP proxies significantly faster than the original, making it ideal for rapid internal network enumeration through proxy servers.

## Features

- Written in Go for improved performance and concurrency
- Scan ports through HTTP proxy servers
- Concurrent scanning with controlled goroutines (15 simultaneous connections)
- Customizable port range scanning
- Built-in timeout handling
- Simple command-line interface

## Installation
```bash
# Clone the repository
git clone https://github.com/karthikuj/gopose.git
cd gopose
# Build the binary
go build -o gopose
```

## Usage

```bash
./gopose -target <target_host> -start <start_port> -end <end_port> -proxy <proxy_url>
```

### Parameters

- `-target <target_host>`: The target host to scan.
- `-start <start_port>`: The start port of the port range to scan.
- `-end <end_port>`: The end port of the port range to scan.
- `-proxy <proxy_url>`: The URL of the proxy server to use.

### Example

```bash
./gopose -target 10.10.11.131 -start 1 -end 1024 -proxy http://10.10.11.131:3128
```

This command will scan ports 1 to 1024 on the target host 10.10.11.131 through the proxy server at http://10.10.11.131:3128.

# Credits

- Original Spose tool: [aancw/spose](https://github.com/aancw/spose)
- Original author: [Aan](https://github.com/aancw)
