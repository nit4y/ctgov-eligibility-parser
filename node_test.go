package ctgov

import (
	"testing"
)

func TestCalcLevel(t *testing.T) {
	tests := []struct {
		line      []byte
		testLevel int
	}{
		{[]byte("  3. Prior Treatment"), 2},
		{[]byte("       - Prior Treatment"), 7},
		{[]byte("Prior Treatment"), 0},
		{[]byte("2 Prior Treatments"), 0},
		{[]byte(" -Prior Treatment"), 1},
	}

	for _, test := range tests {
		if calcLevel(test.line) != test.testLevel {
			t.Errorf("Bad calculation, got: %d, should be: %d.", test.line, test.testLevel)
		}
	}

}

func TestCalcNodeProps(t *testing.T) {
	tests := []struct {
		line       []byte
		lastNode   *node
		testType   lineType
		testStart  int
		testNumber int
	}{
		{[]byte("  3. Prior Treatment"), newNode(0, emptyLine, unk), numberLine, 5, 3},
		{[]byte("       - Prior Treatment"), newNode(0, emptyLine, unk), dashLine, 9, 0},
		{[]byte("  3. Prior Treatment"), newNode(0, textLine, unk), textLine, 5, 0},
		{[]byte("       - Prior Treatment"), newNode(0, textLine, unk), textLine, 9, 0},
		{[]byte("Prior Treatment"), newNode(0, emptyLine, unk), textLine, 0, 0},
		{[]byte("2 Prior Treatments"), newNode(0, emptyLine, unk), textLine, 2, 0},
		{[]byte("-Prior Treatment"), newNode(0, emptyLine, unk), textLine, 1, 0},
	}
	counter := 0
	for _, test := range tests {
		lType, tStart, nValue := calcNodeProps(test.line, test.lastNode)
		if tStart != test.testStart {
			t.Errorf("Bad calculation, got: %d, should be: %d. Test: %d", tStart, test.testStart, counter)
		}

		if lType != test.testType {
			t.Errorf("Bad calculation, got: %d, should be: %d. Test: %d", lType, test.testType, counter)
		}

		if nValue != test.testNumber {
			t.Errorf("Bad calculation, got: %d, should be: %d. Test: %d", lType, test.testNumber, counter)
		}
		counter++
	}

}

func TestCalcHTMLType(t *testing.T) {
	tests := []struct {
		t            lineType
		testHTMLType htmlType
	}{
		{numberLine, ol},
		{dashLine, ul},
	}

	for _, test := range tests {
		if calcHTMLType(test.t) != test.testHTMLType {
			t.Errorf("Bad calculation, got: %d, should be: %d.", test.t, test.testHTMLType)
		}
	}
}
