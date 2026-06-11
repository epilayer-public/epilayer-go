package codegen

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config/types.yaml https://public-api.krs-1.epilayer.eu/compute/v1/openapi.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config/client.yaml https://public-api.krs-1.epilayer.eu/compute/v1/openapi.yaml
