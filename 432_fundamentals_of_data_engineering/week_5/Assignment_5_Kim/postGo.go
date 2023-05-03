package main

import (
	"math/rand"

	"github.com/mactsouk/post05"
	"main.go/utils"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {
	post05.Hostname = "localhost"
	post05.Port = 5433
	post05.Username = "postgres"
	post05.Password = "root"
	post05.Database = "go"

	/*data, err := post05.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)
	random_username := getString(5)

	t := post05.Userdata{
		Username:    random_username,
		Name:        "Mihalis",
		Surname:     "Tsoukalos",
		Description: "This is me!"}

	id := post05.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	err = post05.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	// Trying to delete it again!
	err = post05.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	id = post05.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	t = post05.Userdata{
		Username:    random_username,
		Name:        "Mihalis",
		Surname:     "Tsoukalos",
		Description: "This might not be me!"}

	err = post05.UpdateUser(t)
	if err != nil {
		fmt.Println(err)
	}*/
	test := utils.MSDSCourse{CID: "MSDS401", CNAME: "Applied Statistics with R", CPREREQ: "none"}
	test1 := utils.MSDSCourse{CID: "MSDS402", CNAME: "Data Science and Research Practice", CPREREQ: "none"}
	test2 := utils.MSDSCourse{CID: "MSDS410", CNAME: "Supervised Learning Methods", CPREREQ: "MSDS400,MSDS401"}
	test3 := utils.MSDSCourse{CID: "MSDS411", CNAME: "Unsupervised Learning Methods", CPREREQ: "MSDS400,MSDS401"}
	test4 := utils.MSDSCourse{CID: "MSDS432", CNAME: "Foundations of Data Engineering", CPREREQ: "MSDS400,MSDS420"}
	utils.AddCourse(test)
	utils.AddCourse(test1)
	utils.AddCourse(test2)
	utils.AddCourse(test3)
	utils.AddCourse(test4)
}
