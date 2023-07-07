# Changelog


## `1.0.0` - 07/05/2023

[//]: # ()
[//]: # (### Added)

[//]: # ()
[//]: # (- PR-#11: Add Magic Connect Admin SDK support for Token Validation )

[//]: # (    - [Security Enhancement]: Validate `aud` using Magic client ID.)

[//]: # (    - Pull client ID from Magic servers in client constructor.)

### Summary
- 🚀 **Added:** Magic Connect developers can now use the Admin SDK to validate DID tokens.
- ⚠️ **Changed:** After creating the Magic instance, it is now necessary to call a new initialize method for Magic Connect developers that want to utilize the Admin SDK.
- 🛡️ **Security:** Additional validation of `aud` (client ID) is now being done during initialization of the SDK.

### Developer Notes

#### 🚀 Added

Magic Connect developers can now use the Admin SDK to validate DID tokens.

**Details**
There is full support for all `Token` SDK methods for MC. This is intended to be used with client side [`magic-js`](#) SDK which will now emit an `id-token-created` event with a DID token upon login via the [`connectWithUI`](#) method.

This functionality is replicated on our other SDKs on Python and Ruby.

#### ⚠️ Changed

To validate tokens, a Magic clientId is now required in the `validate` method. 
Initializing the client will pull clientId from the Magic servers.
Alternatively, you can get client ID from the magic dashboard and pass it in directly.

**Previous Version**
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

**New Version**
```golang
package main

import (
    "log"
    "fmt"

	"github.com/magiclabs/magic-admin-go/client"
    "github.com/magiclabs/magic-admin-go/token"
)

func main() {

	c, err := client.New("<YOUR_API_SECRET_KEY>", magic.NewDefaultClient())

	if err != nil {
		log.Fatalf("Unable to initialize client: %s", err.Error())
	}
	
	tk, err := token.NewToken("<DID_TOKEN>")
    if err != nil {
        log.Fatalf("DID token is malformed: %s", err.Error())
    }
    
    if err := tk.Validate(c.ClientInfo.ClientId); err != nil {
        log.Fatalf("DID token is invalid: %v", err)
    }

    fmt.Println(tk.GetClaim())
    fmt.Println(tk.GetProof())
}
```

### 🛡️ Security

#### Client ID Validation

Additional validation of `aud` (client ID) is now being done while validating DID tokens. This is for both Magic Connect and Magic Auth developers.


### 🚨 Breaking

* Client initialization now makes a call to Magic servers to fetch `clientId` and will now return an error if there is an issue communicating to Magic's servers. 
* The `validate` method now takes in a clientId and validates it against the `aud` field in the DID token.

## `0.2.0`

#### Changed

- <PR-#10>
    Dependency update for module.

#### Added

- <PR-#10>
    Added multi-chain wallet for metadata retrieval calls.