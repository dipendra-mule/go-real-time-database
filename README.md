# Go Real-Time Database

A simple, lightweight real-time database and pub/sub system written in Go, focusing on live data synchronization between clients using WebSocket connections.

## Features

- Real-time data sync across connected clients
- Publish/subscribe data channels (topics)
- In-memory key-value store
- RESTful HTTP API for CRUD operations on data
- WebSocket support for live updates
- Written entirely in Go, no external dependencies

## Use Cases

- Real-time dashboards
- Collaborative editing
- Chat applications
- IoT device state monitoring

## Getting Started

### Prerequisites

- Go 1.18 or higher
- (Optional) Docker for containerized deployment

### Clone the Repository

```bash
git clone <repository-url>
cd go-real-time-database
```

### Running Locally

```bash
go run main.go
```

By default, the server will start on `localhost:8080`.

### Docker

```bash
docker build -t go-real-time-db .
docker run -p 8080:8080 go-real-time-db
```

## API Usage

### HTTP Endpoints

- `GET /data/{key}`: Get value by key
- `POST /data/{key}`: Set value for a key (`{"value": <your-value>}`)
- `DELETE /data/{key}`: Delete key-value pair

### WebSocket Endpoint

- `ws://<host>:8080/ws`
  - Subscribe to topics or keys for live updates.
  - Receive JSON payloads when data changes.
  - Broadcast support: all subscribed clients receive updates in real-time.

#### Example WebSocket Message

```json
{
  "action": "subscribe",
  "key": "temperature"
}
```

When `temperature` changes, clients receive:

```json
{
  "key": "temperature",
  "value": <new-value>
}
```

## Project Structure

- `main.go` - Entry point & server setup
- `internal/`
  - `store.go` - In-memory datastore logic
  - `ws.go` - WebSocket handlers & pub/sub logic
  - `api.go` - HTTP handler implementation

## Environment Variables

| Variable | Default | Description           |
| -------- | ------- | --------------------- |
| PORT     | 8080    | Server listening port |

## Example Client

```go
conn, _, _ := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
defer conn.Close()

subscribe := map[string]string{
    "action": "subscribe",
    "key":    "temperature",
}
conn.WriteJSON(subscribe)
for {
    _, msg, _ := conn.ReadMessage()
    fmt.Println("Update:", string(msg))
}
```

## License

MIT License

