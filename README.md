# pas2-go-health-probes

Run named context-aware health probes with bounded concurrency.

Part of **Pas2 Polyglot** by [1Pas2](https://github.com/1Pas2): independent small projects in PHP, Python, Go and C#. No runtime dependencies outside the language standard library. Initial version: 0.1.0.

## Requirements and installation

Go 1.23+. Use a local checkout with `go test ./...`. After GitHub publication, applications can use `go get github.com/1Pas2/pas2-go-health-probes@main`; no version tag is claimed.

## Example

```go
package main

import (
	"context"
	"fmt"
	lib "github.com/1Pas2/pas2-go-health-probes"
)

func main() {
	results, err := lib.Run(context.Background(), map[string]lib.Probe{"database": func(context.Context) error { return nil }}, 2)
	fmt.Println(results, err)
}
```

## API and limits

Run(ctx,probes,concurrency) returns name-sorted results with per-probe errors; invalid configuration returns a top-level error. Callbacks must honor context and must not panic. No forced termination or default timeout. The probes map must remain unchanged during Run. Panics retain normal Go panic behavior.

## Tests and example

From the repository root:

```sh
go test ./...
go vet ./...
go run ./examples/basic
```

Tests use the standard library or a dependency-free executable check harness. A failing assertion exits nonzero. The example is self-contained and uses local in-memory or temporary data. CI runs the same checks on Linux. No registry releases, downloads, users or production deployments are claimed.

## Contributing

Open a focused issue or pull request with reproduction steps and a behavioral test. Run the commands above before submitting. Do not include credentials or private data in reports.

## Ecosystem

[Browse Pas2 projects](https://github.com/1Pas2?tab=repositories&q=pas2-). Each repository is usable independently. The earlier TypeScript series and this polyglot series have distinct repository names.

## License

[MIT](LICENSE) © 2026 1Pas2.

1Pas2
