package cycleAnalysis

// CycleAnalysis
type CycleAnalysis struct {
	CycleStart  float64
	HigherHigh1 float64
	ChangeCSHH1 float64

	HigherLow1   float64
	ChangeHH1HL1 float64

	HigherHigh2  float64
	ChangeHL1HH2 float64

	HigherLow2   float64
	ChangeHH2HL2 float64

	HigherHigh3  float64
	ChangeHL2HH3 float64

	HigherLow3   float64
	ChangeHH3HL3 float64

	HigherHigh4  float64
	ChangeHL3HH4 float64

	HigherLow4   float64
	ChangeHH4HL4 float64

	CycleHigh   float64
	ChangeHL4CH float64

	LowerLow1   float64
	ChangeCHLL1 float64

	LowerHigh1   float64
	ChangeLL1LH1 float64

	LowerLow2    float64
	ChangeLH1LL2 float64

	LowerHigh2   float64
	ChangeLL2LH2 float64

	LowerLow3    float64
	ChangeLH2LL3 float64

	LowerHigh3   float64
	ChangeLL3LH3 float64

	LowerLow4    float64
	ChangeLH3LL4 float64

	LowerHigh4   float64
	ChangeLL4LH4 float64

	CycleEnd    float64
	ChangeLH4CE float64
}
