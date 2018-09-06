Potter
=======================
> Basic Fixture API service

![Harry Potter](./github/image.jpg?raw=true)

# Quicklinks
## Basic stuff
- [Installation](#installation)
- [Dependencies](#dependencies)
- [Configuration](#configuration)
- [Fixture example](#fixture-example)
- [License](#license)

## Installation
- Go 1.11+
- Go modules (for development)

## Dependencies

see [go.mod](./go.mod)

## Configuration

see [config.yml](./config.yml)

Example:
```yaml
pprof:
  address: :6060

metrics:
  address: :8090

log:
  level: debug
  format: console

api:
  debug: true
  address: :8080
  shutdown_timeout: 10s

fixtures:
  # good example
  - method: GET
    url: /fake-json-example
    fixture: fixture.example.json
  # good example
  - method: POST
    url: /fake-echo-example
    echo: true
  # will be ignored (combine fixture and echo unsupported)
  - method: POST
    url: /fake-echo-example
    fixture: fixture.example.json
    echo: true
  # will be ignored (combine GET and echo unsupported)
  - method: GET
    url: /fake-echo-example
    echo: true
```

### Explain

- echo - puts POST request body to response, copy Content-Length
- fixture - file of fixture
- method - GET / POST / etc..
- url - url 🙂

1. `echo` can be combined with GET
2. `echo` can be combined with `fixture`

## Fixture example

any file contents

## License

see [License](./LICENSE)