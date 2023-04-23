package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

var DataRecords []Data

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

var MIN = 0
var MAX = 26

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return temp
}

func PreRunMem() (uint64, uint64, uint64, uint32) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	//fmt.Printf("Alloc = %v b", m.Alloc)
	//fmt.Printf("\tTotalAlloc = %v b", m.TotalAlloc)
	//fmt.Printf("\tSys = %v b", m.Sys)
	//fmt.Printf("\tNumGC = %v\n", m.NumGC)
	return m.Alloc, m.TotalAlloc, m.Sys, m.NumGC
}

func PostRunMem() (uint64, uint64, uint64, uint32) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	//fmt.Printf("Alloc = %v b", m.Alloc)
	//fmt.Printf("\tTotalAlloc = %v b", m.TotalAlloc)
	//fmt.Printf("\tSys = %v b", m.Sys)
	//fmt.Printf("\tNumGC = %v\n", m.NumGC)
	return m.Alloc, m.TotalAlloc, m.Sys, m.NumGC
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

func main() {
	// Create sample data
	var i int
	var t Data
	for i = 0; i < 2; i++ {
		t = Data{
			Key: getString(5),
			Val: random(1, 100),
		}
		DataRecords = append(DataRecords, t)
	}

	// bytes.Buffer is both an io.Reader and io.Writer
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	preAlloc, preTotalAlloc, preSys, preNumGC := PreRunMem()

	err := Serialize(encoder, DataRecords)
	if err != nil {
		fmt.Println(err)
		return
	}

	postAlloc, postTotalAlloc, postSys, postNumGC := PostRunMem()
	serializeAlloc := postAlloc - preAlloc
	serializeTotalAlloc := postTotalAlloc - preTotalAlloc
	serializeSys := postSys - preSys
	serializeNumGC := postNumGC - preNumGC

	fmt.Println("Allocated for Serialize : ", serializeAlloc, "\n")
	fmt.Println("Total Allocated for Serialize : ", serializeTotalAlloc, "\n")
	fmt.Println("Total Memory from OS : ", serializeSys, "\n")
	fmt.Println("Completed GC Cycles : ", serializeNumGC, "\n")

	fmt.Print("After Serialize:", buf)
	decoder := json.NewDecoder(buf)
	var temp []Data

	err = DeSerialize(decoder, &temp)
	fmt.Println("After DeSerialize:")
	for index, value := range temp {
		fmt.Println(index, value)
	}
}
