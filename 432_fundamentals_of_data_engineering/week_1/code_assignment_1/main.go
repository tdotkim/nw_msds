package main

import (
	"code_assignment_1/code_csv"
	"code_assignment_1/code_gota"
	"code_assignment_1/code_struct"
	"fmt"
	"strconv"
)

func main() {
	/*

	   Welcome to MSDS 432, DE with Go. This assignment introduces some basic concepts of using Go (Golang) for data work.
	   This assignment is designed to be easy to follow. Your primary goal will be to call functions and see Go in action
	   and for experimentation purposes you should change the years, months, and different passing parameters as per the data

	*/

	//added all years function each package
	fmt.Println("\n\n ***********  ASSIGNMENT 1 - STARTING *********** \n\n")

	All_CSV_EX()
	ALL_STRUCT_EX()
	ALL_GOTA_EX()

	fmt.Println("\n\n ***********  ASSIGNMENT 1 - ENDING *********** \n\n")

}

func All_CSV_EX() {
	//load up all years slice

	//this function will hold all of the calls to answers all questions using a csv reader
	//read the data
	//call the ReadData function from the code_csv package
	//Add function here

	mycsvdata := code_csv.ReadData(`DATA/Traffic_Crashes_-_Crashes.csv`) // after running with large data set change it to Traffic_Crashes_Mini_Dataset.csv
	// doing that will result in a parsing error, find the solution, fix the code, and explain why it was failing and how you fixed it.
	yearslice := code_csv.AllYears(mycsvdata)
	fmt.Println(yearslice)
	//1. Print a report showing the total crashes during in the year 2020
	// After running the code with year 2020, look at the data and find which years does it have and then change the code to test for few other years
	// call the CrashesInYear function from the code_csv package. Make sure to use the correct arguments
	// Add function here

	question_1 := code_csv.CrashesInYear(mycsvdata, "2020")

	fmt.Println("\n\n *********** CSV READ - QUESTION 1 - Total crashes during in the year 2020 *********** \n\n")

	fmt.Println(question_1)

	fmt.Println("\n\n***********  Question 1 - Total crashes by Year START ***********\n\n")
	for i, year := range yearslice {
		holder := code_csv.CrashesInYear(mycsvdata, strconv.Itoa(year))
		fmt.Println(i, year, holder, "\n ---------------------------------- \n")
	}
	fmt.Println("\n\n***********  Question 1 - Total crashes by Year END ***********\n\n")

	fmt.Println("\n\n *********** CSV READ - QUESTION 1 ENDS *********** \n\n")

	//2. Print a report to show the number of crashes for every Day of the Week for Year 2021
	//call the CrashesDOW function from the code_csv package. Make sure to use the correct arguments
	//same as in the above experiment, change it for other years, months, and days for experimentation with the code
	//Add function here

	question_2 := code_csv.CrashesDOW(mycsvdata, "m", "2021")

	fmt.Println("\n\n *********** CSV READ - QUESTION 2 - Number of crashes for every Day of the Week for Year 2021 *********** \n\n")

	fmt.Println(question_2)

	fmt.Println("\n\n***********  Question 2 - CrashesDOW START ***********\n\n")
	for i, year := range yearslice {
		holder := code_csv.CrashesDOW(mycsvdata, "m", strconv.Itoa(year))
		fmt.Println(i, year, holder, "\n ---------------------------------- \n")
	}

	for i, year := range yearslice {
		holder := code_csv.CrashesDOW(mycsvdata, "d", strconv.Itoa(year))
		fmt.Println(i, year, holder, "\n ---------------------------------- \n")
	}
	fmt.Println("\n\n***********  Question 2 - CrashesDOW END ***********\n\n")

	fmt.Println("\n\n *********** CSV READ -  QUESTION 2 ENDS *********** \n\n")

	//3. Print a report to show the total number of crashes reported grouped by PRIM_CONTRIBUTORY_CAUSE for the month of December in the year 2020
	//call the TotalCrashesGrouped function from the code_csv package. Make sure to use the correct arguments
	//same as in the above experiment, change it for other years, months, and days for experimentation with the code
	//Add function here

	question_3 := code_csv.TotalCrashesGrouped(mycsvdata, "PRIM_CONTRIBUTORY_CAUSE", "12", "2020")

	fmt.Println("\n\n *********** CSV READ -  QUESTION 3 - Total number of crashes reported grouped by PRIM_CONTRIBUTORY_CAUSE *********** \n\n")

	fmt.Println(question_3)

	question_31 := code_csv.TotalCrashesGrouped(mycsvdata, "PRIM_CONTRIBUTORY_CAUSE", "6", "2020")
	question_32 := code_csv.TotalCrashesGrouped(mycsvdata, "PRIM_CONTRIBUTORY_CAUSE", "8", "2019")
	question_33 := code_csv.TotalCrashesGrouped(mycsvdata, "PRIM_CONTRIBUTORY_CAUSE", "1", "2022")

	fmt.Println(question_31)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_32)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_33)

	fmt.Println("\n\n ***********  CSV READ - QUESTION 3 ENDS *********** \n\n")

	//8. Print a report to show the total number of HIT_AND_RUN_I grouped by ROADWAY_SURFACE_COND for the month of December in the year 2020
	//call the HitNRun function from the code_csv package. Make sure to use the correct arguments
	//same as in the above experiment, change it for other years, months, and days for experimentation with the code
	//Add function here

	question_4 := code_csv.HitNRun(mycsvdata, "HIT_AND_RUN_I", "Y", "ROADWAY_SURFACE_COND", "12", "2020")

	fmt.Println("\n\n ***********  CSV READ - QUESTION 4 - Total number of HIT_AND_RUN_I grouped by ROADWAY_SURFACE_COND *********** \n\n")

	fmt.Println(question_4)

	question_41 := code_csv.HitNRun(mycsvdata, "HIT_AND_RUN_I", "Y", "ROADWAY_SURFACE_COND", "12", "2021")
	question_42 := code_csv.HitNRun(mycsvdata, "HIT_AND_RUN_I", "Y", "ROAD_DEFECT", "6", "2020")
	question_43 := code_csv.HitNRun(mycsvdata, "HIT_AND_RUN_I", "Y", "SEC_CONTRIBUTORY_CAUSE", "7", "2019")

	fmt.Println(question_41)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_42)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_43)

	fmt.Println("\n\n ***********  CSV READ - QUESTION 4 ENDS *********** \n\n")
}

