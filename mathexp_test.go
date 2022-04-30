package mathexp_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/shankeerthan-kasilingam/mathexp"
)

func runMathExpTest(t *testing.T, jsonPath string, test func(t *testing.T, mexp *mathexp.MathExp)) {
	if jsonPath != "" {
		jsonPath = "./test-data/electricity-calc-sample.json"
	}
	byteData, file := openFile(t, "./test-data/electricity-calc-sample.json")
	defer closeFile(t, file)
	mathExp, err := mathexp.New(byteData)
	if err != nil {
		t.Fatalf("Failed to parse the data %v", err)
	}

	test(t, mathExp)

}

func TestMathExpFromJson(t *testing.T) {

	runMathExpTest(t, "", func(t *testing.T, mexp *mathexp.MathExp) {
		if len(mexp.ExpWrapper.SubConditionGroups) != 3 {
			t.Errorf("SubConditionsGroups are not correct, got %d want %d", len(mexp.ExpWrapper.SubConditionGroups), 3)
		}
		if len(mexp.ExpWrapper.Expressions) != 0 {
			t.Errorf("Expressions are not empty, got %d want %d", len(mexp.ExpWrapper.Expressions), 0)
		}
		if strings.Compare(mexp.ExpWrapper.SubConditionGroups[1].Cond.Op, "and") != 0 {
			t.Errorf("Operation doesn't match got %s want %s", mexp.ExpWrapper.SubConditionGroups[1].Cond.Op, "and")
		}

	})
}

func Test


func openFile(t *testing.T, filePath string) ([]byte, *os.File) {
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file %s : %v", filePath, err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read file %s : %v", filePath, err)
	}
	return data, file
}

func closeFile(t *testing.T, file *os.File) {
	err := file.Close()
	if err != nil {
		t.Fatalf("Failed to close file %s : %v", file.Name(), err)
	}
}
