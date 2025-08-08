# Testing

This project uses Go's built-in testing tools. All test files reside in the `testing` directory.

## Run all tests

From the repository root, execute:

```bash
go test ./...
```

## Run specific tests

To run tests for a single package with verbose output:

```bash
go test -v ./testing -run TestParseLocalToUTC
```

Set `-count=1` to avoid cached results when re-running tests.