func ALL_STRUCT_EX() {
	//this function will hold all of the calls to answer all questions using structs
	//read the data
	//call the ReadData function from the code_struct package
	//Add function here

	//mystruct_data := code_struct.ReadData(`DATA/Traffic_Crashes_-_Crashes.csv`) // after running with large data set change it to Traffic_Crashes_Mini_Dataset.csv
	// doing that will result in a parsing error, find the solution, fix the code, and explain why it was failing and how you fixed it.
	mystruct_data := code_struct.ReadData(`DATA/Traffic_Crashes_Mini_Dataset.csv`)
	//5. Print a report showing the total crashes during in the year 2020
	//call the CrashesInYear function from the code_struct package. Make sure to use the correct arguments
	//same as in the above experiment, change it for other years, months, and days for experimentation with the code
	//Add function here

	question_1 := code_struct.CrashesInYear(mystruct_data, 2020)
	question_11 := code_struct.CrashesInYear(mystruct_data, 2019)
	question_12 := code_struct.CrashesInYear(mystruct_data, 2021)
	question_13 := code_struct.CrashesInYear(mystruct_data, 2022)

	fmt.Println("\n\n ***********  STRUCT READ - QUESTION 1 - Total crashes during in the year 2020 *********** \n\n")

	fmt.Println(question_1)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_11)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_12)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_13)

	fmt.Println("\n\n ***********  STRUCT READ - QUESTION 1 ENDS *********** \n\n")

	//6. Print a report to show the number of crashes for every Day of the Week for Year 2021
	//call the CrashesDOW function from the code_struct package. Make sure to use the correct arguments
	//same as in the above experiment, change it for other years, months, and days for experimentation with the code
	//Add function here

	question_2 := code_struct.CrashesDOW(mystruct_data, "m", "2021")
	question_21 := code_struct.CrashesDOW(mystruct_data, "d", "2020")
	question_22 := code_struct.CrashesDOW(mystruct_data, "m", "2022")

	fmt.Println("\n\n ***********  STRUCT READ - QUESTION 2 - Number of crashes for every Day of the Week for Year 2021 *********** \n\n")

	fmt.Println(question_2)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_21)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_22)

	fmt.Println("\n\n ***********  STRUCT READ - QUESTION 2 ENDS *********** \n\n")

	//3. Print a report to show the total number of crashes reported grouped by PRIM_CONTRIBUTORY_CAUSE for the month of December in the year 2020
	//call the TotalCrashesGrouped function from the code_struct package.. Make sure to use the correct arguments
	//same as in the above experiment, change it for other years, months, and days for experimentation with the code
	//Add function here

	question_3 := code_struct.TotalCrashesGrouped(mystruct_data, "12", "2020")
	question_31 := code_struct.TotalCrashesGrouped(mystruct_data, "6", "2021")
	question_32 := code_struct.TotalCrashesGrouped(mystruct_data, "3", "2020")

	fmt.Println("\n\n ***********  STRUCT READ - QUESTION 3 - Total number of crashes reported grouped by PRIM_CONTRIBUTORY_CAUSE *********** \n\n")

	fmt.Println(question_3)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_31)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_32)

	fmt.Println("\n\n ***********  STRUCT READ - QUESTION 3 ENDS *********** \n\n")

	//4. Print a report to show the total number of HIT_AND_RUN_I grouped by ROADWAY_SURFACE_COND for the month of December in the year 2020
	//call the HitNRun function from the code_struct package.. Make sure to use the correct arguments
	//same as in the above experiment, change it for other years, months, and days for experimentation with the code
	//Add function here

	question_4 := code_struct.HitNRun(mystruct_data, "Y", "12", "2020")
	question_41 := code_struct.HitNRun(mystruct_data, "Y", "11", "2020")
	question_42 := code_struct.HitNRun(mystruct_data, "Y", "10", "2020")

	fmt.Println("\n\n ***********  STRUCT READ - QUESTION 4 - Total number of HIT_AND_RUN_I grouped by ROADWAY_SURFACE_COND *********** \n\n")

	fmt.Println(question_4)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_41)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_42)

	fmt.Println("\n\n ***********  STRUCT READ - QUESTION 4 ENDS *********** \n\n")

}

