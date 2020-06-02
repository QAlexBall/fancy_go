package main

import "fmt"

func main() {
	var contryCapitalMap map[string]string
	contryCapitalMap = make(map[string]string)

	contryCapitalMap["France"] = "Paris"

	for contry := range contryCapitalMap {
		fmt.Println(contry, "capital is", contryCapitalMap[contry])
	}

	capital, ok := contryCapitalMap["American"]
	if ok {
		fmt.Println("American capitial is", capital, "!")
	} else {
		fmt.Println("American capital is not exist!")
	}
}
