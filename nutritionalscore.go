package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore  struct{
	Value		int
	Positive	int
	Negative	int
	ScoreType 	ScoreType
}

type EnergyKJ float64
type SugarGram float64
type StauratedFattyAcids float64
type SodiumMiligram float64
type FruitPercent float64
type FiberGram float64
type ProteinGram float64

type NutritionalData struct{
	Energy					EnergyKJ
	Sugars					SugarGram
	StauratedFattyAcids		StauratedFattyAcids
	Sodium					SodiumMiligram
	Fruits					FruitPercent
	Fiber					FiberGram
	Protein					ProteinGram
	IsWater					bool

}

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugarsLevels = []float64{45, 40, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5}
var stauratedFattyAcidsLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90}
var fiberLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

var energyLevelsBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarLevelsBeverage = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}


func (e EnergyKJ)GetPoints(st ScoreType) int {

}

func (s SugarGram)GetPoints(st ScoreType) int {

}

func (sfa StauratedFattyAcids)GetPoints(st ScoreType) int {

}

func (s SodiumMiligram)GetPoints(st ScoreType) int{

}

func (f FruitPercent)GetPoints(st ScoreType) int {

}

func (f FiberGram)GetPoints(st ScoreType) int {

}


func (p ProteinGram)GetPoints(st ScoreType) int {

}

func EnergyFromKcal(kcal float64) EnergyKJ{
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(SaltMg float64) SodiumMiligram{
	return SodiumMiligram(SaltMg/2.5)
}

func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore{

	value :=0
	positive :=0
	negative :=0
	
	if st != Water {
		fruitsPoints := n.Fruits.GetPoints(st)
		fiberPoints := n.Fiber.GetPoints(st)

		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.StauratedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = fruitsPoints + fiberPoints + n.Protein.GetPoints(st)

	}

	return NutritionalScore{
		Value: value,
		Positive: positive,
		Negative: negative,
		ScoreType: st,
	}
}

func getPointsFromRange( v float64, steps []float64) int{
	lenSteps := len(steps)
	for i,l := range steps{
		if v > l {
			return lenSteps - 1
		}
	}
	return 0
}