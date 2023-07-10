# code-uv
A code scanner to scan vulnerabilities

# Build and Test

```
go build -o ./bin/ ./cmd/go-scanner
go vet  -vettool=./bin/go-scanner  ./cmd/ast/example/...
```