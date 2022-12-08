package simplemath

func New2DMap[T any](rows int, columns int) map[int]map[int]T {
	m := map[int]map[int]T{}
	for i := 0; i < rows; i++ {
		m[i] = map[int]T{}
	}
	return m
}

func New2DSlice[T any](rows int, columns int) [][]T {
	m := make([][]T, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]T, columns)
	}
	return m
}
