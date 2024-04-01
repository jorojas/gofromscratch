package maps

import "fmt"

func ShowMaps() {
	countries := make(map[string]string)

	countries["Mexico"] = "D.F."
	countries["Colombia"] = "Bogota"

	fmt.Println(countries)
	fmt.Println(countries["Colombia"])

	championship := map[string]int{
		"Madrid":     39,
		"Caracas FC": 20,
		"Boca":       19,
	}

	fmt.Println(championship)

	for team, points := range championship {
		fmt.Printf("Team %s, has %d points", team, points)
		fmt.Println()
	}

	delete(championship, "Boca")
	fmt.Println("-----")

	for team, points := range championship {
		fmt.Printf("Team %s, has %d points", team, points)
		fmt.Println()
	}

	fmt.Println("-----")

	points, exists := championship["Juventus"]

	fmt.Printf("Team exists %t, and has %d points", exists, points)

}
