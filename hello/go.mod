module example/hello

go 1.19

// This added here using go mod tidy, caches to pkg/mod in GOPATH
// to remove the cache, do go clean -modcache
require github.com/google/go-cmp v0.5.9
