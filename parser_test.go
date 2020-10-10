package ctgov

import (
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
		{[]interface{}{1, 2, 3, 4, 5}, 6, true},
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
		input  io.Reader
		output []byte
	}{
		{
			`        1. Documentation of Disease

               -  Histologic Documentation: Confirmation of GCT histology (both seminoma and
                  nonseminoma) on pathologic review at the center of enrollment.

               -  Tumor may have originated in any primary site. NOTE: In rare circumstances,
                  patients will be allowed to enroll even if a pathologic diagnosis may not have
                  been established.

               -  This would require a clinical situation consistent with the diagnosis of GCT
                  (testicular, peritoneal, retroperitoneal or mediastinal mass, elevated tumor
                  marker levels {HCG ≥ 500; AFP ≥ 500} and typical pattern of metastases)

          2. Evidence of Disease

               -  Must have evidence of progressive or recurrent GCT (measurable or non-measurable)
                  following one line of cisplatin-based chemotherapy, defined as meeting at least
                  one of the following criteria:

                    -  Tumor biopsy of new or growing or unresectable lesions demonstrating viable
                       non-teratomatous GCT (enrollment on this study for adjuvant treatment after
                       macroscopically complete resection of viable GCT is not allowed). In the
                       event of an incomplete gross resection where viable GCT is found, patients
                       will be considered eligible for the study.

                    -  Consecutive elevated serum tumor markers (HCG or AFP) that are increasing.
                       Increase of an elevated LDH alone does not constitute progressive disease.

                    -  Development of new or enlarging lesions in the setting of persistently
                       elevated HCG or AFP, even if the HCG and AFP are not continuing to increase.

          3. Prior Treatment

               -  Must have received 3-6 cycles of cisplatin-based chemotherapy as part of
                  first-line (initial) chemotherapy.

                    -  Prior POMBACE, CBOP-BEP, or GAMEC are allowed.

                    -  Note: For patients requiring immediate treatment, 1 cycle of
                       conventional-dose salvage chemotherapy is allowed. Therefore, these patients
                       may have received 7 prior cycles of chemotherapy. 6 cycles as part of
                       first-line chemotherapy and 1 cycle of salvage conventional chemotherapy.

               -  No more than one prior line of chemotherapy for GCT (other than the 1 cycle of
                  salvage chemotherapy as defined in the protocol)

                    -  Definition of one line of chemotherapy: One line of therapy can in some
                       cases consist of 2 different cisplatin-based treatment combinations,
                       provided there is no disease progression between these two regimens.

                    -  Prior treatment with carboplatin as adjuvant therapy is allowed, provided
                       patients meet other eligibility criteria (e.g., the patient has also
                       received 3-4 cycles of cisplatin-based chemotherapy).

                    -  Prior treatment with 1-2 cycles of BEP or EP as adjuvant chemotherapy for
                       early stage GCT is allowed, provided the patient also received 3-4 cycles of
                       BEP or EP again at relapse. Patients treated with 3-4 cycles of VIP at
                       relapse following 1-2 cycles of BEP/EP are not eligible as this would be
                       considered more than 1 line of prior therapy.

               -  No prior treatment with high-dose chemotherapy (defined as treatment utilizing
                  stem cell rescue)

               -  No prior treatment with TIP with the exception when given as a bridge to
                  treatment on protocol for patients with rapidly progressive disease who cannot
                  wait to complete the eligibility screening process. Only one cycle is allowed.

               -  No concurrent treatment with other cytotoxic drugs or targeted therapies.

               -  No radiation therapy (other than to the brain) within 14 days of day 1 of
                  protocol chemotherapy except radiation to brain metastases, which must be
                  completed 7 days prior to start of chemotherapy.

               -  No previous chemotherapy within 17 days prior to enrollment. A minimum of three
                  weeks after the last day of the start of the previous chemotherapy regimen before
                  the first day of chemotherapy on study protocol.

               -  Must have adequate recovery from prior surgery (eg, healed scar, resumption of
                  diet)

          4. Age ≥ 14 years (≥ 18 years in Germany)

          5. ECOG Performance Status 0 to 2

          6. Male gender

          7. Required Initial Laboratory Values:

               -  Absolute Neutrophil Count (ANC) ≥ 1,500/mm^3

               -  Platelet Count ≥ 100,000/mm^3

               -  Calculated creatinine clearance ≥ 50 mL/min

               -  Bilirubin ≤ 2.0 x upper limits of normal (ULN)

               -  AST/ALT ≤ 2.5 x upper limits of normal (ULN)

          8. No concurrent malignancy other than non-melanoma skin cancer, superficial noninvasive
             (pTa or pTis) TCC of the bladder, contralateral GCT, or intratubular germ cell
             neoplasia. Patients with a prior malignancy, but at least 2 years since any evidence
             of disease are allowed.

          9. Negative Serology (antibody test) for the following infectious diseases:

               -  Human Immunodeficiency Virus (HIV) type 1 and 2

               -  Human T-cell Leukemia Virus (HTLV) type 1 and 2 (mandatory in US but optional in
                  Canada and Europe)

               -  Hepatitis B surface antigen

               -  Hepatitis C antibody

         10. No late relapse with completely surgically resectable disease. Patients with late
             relapses (defined as relapse ≥ 2 years from the date of completion of the last
             chemotherapy regimen) whose disease is completely surgically resectable are not
             eligible. Patients with late relapses who have unresectable disease are eligible.

         11. No large (≥ 2 cm) hemorrhagic or symptomatic brain metastases until local treatment
             has been administered (radiation therapy or surgery). Treatment may begin ≥ 7 days
             after completion of local treatment. Patients with small (&lt; 2 cm) and asymptomatic
             brain metastases are allowed and may be treated with radiation therapy and/or surgery
             concurrently with Arm A or cycles 1 and 2 of Arm B if deemed medically indicated.

             Radiation therapy should not be given concurrently with high-dose carboplatin or
             etoposide.

         12. No secondary somatic malignancy arising from teratoma (e.g., teratoma with malignant
             transformation) when it is actively part of the disease recurrence or progression.
`
		}
	}
	for _, test := range tests {
		ret := itemExists(test.slice, test.item)
		if itemExists(test.slice, test.item) != test.result {
			t.Errorf("Bad output, got: %t, should be: %t.", ret, test.result)
		}
	}

}
