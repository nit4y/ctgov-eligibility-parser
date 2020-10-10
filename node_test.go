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
		line      []byte
		lastNode  *node
		testType  lineType
		testStart int
	}{
		{[]byte("  3. Prior Treatment"), newNode(0, emptyLine, unk), numberLine, 5},
		{[]byte("       - Prior Treatment"), newNode(0, emptyLine, unk), dashLine, 9},
		{[]byte("  3. Prior Treatment"), newNode(0, unkLine, unk), textLine, 5},
		{[]byte("       - Prior Treatment"), newNode(0, unkLine, unk), textLine, 9},
		{[]byte("Prior Treatment"), newNode(0, emptyLine, unk), textLine, 0},
		{[]byte("2 Prior Treatments"), newNode(0, emptyLine, unk), textLine, 2},
		{[]byte("-Prior Treatment"), newNode(0, emptyLine, unk), textLine, 1},
	}

	for _, test := range tests {
		lType, tStart := calcNodeProps(test.line, test.lastNode)
		if tStart != test.testStart {
			t.Errorf("Bad calculation, got: %d, should be: %d.", tStart, test.testStart)
		}

		if lType != test.testType {
			t.Errorf("Bad calculation, got: %d, should be: %d.", lType, test.testType)
		}
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
