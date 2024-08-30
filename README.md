# go-channels
Demonstration of go concurrency

## Examples

### No concurrency
```shell
go run 1-nochan/main.go
```

### Using unbuffered go channels
```shell
go run 2-chans/main.go
```

### Using buffered go channels
```shell
go run 3-bufferedchans/main.go
```

### Using wait-groups for multiple parallel instances
```shell
go run 4-multichan/main.go
```
