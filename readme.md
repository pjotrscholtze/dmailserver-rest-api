# DMailserver-rest-api
A simple rest api around `setup.sh` for Docker mail server (https://github.com/docker-mailserver/docker-mailserver).
API has a Swagger/OpenAPI definition, both in version 2.0 and 3.0.

## Features
- [x] API
  - [x] Email
    - [x] List (including Quota)
    - [x] Create
    - [x] Update
    - [x] Delete
  - [x] Fail2ban
    - [x] List (Dovecot, Postfix, and Postfix Sasl)
    - [x] Add
    - [x] Delete
- [x] API requests require an API key, [x] which can be configured.
- [x] Logging is done in JSON format.
- [x] Single executable, so can be dropped on the image and be exposed for easy usage.
- [x] Command prefix can be given allowing it to access containers running in Docker/Podman/Kubernetes.
- [x] Memory consumption is less than 10MB

## Rationale
Docker mailserver has a nice CLI tool to manage it, but this cannot be done
across a network easily, this solves this issue.

## Technical information
Since this is a small project, there are only four not generated source code
files. Everything else is generated with Go-swagger (http://goswagger.io/go-swagger/).
Since Go-swagger only has Swagger 2.0 (at this time) support, this file has been
created (see `openapi2.0.yaml`). However, a 3.0 version is also available (see 
`openapi3.0.yaml`).

## API
A description on how to call the API or generate a client for the API can be
found in the OpenAPI (https://swagger.io/resources/open-api/) yaml files (`openapi2.0.yaml`, and `openapi3.0.yaml`).

## Building the API
```
go build ./cmd/dmailserver-rest-api/main.go
```

## Running the executable
```
./main ./config/config.yaml
```

## Config
An example config can be found at `config/config.yaml`. A description of the
different parts of the configuration can be found here:
| Name          | Description                                                                 |
|---------------|-----------------------------------------------------------------------------|
| Port          | The port the API server needs to listen on.                                 |
| Host          | The interface the API server needs to listen on.                            |
| APIKey        | APIKey that needs to be used for connecting to the server.                  |
| CommandPrefix | String that will be put before all commands, use `""` when running locally. |

Since Gookit/config is used, everything that is defined in the config can also
be retrieved from an environment variable (and allows a fallback too).
How this can be done can be found here: https://github.com/gookit/config?tab=readme-ov-file#usage

An example with loading the API key from the environmental variable `API_KEY`
would be:
```
APIKey: ${API_KEY}
```

With a fallback/default value of `1234` if no environmental variable is set:
```
APIKey: ${API_KEY|1234}
```