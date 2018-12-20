package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	env := "zyx"

	switch env {
	case "uat":
		fmt.Println("Apply UAT environment config")
	case "prod":
		fmt.Println("Apply Production config")
	case "dev":
		fmt.Println("Apply Dev config")
	default:
		fmt.Println("Invalid Environment")
	}

	fmt.Print("When is Saturday? : ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}








switch env_ci := strings.ToLower(env); env_ci {
	case "uat":
		fmt.Println("Apply UAT environment config")
	case "prod":
		fmt.Println("Apply Production config")
	case "dev":
		fmt.Println("Apply Dev config")
	default:
		fmt.Println("Invalid Environment")
	}




}


func getGrade(marks int) string {

	var grade string
	switch {
	case marks >= 90:
		grade = "Passed : A+"
	case marks >= 75 && marks < 90:
		grade = "Passed : B+"
	case marks >= 60 && marks < 75:
		grade = "Passed : C+"
	case marks >= 40 && marks < 60:
		grade = "Passed : D+"
	default:
		grade = "Failed"
	}

	return grade;

}
