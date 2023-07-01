package main

func MapToSlice(mapData map[string]string) [][]string {
	slice := [][]string{}
	pair := []string{}

	for key, value := range mapData {
		pair = append(pair, key)
		pair = append(pair, value)
		slice = append(slice, pair)
		pair = []string{}
	}
	return slice
}
