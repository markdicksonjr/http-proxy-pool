# http-proxy-pool

A super simple pool of proxies for HTTP requests in Go

## Usage

```go

proxyPool := pool.ProxyPool{}.Init([]string{"http://localhost:3521"}, &tls.Config{InsecureSkipVerify: true})

client := proxyPool.GetClient()
res, err := client.Get(url)

// process res, err accordingly
```

Clients are currently returned at random only
