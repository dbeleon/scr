package scr

type Scrambler struct {
	Length     int
	Feedbacks  []int
	Polynomial uint64

	val uint64
}

func New(length int, feedbacks []int, polynomial uint64) *Scrambler {
	if length < 1 || length > 64 {
		panic("invalid scrambler polynomial length")
	}
	if len(feedbacks) == 0 || len(feedbacks) > length {
		panic("invalid scrambler feedbacks length")
	}
	for _, f := range feedbacks {
		if f < 0 || f >= length {
			panic("invalid scrambler feedback")
		}
	}
	s := &Scrambler{
		Length:     length,
		Feedbacks:  feedbacks,
		Polynomial: polynomial,
		val:        polynomial,
	}
	return s
}

func (s *Scrambler) ScrambleAdditive(data []byte) {
	for i := 0; i < len(data); i++ {
		d := data[i]
		for b := byte(0); b < 8; b++ {
			bit := (d >> b) & 1
			for _, f := range s.Feedbacks {
				bit ^= byte((s.val >> f) & 1)
			}
			d = d&(^(1 << b)) | (bit << b)
			s.val = ((s.val >> (s.Length - 1) & 1) | (s.val << 1)) & (uint64(0xFFFF_FFFF_FFFF_FFFF) >> (64 - s.Length))
		}
		data[i] = d
	}
}

func (s *Scrambler) DescrambleAdditive(data []byte) {
	s.Reset()
	s.ScrambleAdditive(data)
}

func (s *Scrambler) Reset() {
	s.val = s.Polynomial
}
