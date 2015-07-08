package idmask

var TotalIDs uint = 63

func Mask(ids []uint) uint {
	var mask uint
	for _, id := range ids {
		mask = setBit(mask, id)
	}
	return mask
}

func Unmask(n uint) []uint {
	var ids []uint

	// not using 0 because this use case is for ids, and id 0 is just REQUIRED.
	for i := uint(0); i <= TotalIDs; i++ {
		if hasBit(n, i) {
			ids = append(ids, i)
		}
	}
	return ids
}

func setBit(n uint, pos uint) uint {
	n |= (1 << pos)
	return n
}

func clearBit(n uint, pos uint) uint {
	mask := ^(uint(1) << pos)
	n &= mask
	return n
}

func hasBit(n uint, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}
