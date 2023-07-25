# Nerinyan Core
Nerinyan core is a library for the Nerinyan API.
It contains functions for managing Nerinyan websites.

## Examples
#### User Authentication
```go
a, err := auth.LoginWithAuth("username", "password") // IMPORTANT) DO NOT PUSH YOUR PASSWORD
err = a.Login()
err = a.Refresh()

exp := a.ExpiredAt()

fmt.Println(a.Token.AccessToken) // Print Access Token
```

## License
Copyright © [thftgr](https://github.com/thftgr), [ZEEE](https://github.com/zeee2)\
This project is licensed under the [GPL-3.0 License](https://tldrlegal.com/license/gnu-general-public-license-v3-(gpl-3)).  Please see [the license file](LICENSE) for more information.
