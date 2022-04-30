package mathexp_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/shankeerthan-kasilingam/mathexp"
)

func TestMathExpFromJson(t *testing.T) {

	testExpElecCalJson, err := os.Open("./test-data/electricity-calc-sample.json")
	defer func() {
		if err := testExpElecCalJson.Close(); err != nil {
			t.Fatalf("Failed to close the file %v", err)
		}
	}()
	if err != nil {
		t.Fatalf("Failed to open test data %v", err)
	}

	byteData, err := ioutil.ReadAll(testExpElecCalJson)
	if err != nil {
		t.Fatalf("Failed to read test data %v", err)
	}
	mathExp, err := mathexp.New(byteData)
	if err != nil {
		t.Fatalf("Failed to parse the data %v", err)
	}
	fmt.Println(mathExp.ExpWrapper)

}
