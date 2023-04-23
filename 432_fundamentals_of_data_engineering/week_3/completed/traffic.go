package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const TimeFormat = "01/02/2006 03:04:05 PM"

type TrafficCrashes struct {
	CRASH_RECORD_ID               string
	RD_NO                         string
	CRASH_DATE_EST_I              string
	CRASH_DATE                    time.Time
	POSTED_SPEED_LIMIT            int
	TRAFFIC_CONTROL_DEVICE        string
	DEVICE_CONDITION              string
	WEATHER_CONDITION             string
	LIGHTING_CONDITION            string
	FIRST_CRASH_TYPE              string
	TRAFFICWAY_TYPE               string
	LANE_CNT                      int
	ALIGNMENT                     string
	ROADWAY_SURFACE_COND          string
	ROAD_DEFECT                   string
	REPORT_TYPE                   string
	CRASH_TYPE                    string
	INTERSECTION_RELATED_I        string
	NOT_RIGHT_OF_WAY_I            string
	HIT_AND_RUN_I                 string
	DAMAGE                        string
	DATE_POLICE_NOTIFIED          time.Time
	PRIM_CONTRIBUTORY_CAUSE       string
	SEC_CONTRIBUTORY_CAUSE        string
	STREET_NO                     int
	STREET_DIRECTION              string
	STREET_NAME                   string
	BEAT_OF_OCCURRENCE            int
	PHOTOS_TAKEN_I                string
	STATEMENTS_TAKEN_I            string
	DOORING_I                     string
	WORK_ZONE_I                   string
	WORK_ZONE_TYPE                string
	WORKERS_PRESENT_I             string
	NUM_UNITS                     int
	MOST_SEVERE_INJURY            string
	INJURIES_TOTAL                int
	INJURIES_FATAL                int
	INJURIES_INCAPACITATING       int
	INJURIES_NON_INCAPACITATING   int
	INJURIES_REPORTED_NOT_EVIDENT int
	INJURIES_NO_INDICATION        int
	INJURIES_UNKNOWN              int
	CRASH_HOUR                    int
	CRASH_DAY_OF_WEEK             int
	CRASH_MONTH                   int
	LATITUDE                      float64
	LONGITUDE                     float64
	LOCATION                      string
	E_Day_of_Week                 string
	E_ZIP_CODE                    string
}

