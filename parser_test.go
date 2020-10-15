package ctgov

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {

	tests := []struct {
		input      string
		testOutput string
	}{
		{ // numberline with nested dashline
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
		{ // mixed tabs and space indetation
			`
        Inclusion Criteria:

          -  Agreement to avoid pregnancy or fathering children.

        Exclusion Criteria:

          -  Any prior allogeneic stem cell transplantation or a candidate for such
             transplantation.

		  -  Diagnosis of chronic liver disease.
`, "<p>Inclusion Criteria:</p><ul><li>Agreement to avoid pregnancy or fathering children.</li></ul><p>Exclusion Criteria:</p><ul><li>Any prior allogeneic stem cell transplantation or a candidate for such transplantation.</li><li>Diagnosis of chronic liver disease.</ul>",
		},
		{ // numbering exceeds 10
			`        Exclusion Criteria:

          1. Central nervous system metastasis (including brain metastasis, meningeal metastasis,
             etc.);

          2. Other immunosuppressive drugs used within 14 days before before study drug
             administration, excluding nasal sprays and inhaled corticosteroids or physiological
             doses of systemic steroids (ie not more than 10 mg/day of prednisolone or Other
             corticosteroids of equivalent pharmacological physiological dose);

          3. Hypertension and unable to be controlled within normal level following treatment of
             anti-hypertension agents: systolic blood pressure ≥140 mmHg, diastolic blood pressure
             ≥ 90 mmHg;

          4. Clinically significant cardiovascular diseases：Myocardial ischemia or myocardial
             infarction above grade II, ventricular arrhythmia which poorly
             controlled，QTc&gt;450ms（male）/QTc&gt;470ms (female)；Congestive heart failure (New York heart
             association (NYHA) class is Ⅲ～Ⅳ)；or cardiac color Doppler ultrasound examination
             revealed that the left ventricular ejection fraction (LVEF) &lt;50%;

          5. Accompanied by uncontrolled pleural effusion, pericardial effusion, or ascites that
             requires repeated drainage;

          6. Patients with any active autoimmune disease or history of autoimmune disease,
             including but not limited to the following: hepatitis, pneumonitis, uveitis, colitis
             (inflammatory bowel disease), hypophysitis, vasculitis, nephritis, hyperthyroidism,
             and hypothyroidism;

          7. Asthma that requires intermittent use of bronchodilators or other medical intervention
             should be excluded（Asthma has been completely relieved in childhood, and those without
             any intervention after adulthood can be included);

          8. Coagulation abnormalities (INR&gt;1.5、PT&gt;ULN+4s、APTT &gt;1.5 ULN）, with bleeding tendency or
             are receiving thrombolytic or anticoagulant therapy;

          9. Proteinuria ≥ (++) and 24 hours total urine protein &gt; 1.0 g;

         10. Received major surgery or suffered severe traumatic injury, fracture or ulcer within 4
             weeks before enrollment;

         11. Severe infections (such as intravenous infusion of antibiotics, antifungal or
             antiviral drugs) within 4 weeks before the first administration, or unexplained fever&gt;
             38.5°C during the screening period/before the first administration;
`, "<p>Exclusion Criteria:</p><ol><li>Central nervous system metastasis (including brain metastasis, meningeal metastasis, etc.);</li><li>Other immunosuppressive drugs used within 14 days before before study drug administration, excluding nasal sprays and inhaled corticosteroids or physiological doses of systemic steroids (ie not more than 10 mg/day of prednisolone or Other corticosteroids of equivalent pharmacological physiological dose);</li><li>Hypertension and unable to be controlled within normal level following treatment of anti-hypertension agents: systolic blood pressure ≥140 mmHg, diastolic blood pressure ≥ 90 mmHg;</li><li>Clinically significant cardiovascular diseases：Myocardial ischemia or myocardial infarction above grade II, ventricular arrhythmia which poorly controlled，QTc&gt;450ms（male）/QTc&gt;470ms (female)；Congestive heart failure (New York heart association (NYHA) class is Ⅲ～Ⅳ)；or cardiac color Doppler ultrasound examination revealed that the left ventricular ejection fraction (LVEF) &lt;50%;</li><li>Accompanied by uncontrolled pleural effusion, pericardial effusion, or ascites that requires repeated drainage;</li><li>Patients with any active autoimmune disease or history of autoimmune disease, including but not limited to the following: hepatitis, pneumonitis, uveitis, colitis (inflammatory bowel disease), hypophysitis, vasculitis, nephritis, hyperthyroidism, and hypothyroidism;</li><li>Asthma that requires intermittent use of bronchodilators or other medical intervention should be excluded（Asthma has been completely relieved in childhood, and those without any intervention after adulthood can be included);</li><li>Coagulation abnormalities (INR&gt;1.5、PT&gt;ULN+4s、APTT &gt;1.5 ULN）, with bleeding tendency or are receiving thrombolytic or anticoagulant therapy;</li><li>Proteinuria ≥ (++) and 24 hours total urine protein &gt; 1.0 g;</li><li>Received major surgery or suffered severe traumatic injury, fracture or ulcer within 4 weeks before enrollment;</li><li>Severe infections (such as intravenous infusion of antibiotics, antifungal or antiviral drugs) within 4 weeks before the first administration, or unexplained fever&gt; 38.5°C during the screening period/before the first administration;</ol>",
		},
		{ // a, b, c ordered list
			`        Exclusion Criteria:

          1. De novo metastatic patients who needs immediate docetaxel therapy;

          2. Inadequate laboratory function:

               1. Absolute neutrophil count &lt;1.5 x 109 /L (1,500 per mm3),

               2. Platelet count &lt; 100 x 109 /L (100 000 per mm3),

               3. Hemoglobin ≤9.0 g/dL,

               4. Serum bilirubin &gt; ULN,

               5. AST or ALT

             i.&gt;2.5 x ULN in patient without liver metastases, ii.&gt;5x ULN in patients with liver
             metastases.

          3. Cardiological status
`, "<p>Exclusion Criteria:</p><ol><li>De novo metastatic patients who needs immediate docetaxel therapy;</li><li>Inadequate laboratory function:</li><ol type=\"a\"><li>Absolute neutrophil count &lt;1.5 x 109 /L (1,500 per mm3),</li><li>Platelet count &lt; 100 x 109 /L (100 000 per mm3),</li><li>Hemoglobin ≤9.0 g/dL,</li><li>Serum bilirubin &gt; ULN,</li><li>AST or ALT</li></ol> i.&gt;2.5 x ULN in patient without liver metastases, ii.&gt;5x ULN in patients with liver metastases.</li><li>Cardiological status</ol>",
		},
		{ // 4 levels nested, commentLine included, back to level 2
			`        Inclusion Criteria:

        Patients eligible for inclusion in this study have to meet all of the following criteria:

          1. Provide informed consent voluntarily.

          2. Male and female patients ≥ 18 years of age (or having reached the age of majority
             according to local laws and regulations, if the age is &gt; 18 years).

          3. Availability of tumor tissue sample (either fresh tumor biopsy or archival tumor
             tissue sample) or blood samples.

          4. Eastern Cooperative Oncology Group (ECOG) performance status ≤ 1.

          5. Patient must meet the following laboratory values:

               1. Serum total bilirubin ≤ 1.5 × ULN, (≤ 3.0 mg/dL for patients with Gilbert's
                  syndrome);

               2. Aspartate aminotransferase (AST) and alanine aminotransferase (ALT) ≤ 2.5 × ULN
                  (≤ 5.0 × ULN for patients with hepatic metastasis);

               3. Serum creatinine &lt; 1.5 x ULN or creatinine clearance (calculated* or measured
                  value)

                  ≥ 50 mL/min

                  *For calculated creatinine clearance (Clcr) value, the eligibility should be
                  determined using the Cockcroft-Gault formula:

                    -  Male Clcr (mL/mim) = body weight (kg) x (140-age) / [72 x creatinine
                       (mg/dL)]

                        -  Female Clcr (mL/min) = male Clcr x 0.85

               4. Calcium (corrected for serum albumin) and magnesium within normal limits or ≤
                  grade 1 according to NCI-CTCAE version 5.0 if judged clinically not significant
                  by the Investigator;`, "<p>Inclusion Criteria:</p><p>Patients eligible for inclusion in this study have to meet all of the following criteria:</p><ol><li>Provide informed consent voluntarily.</li><li>Male and female patients ≥ 18 years of age (or having reached the age of majority according to local laws and regulations, if the age is &gt; 18 years).</li><li>Availability of tumor tissue sample (either fresh tumor biopsy or archival tumor tissue sample) or blood samples.</li><li>Eastern Cooperative Oncology Group (ECOG) performance status ≤ 1.</li><li>Patient must meet the following laboratory values:</li><ol type=\"a\"><li>Serum total bilirubin ≤ 1.5 × ULN, (≤ 3.0 mg/dL for patients with Gilbert's syndrome);</li><li>Aspartate aminotransferase (AST) and alanine aminotransferase (ALT) ≤ 2.5 × ULN (≤ 5.0 × ULN for patients with hepatic metastasis);</li><li>Serum creatinine &lt; 1.5 x ULN or creatinine clearance (calculated* or measured value)</li> ≥ 50 mL/min</li><br>*For calculated creatinine clearance (Clcr) value, the eligibility should be determined using the Cockcroft-Gault formula:</li><ul><li>Male Clcr (mL/mim) = body weight (kg) x (140-age) / [72 x creatinine (mg/dL)]</li><ul><li>Female Clcr (mL/min) = male Clcr x 0.85</li></ul></ul><li>Calcium (corrected for serum albumin) and magnesium within normal limits or ≤ grade 1 according to NCI-CTCAE version 5.0 if judged clinically not significant by the Investigator;</ol></ol>",
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
