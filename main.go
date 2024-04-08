package main

import "fmt"


func main() {

	ns := GetNutritionalScore(NutritionalData{
		Energy: EnergyFromKcal(),
		Sugars: SugarGram(),
		StauratedFattyAcids: StauratedFattyAcids(),
		Sodium: SodiumMiligram(),
		Fruits: FruitPercent(),
		Fiber: FiberGram(), 
		Protein: ProteinGram(),
	}, Food)

	fmt.Printf("Nutritional Score:%d\n", ns.Value)
}