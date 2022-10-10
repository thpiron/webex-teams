# webex-teams

webex-teams is a Go client library for the [Cisco Webex Teams API](https://developer.webex.com/index.html).

## Usage

```go
import webexteams "github.com/thpiron/webex-teams/sdk"
```

## Documentation

Documentation for the library can be found [here](https://godoc.org/github.com/thpiron/webex-teams/sdk)

## Changes

- 2022-10-10: **Tag v0.1.0**: forked from jbogarin/go-cisco-webex-teams, added auto retry on 429, fixed pagination (RequestBy), updating schemas

## Authorization Token

Authorization token can be defined in environment variable as WEBEX_TEAMS_ACCESS_TOKEN or within the code:

```go
Client = webexteams.NewClient()
Client.SetAuthToken("<WEBEX TEAMS TOKEN>")
```

## TODO

1. Documentation
   1.1. In the code files
   1.2. In the README
2. Examples
3. Testing

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.

## Inspiration

This library is inspired by the following ones:

- [godo](https://github.com/digitalocean/godo)
- [go-github](https://github.com/google/go-github)
