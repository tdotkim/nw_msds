package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

var (
	CourseArray  [5]MSDSCourse
	CourseSlice  []MSDSCourse
	CourseMap    [5]map[string]MSDSCourse
	CourseIds    []string
	CourseNames  []string
	CoursePreReq []string
)

type Entry struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Tel        string `json:"tel"`
	LastAccess string `json:"lastaccess"`
}

// Ignoring empty fields in JSON
type MSDSCourse struct {
	CID string `json:"course_ID"`

	CNAME string `json:"course_name"`

	CPREREQ string `json:"prerequisite"`
}

func main() {

	//add course raw data
	MyGenData(5)

	for i := 0; i < 5; i++ {
		CourseArray[i] = MSDSCourse{CID: CourseIds[i], CNAME: CourseNames[i], CPREREQ: CoursePreReq[i]}
		CourseSlice = append(CourseSlice, MSDSCourse{CID: CourseIds[i], CNAME: CourseNames[i], CPREREQ: CoursePreReq[i]})
		CourseMap[i] = map[string]MSDSCourse{fmt.Sprintf("Course #%d", i+1): {CID: CourseIds[i], CNAME: CourseNames[i], CPREREQ: CoursePreReq[i]}}
	}

	CourseArrayJson, err := json.MarshalIndent(&CourseArray, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nCourse Json from Array:\n%v", string(CourseArrayJson))
	CourseSliceJson, err := json.MarshalIndent(&CourseSlice, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nCourse Json from Slice:\n%v", string(CourseSliceJson))
	CourseMapJson, err := json.MarshalIndent(&CourseMap, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nCourse Json from Map:\n%v", string(CourseMapJson))
}

// a helper function to generate some data
func MyGenData(recordcount int) {

	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= recordcount; i++ {
		CourseIds = append(CourseIds, fmt.Sprintf("%d", rand.Intn(rand.Intn(999-100+1)+100)))
		CourseNames = append(CourseNames, fmt.Sprintf("MSDS %d", rand.Intn(rand.Intn(500-100+1)+100)))
		CoursePreReq = append(CoursePreReq, fmt.Sprintf("MSDS %d", rand.Intn(rand.Intn(400-50+1)+100)))

	}

}
