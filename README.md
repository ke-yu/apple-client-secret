# Generate client secret for Apple Sign In
A util to generate client secret used in Apple Sign In.

Create a config.json file with the following information:
```json
{
	"service_id": "<service_id>",
	"key_id": "<key_id>",
	"team_id": "<team_id>"
}
```
Place private key under root folder as `apple.p8`:
```
-----BEGIN PRIVATE KEY-----
<private key here>
-----END PRIVATE KEY-----
```
And then run `go run main.go`