func ALL_GOTA_EX() {
	//this function will hold all of the calls to answer all questions using dataframes from gota
	//read the data
	//call the ReadData function from the code_gota package
	//Add function here
	//mydataframe := code_gota.ReadData(`DATA/Traffic_Crashes_-_Crashes.csv`) // after running with large data set change it to Traffic_Crashes_Mini_Dataset.csv
	// doing that will result in a parsing error, find the solution, fix the code, and explain why it was failing and how you fixed it.
	mydataframe := code_gota.ReadData(`DATA/Traffic_Crashes_Mini_Dataset.csv`)
	//1. Print a report showing the total crashes during in the year 2020
	//call the CrashesInYear function from the code_gota package. Make sure to use the correct arguments
	//same as in the above experiment, change it for other years, months, and days for experimentation with the code
	//Add function here

	question_1 := code_gota.CrashesInYear(mydataframe, 2020)
	question_11 := code_gota.CrashesInYear(mydataframe, 2021)
	question_12 := code_gota.CrashesInYear(mydataframe, 2019)

	fmt.Println("\n\n ***********  GOTA DATAFRAME - QUESTION 1 - Total crashes during in the year 2020 *********** \n\n")

	fmt.Println(question_1)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_11)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_12)

	fmt.Println("\n\n ***********  GOTA DATAFRAME - QUESTION 1 ENDS *********** \n\n")

	//2. Print a report to show the number of crashes for every Day of the Week for Year 2021
	//call the CrashesDOW function from the code_gota package. Make sure to use the correct arguments
	//same as in the above experiment, change it for other years, months, and days for experimentation with the code
	//Add function here

	question_2 := code_gota.CrashesDOW(mydataframe, "m", "2021")
	question_21 := code_gota.CrashesDOW(mydataframe, "d", "2022")
	question_22 := code_gota.CrashesDOW(mydataframe, "d", "2020")

	fmt.Println("\n\n ***********  GOTA DATAFRAME - QUESTION 2 - Number of crashes for every Day of the Week for Year 2021 ***********")

	fmt.Println(question_2)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_21)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_22)

	fmt.Println("\n\n ***********  GOTA DATAFRAME - QUESTION 2 ENDS *********** \n\n")

	//3. Print a report to show the total number of crashes reported grouped by PRIM_CONTRIBUTORY_CAUSE for the month of December in the year 2020
	//call the TotalCrashesGrouped function from the code_gota package.. Make sure to use the correct arguments
	//same as in the above experiment, change it for other years, months, and days for experimentation with the code
	//Add function here

	question_3 := code_gota.TotalCrashesGrouped(mydataframe, "December", "2020")
	question_31 := code_gota.TotalCrashesGrouped(mydataframe, "June", "2020")
	question_32 := code_gota.TotalCrashesGrouped(mydataframe, "June", "2021")

	fmt.Println("\n\n ***********  GOTA DATAFRAME - QUESTION 3 - Total number of crashes reported grouped by PRIM_CONTRIBUTORY_CAUSE *********** \n\n")

	fmt.Print(question_3)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Print(question_31)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Print(question_32)

	fmt.Println("\n\n ***********  GOTA DATAFRAME - QUESTION 3 ENDS *********** \n\n")

	//4. Print a report to show the total number of HIT_AND_RUN_I grouped by ROADWAY_SURFACE_COND for the month of December in the year 2020
	//call the HitNRun function from the code_gota package.. Make sure to use the correct arguments
	//same as in the above experiment, change it for other years, months, and days for experimentation with the code
	//Add function here

	question_4 := code_gota.HitNRun(mydataframe, "December", "2020")
	question_41 := code_gota.HitNRun(mydataframe, "December", "2021")
	question_42 := code_gota.HitNRun(mydataframe, "December", "2019")

	fmt.Println("\n\n ***********  GOTA DATAFRAME - QUESTION 4 - Total number of HIT_AND_RUN_I grouped by ROADWAY_SURFACE_COND *********** \n\n")

	fmt.Println(question_4)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_41)
	fmt.Print("\n ---------------------------------- \n")
	fmt.Println(question_42)

	fmt.Println("\n\n ***********  GOTA DATAFRAME - QUESTION 4 ENDS *********** \n\n")
}
