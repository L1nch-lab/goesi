module github.com/L1nch-lab/goesi

go 1.23.0

toolchain go1.24.2

require (
	github.com/antihax/goesi v0.0.0-20250326124837-837c9408dfa4
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/mailru/easyjson v0.9.0
	golang.org/x/oauth2 v0.29.0
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
)

replace github.com/antihax/goesi/optional => ./optional

replace github.com/antihax/goesi/esi => ./esi

replace github.com/antihax/goesi/meta => ./meta
