# Redis
Here's a README for your Redis clone project with Go, based on your description:

---

# Redis Clone in Go

This project is a simple Redis clone built using Go, supporting basic data structures like strings and hashes. It implements the Redis Serialization Protocol (RESP) for command parsing and communication, handles concurrent TCP connections with Go routines, and includes a data persistence mechanism using the Append-Only File (AOF) method for crash recovery.

## Features

- **Support for Strings and Hashes:** Efficiently store and retrieve data using strings and hashes.
- **RESP Parsing:** Implements the Redis Serialization Protocol (RESP) for accurate command parsing and execution.
- **Concurrency with Go Routines:** Leverages Go routines to handle multiple TCP connections concurrently, minimizing latency.
- **Data Persistence with AOF:** Implements Append-Only File (AOF) method for persistent storage, ensuring reliable data recovery after crashes.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/redis-clone-go.git
    cd redis-clone-go
    ```

2. Run the Go module to install dependencies:
    ```bash
    go mod tidy
    ```

## Running the Server

To start the Redis clone server, simply run:

```bash
go run main.go
```

The server will start on port `6379` by default (standard Redis port).

## Connecting to the Server

You can use any Redis client to connect to the server. Here's how you can do it using the `redis-cli` (Redis command-line interface):

1. Open a terminal and use the Redis CLI to connect to your server:
    ```bash
    redis-cli -h 127.0.0.1 -p 6379
    ```

2. Once connected, you can run standard Redis commands like:
    ```bash
    SET key "value"
    GET key
    HSET myhash field1 "value1"
    HGET myhash field1
    ```

## AOF Persistence

The server will save data to an Append-Only File (AOF) for persistence. If the server crashes, the data can be recovered when the server restarts. The AOF file is stored in the `data` directory.

## Contributing

Feel free to fork the repository, open issues, or create pull requests for improvements or fixes!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### Notes:
- **Go Mod:** The `go mod tidy` command ensures all dependencies are fetched, and `go run main.go` is used to start the server.
- **Connecting:** The Redis client (`redis-cli`) connects to the server on the default Redis port (6379). You can also create your own client or use a GUI like RedisInsight.

Let me know if you'd like to add or modify anything!
