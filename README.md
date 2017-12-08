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
- Go 1.9
- Dep (for development) [Vendoring tool](https://github.com/golang/dep#usage)

## Dependencies

see [Gopkg.lock](./Gopkg.lock)

## Configuration

see [config.yml](./config.example.yml)

Example:
```yaml
# Uncomment if needed:
#proxy: http://proxy.url
api:
  - method: GET
    url: /fake-example
    fixture: fixture.example.json
  - method: POST
    url: /fake-example
    fixture: fixture.example.json
  - method: HEAD
    url: /fake-example
    fixture: fixture.example.json

```

## Fixture example

any file contents

## License

see [License](./LICENSE)