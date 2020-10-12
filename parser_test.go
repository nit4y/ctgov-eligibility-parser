package ctgov

import (
	"strings"
	"testing"
)

func TestItemExists(t *testing.T) {
	tests := []struct {
		slice  []interface{}
		item   interface{}
		result bool
	}{
		{[]interface{}{"This", "is", "a", "test"}, "is", true},
		{[]interface{}{"This", "is", "a", "test"}, "tes", false},
		{[]interface{}{1, 2, 3, 4, 5}, 1, true},
		{[]interface{}{1, 2, 3, 4, 5}, 6, false},
	}
	for _, test := range tests {
		ret := itemExists(test.slice, test.item)
		if itemExists(test.slice, test.item) != test.result {
			t.Errorf("Bad output, got: %t, should be: %t.", ret, test.result)
		}
	}

}

func TestParse(t *testing.T) {

	tests := []struct {
		input      string
		testOutput string
	}{
		{
			`          1. Prior Treatment

               -  Must have received 3-6 cycles of cisplatin-based chemotherapy as part of
                  first-line (initial) chemotherapy.

                    -  Prior POMBACE, CBOP-BEP, or GAMEC are allowed.

               -  No more than one prior line of chemotherapy for GCT (other than the 1 cycle of
                  salvage chemotherapy as defined in the protocol)

                    -  Definition of one line of chemotherapy: One line of therapy can in some

               -  No prior treatment with high-dose chemotherapy (defined as treatment utilizing
                  stem cell rescue)

               -  Must have adequate recovery from prior surgery (eg, healed scar, resumption of
                  diet)`, "<ol><li>Prior Treatment</li><ul><li>Must have received 3-6 cycles of cisplatin-based chemotherapy as part of first-line (initial) chemotherapy.</li><ul><li>Prior POMBACE, CBOP-BEP, or GAMEC are allowed.</li></ul><li>No more than one prior line of chemotherapy for GCT (other than the 1 cycle of salvage chemotherapy as defined in the protocol)</li><ul><li>Definition of one line of chemotherapy: One line of therapy can in some</li></ul><li>No prior treatment with high-dose chemotherapy (defined as treatment utilizing stem cell rescue)</li><li>Must have adequate recovery from prior surgery (eg, healed scar, resumption of diet)</ul></ol>"},
		{ // mixed tabs and space intented
			`
        Inclusion Criteria:

          -  Agreement to avoid pregnancy or fathering children.

        Exclusion Criteria:

          -  Any prior allogeneic stem cell transplantation or a candidate for such
             transplantation.

		  -  Diagnosis of chronic liver disease.
`, "<p>Inclusion Criteria:</p><ul><li>Agreement to avoid pregnancy or fathering children.</li></ul><p>Exclusion Criteria:</p><ul><li>Any prior allogeneic stem cell transplantation or a candidate for such transplantation.</li><li>Diagnosis of chronic liver disease.</ul>",
		},
	}
	for _, test := range tests {
		p := NewParser()
		st := strings.NewReader(test.input)
		ret := p.Parse(st)
		if string(ret) != test.testOutput {
			t.Errorf("Parsing output is wrong. Please check output with a browser.")
		}
	}
}
