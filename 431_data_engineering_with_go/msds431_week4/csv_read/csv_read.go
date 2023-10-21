package csv_read

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

type Houses struct {
	VALUE    float64
	INCOME   float64
	AGE      float64
	ROOMS    float64
	BEDROOMS float64
	POP      float64
	HH       float64
}

func ReadData(filepath string) []Houses {
	//this function will read in the target data with a struct
	var TempStruct []Houses

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err) //replace with log
	}
	defer file.Close()
	MyReader := csv.NewReader(file)
	MyReader.FieldsPerRecord = -1

	row := 0
	var RecordSlice Houses
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
				switch i {
				case 0:
					floatParsed, err := strconv.ParseFloat(val, 64)
					if err != nil {
						floatParsed = 12345678.0 //defaulting to 12345678.0 for troubleshooting
					}
					RecordSlice.VALUE = floatParsed
				case 1:
					floatParsed, err := strconv.ParseFloat(val, 64)
					if err != nil {
						floatParsed = 12345678.0 //defaulting to 12345678.0 for troubleshooting
					}
					RecordSlice.INCOME = floatParsed
				case 2:
					floatParsed, err := strconv.ParseFloat(val, 64)
					if err != nil {
						floatParsed = 12345678.0 //defaulting to 12345678.0 for troubleshooting
					}
					RecordSlice.AGE = floatParsed
				case 3:
					floatParsed, err := strconv.ParseFloat(val, 64)
					if err != nil {
						floatParsed = 12345678.0 //defaulting to 12345678.0 for troubleshooting
					}
					RecordSlice.ROOMS = floatParsed
				case 4:
					floatParsed, err := strconv.ParseFloat(val, 64)
					if err != nil {
						floatParsed = 12345678.0 //defaulting to 12345678.0 for troubleshooting
					}
					RecordSlice.BEDROOMS = floatParsed
				case 5:
					floatParsed, err := strconv.ParseFloat(val, 64)
					if err != nil {
						floatParsed = 12345678.0 //defaulting to 12345678.0 for troubleshooting
					}
					RecordSlice.POP = floatParsed
				case 6:
					floatParsed, err := strconv.ParseFloat(val, 64)
					if err != nil {
						floatParsed = 12345678.0 //defaulting to 12345678.0 for troubleshooting
					}
					RecordSlice.HH = floatParsed
				}
			}

		}
		TempStruct = append(TempStruct, RecordSlice)
		row += 1
	}
	return TempStruct
}

func Sums(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}

	return sum
}

func Avg(data []float64) float64 {
	return Sums(data) / float64(len(data))
}

func StandardDeviation(data []float64) float64 {
	var mean, holder, sd float64
	mean = Avg(data)
	for i := 0; i < len(data); i++ {
		holder += math.Pow(data[i]-mean, 2)
	}
	sd = math.Sqrt(holder / float64(len(data)-1))
	return sd
}

func MinMax(data []float64) (float64, float64) {
	var min, max float64
	min = data[0]
	max = data[len(data)-1]
	return min, max
}

func Median(data []float64) float64 {
	length := len(data)
	var q2 float64
	if length%2 == 0 {
		midpointone := length/2 - 1 // since index = 0 for array
		midpointtwo := midpointone + 1
		q2 = (data[midpointone] + data[midpointtwo]) / 2
	} else {
		midpoint := ((length + 1) / 2)
		q2 = data[midpoint]
	}

	return q2
}

func Percentile(data []float64, percentile float64) float64 {
	length := len(data)

	holder := float64(length-1) * percentile
	fraction := math.Abs(holder - math.Floor(holder))
	ival := data[int(math.Floor(holder))]
	jval := data[int(math.Ceil(holder))]
	quartileval := ival + ((jval - ival) * fraction)
	//fmt.Println(fraction)
	//fmt.Println(ival)
	//fmt.Println(jval)
	//fmt.Println(quartileval)
	return quartileval

}
