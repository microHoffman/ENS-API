# ENS API
Simple API for basic ENS operations.

## Available endpoints
Base URL of the server is http://localhost:1488.
- `/get-name/:address` gets ENS for given address.
- `/get-address/:name` gets address for given ENS.
- `/get-avatar/:name`  gets avatar for given ENS.
- `/get-all/:param` gets `{ address, ens, avatar }` for given `address` or `ens`.

## Run project
To run the project: `go run *.go`.