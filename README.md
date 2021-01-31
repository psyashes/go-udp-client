# Go UDP Client
Simple UDP Client written in Go

## Usage

### Listen

```
$ nc -ul 127.0.0.1 34254
Connect client: 127.0.0.1:530901and2and3
```

### Send

```
$ go run main.go
The UDP server is 127.0.0.1:34254
> 1
> and
> 2
> and
> 3
```
