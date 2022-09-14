package main

import (
	"encoding/hex"
	"fmt"
)

func IDplusDLC(frames [][16]byte, pos int) (byte, byte) {
	var id byte
	var DLC byte
	id = frames[pos][0] + frames[pos][1] // 11 id bitar DLC = frames[i][4]
	DLC = frames[pos][4]                 // 5:e byten är data length code byten
	return id, DLC
}

func getData(frames [][16]byte, DLC byte, pos int) string {
	var FrameDataByte []byte
	for j := 0; j < int(DLC); j++ {
		FrameDataByte = append(FrameDataByte, frames[pos][8+j])
	}
	FrameDataString := hex.EncodeToString(FrameDataByte)
	return FrameDataString
}

func getSpeedData(frames [][16]byte, pos int) float64 {
	var SpeedData []byte
	for j := 12; j != 10; j-- { // little-endian
		SpeedData = append(SpeedData, frames[pos][j])
	}
	SpeedDataInt := (int(SpeedData[0]) << 8) + (int(SpeedData[1]))
	// Speed = value * scale + offset
	Speed := float64(SpeedDataInt)*0.00391 + 0
	return Speed
}

func main() {

	var encodedCANFrames = [][16]byte{
		{0xf8, 0x02, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x7b, 0x82, 0x1a, 0x2a, 0x68, 0xd0, 0xb0, 0x00},
		{0x88, 0x04, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x50, 0x68, 0xc0, 0x00, 0x00, 0x00, 0x00},
		{0x68, 0x03, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xc0, 0xde, 0x0f, 0xee, 0x28, 0x09, 0x00, 0x23},
		{0x08, 0x01, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0xe3, 0xd2, 0x81, 0x71, 0xa6, 0x00, 0xff, 0xbb},
	}

	fmt.Println("Part1 Start")
	for i := 0; i < len(encodedCANFrames); i++ {
		id, DLC := IDplusDLC(encodedCANFrames, i)
		FrameData := getData(encodedCANFrames, DLC, i)

		fmt.Println("Frame", i, ":", id, "\t", DLC, "Data bytes:", FrameData)
	}
	fmt.Println("End Part1")

	fmt.Println("Part2 Start")
	for i := 0; i < len(encodedCANFrames); i++ {
		id, DLC := IDplusDLC(encodedCANFrames, i)
		if DLC > 5 {
			Speed := getSpeedData(encodedCANFrames, i)
			fmt.Println("Frame", i, ":", id, "\t", "Speed:", Speed, "km/h")
		}
	}
	fmt.Println("End Part2")
}
