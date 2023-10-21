package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"week4/csv_read"
)

type ColumnHolder struct {
	DATA  []float64
	COUNT float64
	MEAN  float64
	STD   float64
	MIN   float64
	Q1    float64
	Q2    float64
	Q3    float64
	MAX   float64
}

func main() {
	f, err := os.Create("housesOutputGo.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for i := 1; i <= 100; i++ {
		data := csv_read.ReadData("housesInput.csv")
		var ValueHolder ColumnHolder
		var IncomeHolder ColumnHolder
		var AgeHolder ColumnHolder
		var RoomsHolder ColumnHolder
		var BedRoomsHolder ColumnHolder
		var PopHolder ColumnHolder
		var HhHolder ColumnHolder

		for _, row := range data {
			ValueHolder.DATA = append(ValueHolder.DATA, row.VALUE)
			IncomeHolder.DATA = append(IncomeHolder.DATA, row.INCOME)
			AgeHolder.DATA = append(AgeHolder.DATA, row.AGE)
			RoomsHolder.DATA = append(RoomsHolder.DATA, row.ROOMS)
			BedRoomsHolder.DATA = append(BedRoomsHolder.DATA, row.BEDROOMS)
			PopHolder.DATA = append(PopHolder.DATA, row.POP)
			HhHolder.DATA = append(HhHolder.DATA, row.HH)

		}

		bigHolder := []*ColumnHolder{
			&ValueHolder,
			&IncomeHolder,
			&AgeHolder,
			&RoomsHolder,
			&BedRoomsHolder,
			&PopHolder,
			&HhHolder,
		}

		for i := range bigHolder {
			sort.Float64s(bigHolder[i].DATA)
			bigHolder[i].COUNT = float64(len(bigHolder[i].DATA))
			bigHolder[i].MEAN = csv_read.Avg(bigHolder[i].DATA)
			bigHolder[i].STD = csv_read.StandardDeviation(bigHolder[i].DATA)
			bigHolder[i].MIN, bigHolder[i].MAX = csv_read.MinMax(bigHolder[i].DATA)
			bigHolder[i].Q1 = csv_read.Percentile(bigHolder[i].DATA, .25)
			bigHolder[i].Q2 = csv_read.Median(bigHolder[i].DATA)
			bigHolder[i].Q3 = csv_read.Percentile(bigHolder[i].DATA, .75)
		}

		fmt.Fprintln(f, "\tvalue \t\t income \t age \t\t rooms \t\t bedrooms \t\t pop \t\t hh")
		fmt.Fprintf(f, "count\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\n",
			ValueHolder.COUNT,
			IncomeHolder.COUNT,
			AgeHolder.COUNT,
			RoomsHolder.COUNT,
			BedRoomsHolder.COUNT,
			PopHolder.COUNT,
			HhHolder.COUNT)

		fmt.Fprintf(f, "mean\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\n",
			ValueHolder.MEAN,
			IncomeHolder.MEAN,
			AgeHolder.MEAN,
			RoomsHolder.MEAN,
			BedRoomsHolder.MEAN,
			PopHolder.MEAN,
			HhHolder.MEAN)

		fmt.Fprintf(f, "std\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\n",
			ValueHolder.STD,
			IncomeHolder.STD,
			AgeHolder.STD,
			RoomsHolder.STD,
			BedRoomsHolder.STD,
			PopHolder.STD,
			HhHolder.STD)

		fmt.Fprintf(f, "min\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\n",
			ValueHolder.MIN,
			IncomeHolder.MIN,
			AgeHolder.MIN,
			RoomsHolder.MIN,
			BedRoomsHolder.MIN,
			PopHolder.MIN,
			HhHolder.MIN)

		fmt.Fprintf(f, "25%%\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\n",
			ValueHolder.Q1,
			IncomeHolder.Q1,
			AgeHolder.Q1,
			RoomsHolder.Q1,
			BedRoomsHolder.Q1,
			PopHolder.Q1,
			HhHolder.Q1)

		fmt.Fprintf(f, "50%%\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\n",
			ValueHolder.Q2,
			IncomeHolder.Q2,
			AgeHolder.Q2,
			RoomsHolder.Q2,
			BedRoomsHolder.Q2,
			PopHolder.Q2,
			HhHolder.Q2)

		fmt.Fprintf(f, "75%%\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\n",
			ValueHolder.Q3,
			IncomeHolder.Q3,
			AgeHolder.Q3,
			RoomsHolder.Q3,
			BedRoomsHolder.Q3,
			PopHolder.Q3,
			HhHolder.Q3)

		fmt.Fprintf(f, "max\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\t %10.6f\n",
			ValueHolder.MAX,
			IncomeHolder.MAX,
			AgeHolder.MAX,
			RoomsHolder.MAX,
			BedRoomsHolder.MAX,
			PopHolder.MAX,
			HhHolder.MAX)

		fmt.Fprintln(f, "-------------------------------")

	}
}
