package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ksankeerth/mathexp"
)

func main() {
	jsonfile, err := os.Open("electricity-calc-sample.json")
	defer func() {
		err := jsonfile.Close()
		if err != nil {
			fmt.Errorf("Error while closing file %v", err)
		}
	}()
	if err != nil {
		fmt.Errorf("Unable to open json %v", err)
	}
	data, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		fmt.Errorf("Unable to read json %v", err)
	}
	mexp, err := mathexp.New(data)
	if err != nil {
		fmt.Errorf("Errors while parsing expressions %v", err)
	}
	args := make(map[string]interface{})
	args["elec_consumption"] = 100.0
	out, err := mexp.Evaluate(args)
	if err != nil {
		fmt.Errorf("Error while evaluating expressions")
	}
	bill, ok := out["elec_cost"]
	if ok {
		fmt.Printf("The cost of electricitiy is %f.\n", bill)
	}
}
