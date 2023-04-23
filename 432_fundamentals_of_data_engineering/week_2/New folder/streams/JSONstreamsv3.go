package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"
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

func CalcDifference(funcname string, premem runtime.MemStats, postmem runtime.MemStats) {
	fmt.Println("\n\n ***********", funcname, " Mem Usage START ***********")
	fmt.Println("Allocated for ", funcname, " : ", postmem.Alloc-premem.Alloc)
	fmt.Println("Total Allocated for ", funcname, " : ", postmem.TotalAlloc-premem.TotalAlloc)
	fmt.Println("Total Memory from OS : ", postmem.Sys-premem.Sys)
	fmt.Println("Completed GC Cycles : ", postmem.NumGC-premem.NumGC)
	fmt.Println("\n\n ***********  ", funcname, " Mem Usage END ***********")
}

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(e *json.Decoder, slice interface{}) error {
	defer duration(track("DeSerialize"))
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(e *json.Encoder, slice interface{}) error {
	defer duration(track("Serialize"))
	return e.Encode(slice)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Println("\n\n *********** RUNTIME START ***********")
	fmt.Println("\n", msg, "Runtime : ", time.Since(start), "")
	fmt.Println("\n\n *********** RUNTIME END ***********")
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

	// Create sample data
	var i int
	var t Data
	var preRunSerialize, postRunSerialize, preRunDeserialize, postRunDeserialize runtime.MemStats
	for i = 0; i < 1000000; i++ {
		t = Data{
			Key: getString(5),
			Val: random(1, 100),
		}
		DataRecords = append(DataRecords, t)
	}

	// bytes.Buffer is both an io.Reader and io.Writer
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	///memory snapshot before
	runtime.ReadMemStats(&preRunSerialize)

	err := Serialize(encoder, DataRecords)
	if err != nil {
		fmt.Println(err)
		return
	}

	runtime.ReadMemStats(&postRunSerialize)
	CalcDifference("Serialize", preRunSerialize, postRunSerialize)
	//fmt.Print("After Serialize:", buf)

	decoder := json.NewDecoder(buf)
	var temp []Data

	runtime.ReadMemStats(&preRunDeserialize)
	err = DeSerialize(decoder, &temp)

	runtime.ReadMemStats(&postRunDeserialize)

	CalcDifference("DeSerialize", preRunDeserialize, postRunDeserialize)

	//fmt.Println("After DeSerialize:")
	//for index, value := range temp {
	//fmt.Println(index, value)
	//}

}
