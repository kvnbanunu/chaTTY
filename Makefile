SERVER = cmd/chaTTY-server.go
SERVER_TARGET = bin/chaTTY-server
CLIENT = cmd/chaTTY.go
CLIENT_TARGET = bin/chaTTY

build: clean
	go build -o $(SERVER_TARGET) $(SERVER)
	go build -o $(CLIENT_TARGET) $(CLIENT)

server:
	go run $(SERVER)

client:
	go run $(CLIENT)

clean:
	rm -rf bin