func ReadData(filepath string) []TrafficCrashes {
	//this function will read in the target data with a struct
	var TempStruct []TrafficCrashes

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err) //replace with log
	}
	defer file.Close()
	MyReader := csv.NewReader(file)
	MyReader.FieldsPerRecord = -1
	//loop over records and append
	//For missing string values use M
	//missing dates use 1901-01-01
	//missing floats use 12.34567890
	//missing ints use -6
	row := 0
	var RecordSlice TrafficCrashes
	for {
		record, err := MyReader.Read()
		if err == io.EOF {
			break
		}

		if row == 0 {
			row += 1
			continue
		} else {

			for i, val := range record {
				ValType := fmt.Sprintf("%T", val)

				switch i {
				case 0:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.CRASH_RECORD_ID = val
				case 1:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.RD_NO = val
				case 2:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.CRASH_DATE_EST_I = val
				case 3:
					timeParsed, err := time.Parse(TimeFormat, val)
					if err != nil {
						fmt.Println(err)
						timeParsed = time.Date(1901, time.January, 01, 0, 0, 0, 0, time.UTC)
					}
					RecordSlice.CRASH_DATE = timeParsed
				case 4:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {

					}
					RecordSlice.POSTED_SPEED_LIMIT = int(intParsed)
				case 5:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.TRAFFIC_CONTROL_DEVICE = val
				case 6:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.DEVICE_CONDITION = val
				case 7:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.WEATHER_CONDITION = val
				case 8:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.LIGHTING_CONDITION = val
				case 9:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.FIRST_CRASH_TYPE = val
				case 10:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.TRAFFICWAY_TYPE = val
				case 11:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.LANE_CNT = int(intParsed)
				case 12:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.ALIGNMENT = val
				case 13:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.ROADWAY_SURFACE_COND = val
				case 14:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.ROAD_DEFECT = val
				case 15:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.REPORT_TYPE = val
				case 16:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.CRASH_TYPE = val
				case 17:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.INTERSECTION_RELATED_I = val
				case 18:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.NOT_RIGHT_OF_WAY_I = val
				case 19:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.HIT_AND_RUN_I = val
				case 20:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.DAMAGE = val
				case 21:
					timeParsed, err := time.Parse(TimeFormat, val)
					if err != nil {
						timeParsed = time.Date(1901, time.January, 01, 0, 0, 0, 0, time.UTC)
					}
					RecordSlice.DATE_POLICE_NOTIFIED = timeParsed
				case 22:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.PRIM_CONTRIBUTORY_CAUSE = val
				case 23:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.SEC_CONTRIBUTORY_CAUSE = val
				case 24:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.STREET_NO = int(intParsed)
				case 25:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.STREET_DIRECTION = val
				case 26:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.STREET_NAME = val
				case 27:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.BEAT_OF_OCCURRENCE = int(intParsed)
				case 28:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.PHOTOS_TAKEN_I = val
				case 29:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.STATEMENTS_TAKEN_I = val
				case 30:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.DOORING_I = val
				case 31:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.WORK_ZONE_I = val
				case 32:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.WORK_ZONE_TYPE = val
				case 33:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.WORKERS_PRESENT_I = val
				case 34:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.NUM_UNITS = int(intParsed)
				case 35:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.MOST_SEVERE_INJURY = val
				case 36:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.INJURIES_TOTAL = int(intParsed)
				case 37:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.INJURIES_FATAL = int(intParsed)
				case 38:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.INJURIES_INCAPACITATING = int(intParsed)
				case 39:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.INJURIES_NON_INCAPACITATING = int(intParsed)
				case 40:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.INJURIES_REPORTED_NOT_EVIDENT = int(intParsed)
				case 41:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.INJURIES_NO_INDICATION = int(intParsed)
				case 42:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.INJURIES_UNKNOWN = int(intParsed)
				case 43:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.CRASH_HOUR = int(intParsed)
				case 44:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.CRASH_DAY_OF_WEEK = int(intParsed)
				case 45:
					intParsed, err := strconv.ParseInt(val, 0, 0)
					if err != nil {
						intParsed = -6
					}
					RecordSlice.CRASH_MONTH = int(intParsed)
				case 46:
					floatParsed, err := strconv.ParseFloat(val, 64)
					if err != nil {
						floatParsed = 41.881832 //defaulting to the center of chicago according to https://www.latlong.net/place/chicago-il-usa-1855.html#:~:text=Latitude%20and%20longitude%20coordinates%20are%3A%2041.881832%2C%20-87.623177.%20Famous,most%20multi-cultural%20and%20multi-national%20communities%20in%20the%20country.
					}
					RecordSlice.LATITUDE = floatParsed
				case 47:
					floatParsed, err := strconv.ParseFloat(val, 64)
					if err != nil {
						floatParsed = -87.623177 //defaulting to the center of chicago according to https://www.latlong.net/place/chicago-il-usa-1855.html#:~:text=Latitude%20and%20longitude%20coordinates%20are%3A%2041.881832%2C%20-87.623177.%20Famous,most%20multi-cultural%20and%20multi-national%20communities%20in%20the%20country.
					}
					RecordSlice.LONGITUDE = floatParsed
				case 48:
					if ValType != "string" {
						val = "M"
					}
					RecordSlice.LOCATION = val

				}
			}

			//add weekday for the current record
			RecordSlice.E_Day_of_Week = RecordSlice.CRASH_DATE.Weekday().String()

		}

		TempStruct = append(TempStruct, RecordSlice)
		row += 1

	}
	return TempStruct
}

func contains(somemap map[string]string, str string) bool {
	//fmt.Println(str)
	_, ok := somemap[str]
	// If the key exists
	if ok {
		//fmt.Println("pickle")
		return true
	} else {
		//fmt.Println("no pickle")
		return false
	}
}

func main() {
	mystruct_data := ReadData(`Traffic_Crashes_-_Crashes.csv`)
	conditions_slice := make([]string, 0)
	unique_map := make(map[string]string)
	//results := make(chan string)
	//xx := make([]string, 0)

	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	for i, row := range mystruct_data {
		if i != 0 {
			conditions_slice = append(conditions_slice, row.WEATHER_CONDITION)
			//unique_map[row.WEATHER_CONDITION] = "pickle"
		}
	}

	generator := func(done <-chan interface{}, someslice []string) <-chan string {
		strStream := make(chan string)
		go func() {
			defer close(strStream)
			for i := 0; i < len(someslice); i++ {
				//fmt.Println(&stdoutBuff, "Sending: %d", conditions_slice[i])
				select {
				case <-done:
					return
				case strStream <- someslice[i]:
				}
			}
		}()
		return strStream
	}

	checker := func(
		done <-chan interface{},
		strStream <-chan string,
	) <-chan string {
		unique_stream := make(chan string)
		go func() {
			defer close(unique_stream)
			//newmap := make(map[string]int)
			for i := range strStream {
				select {
				case <-done:
					return
				default:
					if contains(unique_map, i) != true {
						unique_map[i] = "pickle"
						//fmt.Println(i)
						unique_stream <- i
					} else {
						continue
					}
				}

			}
		}()
		return unique_stream
	}

	done := make(chan interface{})
	defer close(done)

	strStream := generator(done, conditions_slice)
	pipeline := checker(done, strStream)
	for v := range pipeline {
		fmt.Println(v)
	}

}
