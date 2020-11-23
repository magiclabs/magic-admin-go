# Magic Admin Golang SDK

The Magic Admin Golang SDK provides convenient ways for developers to interact with Magic API endpoints and an array of utilities to handle [DID Token](https://docs.magic.link/tutorials/decentralized-id).

## Table of Contents

* [Documentation](#documentation)
* [Quick Start](#quick-start)
* [Development](#development)
* [Changelog](#changelog)
* [License](#license)

## Documentation
See the [Magic doc](https://docs.magic.link/admin-sdk/go)!

## Installation
You can use Golang the SDK by specifying it as a dependency:

go.mod:

```
require github.com/magiclabs/magic-admin-go v1.0.0
```

### Prerequisites

- Go v1.10.0

## Command line utility

Command line utility is created for test purposes and can be used for decoding DID token and validating it as well as making api requests to `magic.link` for receiving the data about user, making logout, etc..
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
   decode, d  magic-cli decode --did <DID token>
   user, u    magic-cli -s <secret> user --did <DID token>
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --secret value, -s value  Secret token which will be used for making request to backend api [$MAGIC_API_SECRET_KEY]
   --help, -h                show help (default: false)
```

## Quick Start
Before you start, you will need an API secret key. You can get one from the [Magic Dashboard](https://dashboard.magic.link/). Once you have the API secret key, you can instantiate a Magic object.

Sample code for making a request of receiving metadata from the magic.link:

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
    meta, err := m.User.GetMetadataByToken("<DID_TOKEN>")
    if err != nil {
        log.Fatalf("error in processing requests: %s", err.Error())
    }

    fmt.Println(meta)
}
```

How to simply decode and validate DID token separately and read claim encoded into it:
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
        log.Fatalf("error in processing requests: %s", err.Error())
    }
    
    if err := tk.Validate(); err != nil {
        log.Fatalf("token is not valid: %v", err)
    }

    fmt.Println(tk.GetClaim())
}
```

### Configure Network Strategy
The `NewClientWithRetry` method create client with `retries`,  `retryWait`, `timeout` arguments for configuring retry options.
`NewClientWithRetry` returns `*resty.Client` instance which could be configured with any different configuration according official documentation of this library.

```golang
cl := magic.NewClientWithRetry(5, time.Second, 10 * time.Second)
m := client.New("<YOUR_API_SECRET_KEY>", cl)
```

## Development
We would love to have you contributing to this SDK. To get started, you need clone this repository and fetch dependencies.
To make sure your new code works with the existing SDK, run the test, for the new code test covering is mandatory.

```bash
make test
```

To build and install magic-cli internally in the system you need to run:

```bash
make install
```

For the building command line utility separately as binary:

```bash
make build
```

This repository is installed with [pre-commit](https://pre-commit.com/). All of the pre-commit hooks are run automatically with every new commit. This is to keep the codebase styling and format consistent.

You can also run the pre-commit manually. You can find all the pre-commit hooks [here](.pre-commit-config.yaml).

```bash
pre-commit run
```

Please also see our [CONTRIBUTING](CONTRIBUTING.md) guide for other information.

## Changelog
See [Changelog](CHANGELOG.md)

## License
See [License](LICENSE.txt)