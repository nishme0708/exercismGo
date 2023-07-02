package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodleQ := 0
	sauceQ := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodleQ += 50
		} else if layer == "sauce" {
			sauceQ += 0.2
		}
	}
	return noodleQ, sauceQ
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendLayers []string, myLayers []string) {
	myLayers[len(myLayers)-1] = friendLayers[len(friendLayers)-1]
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(quantities []float64, scaledQuan int) []float64 {
    actualScaled := float64(scaledQuan) / 2
	var newQuan []float64
	for _, quan := range quantities {
		newQuan = append(newQuan,actualScaled * quan)
	}
	return newQuan
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
