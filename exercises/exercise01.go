package exercises

import "strconv"

func ConvertToInt(numStr string) (int, string) {

	if num, err := strconv.Atoi(numStr); err != nil {
		return -1, "No se pudo convertir correctamente -> " + numStr
	} else if num > 100 {
		return num, "Es mayor a 100"
	} else if num == 100 {
		return num, "Es igual a 100"
	} else if num < 100 {
		return num, "Es menor a 100"
	}

	return -1, "No se pudo convertir correctamente -> " + numStr
}
