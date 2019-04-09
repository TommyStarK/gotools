package bit

func SetBit(target, position int) int {
	return target | (1 << uint8(position))
}

func GetBit(target, position int) bool {
	return target&(1<<uint8(position)) != 0
}
