package temps

import (
	"fmt"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

*/

func UniteTempsInGroupsByTenDegree(temps []float64) (map[int][]float64, error) {
	if len(temps) == 0 {
		return nil, fmt.Errorf("temps length equals zero")
	}

	result := make(map[int][]float64, len(temps))

	for _, temp := range temps {
		step := getTempStep(temp)
		updatedTemps := append(result[step], temp)
		result[step] = updatedTemps
	}
	return result, nil
}

func getTempStep(temp float64) int {
	return int(temp/10) * 10
}
