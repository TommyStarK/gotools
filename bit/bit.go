package bit

// SetBit sets the bit at the given position
func SetBit(target, position int) int {
	return target | (1 << uint8(position))
}

// GetBit returns whether the bit at the given position is set or not
func GetBit(target, position int) bool {
	return target&(1<<uint8(position)) != 0
}
