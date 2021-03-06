package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		result[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			result[y][x] = uint8(f(x, y))
		}
	}

	return result
}

func f(x, y int) int {
	return x ^ y
}

func main() {
	pic.Show(Pic)
}
