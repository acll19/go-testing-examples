```bash
# all tests
go test ./... -v

# filter
go test ./... -v -run <pattern>
```

### Test Coverage

```bash
go test ./... -v -cover

# See profile in the browser
go test ./... -v -coverprofile coverage.out

go tool cover -html=coverage.out
```