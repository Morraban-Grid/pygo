// frequency.go

package main

// Función para contar la frecuencia de elementos en un slice de strings
func CountFrequency(items []string) map[string]int {

	// Creamos un mapa para almacenar la frecuencia de cada elemento
	freq := make(map[string]int)

	// Iteramos sobre cada elemento en el slice y contamos su frecuencia
	for _, item := range items {
		freq[item]++
	}

	return freq
}
