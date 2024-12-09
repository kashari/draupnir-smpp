package main

import (
	"log"
	"net"

	"github.com/kashari/draupnir/pdu"
)

func main() {
	// Start TCP server on port 9999
	listen, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	defer listen.Close()

	log.Println("SMPP client-handler listening on port 4444")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("New connection established, handling SMPP session: ", conn.RemoteAddr().String())
	bp, err := pdu.Parse(conn)
	if err != nil {
		log.Println("An error occurs on decode: ", err)
	}

	log.Println(bp)
	buff := pdu.NewBuffer(make([]byte, 0, 64))
	bp.GetResponse().Marshal(buff)

	conn.Write(buff.Bytes())
}
