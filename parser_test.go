package ctgov

import (
	"strings"
	"testing"
)

type sample struct {
	input  string
	output string
}

func TestParse(t *testing.T) {
	tests := getParseTests()
	counter := 1
	for _, test := range tests {
		p := NewParser()
		st := strings.NewReader(test.input)
		ret := p.Parse(st)
		if string(ret) != test.output {
			t.Errorf("Parsing output is wrong in test %d. Please test output with browser. Test number:", counter)
		}
		counter = counter + 1
	}
}

func getParseTests() []sample {
	return []sample{
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
                  diet)`, "<ol><li>Prior Treatment</li><ul><li>Must have received 3-6 cycles of cisplatin-based chemotherapy as part of first-line (initial) chemotherapy.</li><ul><li>Prior POMBACE, CBOP-BEP, or GAMEC are allowed.</li></ul><li>No more than one prior line of chemotherapy for GCT (other than the 1 cycle of salvage chemotherapy as defined in the protocol)</li><ul><li>Definition of one line of chemotherapy: One line of therapy can in some</li></ul><li>No prior treatment with high-dose chemotherapy (defined as treatment utilizing stem cell rescue)</li><li>Must have adequate recovery from prior surgery (eg, healed scar, resumption of diet)</li></ul></ol>"},
		{ // mixed tabs and space indetation
			`
        Inclusion Criteria:

          -  Agreement to avoid pregnancy or fathering children.

        Exclusion Criteria:

          -  Any prior allogeneic stem cell transplantation or a candidate for such
             transplantation.

		  -  Diagnosis of chronic liver disease.
`, "<p>Inclusion Criteria:</p><ul><li>Agreement to avoid pregnancy or fathering children.</li></ul><p>Exclusion Criteria:</p><ul><li>Any prior allogeneic stem cell transplantation or a candidate for such transplantation.</li><li>Diagnosis of chronic liver disease.</li></ul>",
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
`, "<p>Exclusion Criteria:</p><ol><li>Central nervous system metastasis (including brain metastasis, meningeal metastasis, etc.);</li><li>Other immunosuppressive drugs used within 14 days before before study drug administration, excluding nasal sprays and inhaled corticosteroids or physiological doses of systemic steroids (ie not more than 10 mg/day of prednisolone or Other corticosteroids of equivalent pharmacological physiological dose);</li><li>Hypertension and unable to be controlled within normal level following treatment of anti-hypertension agents: systolic blood pressure ≥140 mmHg, diastolic blood pressure ≥ 90 mmHg;</li><li>Clinically significant cardiovascular diseases：Myocardial ischemia or myocardial infarction above grade II, ventricular arrhythmia which poorly controlled，QTc&gt;450ms（male）/QTc&gt;470ms (female)；Congestive heart failure (New York heart association (NYHA) class is Ⅲ～Ⅳ)；or cardiac color Doppler ultrasound examination revealed that the left ventricular ejection fraction (LVEF) &lt;50%;</li><li>Accompanied by uncontrolled pleural effusion, pericardial effusion, or ascites that requires repeated drainage;</li><li>Patients with any active autoimmune disease or history of autoimmune disease, including but not limited to the following: hepatitis, pneumonitis, uveitis, colitis (inflammatory bowel disease), hypophysitis, vasculitis, nephritis, hyperthyroidism, and hypothyroidism;</li><li>Asthma that requires intermittent use of bronchodilators or other medical intervention should be excluded（Asthma has been completely relieved in childhood, and those without any intervention after adulthood can be included);</li><li>Coagulation abnormalities (INR&gt;1.5、PT&gt;ULN+4s、APTT &gt;1.5 ULN）, with bleeding tendency or are receiving thrombolytic or anticoagulant therapy;</li><li>Proteinuria ≥ (++) and 24 hours total urine protein &gt; 1.0 g;</li><li>Received major surgery or suffered severe traumatic injury, fracture or ulcer within 4 weeks before enrollment;</li><li>Severe infections (such as intravenous infusion of antibiotics, antifungal or antiviral drugs) within 4 weeks before the first administration, or unexplained fever&gt; 38.5°C during the screening period/before the first administration;</li></ol>",
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
`, "<p>Exclusion Criteria:</p><ol><li>De novo metastatic patients who needs immediate docetaxel therapy;</li><li>Inadequate laboratory function:</li><ol type=\"a\"><li>Absolute neutrophil count &lt;1.5 x 109 /L (1,500 per mm3),</li><li>Platelet count &lt; 100 x 109 /L (100 000 per mm3),</li><li>Hemoglobin ≤9.0 g/dL,</li><li>Serum bilirubin &gt; ULN,</li><li>AST or ALT</li></ol><p>i.&gt;2.5 x ULN in patient without liver metastases, ii.&gt;5x ULN in patients with liver metastases.</p><li>Cardiological status</li></ol>",
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
                  by the Investigator;`, "<p>Inclusion Criteria:</p><p>Patients eligible for inclusion in this study have to meet all of the following criteria:</p><ol><li>Provide informed consent voluntarily.</li><li>Male and female patients ≥ 18 years of age (or having reached the age of majority according to local laws and regulations, if the age is &gt; 18 years).</li><li>Availability of tumor tissue sample (either fresh tumor biopsy or archival tumor tissue sample) or blood samples.</li><li>Eastern Cooperative Oncology Group (ECOG) performance status ≤ 1.</li><li>Patient must meet the following laboratory values:</li><ol type=\"a\"><li>Serum total bilirubin ≤ 1.5 × ULN, (≤ 3.0 mg/dL for patients with Gilbert's syndrome);</li><li>Aspartate aminotransferase (AST) and alanine aminotransferase (ALT) ≤ 2.5 × ULN (≤ 5.0 × ULN for patients with hepatic metastasis);</li><li>Serum creatinine &lt; 1.5 x ULN or creatinine clearance (calculated* or measured value)</li><p>≥ 50 mL/min</p><p>*For calculated creatinine clearance (Clcr) value, the eligibility should be determined using the Cockcroft-Gault formula:<ul><li>Male Clcr (mL/mim) = body weight (kg) x (140-age) / [72 x creatinine (mg/dL)]</li><ul><li>Female Clcr (mL/min) = male Clcr x 0.85</li></ul></ul></p><li>Calcium (corrected for serum albumin) and magnesium within normal limits or ≤ grade 1 according to NCI-CTCAE version 5.0 if judged clinically not significant by the Investigator;</li></ol></ol>",
		},
		{ // heavy indented and robust
			`        Inclusion Criteria:

        -

        Patients eligible for inclusion in this study have to meet all of the following criteria:

          1. Provide informed consent voluntarily.

          2. Male and female patients ≥ 18 years of age (or having reached the age of majority
             according to local laws and regulations, if the age is &gt; 18 years).

          3. Patients with advanced solid tumor who have failed at least one line of prior systemic
             therapy or for whom standard therapy do not exist and meet the following eligibility
             for the corresponding part of the study:

               1. Patient must have a histologically or cytologically confirmed diagnosis of
                  advanced recurrent or metastatic solid tumor.

               2. At least one measurable lesion as per RECIST 1.1. (Ovarian cancer participants
                  must have measurable disease by RECIST 1.1 criteria or evaluable cancer via CA125
                  GCIG criteria; Prostate cancer participants must have measurable disease by
                  RECIST 1.1 criteria or evaluable cancer via PSA response).

               3. Population eligibility:

                  • Patients eligible for Part 1 dose escalation: Advanced solid tumors with any
                  DDR gene 1) or PIK3CA 2) mutation who have failed or cannot tolerate standard
                  treatment or currently have no standard treatment.

                  Note:

                    1. DDRa panel: BRCA1, BRCA2, ATM, CDK12, CHEK1, CHEK2, BARD1, BRIP1, FANCL,
                       PALB2, PPP2R2A, RAD51B, RAD51C, RAD51D, RAD54L.

                    2. PIK3CA hotspot mutation: E545X, H1047X, or E542K. For the gene mutation
                       testing result either tumor tissue samples or circulating free tumor DNA
                       (ctDNA) test can be accppted.

                       • Patients eligible for Part 2 dose expansion:

                       - Cohort 1: Advanced solid tumors with any selected DDR3) gene mutation

                       - Cohort 2: Advanced solid tumors with PIK3CA hotspot mutation (E545X,
                       H1047X, E542K).

                         -  Cohort 3: Advanced high grade serous ovarian, fallopian tube or primary
                            peritoneal cancer patients with acquired PARP inhibitor resistance4)

                         -  Cohort 4: Advanced solid tumors with any selected DDR3) gene mutation
                            with acquired PARP inhibitor resistance4) .

                         -  Cohort 5: Platinum resistant/refractory5) recurrent high grade serous
                            ovarian, fallopian tube, or primary peritoneal cancer.

                       Note:

                    3. DDRb panel: BRCA1, BRCA2, BARD1, BRIP1, FANCL, PALB2, PPP2R2A, RAD51B,
                       RAD51C, RAD51D, RAD54L. For the gene mutation testing result either tumor
                       tissue samples or circulating free tumor DNA (ctDNA) test can be accepted

                    4. Acquired resistance to prior PARP inhibitor (PARPi) is defined meeting the
                       following criteria: - Patients with advanced solid tumor, with progression
                       on PARPi therapy prior trial consent - Patients should have responded to
                       their prior PARPi therapy (radiology evaluation SD&gt;4months, or CR/PR). -
                       Patients must not have received another antitumor therapy between stopping
                       their previous PARPi therapy and initiating therapy on this trial - Patients
                       must be off prior PARPi for at least 3 weeks or 5 half-lives, whichever is
                       shorter.&quot; -

                    5. Platinum refractory or resistance is defined meeting the following criteria:

                         -  Platinum refractory is defined as either relapse less than 2 months
                            after the last platinum-based therapy or relapse during platinum
                            therapy.

                         -  Platinum-resistance is defined as relapse within 2 to 6 months after
                            last dose of platinum-based chemotherapy

          4. Availability of tumor tissue sample (either fresh tumor biopsy or archival tumor
             tissue sample) or blood samples.

          5. Eastern Cooperative Oncology Group (ECOG) performance status ≤ 1.

          6. Patient must meet the following laboratory values:

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
                  by the Investigator;

               5. Potassium within normal limits, or corrected with supplements;

               6. Platelets ≥ 100 x 109/L;

               7. Hemoglobin (Hgb) ≥ 10 g/dL;

               8. Absolute neutrophil count (ANC) ≥ 1.5 x 109/L;

               9. International normalized ratio (INR) &lt; 1.5 (or &lt; 3.0 if on anticoagulation);

              10. Fasting plasma glucose (FPG) ≤ 100 mg/dL (6.1 mmol/L) and Glycosylated Hemoglobin
                  (HbA1c) ≤ 5.7% (both criteria have to be met);

              11. Fasting serum amylase ≤ 2 × ULN;

              12. Fasting serum lipase ≤ ULN

        Exclusion Criteria:

          -  Patients eligible for this study must not meet any of the following criteria:

               1. Patient has received any anticancer therapy (including chemotherapy, targeted
                  therapy, hormonal therapy, biotherapy, immunotherapy, or other investigational
                  agents.) within 28 days or 5 times of half-lives (whichever is shorter) prior to
                  the first dose of the study treatment or who have not recovered from the side
                  effect of such therapy.

               2. Patients with contraindication to olaparib treatment or who did not tolerate
                  olaparib previously.

               3. Patients who had prior treatment with PARP inhibitor, PI3Kα inhibitor, AKT
                  inhibitor or mTOR inhibitor (Part 2 dose expansion cohort 1&amp; 2 only).

               4. Radical radiation therapy (including radiation therapy for over 25% bone marrow)
                  within 4 weeks prior to the first dose of the investigational product or received
                  local palliative radiation therapy for bone metastases within 2 weeks.

               5. Any toxicities from prior treatment that have not recovered to baseline or ≤
                  CTCAE Grade 1 before the start of study treatment, with exception of hair loss.

               6. Patients who have been treated with any hematopoietic colony-stimulating growth
                  factors (e.g., G-CSF, GM-CSF) ≤ 2 weeks prior to starting study drug.
                  Erythropoietin or darbepoetin therapy, if initiated at least 2 weeks prior to
                  enrollment, may be continued (Part 1 dose escalation only).

               7. Patients who have symptomatic CNS metastasis which is neurologically unstable or
                  those who have CNS disease requiring increase in the dose of steroid. (Note:
                  Patients withcontrolled CNS metastasis can participate in the trial. Before
                  entering the study, patientsshould have finished radiotherapy, or have received
                  operation for CNS tumor metastasis at least two weeks before. Patients'
                  neurological function must be in a stable state; no newneurological deficit is
                  found during clinical examination and no new problem is found during CNS imaging
                  examinations. If patients need to use steroids to treat CNS metastasis, the
                  therapeutic dose of steroid should be stable for ≥ 3 months at least two weeks
                  prior to entering the study with treatment dose no more than dexamethasone 4 mg
                  daily or anequivalent dose of steroids).

               8. Patients with an established diagnosis of diabetes mellitus including
                  steroid-induced diabetes mellitus.

               9. Major surgery or had significant traumatic injury within 28 days prior to the
                  first dose of the investigational product or has not recovered from major side
                  effects.

              10. Known HIV infection with a history of acquired immunodeficiency syndrome (AIDS)-
                  defining opportunity infection within the past 12 months; active hepatitis B and
                  hepatitis

                  C. Patients whose test results meet one of the following will not be enrolled:

                    -  For patients in China, confirmed HIV antibody positive. For patients in the
                       US, patients with a history of HIV but no history of AIDS or an
                       AIDS-defining opportunistic infection are allowed to be enrolled;

                    -  Serum HBsAg positive and HBV DNA&gt;200 IU/ml or 1000 copies/mL;

                    -  Serum HCV antibody and HCV RNA positive.

              11. Patient has any other concurrent disease which had potential risk of insulin
                  resistance (e.g., pancreatic disorders, acromegaly, Cushing's syndrome) or
                  current use of medication with potential risk of insulin resistance.

              12. Patient with pancreatic cancer.

              13. Patient is currently receiving or has received systemic corticosteroids ≤ 2 weeks
                  prior to starting study treatment, or who have not fully recovered from side
                  effects of such treatment.

                  Note: The following uses of corticosteroids are permitted: single doses, topical
                  applications (e.g., for rash), inhaled sprays (e.g., for obstructive airways
                  diseases), eye drops or local injections (e.g., intra-articular).

              14. Use of therapeutic doses of warfarin sodium (Coumadin®), or any other
                  coumarinderivative anticoagulants. The administration of low-molecular weight
                  heparin is allowed.

              15. History of acute pancreatitis within 1 year of screening or past medical history
                  of chronic pancreatitis.

              16. Gastrointestinal condition which could impair absorption of study medication
                  (e.g., ulcerative diseases, uncontrolled nausea, vomiting, diarrhea,
                  malabsorption syndrome, or small bowel resection).

              17. Patients with clinically significant cardiovascular disease, including:

                    -  NYHA Class III or higher congestive heart failure;

                    -  History or current evidence of serious uncontrolled ventricular arrhythmias
                       requiring drug therapy;

                    -  Acute myocardial infarction, severe or unstable angina pectoris, coronary
                       artery or peripheral artery bypass graft received within 6 months prior to
                       the first dose;

                    -  Left ventricular ejection fraction (LVEF) &lt; 50%;

                    -  Fridericia's corrected QT interval (QTcF) &gt; 460 ms on ECG conducted during
                       screening;

                    -  Congenital long QT syndrome, or any known history of torsade de pointes
                       (TdP), or family history of unexplained sudden death;

                    -  Clinically uncontrolled hypertension (after standard antihypertensive
                       treatment, systolic blood pressure ≥ 140 mmHg and/or diastolic blood
                       pressure ≥ 90 mmHg).

              18. Any diseases or medical conditions, at the Investigator's discretion, that may be
                  unstable or influence their safety or study compliance, including organ
                  transplantation, abuse of psychotropic medication, alcohol abuse or history of
                  drug abuse.

              19. Other serious illness or medical conditions at the Investigator's discretion,
                  that may influence study results, including but not limited to serious infection,
                  diabetes, cardiovascular and cerebrovascular diseases or lung disease.

              20. Participation in a prior investigational treatment within 28 days prior to the
                  start of study treatment or within 5 half-lives of the investigational product
                  (whichever is no longer than 28 days).

              21. Pregnant or breast-feeding patients. Pregnancy refers to the state of a woman
                  between fertilization and the end of pregnancy confirmed by positive laboratory
                  hCG test (&gt; 5 mIU/mL). Breast-feeding woman can become eligible for this study if
                  she stops breastfeeding, however, cannot restart the breast-feeding on/after the
                  completion of the study treatment.

              22. Male and female of childbearing potential not using effective contraception (e.g.

        intrauterine device (IUD), diaphragm with spermicide, cervical cap with spermicide, male
        condoms, female condoms with spermicide, oral hormonal contraceptive) during the trial and
        within 6 months after the end of treatment. Definition of child-bearing potential: a female
        that fulfills the one of the following criteria is considered to be without childbearing
        potential: spontaneous menopause for 12 consecutive months with appropriate clinical
        features (e.g. proper age, a history of vasomotor diseases, etc.), or a history of
        bilateral ovariectomy (with or without hysterectomy) or tubal ligation performed at least 6
        weeks. For patients with amenorrhea due to anti-tumor agents, even amenorrhea over 12
        months, a pregnancy test is necessary. If a female only received an ovariectomy, she will
        be considered as no childbearing potential only after confirmation by hormone levels.
`, "<p>Inclusion Criteria:</p><ul><li>        -</li><p>Patients eligible for inclusion in this study have to meet all of the following criteria:</p></li></ul><ol><li>Provide informed consent voluntarily.</li><li>Male and female patients ≥ 18 years of age (or having reached the age of majority according to local laws and regulations, if the age is &gt; 18 years).</li><li>Patients with advanced solid tumor who have failed at least one line of prior systemic therapy or for whom standard therapy do not exist and meet the following eligibility for the corresponding part of the study:</li><ol type=\"a\"><li>Patient must have a histologically or cytologically confirmed diagnosis of advanced recurrent or metastatic solid tumor.</li><li>At least one measurable lesion as per RECIST 1.1. (Ovarian cancer participants must have measurable disease by RECIST 1.1 criteria or evaluable cancer via CA125 GCIG criteria; Prostate cancer participants must have measurable disease by RECIST 1.1 criteria or evaluable cancer via PSA response).</li><li>Population eligibility:</li><p>• Patients eligible for Part 1 dose escalation: Advanced solid tumors with any DDR gene 1) or PIK3CA 2) mutation who have failed or cannot tolerate standard treatment or currently have no standard treatment.</p><p>Note:</p><ol type=\"a\"><li>DDRa panel: BRCA1, BRCA2, ATM, CDK12, CHEK1, CHEK2, BARD1, BRIP1, FANCL, PALB2, PPP2R2A, RAD51B, RAD51C, RAD51D, RAD54L.</li><li>PIK3CA hotspot mutation: E545X, H1047X, or E542K. For the gene mutation testing result either tumor tissue samples or circulating free tumor DNA (ctDNA) test can be accppted.</li><p>• Patients eligible for Part 2 dose expansion:</p><ul><li>Cohort 1: Advanced solid tumors with any selected DDR3) gene mutation</li><li>Cohort 2: Advanced solid tumors with PIK3CA hotspot mutation (E545X, H1047X, E542K).<ul><li>Cohort 3: Advanced high grade serous ovarian, fallopian tube or primary peritoneal cancer patients with acquired PARP inhibitor resistance4)</li><li>Cohort 4: Advanced solid tumors with any selected DDR3) gene mutation with acquired PARP inhibitor resistance4) .</li><li>Cohort 5: Platinum resistant/refractory5) recurrent high grade serous ovarian, fallopian tube, or primary peritoneal cancer.</li></ul><p>Note:</p></li></ul><li>DDRb panel: BRCA1, BRCA2, BARD1, BRIP1, FANCL, PALB2, PPP2R2A, RAD51B, RAD51C, RAD51D, RAD54L. For the gene mutation testing result either tumor tissue samples or circulating free tumor DNA (ctDNA) test can be accepted</li><li>Acquired resistance to prior PARP inhibitor (PARPi) is defined meeting the following criteria: - Patients with advanced solid tumor, with progression on PARPi therapy prior trial consent - Patients should have responded to their prior PARPi therapy (radiology evaluation SD&gt;4months, or CR/PR). - Patients must not have received another antitumor therapy between stopping their previous PARPi therapy and initiating therapy on this trial - Patients must be off prior PARPi for at least 3 weeks or 5 half-lives, whichever is shorter.&quot; -</li><li>Platinum refractory or resistance is defined meeting the following criteria:</li><ul><li>Platinum refractory is defined as either relapse less than 2 months after the last platinum-based therapy or relapse during platinum therapy.</li><li>Platinum-resistance is defined as relapse within 2 to 6 months after last dose of platinum-based chemotherapy</li></ul></ol></ol><li>Availability of tumor tissue sample (either fresh tumor biopsy or archival tumor tissue sample) or blood samples.</li><li>Eastern Cooperative Oncology Group (ECOG) performance status ≤ 1.</li><li>Patient must meet the following laboratory values:</li><ol type=\"a\"><li>Serum total bilirubin ≤ 1.5 × ULN, (≤ 3.0 mg/dL for patients with Gilbert's syndrome);</li><li>Aspartate aminotransferase (AST) and alanine aminotransferase (ALT) ≤ 2.5 × ULN (≤ 5.0 × ULN for patients with hepatic metastasis);</li><li>Serum creatinine &lt; 1.5 x ULN or creatinine clearance (calculated* or measured value)</li><p>≥ 50 mL/min</p><p>*For calculated creatinine clearance (Clcr) value, the eligibility should be determined using the Cockcroft-Gault formula:<ul><li>Male Clcr (mL/mim) = body weight (kg) x (140-age) / [72 x creatinine (mg/dL)]</li><li>Female Clcr (mL/min) = male Clcr x 0.85</li></ul></p><li>Calcium (corrected for serum albumin) and magnesium within normal limits or ≤ grade 1 according to NCI-CTCAE version 5.0 if judged clinically not significant by the Investigator;</li><li>Potassium within normal limits, or corrected with supplements;</li><li>Platelets ≥ 100 x 109/L;</li><li>Hemoglobin (Hgb) ≥ 10 g/dL;</li><li>Absolute neutrophil count (ANC) ≥ 1.5 x 109/L;</li><li>International normalized ratio (INR) &lt; 1.5 (or &lt; 3.0 if on anticoagulation);</li><li>Fasting plasma glucose (FPG) ≤ 100 mg/dL (6.1 mmol/L) and Glycosylated Hemoglobin (HbA1c) ≤ 5.7% (both criteria have to be met);</li><li>Fasting serum amylase ≤ 2 × ULN;</li><li>Fasting serum lipase ≤ ULN</li></ol></ol><p>Exclusion Criteria:</p><ul><li>Patients eligible for this study must not meet any of the following criteria:</li><ol><li>Patient has received any anticancer therapy (including chemotherapy, targeted therapy, hormonal therapy, biotherapy, immunotherapy, or other investigational agents.) within 28 days or 5 times of half-lives (whichever is shorter) prior to the first dose of the study treatment or who have not recovered from the side effect of such therapy.</li><li>Patients with contraindication to olaparib treatment or who did not tolerate olaparib previously.</li><li>Patients who had prior treatment with PARP inhibitor, PI3Kα inhibitor, AKT inhibitor or mTOR inhibitor (Part 2 dose expansion cohort 1&amp; 2 only).</li><li>Radical radiation therapy (including radiation therapy for over 25% bone marrow) within 4 weeks prior to the first dose of the investigational product or received local palliative radiation therapy for bone metastases within 2 weeks.</li><li>Any toxicities from prior treatment that have not recovered to baseline or ≤ CTCAE Grade 1 before the start of study treatment, with exception of hair loss.</li><li>Patients who have been treated with any hematopoietic colony-stimulating growth factors (e.g., G-CSF, GM-CSF) ≤ 2 weeks prior to starting study drug. Erythropoietin or darbepoetin therapy, if initiated at least 2 weeks prior to enrollment, may be continued (Part 1 dose escalation only).</li><li>Patients who have symptomatic CNS metastasis which is neurologically unstable or those who have CNS disease requiring increase in the dose of steroid. (Note: Patients withcontrolled CNS metastasis can participate in the trial. Before entering the study, patientsshould have finished radiotherapy, or have received operation for CNS tumor metastasis at least two weeks before. Patients' neurological function must be in a stable state; no newneurological deficit is found during clinical examination and no new problem is found during CNS imaging examinations. If patients need to use steroids to treat CNS metastasis, the therapeutic dose of steroid should be stable for ≥ 3 months at least two weeks prior to entering the study with treatment dose no more than dexamethasone 4 mg daily or anequivalent dose of steroids).</li><li>Patients with an established diagnosis of diabetes mellitus including steroid-induced diabetes mellitus.</li><li>Major surgery or had significant traumatic injury within 28 days prior to the first dose of the investigational product or has not recovered from major side effects.</li><li>Known HIV infection with a history of acquired immunodeficiency syndrome (AIDS)- defining opportunity infection within the past 12 months; active hepatitis B and hepatitis</li><p>C. Patients whose test results meet one of the following will not be enrolled:</p><ul><li>For patients in China, confirmed HIV antibody positive. For patients in the US, patients with a history of HIV but no history of AIDS or an AIDS-defining opportunistic infection are allowed to be enrolled;</li><li>Serum HBsAg positive and HBV DNA&gt;200 IU/ml or 1000 copies/mL;</li><li>Serum HCV antibody and HCV RNA positive.</li></ul><li>Patient has any other concurrent disease which had potential risk of insulin resistance (e.g., pancreatic disorders, acromegaly, Cushing's syndrome) or current use of medication with potential risk of insulin resistance.</li><li>Patient with pancreatic cancer.</li><li>Patient is currently receiving or has received systemic corticosteroids ≤ 2 weeks prior to starting study treatment, or who have not fully recovered from side effects of such treatment.</li><p>Note: The following uses of corticosteroids are permitted: single doses, topical applications (e.g., for rash), inhaled sprays (e.g., for obstructive airways diseases), eye drops or local injections (e.g., intra-articular).</p><li>Use of therapeutic doses of warfarin sodium (Coumadin®), or any other coumarinderivative anticoagulants. The administration of low-molecular weight heparin is allowed.</li><li>History of acute pancreatitis within 1 year of screening or past medical history of chronic pancreatitis.</li><li>Gastrointestinal condition which could impair absorption of study medication (e.g., ulcerative diseases, uncontrolled nausea, vomiting, diarrhea, malabsorption syndrome, or small bowel resection).</li><li>Patients with clinically significant cardiovascular disease, including:</li><ul><li>NYHA Class III or higher congestive heart failure;</li><li>History or current evidence of serious uncontrolled ventricular arrhythmias requiring drug therapy;</li><li>Acute myocardial infarction, severe or unstable angina pectoris, coronary artery or peripheral artery bypass graft received within 6 months prior to the first dose;</li><li>Left ventricular ejection fraction (LVEF) &lt; 50%;</li><li>Fridericia's corrected QT interval (QTcF) &gt; 460 ms on ECG conducted during screening;</li><li>Congenital long QT syndrome, or any known history of torsade de pointes (TdP), or family history of unexplained sudden death;</li><li>Clinically uncontrolled hypertension (after standard antihypertensive treatment, systolic blood pressure ≥ 140 mmHg and/or diastolic blood pressure ≥ 90 mmHg).</li></ul><li>Any diseases or medical conditions, at the Investigator's discretion, that may be unstable or influence their safety or study compliance, including organ transplantation, abuse of psychotropic medication, alcohol abuse or history of drug abuse.</li><li>Other serious illness or medical conditions at the Investigator's discretion, that may influence study results, including but not limited to serious infection, diabetes, cardiovascular and cerebrovascular diseases or lung disease.</li><li>Participation in a prior investigational treatment within 28 days prior to the start of study treatment or within 5 half-lives of the investigational product (whichever is no longer than 28 days).</li><li>Pregnant or breast-feeding patients. Pregnancy refers to the state of a woman between fertilization and the end of pregnancy confirmed by positive laboratory hCG test (&gt; 5 mIU/mL). Breast-feeding woman can become eligible for this study if she stops breastfeeding, however, cannot restart the breast-feeding on/after the completion of the study treatment.</li><li>Male and female of childbearing potential not using effective contraception (e.g.</li></ol></ul><p>intrauterine device (IUD), diaphragm with spermicide, cervical cap with spermicide, male condoms, female condoms with spermicide, oral hormonal contraceptive) during the trial and within 6 months after the end of treatment. Definition of child-bearing potential: a female that fulfills the one of the following criteria is considered to be without childbearing potential: spontaneous menopause for 12 consecutive months with appropriate clinical features (e.g. proper age, a history of vasomotor diseases, etc.), or a history of bilateral ovariectomy (with or without hysterectomy) or tubal ligation performed at least 6 weeks. For patients with amenorrhea due to anti-tumor agents, even amenorrhea over 12 months, a pregnancy test is necessary. If a female only received an ovariectomy, she will be considered as no childbearing potential only after confirmation by hormone levels.</p>",
		},
		{
			`        Inclusion Criteria:

        Inclusion Criteria For All Arms:

          1. Diagnosis of relapsed/refractory aggressive Non Hodgkin's Lymphoma (NHL) with
             histology based on established World Health Organization (WHO) criteria.

          2. Must have received ≥1 prior line of therapy for the treatment of current histology,
             there are no known curative treatment options available, or subject ineligible for
             potential curative options.

          3. Presence of radiographically measurable lymphadenopathy or extranodal lymphoid
             malignancy. Not applicable for cutaneous lesions.

          4. ECOG performance status of ≤2.

        Inclusion Criteria for Arm 1:

        1. Must have previously received rituximab, cyclophosphamide, doxorubicin hydrochloride
        (hydroxydaunorubicin), vincristine sulfate, and prednisone or equivalent regimen with
        stem-cell rescue. Or who are ineligible for highdose chemotherapy with stem-cell rescue
        and/or chimeric antigen receptor (CAR) T cell therapy.

        Ineligibility for high-dose therapy with stem cell rescue and/or CAR T-cell therapy will be
        determined by the investigator.

        Inclusion Criteria for Arm 2:

        1. Must have previously received rituximab, cyclophosphamide, doxorubicin hydrochloride
        (hydroxydaunorubicin), vincristine sulfate, and prednisone or equivalent regimen with
        stem-cell rescue. Or who are ineligible for highdose chemotherapy with stem-cell rescue
        and/or chimeric antigen receptor (CAR) T-cell therapy.

        Ineligibility for high-dose therapy with stem cell rescue and/or CAR T-cell therapy will be
        determined by the investigator.

        Inclusion Criteria for Arm 3:

        1. Must have previously received rituximab, cyclophosphamide, doxorubicin hydrochloride
        (hydroxydaunorubicin), vincristine sulfate, and prednisone or equivalent regimen with
        stem-cell rescue. Or who are ineligible for highdose chemotherapy with stem-cell rescue
        and/or chimeric antigen receptor (CAR) T cell therapy.

        Ineligibility for high-dose therapy with stem cell rescue and/or CAR T-cell therapy will be
        determined by the investigator.

        Inclusion Criteria for Arm 4:

        1. Must have previously received rituximab, cyclophosphamide, doxorubicin hydrochloride
        (hydroxydaunorubicin), vincristine sulfate, and prednisone or equivalent regimen with
        stem-cell rescue. Or who are ineligible for highdose chemotherapy with stem-cell rescue
        and/or chimeric antigen receptor (CAR) T cell therapy.

        Ineligibility for high-dose therapy with stem cell rescue and/or CAR T-cell therapy will be
        determined by the investigator.

        Exclusion Criteria:

        Exclusion Criteria For All Arms:

          1. History of prior malignancy except for the following: a) Malignancy treated with
             curative intent and with no evidence of active disease present for more than 2 years
             before screening and felt to be at low risk for recurrence by treating physician, b)
             Adequately treated lentigo maligna melanoma without current evidence of disease or
             adequately controlled nonmelanomatous skin cancer, c) Adequately treated carcinoma in
             situ without current evidence of disease, d) Evidence of severe or uncontrolled
             systemic disease, or current unstable or uncompensated respiratory or cardiac
             conditions, or uncontrolled hypertension, history of, or active, bleeding diatheses or
             uncontrolled active systemic fungal, bacterial, viral, or other infection, or
             intravenous anti-infective treatment within 2 weeks before first dose of study
             treatment.

          2. Serologic status reflecting active hepatitis B or C infection.

          3. Prior use of standard antilymphoma therapy or radiation therapy within 14 days of
             receiving the first dose of study treatment (not including palliative radiotherapy or
             palliative corticosteroid use).

          4. Requires ongoing immunosuppressive therapy, including systemic corticosteroids for
             treatment of lymphoid cancer or other conditions.

          5. For subjects under DLT review only: Any haematopoietic growth factors or darbepoetin
             within 14 days of the first dose of study treatment.

        Exclusion Criteria for Arm 1:

          1. Presence of central nervous system (CNS) lymphoma or leptomeningeal disease.

          2. Current refractory nausea and vomiting, malabsorption syndrome, disease significantly
             affecting gastrointestinal (GI) function, resection of the stomach, extensive small
             bowel resection that is likely to affect absorption, symptomatic inflammatory bowel
             disease, partial or complete bowel obstruction, or gastric restrictions and bariatric
             surgery, such as gastric bypass.

          3. Requires treatment with proton-pump inhibitors.

          4. Requires treatment with strong CYP3A inhibitors or inducers.

        Exclusion Criteria for Arm 2:

          1. Relative hypotension (&lt; 90/60 mm Hg) or clinically relevant orthostatic hypotension,
             including a fall in blood pressure of &gt;20 mm Hg.

          2. Uncontrolled hypertension requiring clinical intervention.

          3. At risk for brain perfusion problems based on medical history.

          4. Mean resting QT interval (QTc) calculated using Fridericia's formula (QTcF) &gt;470 msec
             for female subjects and &gt;450 msec for male subjects obtained from 3 electrocardiograms
             (ECGs), or congenital long QT syndrome.

          5. Presence of central nervous system (CNS) lymphoma or leptomeningeal disease.

          6. Known to have tested positive for human immunodeficiency virus (HIV) &amp; requires
             treatment with restricted medications.

          7. Current refractory nausea and vomiting, malabsorption syndrome, disease significantly
             affecting gastrointestinal (GI) function, resection of the stomach, extensive small
             bowel resection that is likely to affect absorption, symptomatic inflammatory bowel
             disease, partial or complete bowel obstruction, or gastric restrictions and bariatric
             surgery, such as gastric bypass.

          8. Requires treatment with proton-pump inhibitors.

        Exclusion Criteria for Arm 3:

          1. Presence of central nervous system (CNS) lymphoma or leptomeningeal disease.

          2. Current refractory nausea and vomiting, malabsorption syndrome, disease significantly
             affecting gastrointestinal (GI) function, resection of the stomach, extensive small
             bowel resection that is likely to affect absorption, symptomatic inflammatory bowel
             disease, partial or complete bowel obstruction, or gastric restrictions and bariatric
             surgery, such as gastric bypass.

          3. Requires treatment with proton-pump inhibitors.

          4. Red blood cell (RBC) transfusion dependence, defined as requiring more than 2 units of
             RBC transfusions during the 4-week period prior to screening.

          5. History of haemolytic anaemia or Evans syndrome in the last 3 months before enrolment.

          6. Positive IgG component of the direct antiglobulin test (DAT).

          7. Prior treatment with CD47 or SIRPα-targeting agents.

          8. Hypersensitivity to the active substance or to murine proteins, or to any of the other
             excipients of rituximab

        Exclusion Criteria for Arm 4:

          1. Presence of central nervous system (CNS) lymphoma or leptomeningeal disease.

          2. Current refractory nausea and vomiting, malabsorption syndrome, disease significantly
             affecting gastrointestinal (GI) function, resection of the stomach, extensive small
             bowel resection that is likely to affect absorption, symptomatic inflammatory bowel
             disease, partial or complete bowel obstruction, or gastric restrictions and bariatric
             surgery, such as gastric bypass.

          3. Requires treatment with proton-pump inhibitors.

          4. Requires treatment with CYP3A inhibitors or inducers or substrates of drug
             transporters.

          5. History of tuberculosis.

          6. Mean resting corrected QT interval (QTcF) &gt;450 msec obtained from 3 electrocardiograms
             (ECGs); clinically important ECG findings, or risk factors for QTc prolongation.

          7. Subjects receiving antiplatelet or anticoagulant therapies within 28 days of first
             dose of study drug.`, "<p>Inclusion Criteria:</p><p>Inclusion Criteria For All Arms:</p><ol><li>Diagnosis of relapsed/refractory aggressive Non Hodgkin's Lymphoma (NHL) with histology based on established World Health Organization (WHO) criteria.</li><li>Must have received ≥1 prior line of therapy for the treatment of current histology, there are no known curative treatment options available, or subject ineligible for potential curative options.</li><li>Presence of radiographically measurable lymphadenopathy or extranodal lymphoid malignancy. Not applicable for cutaneous lesions.</li><li>ECOG performance status of ≤2.</li></ol><p>Inclusion Criteria for Arm 1:</p><ol><li>Must have previously received rituximab, cyclophosphamide, doxorubicin hydrochloride (hydroxydaunorubicin), vincristine sulfate, and prednisone or equivalent regimen with stem-cell rescue. Or who are ineligible for highdose chemotherapy with stem-cell rescue and/or chimeric antigen receptor (CAR) T cell therapy.<p>Ineligibility for high-dose therapy with stem cell rescue and/or CAR T-cell therapy will be determined by the investigator.</p></li></ol><p>Inclusion Criteria for Arm 2:</p><ol><li>Must have previously received rituximab, cyclophosphamide, doxorubicin hydrochloride (hydroxydaunorubicin), vincristine sulfate, and prednisone or equivalent regimen with stem-cell rescue. Or who are ineligible for highdose chemotherapy with stem-cell rescue and/or chimeric antigen receptor (CAR) T-cell therapy.<p>Ineligibility for high-dose therapy with stem cell rescue and/or CAR T-cell therapy will be determined by the investigator.</p></li></ol><p>Inclusion Criteria for Arm 3:</p><ol><li>Must have previously received rituximab, cyclophosphamide, doxorubicin hydrochloride (hydroxydaunorubicin), vincristine sulfate, and prednisone or equivalent regimen with stem-cell rescue. Or who are ineligible for highdose chemotherapy with stem-cell rescue and/or chimeric antigen receptor (CAR) T cell therapy.<p>Ineligibility for high-dose therapy with stem cell rescue and/or CAR T-cell therapy will be determined by the investigator.</p></li></ol><p>Inclusion Criteria for Arm 4:</p><ol><li>Must have previously received rituximab, cyclophosphamide, doxorubicin hydrochloride (hydroxydaunorubicin), vincristine sulfate, and prednisone or equivalent regimen with stem-cell rescue. Or who are ineligible for highdose chemotherapy with stem-cell rescue and/or chimeric antigen receptor (CAR) T cell therapy.<p>Ineligibility for high-dose therapy with stem cell rescue and/or CAR T-cell therapy will be determined by the investigator.</p></li></ol><p>Exclusion Criteria:</p><p>Exclusion Criteria For All Arms:</p><ol><li>History of prior malignancy except for the following: a) Malignancy treated with curative intent and with no evidence of active disease present for more than 2 years before screening and felt to be at low risk for recurrence by treating physician, b) Adequately treated lentigo maligna melanoma without current evidence of disease or adequately controlled nonmelanomatous skin cancer, c) Adequately treated carcinoma in situ without current evidence of disease, d) Evidence of severe or uncontrolled systemic disease, or current unstable or uncompensated respiratory or cardiac conditions, or uncontrolled hypertension, history of, or active, bleeding diatheses or uncontrolled active systemic fungal, bacterial, viral, or other infection, or intravenous anti-infective treatment within 2 weeks before first dose of study treatment.</li><li>Serologic status reflecting active hepatitis B or C infection.</li><li>Prior use of standard antilymphoma therapy or radiation therapy within 14 days of receiving the first dose of study treatment (not including palliative radiotherapy or palliative corticosteroid use).</li><li>Requires ongoing immunosuppressive therapy, including systemic corticosteroids for treatment of lymphoid cancer or other conditions.</li><li>For subjects under DLT review only: Any haematopoietic growth factors or darbepoetin within 14 days of the first dose of study treatment.</li></ol><p>Exclusion Criteria for Arm 1:</p><ol><li>Presence of central nervous system (CNS) lymphoma or leptomeningeal disease.</li><li>Current refractory nausea and vomiting, malabsorption syndrome, disease significantly affecting gastrointestinal (GI) function, resection of the stomach, extensive small bowel resection that is likely to affect absorption, symptomatic inflammatory bowel disease, partial or complete bowel obstruction, or gastric restrictions and bariatric surgery, such as gastric bypass.</li><li>Requires treatment with proton-pump inhibitors.</li><li>Requires treatment with strong CYP3A inhibitors or inducers.</li></ol><p>Exclusion Criteria for Arm 2:</p><ol><li>Relative hypotension (&lt; 90/60 mm Hg) or clinically relevant orthostatic hypotension, including a fall in blood pressure of &gt;20 mm Hg.</li><li>Uncontrolled hypertension requiring clinical intervention.</li><li>At risk for brain perfusion problems based on medical history.</li><li>Mean resting QT interval (QTc) calculated using Fridericia's formula (QTcF) &gt;470 msec for female subjects and &gt;450 msec for male subjects obtained from 3 electrocardiograms (ECGs), or congenital long QT syndrome.</li><li>Presence of central nervous system (CNS) lymphoma or leptomeningeal disease.</li><li>Known to have tested positive for human immunodeficiency virus (HIV) &amp; requires treatment with restricted medications.</li><li>Current refractory nausea and vomiting, malabsorption syndrome, disease significantly affecting gastrointestinal (GI) function, resection of the stomach, extensive small bowel resection that is likely to affect absorption, symptomatic inflammatory bowel disease, partial or complete bowel obstruction, or gastric restrictions and bariatric surgery, such as gastric bypass.</li><li>Requires treatment with proton-pump inhibitors.</li></ol><p>Exclusion Criteria for Arm 3:</p><ol><li>Presence of central nervous system (CNS) lymphoma or leptomeningeal disease.</li><li>Current refractory nausea and vomiting, malabsorption syndrome, disease significantly affecting gastrointestinal (GI) function, resection of the stomach, extensive small bowel resection that is likely to affect absorption, symptomatic inflammatory bowel disease, partial or complete bowel obstruction, or gastric restrictions and bariatric surgery, such as gastric bypass.</li><li>Requires treatment with proton-pump inhibitors.</li><li>Red blood cell (RBC) transfusion dependence, defined as requiring more than 2 units of RBC transfusions during the 4-week period prior to screening.</li><li>History of haemolytic anaemia or Evans syndrome in the last 3 months before enrolment.</li><li>Positive IgG component of the direct antiglobulin test (DAT).</li><li>Prior treatment with CD47 or SIRPα-targeting agents.</li><li>Hypersensitivity to the active substance or to murine proteins, or to any of the other excipients of rituximab</li></ol><p>Exclusion Criteria for Arm 4:</p><ol><li>Presence of central nervous system (CNS) lymphoma or leptomeningeal disease.</li><li>Current refractory nausea and vomiting, malabsorption syndrome, disease significantly affecting gastrointestinal (GI) function, resection of the stomach, extensive small bowel resection that is likely to affect absorption, symptomatic inflammatory bowel disease, partial or complete bowel obstruction, or gastric restrictions and bariatric surgery, such as gastric bypass.</li><li>Requires treatment with proton-pump inhibitors.</li><li>Requires treatment with CYP3A inhibitors or inducers or substrates of drug transporters.</li><li>History of tuberculosis.</li><li>Mean resting corrected QT interval (QTcF) &gt;450 msec obtained from 3 electrocardiograms (ECGs); clinically important ECG findings, or risk factors for QTc prolongation.</li><li>Subjects receiving antiplatelet or anticoagulant therapies within 28 days of first dose of study drug.</li></ol>",
		},
	}
}
