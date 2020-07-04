package main

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/goburrow/modbus"
	"log"
	"time"
)

func main() {
	// Modbus RTU/ASCII
	handler := modbus.NewRTUClientHandler(`\\.\COM5`)
	handler.BaudRate = 9600 // Bit/秒
	handler.DataBits = 8    // データビット
	handler.Parity = "N"    // パリティ
	handler.StopBits = 1    // トップビット
	handler.SlaveId = 1     // スレーブID
	handler.Timeout = 5 * time.Second

	if err := handler.Connect(); err != nil {
		// ポートの指定間違いや、他のプロセスがポートを開いている場合
		log.Fatal("connect: ", err)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(0, 2) // データサイズ
	if err != nil {
		log.Fatal(err)
	}

	log.Println("len: ", len(results))

	humidity := results[0:2]
	temperature := results[2:4]

	log.Println(hex.EncodeToString(humidity), hex.EncodeToString(temperature))

	humidityNum := float32(int16(binary.BigEndian.Uint16(humidity))) * 1/10

	temperatureFNum := float32(int16(binary.BigEndian.Uint16(temperature))) * 1/10
	temperatureCNum := (temperatureFNum -32) * 5 / 9

	log.Fatal(humidityNum, temperatureCNum)



}
