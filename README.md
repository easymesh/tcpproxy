# tcpproxy

## Function
A simple tcpproxy proxy, a tcp transparent proxy that supports covert tls protocol <-> notls protocol.

## Usage

```
./tcpproxy --help

Usage of ./tcpproxy:
  -help
        this help
  -local string
        set local tcp protocol listen address. (default "0.0.0.0:8080")
  -protocol string
        set remote tcp protocol tcp or tcp6. (default "tcp")
  -remote string
        set remote tcp protocol proxy address.
  -tls
        enable remote tcp with tls encryption.
```

## Example
```
./tcpproxy -local 0.0.0.0:8080 -remote proxy.domain.com:8080 -protocol tcp6 -tls
```
