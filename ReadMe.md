# DistributedFS

A simple distributed filesystem prototype in Go, featuring a pluggable peer-to-peer (P2P) transport layer. The current implementation uses TCP for node communication.

## Features

- **Peer-to-Peer Architecture:** Nodes communicate directly using a pluggable transport interface.
- **TCP Transport:** Out-of-the-box support for TCP-based networking.
- **Extensible Design:** Easily add new transport types (e.g., UDP, WebSockets) by implementing the `Transport` interface.
- **Simple API:** Minimal setup to start a node and accept connections.

## Getting Started

### Prerequisites

- [Go 1.24+](https://golang.org/dl/)

### Build

```sh
make build
```

### Run

```sh
make run
```

By default, the node listens on port `:3001`. You can modify the port in `main.go`.

### Test

```sh
make test
```

## Project Structure

- `main.go` — Entry point; starts a TCP listener node.
- `p2p/` — Peer-to-peer networking logic:
  - `transport.go` — Defines the `Transport` and `Peer` interfaces.
  - `tcp_transport.go` — Implements TCP-based transport.
  - `handshaker.go` — Handshake logic for new connections.
  - `tcp_transport_test.go` — Basic tests for TCP transport.

## Extending

To add a new transport (e.g., UDP), implement the `Transport` interface:

```go
type Transport interface {
    ListenAndAccept() error
}
```

## License

MIT
