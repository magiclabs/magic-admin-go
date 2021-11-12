# Magic Admin Golang SDK

The Magic Admin Golang SDK provides convenient ways for developers to interact with Magic API endpoints and an array of utilities to handle [DID Token](https://magic.link/docs/introduction/decentralized-id).

## Table of Contents

* [Documentation](#documentation)
* [Quick Start](#quick-start)
* [Development](#development)
* [Changelog](#changelog)
* [License](#license)

## Documentation
See the [Magic doc](https://magic.link/docs/api-reference/server-side-sdks/go)!

## Installation

The SDK requires `Golang 1.13+` and Go Modules. To make sure your project is using Go Modules, you can look for `go.mod` file in your project's root directory. If it exits, then you are already using the Go Modules. If not, you can follow [this guide](https://blog.golang.org/migrating-to-go-modules) to migrate to Go Modules.

Simply reference `magic-admin-go` in a Go program with an `import` of the SDK:

``` golang
import (
    ...
    "github.com/magiclabs/magic-admin-go"
    ...
)
```

Run any of the normal `go` commands (ex: `build`/`install`). The Go toolchain will take care of fetching the SDK automatically.

Alternatively, you can explicitly `go get` the package into a project:

```sh
go get github.com/magiclabs/magic-admin-go
```

## Command line utility

Command line utility is created for testing purposes and can be used for decoding and validating DID tokens. It also provides functionality to retrieve user info.

You can simply install it by the command:
```bash
go install github.com/magiclabs/magic-admin-go/cmd/magic-cli
```

Current available command supported:

```bash
$ magic-cli -h
NAME:
   magic-cli - command line utility to make requests to api and validate tokens

USAGE:
   magic-cli [global options] command [command options] [arguments...]

COMMANDS:
   token, t   magic-cli token [decode|validate] --did <DID token>
   user, u    magic-cli -s <secret> user --did <DID token>
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --secret value, -s value  Secret token which will be used for making request to backend api [$MAGIC_API_SECRET_KEY]
   --help, -h                show help (default: false)
```

## Quick Start

Before you start, you will need an API secret key. You can get one from the [Magic Dashboard](https://dashboard.magic.link/). Once you have the API secret key, you can instantiate a Magic object.

Sample code to retrieve user info by a [DID token](https://docs.magic.link/decentralized-id):
```golang
package main

import (
    "log"
    "fmt"

    "github.com/magiclabs/magic-admin-go"
    "github.com/magiclabs/magic-admin-go/client"
)

func main() {
    m := client.New("<YOUR_API_SECRET_KEY>", magic.NewDefaultClient())
    userInfo, err := m.User.GetMetadataByToken("<DID_TOKEN>")
    if err != nil {
        log.Fatalf("Error: %s", err.Error())
    }

    fmt.Println(userInfo)
}
```

Sample code to validate a [DID token](https://docs.magic.link/decentralized-id) and retrieve the `claim` and `proof` from the token:
```golang
package main

import (
    "log"
    "fmt"

    "github.com/magiclabs/magic-admin-go/token"
)

func main() {
    tk, err := token.NewToken("<DID_TOKEN>")
    if err != nil {
        log.Fatalf("DID token is malformed: %s", err.Error())
    }
    
    if err := tk.Validate(); err != nil {
        log.Fatalf("DID token is invalid: %v", err)
    }

    fmt.Println(tk.GetClaim())
    fmt.Println(tk.GetProof())
}
```

### Configure Network Strategy

The `NewClientWithRetry` method creates a client with `retries`,  `retryWait`, `timeout` options. `NewClientWithRetry` returns a `*resty.Client` instance which can be used with the Magic client.

```golang
cl := magic.NewClientWithRetry(5, time.Second, 10 * time.Second)
m := client.New("<YOUR_API_SECRET_KEY>", cl)
```

## Development

We would love to have you contribute to the SDK. To get started, you will need to clone this repository and fetch the dependencies.

To run the existing tests:

```bash
make test
```

To build and install magic-cli utility tool, you can run:

```bash
make install
```

To build magic-cli utility tool separately as a binary, you can run:

```bash
make build
```

Please also see our [CONTRIBUTING](CONTRIBUTING.md) guide for more information.

## Changelog
See [Changelog](CHANGELOG.md)

## License
See [License](LICENSE.txt)
