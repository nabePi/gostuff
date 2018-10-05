package raindrops

import "strconv"

func Convert(number int) string {
	var retrunValue string
	var rainOutput = map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			if value, isTrue := rainOutput[i]; isTrue {
				retrunValue += value
			}
		}
	}
	if retrunValue == "" {
		retrunValue = strconv.Itoa(number)
	}
	return retrunValue
}
