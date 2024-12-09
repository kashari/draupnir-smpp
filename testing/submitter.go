package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/kashari/draupnir/data"
	"github.com/kashari/draupnir/pdu"
)

func main() {

	// Establish TCP connection to SMPP server
	dlr, err := net.Dial("tcp", "localhost:4444")
	if err != nil {
		log.Println("Error on connection:", err)
		return
	}
	defer dlr.Close()

	// Step 1: Bind as a Transmitter
	tx := pdu.NewBindTransmitter()

	// Step 2: Prepare the Bind Transmitter request
	buff := pdu.NewBuffer(make([]byte, 0, 16))
	tx.Marshal(buff)

	PrintHexOctets(buff.Bytes())

	// Step 3: Send Bind Transmitter request to the server
	_, err = dlr.Write(buff.Bytes())
	if err != nil {
		log.Println("Error sending BindTransmitter:", err)
		return
	}

	// Step 4: Read Bind Transmitter Response from the server
	response := make([]byte, 16)
	_, err = dlr.Read(response)
	if err != nil {
		log.Println("Error reading response:", err)
		return
	}
	PrintHexOctets(response)
}

func PrintHexOctets(data []byte) {
	// The PDU Header structure is assumed to start with the following:
	// [Length (4 bytes)] [CommandID (4 bytes)] [CommandStatus (4 bytes)] [Other fields ...]

	// Extract Command ID (bytes 4-7)
	commandID := binary.BigEndian.Uint32(data[4:8])
	// Extract Command Status (bytes 8-11)
	commandStatus := binary.BigEndian.Uint32(data[8:12])

	fmt.Printf("Command ID: 0x%08X\n", commandID)
	fmt.Printf("Command Status: 0x%08X\n", commandStatus)
}

func GenerateSubmitSm() *pdu.SubmitSm {
        // build up submitSM
        srcAddr := pdu.NewAddress()
        srcAddr.SetTon(5)
        srcAddr.SetNpi(0)
        _ = srcAddr.SetAddress("00" + "522241")

        destAddr := pdu.NewAddress()
        destAddr.SetTon(1)
        destAddr.SetNpi(1)
        _ = destAddr.SetAddress("99" + "522241")

        submitSM := pdu.NewSubmitSm().(*pdu.SubmitSm)
        submitSM.SourceAddr = srcAddr
        submitSM.DestAddr = destAddr
        _ = submitSM.Message.SetMessageWithEncoding("Test base latin charset.", data.UCS2)
        submitSM.ProtocolID = 0
        submitSM.RegisteredDelivery = 1
        submitSM.ReplaceIfPresentFlag = 0
        submitSM.EsmClass = 0

        return submitSM
}