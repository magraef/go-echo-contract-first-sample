# go-echo-contract-first-sample

This is a go sample application that uses the contract first approach to provide an [echo](https://github.com/labstack/echo) api using [oapi-codegen](https://github.com/deepmap/oapi-codegen).

This app uses [embed](https://pkg.go.dev/embed) to embed a swagger-ui and the openapi spec into the binary and [wire](https://github.com/google/wire) for dependency injection.

## Build and Run

### Build and run app using docker

```shell
docker build -t go-echo-contract-first-sample:latest -f build/Dockerfile .
docker run --rm -p 8080:8080 -t go-echo-contract-first-sample:latest
```

### Swagger UI / Open Api V3
Swagger UI / Open Api V3 can be accessed via http://localhost:8080/q/swagger-ui.

### Health endpoint
The app exposes a health endpoint at http://localhost:8080/q/health. 
## Configure the application

The application can be configured via [envconfig](https://github.com/kelseyhightower/envconfig). The following list shows the corresponding env var with default value.

Envvar name | default value |^
:--- | ---: | 
`APP_PORT` | `8080`
`APP_API_BASE_URI` | `api`
`APP_DEBUG` | `false`
`APP_JSON_LOG` | `false`

# Local Dev

The app uses Golang >= 1.17. It is build as a single golang module.

## Generate REST endpoints

This app uses a _contract first_ approach and write OpenAPI specs and use these to generate server side 
implementations using [github.com/deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen).

Run 

```
$ go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.8.2
```

to install the generator.

If you make changes to any OpenAPI specs, re-run the generator by pointing `go generate` to the respective
`router.go` file. Make sure you commit all generated code.

## Generate dependency injection code

This app uses [github.com/google/wire](https://github.com/google/wire) to generate dependency injection code. Use

```
$ go install github.com/google/wire/cmd/wire@v0.5.0
```

to install the code generator and run `go generate` on all respective `wire.go` files. As stated above: 
Make sure you commit all generated files.
