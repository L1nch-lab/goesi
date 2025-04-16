# GoESI Fork by L1nch-lab

This is a maintained and structured fork of the original [antihax/goesi](https://github.com/antihax/goesi) client for the EVE Online ESI API. It includes modernizations, bugfixes, and tooling to support secure, maintainable and extensible use in production or research environments.

## Features

- Full OpenAPI 3 coverage of the ESI API
- OAuth2 authentication using EVE SSO
- Versioned and modular API clients
- CLI and Web-ready usage
- Tools for code hygiene and automation (`golangci-lint`, `goimports`, fix scripts)
- Beta-stage quality: actively maintained and extended

## Setup

```bash
git clone https://github.com/L1nch-lab/goesi.git
cd goesi
go mod tidy
```

> Optional: Use `tools/fix_advanced.go` and `tools/replace_replace.go` to perform automated code refactors and hygiene passes.

## Example: Creating a New Client

```go
import (
  "context"
  "fmt"
  "net/http"

  "github.com/L1nch-lab/goesi"
)

func main() {
  client := goesi.NewAPIClient(nil, "MyApp (contact@domain.com)")
  status, _, err := client.ESI.StatusApi.GetStatus(context.Background(), nil)
  if err != nil {
    panic(err)
  }
  fmt.Println("Players online:", status.Players)
}
```

## OAuth2 Support

Supports full OAuth2 flow including token refresh. Tokens are passed via context:

```go
ctx := context.WithValue(context.Background(), goesi.ContextOAuth2, tokenSource.Token)
data, _, err := client.V1.CharacterApi.GetCharactersCharacterId(ctx, characterID, nil)
```

## Tools Directory

Contains CLI tooling to clean, refactor and modernize the codebase. Key tools:

- `fix_advanced.go`: Performs AST-based automated fixes (e.g. `defer Close`, `ReplaceAll`, `EqualFold`)
- `replace_replace.go`: Legacy string-replacer for common issues

Run with:

```bash
go run tools/fix_advanced.go
```
## Etiquette

* Create a descriptive user agent so CCP can contact you (preferably on devfleet slack).
* Obey Cache Timers.
* Obey error rate limits: https://developers.eveonline.com/blog/article/error-limiting-imminent

## Authenticating

Register your application at https://developers.eveonline.com/ to get your secretKey, clientID, and scopes.

Obtaining tokens for client requires two HTTP handlers. One to generate and redirect
to the SSO URL, and one to receive the response.

It is mandatory to create a random state and compare this state on return to prevent token injection attacks on the application.

## Development

```bash
golangci-lint run
```

CI/CD workflows include:
- Linting (`golangci-lint`)
- Security checks (`gosec`)
- Dependency updates (`dependabot`)
- Release automation via tags

## Contributing

PRs and forks welcome. Please lint and format your code before submitting. All tools are included.

## License

MIT – See original license under [antihax/goesi](https://github.com/antihax/goesi).

## Maintainer

L1nch-lab – https://github.com/L1nch-lab
> 💬 `l1nch_` on [Discord]