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