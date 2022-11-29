# Summarize REGAPIC roll out

```bash
cd internal/gapicgen
# emit those packages that are in a terminal state e.g. explicitly transport
# could be gRPC only, REST only, or both.
go run ./cmd/sumregapic -print_done

# emit those packages that do not have a transport selected explicltly yet.
go run ./cmd/sumregapic -print_ready
```

Note: Always pull from upstream `main` to ensure the most up to date config.