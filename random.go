package uniformrandom

import "math"

const (
	NTAB = 32
	IA = 16807
	IM = 2147483647
	IQ = 127773
	IR = 2836
	NDIV = 1+(IM-1)/NTAB
	MAX_RANDOM_RANGE = 0x7FFFFFFF
	AM = 1.0/IM
	EPS = 1.2e-7
	RNMX = 1.0-EPS
)

// Uniform Random Stream Based on Valve's Source SDK
type Stream struct {
	mIdum int
	mIy int
	mIv [NTAB]int
}

func (s *Stream) SetSeed(iSeed int) {
	s.mIdum = iSeed
	if iSeed >= 0 {
		s.mIdum = -iSeed
	}

	s.mIy = 0
}

func (s *Stream) GenerateRandomNumber() int {
	var j, k int

	if s.mIdum <= 0 || s.mIy == 0 {
		if -s.mIdum < 1 {
			s.mIdum = 1
		} else {
			s.mIdum = -s.mIdum
		}

		for j := NTAB+7; j >= 0; j-- {
			k = s.mIdum/IQ
			s.mIdum = IA*(s.mIdum-k*IQ)-IR*k
			if s.mIdum < 0 {
				s.mIdum += IM
			}
			if j < NTAB {
				s.mIv[j] = s.mIdum
			}
		}
		s.mIy = s.mIv[0]
	}

	k = s.mIdum/IQ
	s.mIdum = IA * (s.mIdum - k * IQ) - IR*k
	if s.mIdum < 0 {
		s.mIdum += IM
	}

	j = s.mIy/NDIV

	s.mIy = s.mIv[j]
	s.mIv[j] = s.mIdum

	return s.mIy
}

func (s *Stream) RandomFloat(flLow, flHigh float32) float32 {
	// float in [0,1)
	fl := AM * float32(s.GenerateRandomNumber())
	if fl > RNMX {
		fl = RNMX
	}

	return (fl * (flHigh - flLow)) + flLow // float in [low,high)
}

func (s *Stream) RandomFloatExp(flMinVal, flMaxVal, flExponent float32) float32 {
	// float in [0,1)
	fl := AM * float32(s.GenerateRandomNumber())
	if fl > RNMX {
		fl = RNMX
	}

	if flExponent != 1.0 {
		fl = float32(math.Pow(float64(fl), float64(flExponent)))
	}

	return (fl * (flMaxVal - flMinVal)) + flMinVal // float in [low,high)
}

func (s *Stream) RandomInt(iLow, iHigh int) int {
	x := iHigh - iLow + 1

	if x <= 1 || MAX_RANDOM_RANGE < x-1 {
		return iLow
	}

	// From Source Engine 2007:
	// The following maps a uniform distribution on the interval [0,MAX_RANDOM_RANGE]
	// to a smaller, client-specified range of [0,x-1] in a way that doesn't bias
	// the uniform distribution unfavorably. Even for a worst case x, the loop is
	// guaranteed to be taken no more than half the time, so for that worst case x,
	// the average number of times through the loop is 2. For cases where x is
	// much smaller than MAX_RANDOM_RANGE, the average number of times through the
	// loop is very close to 1.
	maxAcceptable := MAX_RANDOM_RANGE - ((MAX_RANDOM_RANGE+1) % x)

	var n int

	for {
		n = s.GenerateRandomNumber()

		if n <= maxAcceptable {
			break
		}
	}

	return iLow + (n % x)
}
