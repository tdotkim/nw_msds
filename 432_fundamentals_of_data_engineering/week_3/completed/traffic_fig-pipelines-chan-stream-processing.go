package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func contains(somemap map[string]string, chnl chan string) {
	var str string
	str = <-chnl
	fmt.Println(str)
	_, ok := somemap[str]
	// If the key exists
	if ok {
		// Do something
	} else {
		fmt.Println("contains", str)
	}
}

func writeChannel(str string, c chan string) {
	c <- str
	fmt.Println(str)
}

func main() {
	conditions_slice := make([]string, 0)
	dataChannel := make(chan string)
	//uniqueChannel := make(chan string)
	unique_map := make(map[string]string)
	//read the file line by line
	//open the file
	f, err := os.Open("Traffic_Crashes_-_Crashes.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// start reading line by line
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(rec[7])

		unique_map[rec[7]] = "pickle"
		conditions_slice = append(conditions_slice, rec[7])
		//go writeChannel(rec[7], dataChannel)
		//go contains(unique_map, dataChannel)
		//break
	}

	//fmt.Println(len(unique_map))
	go contains(unique_map, dataChannel)

}
