# Editor-api

## About the project
DrawFlow is a tool to learn to program visually

## Client
* [editor-client](https://github.com/24Roger/editor-client.git)
## Getting Started
### Prerequisites:
To run the api on your local system, you need the following:
* golang
* Docker

### Try local
1. Clone the repo
```sh
git clone https://github.com/24Roger/editor-api.git
```
2. Replace in graph.go file line 21 dgo.DialCloud for grpc.Dial
```go
conn, err := dgo.DialCloud(DATABASE_URL, API_KEY)

conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
```
3. Run the following command
```sh
docker run --rm -it -p "8080:8080" -p "9080:9080" -p "8000:8000" -v ~/dgraph:/dgraph "dgraph/standalone:v21.12.0"
```
4. Run the following command
```sh
go run cmd/main.go
```