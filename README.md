# go-whitelist
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go whitelist is a package that allows you to limit request accessing your service into your golang applications.


## Installation
Get Go Whitelist package on your project:

```bash
go get github.com/ihsanardanto-djoin/go-whitelist
```

## Usage
This packages provides a middleware using Echo which can be added as a global middleware or as a single route.

```go
// in server file or anywhere middleware should be registered
allowedIPs := []string{"exampleip1", "exampleip2"} // list of ip allowed in service
e.Use(gowhitelist.IPWhitelistMiddleware(allowedIPs))
```

```go
// in route file or anywhere route should be registered
router.Echo.GET("api/v1/posts", handler, gowhitelist.IPWhitelistMiddleware(allowedIPs))
```

## License
This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/MarketingPipeline/README-Quotes/blob/main/LICENSE) file for details.

## Contributors
<a href="https://github.com/ihsanardanto-djoin/go-whitelist/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=ihsanardanto-djoin/go-whitelist" />
</a>
