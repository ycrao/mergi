package ease

import "math"

type Move func(t float64) float64


func mapVlaues(value, start1, stop1, start2, stop2 float64) float64 {
	return start2 + (stop2-start2)*((value-start1)/(stop1-start1))
}

func AnimateXY(move Move, px, minX, maxX, minY, maxY float64) (y float64) {
	t := mapVlaues(px, minX, maxX, 0, 1)
	my := move(t)
	return mapVlaues(my, 0, 1, minY, maxY)
}

func AnimateYX(move Move, py, minX, maxX, minY, maxY float64) (x float64) {
	t := mapVlaues(py, minY, maxY, 0, 1)
	mx := move(t)
	return mapVlaues(mx, 0, 1, minX, maxX)
}

// this contains easing animations algorithms
// The visualization of theses algorothms can be found on > http://easings.net/

func InBack(t float64) float64 {
	s := 1.70158
	return t * t * ((s+1)*t - s)
}

func InBounce(t float64) float64 {
	return 1 - OutBounce(1-t)
}

func InCirc(t float64) float64 {
	if t >= 1 {
		return t
	}
	return -1 * (math.Sqrt(1-t*t) - 1)
}

func InCubic(t float64) float64 {
	return t * t * t
}

func InElastic(t float64) float64 {
	s := 1.70158
	p := 0.0
	a := 1.0
	if t == 0 {
		return 0.0
	}
	if t == 1 {
		return 1.0
	}
	if p == 0 {
		p = 1 * .3
	}
	if a < 1 {
		a = 1.0
		s = p / 4
	} else {
		s = p / (2 * math.Pi) * math.Asin(1/a)
	}
	t = t - 1
	return -(a * math.Pow(2, 10*(t)) * math.Sin((t*1-s)*(2*math.Pi)/p))
}

func InExpo(t float64) float64 {
	if t == 0 {
		return 1
	}
	return 1 * math.Pow(2, 10*(t/1-1))
}

func InOutBack(t float64) float64 {
	s := 1.70158
	t = t / 0.5
	s = s * 1.525
	if (t) < 1 {
		return 1 / 2 * (t * t * ((s+1)*t - s))
	}
	t = t - 2
	return 1 / 2 * ((t*t*(s+1)*t + s) + 2)
}

func InOutBounce(t float64) float64 {
	if t < 1/2 {
		return InBounce(t*2) * .5
	}
	return OutBounce(t*2-1)*.5 + 1*.5
}

func InOutCirc(t float64) float64 {
	t = t / 0.5
	if t < 1 {
		return -1 / 2 * (math.Sqrt(1-t*t) - 1)
	}
	t = t - 2
	return 1 / 2 * (math.Sqrt(1-t*t) + 1)
}

func InOutCubic(t float64) float64 {
	t = t / (1 / 2)
	if t < 1 {
		return 1 / 2 * t * t * t
	}
	t = t - 2
	return 1 / 2 * (t*t*t + 2)
}

func InOutElastic(t float64) float64 {
	s := 1.70158
	p := 0.0
	a := 1.0
	if t == 0 {
		return 0.0
	}
	t = t / (1 / 2)
	if t == 2 {
		return 1.0
	}
	if p == 0.0 {
		p = 1 * (.3 * 1.5)
	}
	if a < 1 {
		a = 1.0
		// can cause a bug
		s = p / 4
	} else {
		s = p / (2 * math.Pi) * math.Asin(1/a)
	}
	t = t - 1
	if t < 1 {
		return -.5 * (a * math.Pow(2, 10*t) * math.Sin((t*1-s)*(2*math.Pi)/p))
	}
	return a*math.Pow(2, -10*t)*math.Sin((t*1-s)*(2*math.Pi)/p)*.5 + 1
}

func InOutExpo(t float64) float64 {
	if t == 0 {
		return 0.0
	}
	if t == 1 {
		return 1.0
	}
	t = t / (1 / 2)
	if t < 1 {
		return 1 / 2 * math.Pow(2, 10*(t-1))
	}
	t = t - 1
	return 1 / 2 * (-math.Pow(2, -10*t) + 2)
}

func InOutQuad(t float64) float64 {
	t = t / (1 / 2)
	if t < 1 {
		return 1 / 2 * t * t
	}
	t = t - 1
	return -1 / 2 * (t*(t-2) - 1)
}

func InOutQuart(t float64) float64 {
	t = t / (1 / 2)
	if t < 1 {
		return 1 / 2 * t * t * t * t
	}
	t = t - 2
	return -1 / 2 * (t*t*t*t - 2)
}

func InOutQuint(t float64) float64 {
	t = t / (1 / 2)
	if t < 1 {
		return (1 / 2) * t * t * t * t * t
	}
	t = t - 2
	return 1 / 2 * (t*t*t*t*t + 2)
}

func InOutSine(t float64) float64 {
	return -1 / 2 * (math.Cos(math.Pi*t/1) - 1)
}

func InQuad(t float64) float64 {
	return t * t
}

func InQuart(t float64) float64 {
	return t * t * t * t
}

func InQuint(t float64) float64 {
	return t * t * t * t * t
}

func InSine(t float64) float64 {
	return -1*math.Cos(t/1*(math.Pi/2)) + 1
}

func OutBack(t float64) float64 {
	s := 1.70158
	t = t/1 - 1
	return 1 * (t*t*((s+1)*t+s) + 1)
}

func OutBounce(t float64) float64 {
	if t < (1 / 2.75) {
		return 1 * (7.5625 * t * t)
	} else if t < (2 / 2.75) {
		t = t - (1.5 / 2.75)
		return 1 * (7.5625*t*t + .75)
	} else if t < (2.5 / 2.75) {
		t = t - (2.25 / 2.75)
		return 1 * (7.5625*t*t + .9375)
	}
	t = t - (2.625 / 2.75)
	return 1 * (7.5625*t*t + .984375)
}

func OutCirc(t float64) float64 {
	t = t - 1
	return 1 * math.Sqrt(1-(t*t))
}

func OutCubic(t float64) float64 {
	t = t/1 - 1
	return 1 * (t*t*t + 1)
}

func OutElastic(t float64) float64 {
	s := 1.70158
	p := 0.0
	a := 1.0
	if t == 0 {
		return 0.0
	}
	if t == 1 {
		return 1.0
	}
	if p == 0 {
		p = 1 * .3
	}
	if a < 1 {
		a = 1.0
		s = p / 4
	} else {
		s = p / (2 * math.Pi) * math.Asin(1/a)
	}
	return a*math.Pow(2, -10*t)*math.Sin((t*1-s)*(2*math.Pi)/p) + 1
}

func OutExpo(t float64) float64 {
	if t == 1 {
		return 1
	}
	return 1 * (-math.Pow(2, -10*t/1) + 1)
}

func OutQuad(t float64) float64 {
	return -1 * t * (t - 2)
}

func OutQuart(t float64) float64 {
	t = t/1 - 1
	return -1 * (t*t*t*t - 1)
}

func OutQuint(t float64) float64 {
	t = t - 1
	return (t*t*t*t*t + 1)
}

func OutSine(t float64) float64 {
	return 1 * math.Sin(t/1*(math.Pi/2))
}

func Linear(t float64) float64 {
	return t
}