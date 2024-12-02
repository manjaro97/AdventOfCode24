package MathHelper

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
