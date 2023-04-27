package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type MSDSCourse struct {
	CID string `json:"courseI_D`

	CNAME string `json:"course_name"`

	CPREREQ string `json:"prerequisite"`
}

// JSONFILE resides in the current directory
var CSVFILE = "./data.csv"

type MSDSCourseCatalog []MSDSCourse

var data = MSDSCourseCatalog{}
var index map[string]int

const PORT = ":1234"

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Get telephone
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramStr)
	if len(paramStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not found: "+r.URL.Path)
		return
	}

	log.Println("Serving:", r.URL.Path, "from", r.Host)

	courseID := paramStr[2]
	err := DeleteEntry(courseID)
	if err != nil {
		fmt.Println(err)
		Body := err.Error() + "\n"
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", Body)
		return
	}

	Body := courseID + " deleted!\n"
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", Body)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := List()
	fmt.Fprintf(w, "%s", Body)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := fmt.Sprintf("Total entries: %d\n", len(data))
	fmt.Fprintf(w, "%s", Body)
}

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	// Split URL
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramStr)

	if len(paramStr) < 5 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not enough arguments: "+r.URL.Path)
		return
	}

	courseID := paramStr[2]
	courseName := paramStr[3]
	coursePreReq := paramStr[4]

	if !MatchID(courseID) {
		fmt.Println("Not a valid course id: ", courseID)
		return
	}

	temp := &MSDSCourse{CID: courseID, CNAME: courseName, CPREREQ: coursePreReq}
	err := Insert(temp)

	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		Body := "Failed to add record\n"
		fmt.Fprintf(w, "%s", Body)
	} else {
		log.Println("Serving:", r.URL.Path, "from", r.Host)
		Body := "New record added successfully\n"
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", Body)
	}

	log.Println("Serving:", r.URL.Path, "from", r.Host)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Get Search value from URL
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramStr)

	if len(paramStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not found: "+r.URL.Path)
		return
	}

	var Body string
	courseid := paramStr[0]
	t := Search(courseid)
	if t == nil {
		w.WriteHeader(http.StatusNotFound)
		Body = "Could not be found: " + courseid + "\n"
	} else {
		w.WriteHeader(http.StatusOK)
		Body = t.CID + " " + t.CNAME + " " + t.CPREREQ + "\n"
	}

	fmt.Println("Serving:", r.URL.Path, "from", r.Host)
	fmt.Fprintf(w, "%s", Body)
}

func ReadCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	// CSV file read all at once
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := MSDSCourse{
			CID:     line[0],
			CNAME:   line[1],
			CPREREQ: line[2],
		}
		// Storing to global variable
		data = append(data, temp)
	}

	return nil
}

func SaveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range data {
		temp := []string{row.CID, row.CNAME, row.CPREREQ}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func CreateIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		key := k.CID
		index[key] = i
	}
	return nil
}

// Initialized by the user – returns a pointer
// If it returns nil, there was an error
func InitS(N, S, T string) *MSDSCourse {
	// Both of them should have a value
	if N == "" || S == "" {
		return nil
	}
	return &MSDSCourse{CID: N, CNAME: S, CPREREQ: T}
}

func Insert(pS *MSDSCourse) error {
	// If it already exists, do not add it
	_, ok := index[(*pS).CID]
	if ok {
		return fmt.Errorf("%s already exists", pS.CID)
	}

	data = append(data, *pS)
	// Update the index
	_ = CreateIndex()

	err := SaveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	data = append(data[:i], data[i+1:]...)
	// Update the index - key does not exist any more
	delete(index, key)

	err := SaveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func Search(key string) *MSDSCourse {
	i, ok := index[key]
	if !ok {
		return nil
	}
	return &data[i]
}

func MatchID(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`[A-Za-z]+ [0-9]+-[A-Za-z]+`)
	return re.Match(t)
}

func List() string {
	var all string
	for _, k := range data {
		all = all + k.CID + " " + k.CNAME + " " + k.CPREREQ + "\n"
	}
	return all
}
