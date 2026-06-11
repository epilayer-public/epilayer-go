# EpiLayer Go Client

The Genesis Cloud Go client provides the ability to manage each services resources programmatically from scripts or applications.

- [Godoc](https://pkg.go.dev/github.com/epilayer-public/epilayer-go?tab=doc)
- [Genesis Cloud API Documentation](https://developers.genesiscloud.com/)
- [How to generate an API key?](https://support.genesiscloud.com/support/solutions/articles/47001126146-how-to-generate-an-api-token-)

This repository is licensed under Mozilla Public License 2.0 (no copyleft exception) (see [LICENSE.txt](./LICENSE.txt)).

# Maintainers

This client is maintained by EpiLayer. If you encounter any problems, feel free to create issues or pull requests.

## Requirements

- [Go](https://golang.org/doc/install) >= 1.25

## Installation

```bash
go get github.com/epilayer-public/epilayer-go
```

## Getting Started

```go
import "github.com/epilayer-public/epilayer-go"

client, err := epilayer.NewEpiLayerClient(epilayer.ClientConfig{
	Token: os.Getenv("EPILAYER_TOKEN"),
})
if err != nil {
	// ...
}

// Pass nil for default options, or provide query parameters
resp, err := client.ListInstancesPaginated(ctx, nil)
if err != nil {
	// ...
}

for _, instance := range resp.Instances {
	fmt.Printf("%s %s\n", instance.ID, instance.Name)
}
```

## Examples

You can find additional examples in the [GoDoc](https://pkg.go.dev/github.com/epilayer-public/epilayer-go?tab=doc) or
check the [examples folder](./examples).

```sh
EPILAYER_TOKEN="XXXX" go run ./examples/list-instances
```

## Development or update of OpenAPI document

```sh
# Update openapi.yaml (./codegen/openapi.yaml)

# Generate code
go generate ./...
```
