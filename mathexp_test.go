package mathexp_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/shankeerthan-kasilingam/mathexp"
)

func TestMathExpFromJson(t *testing.T) {

	byteData, file := openFile(t, "./test-data/electricity-calc-sample.json")
	defer closeFile(t, file)

	mathExp, err := mathexp.New(byteData)
	if err != nil {
		t.Fatalf("Failed to parse the data %v", err)
	}
	if len(mathExp.ExpWrapper.SubConditionGroups) != 3 {
		t.Errorf("SubConditionsGroups are not correct, got %d want %d", len(mathExp.ExpWrapper.SubConditionGroups), 3)
	}
	if len(mathExp.ExpWrapper.Expressions) != 0 {
		t.Errorf("Expressions are not empty, got %d want %d", len(mathExp.ExpWrapper.Expressions), 0)
	}
	if strings.Compare(mathExp.ExpWrapper.SubConditionGroups[1].Cond.Op, "and") != 0 {
		t.Errorf("Operation doesn't match got %s want %s", mathExp.ExpWrapper.SubConditionGroups[1].Cond.Op, "and")
	}

}

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
