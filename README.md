# Key-Value Store in Go

This project is an example of a distributed client-server system in Go, implementing a storage service in the form of a key-value store using Remote Procedure Call (RPC). This README will guide you through creating and running the project.

## Server

The code in server.go is the server part. This code creates a TCP server on port 1234 and offers two operations for manipulating data in the storage: Put and Get.

## Client

The code in client.go is the client part. This code is responsible for sending RPC requests to manipulate and retrieve data from the server.

### Steps to run the project

1. **Create the file structure as shown above.**

2. **Place the provided code into the respective files (server.go and client.go).**

3. **In the terminal, navigate to the server directory and execute the command:**

    ```bash
    go run server.go
    ```

4. **In another terminal, navigate to the client directory and execute the commands to manipulate the data:**

    ```bash
    go run client.go put "1" "test"
    go run client.go get "1"
    ```

## Project Structure

```plaintext
kvstore/
├── server/
│   │   └── server.go
├── client/
│   │   └── client.go
├── storage.json
```
