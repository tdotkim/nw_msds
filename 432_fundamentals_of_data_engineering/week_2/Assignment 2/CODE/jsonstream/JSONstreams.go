package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

var 
(
	DataRecords []Data
	MIN = 0
	MAX = 26
	RecordGen = []int{100,1000,10000}
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}


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

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}


func PrimaryTest(recordCount int){
		// Create sample data
		var i int
		var t Data
		for i = 0; i < recordCount; i++ {
			t = Data{
				Key: getString(5),
				Val: random(1, 100),
			}
			DataRecords = append(DataRecords, t)
		}
	
		// bytes.Buffer is both an io.Reader and io.Writer
		buf := new(bytes.Buffer)
	
		encoder := json.NewEncoder(buf)
		err := Serialize(encoder, DataRecords)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("After Serialize:", buf)
	
		decoder := json.NewDecoder(buf)
		var temp []Data
		err = DeSerialize(decoder, &temp)
		fmt.Println("After DeSerialize:")
		for index, value := range temp {
			fmt.Println(index, value)
		}
}

func main() {

	var durationTracker []string
	var memTracker []string
	for _,val := range RecordGen{
		//start tracking time
		start := time.Now()
		PrimaryTest(val)
			//run complete
		duration := fmt.Sprintf("\n\nTime to run %d records was:\n%v",val,time.Since(start))
		durationTracker = append(durationTracker, duration)
		memTracker = append(memTracker, PrintMemUsage(val))
	}

	//loop over and print time to run results problem 1.2
	for _,runtime := range durationTracker{
		//enter println here to see the runtime
		fmt.Println(runtime)
		
	}

	//loop over and print mem needed to run results problem 1.3
	for _,memUsed:= range memTracker{
		//enter println here to see the memory used
		fmt.Println(memUsed)
	}



}


func PrintMemUsage(records int) string{
	var m runtime.MemStats
	var mainString string
	runtime.ReadMemStats(&m)
	mainString = 
	fmt.Sprintf("\nMemory used for %d records:\nAlloc = %v MiB\tTotalAlloc = %v MiB\tSys = %v MiB\tNumGC = %v\n",records, bToMb(m.Alloc),bToMb(m.TotalAlloc),bToMb(m.Sys),m.NumGC)
	return mainString
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}