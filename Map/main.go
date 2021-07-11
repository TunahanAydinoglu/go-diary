package main

import "fmt"

func main() {

	trCities := make(map[int]string)

	trCities[07] = "Antalya"
	trCities[32] = "Isparta"

	fmt.Println(trCities)

	trCitiesKey := make(map[string]string)
	trCitiesKey["ANT"] = "Antalya"
	trCitiesKey["ISP"] = "Isparta"
	trCitiesKey["ANK"] = "Ankara"

	fmt.Println(trCitiesKey)

	delete(trCitiesKey, "ANK")
	fmt.Println(trCitiesKey)

	for index, item := range trCitiesKey {
		fmt.Println(index, item)
	}

	keys := make([]string, len(trCitiesKey))

	i := 0
	for index := range trCitiesKey {
		keys[i] = index
		i++
	}

	fmt.Println(keys)
}
