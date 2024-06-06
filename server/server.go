package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"sync"
)

const storageFile = "storage.json"

// Storage represents the key-value store
type Storage struct {
	sync.Mutex
	data map[string]string
}

// Args represents the arguments for Put and Get operations
type Args struct {
	Key   string
	Value string
}

// Put stores the key-value pair in the storage
func (s *Storage) Put(args *Args, reply *string) error {
	s.Lock()
	defer s.Unlock()

	s.data[args.Key] = args.Value
	*reply = "OK"

	return s.saveToFile()
}

// Get retrieves the value for a given key from the storage
func (s *Storage) Get(key *string, reply *string) error {
	s.Lock()
	defer s.Unlock()

	value, exists := s.data[*key]
	if !exists {
		return fmt.Errorf("key not found")
	}
	*reply = value
	return nil
}

// saveToFile saves the storage data to a file
func (s *Storage) saveToFile() error {
	file, err := os.Create(storageFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(s.data)
}

// loadFromFile loads the storage data from a file
func (s *Storage) loadFromFile() error {
	file, err := os.Open(storageFile)
	if err != nil {
		if os.IsNotExist(err) {
			s.data = make(map[string]string)
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&s.data)
}

func main() {
	storage := new(Storage)
	err := storage.loadFromFile()
	if err != nil {
		fmt.Println("Error loading storage:", err)
		return
	}

	rpc.Register(storage)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started at :1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
