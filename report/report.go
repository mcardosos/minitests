package report

import (
	"fmt"

	"github.com/Azure/Tests/report"
)

func GetReport() string {
	output := ""
	var reportClient = report.NewWithBaseURI("http://localhost:3000", "jsklad")
	res, err := reportClient.GetReport()
	if err != nil {
		output += fmt.Sprintf("Error: %v\n", err)
		fmt.Println("Error:", err)
	}

	count := 0
	for key, val := range *res.Value {
		if *val <= 0 /*strings.Contains(strings.ToUpper(key), "UUID") */ {
			output += fmt.Sprintf("%v %v\n", key, *val)
			fmt.Println(key, *val)
			count++
		}
	}
	total := len(*res.Value)
	output += fmt.Sprintf("Total: %v\n", total)
	fmt.Println("Total: ", total)
	output += fmt.Sprintf("Left: %v\n", count)
	fmt.Println("Left:", count)

	return output
}
