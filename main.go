package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/goburrow/modbus"
	"log"
	"time"
)

func main() {

	h := modbus.NewRTUClientHandler(`COM3`) // RTU, COM3ポート
	h.BaudRate = 9600                       // Bit/秒
	h.DataBits = 8                          // データビット
	h.Parity = "N"                          // パリティ
	h.StopBits = 1                          // トップビット
	h.SlaveId = 1                           // スレーブID
	h.Timeout = 5 * time.Second

	if err := h.Connect(); err != nil {
		// ポートの指定間違いや、他のプロセスがポートを開いている場合
		log.Fatal("connect: ", err)
	}
	defer h.Close()

	mc := modbus.NewClient(h)
	results, err := mc.ReadHoldingRegisters(0, 2) // データサイズ
	if err != nil {
		log.Fatal(err)
	}

	humidity := results[0:2]
	temperature := results[2:4]

	fmt.Println("湿度", "温度")
	fmt.Println(hex.EncodeToString(humidity), hex.EncodeToString(temperature))

	humidityNum := float32(binary.BigEndian.Uint16(humidity)) * 1 / 10
	temperatureFNum := float32(binary.BigEndian.Uint16(temperature)) * 1 / 10
	temperatureCNum := (temperatureFNum - 32) * 5 / 9

	fmt.Println(humidityNum, temperatureCNum)
}
