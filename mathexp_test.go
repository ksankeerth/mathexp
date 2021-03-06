package mathexp

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func runMathExpTest(t *testing.T, jsonPath string, test func(t *testing.T, mexp *MathExp)) {

	if jsonPath == "" {
		jsonPath = "./test-data/electricity-calc-sample.json"
	}
	byteData, file := openFile(t, jsonPath)
	defer closeFile(t, file)
	mathExp, err := New(byteData)
	if err != nil {
		t.Fatalf("Failed to parse the data : %v", err)
	}

	test(t, mathExp)

}

func TestMathExpFromJson(t *testing.T) {

	runMathExpTest(t, "", func(t *testing.T, mexp *MathExp) {
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

func TestMathExpTransversal(t *testing.T) {
	runMathExpTest(t, "", func(t *testing.T, mexp *MathExp) {
		var count = 0
		mexp.ExpWrapper.traverse(new([]*VarSpec), func(cg *ConditionGroupSpec, vars []*VarSpec) bool {
			if len(vars) != 5 {
				t.Errorf("No of vars doesn't match got %d want %d", len(vars), 5)
			}
			count++
			return false
		})
		if count != 4 {
			t.Errorf("Traversal is not complete, stopped at %d but exptcted to stop at %d", count, 4)
		}
	})
}

func TestMathExpIsValid(t *testing.T) {
	runMathExpTest(t, "", func(t *testing.T, mexp *MathExp) {
		if valid, _ := mexp.ExpWrapper.isValid(); !valid {
			t.Errorf("Expected expression to be valid, but invalid")
		}
	})
}

func TestMathExpAllVars(t *testing.T) {
	runMathExpTest(t, "./test-data/electricity-calc-sample-vars-in-subcondg.json", func(t *testing.T, mexp *MathExp) {
		if len(allVars(mexp.ExpWrapper)) != 6 {
			t.Errorf("Expected allVars to return %d vars but returned %d", 6, len(allVars(mexp.ExpWrapper)))
		}
	})
}

func TestMathExpVerifyBeforEvaluate(t *testing.T) {
	runMathExpTest(t, "", func(t *testing.T, mexp *MathExp) {
		// Expect ok to be false
		ok := verifyBeforeEvalute(mexp.ExpWrapper.Vars, make(map[string]interface{}))
		if ok {
			t.Errorf("Expected false  but got true")
		}

		// Expect ok to be true
		args := make(map[string]interface{})
		args["elec_consumption"] = 100.0
		ok = verifyBeforeEvalute(mexp.ExpWrapper.Vars, args)

		if !ok {
			t.Errorf("Expected true  but got false")
		}

		args["test_additional_in"] = 2000
		ok = verifyBeforeEvalute(mexp.ExpWrapper.Vars, args)

		if ok {
			t.Errorf("Expected false  but got true")
		}

	})
}

func TestMathExpEvaluate(t *testing.T) {
	runMathExpTest(t, "", func(t *testing.T, mexp *MathExp) {
		args := make(map[string]interface{})
		args["elec_consumption"] = 99.0
		out, err := mexp.Evaluate(args)
		if err != nil {
			t.Errorf("Failed due to %v", err)
		}
		if len(out) != 1 {
			t.Errorf("Expected output size %d but got %d", 1, len(out))
		}
		if v, ok := out["elec_cost"]; !ok || v != 495.0 {
			t.Errorf("Expected %f but got %f", 495.0, v)
		}
		args["elec_consumption"] = 100.0
		out, _ = mexp.Evaluate(args)
		if v, ok := out["elec_cost"]; !ok || v != 700.0 {
			t.Errorf("Expected %f but got %f", 840.0, v)
		}
		args["elec_consumption"] = 1000.0
		out, _ = mexp.Evaluate(args)
		if v, ok := out["elec_cost"]; !ok || v != 10000.0 {
			t.Errorf("Expected %f but got %f", 10000.0, v)
		}
	})
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
