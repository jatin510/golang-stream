package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type FileServer struct {
}

func (fs *FileServer) start() {
	ln, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go fs.readLoop(conn)
	}
}

func (fs *FileServer) readLoop(conn net.Conn) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		// n = no. of bytes
		file := buf[:n]
		fmt.Println("file", file)
		fmt.Printf("Received %d bytes over the network\n", n)

	}
}

func sendFile(size int) error {
	file := make([]byte, size)

	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		return err
	}

	n, err := conn.Write(file)
	if err != nil {
		return err
	}

	fmt.Printf("written %d bytes over the network\n", n)
	return nil
}

func main() {
	go func() {
		time.Sleep(4 * time.Second)
		sendFile(3000)
	}()

	server := &FileServer{}

	server.start()
}
