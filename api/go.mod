module github.com/openbao/openbao/api

// The Go version directive for the api package should normally only be updated when
// code in the api package requires a newer Go version to build.  It should not
// automatically track the Go version used to build Vault itself.  Many projects import
// the api module and we don't want to impose a newer version on them any more than we
// have to.
go 1.21

toolchain go1.22.2

replace github.com/openbao/openbao/sdk => ../sdk

require (
	github.com/cenkalti/backoff/v3 v3.2.2
	github.com/go-jose/go-jose/v3 v3.0.1
	github.com/go-test/deep v1.1.0
	github.com/hashicorp/errwrap v1.1.0
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/go-hclog v1.4.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-retryablehttp v0.7.1
	github.com/hashicorp/go-rootcerts v1.0.2
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.7
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2
	github.com/hashicorp/hcl v1.0.1-vault-5
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/openbao/openbao/sdk v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.22.0
	golang.org/x/time v0.0.0-20220411224347-583f2d630306
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

retract [v1.0.1, v1.12.0]
