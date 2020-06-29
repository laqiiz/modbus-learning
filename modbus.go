package main

import (
	. "fmt"
	"github.com/goburrow/modbus"
	"log"
	"os"
	"time"
)

func main() {
	client := modbus.TCPClient("localhost:502")
	// Read input register 9
	results, err := client.ReadInputRegisters(8, 1)
	if err != nil {
		log.Fatal(err)
	}
	Println(string(results))

	handler := modbus.NewTCPClientHandler("localhost:502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0xFF
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	if err := handler.Connect(); err != nil {
		log.Fatal(err)
	}

	defer handler.Close()

	c2 := modbus.NewClient(handler)
	results, err = c2.ReadDiscreteInputs(15, 2)
	if err != nil {
		log.Fatal(err)
	}
	Println("[ReadDiscreteInputs]", string(results))

	results, err = c2.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
	if err != nil {
		log.Fatal(err)
	}
	Println("[WriteMultipleRegisters]", string(results))

	results, err = c2.WriteMultipleCoils(5, 10, []byte{4, 3})
	if err != nil {
		log.Fatal(err)
	}
	Println("[WriteMultipleCoils]", string(results))

}
