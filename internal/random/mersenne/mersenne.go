package mersenne

const (
	n         = 624
	m         = 397
	matrixA   = 0x9908B0DF
	upperMask = 0x80000000
	lowerMask = 0x7fffffff
)

type MT19937 struct {
	mt  [n]uint32
	idx int
}

func NewMT19937(seed uint32) *MT19937 {
	mt := &MT19937{}
	mt.mt[0] = seed
	for i := 1; i < n; i++ {
		mt.mt[i] = 1812433253*(mt.mt[i-1]^((mt.mt[i-1])>>30)) + uint32(i)
	}
	mt.idx = n
	return mt
}

func (mt *MT19937) ExtractNumber() uint32 {
	if mt.idx >= n {
		mt.twist()
	}

	y := mt.mt[mt.idx]
	y ^= y >> 11
	y ^= (y << 7) & 0x9D2C5680
	y ^= (y << 15) & 0xEFC60000
	y ^= y >> 18

	mt.idx++
	return y
}

func (mt *MT19937) twist() {
	for i := 0; i < n; i++ {
		y := (mt.mt[i] & upperMask) + (mt.mt[(i+1)%n] & lowerMask)
		mt.mt[i] = mt.mt[(i+m)%n] ^ (y >> 1)
		if y%2 != 0 {
			mt.mt[i] ^= matrixA
		}
	}
	mt.idx = 0
}
