// Package easing contains easing functions.
package easing

import "math"

// LinearEase interpolates between a start and end value linearly.
func LinearEase(t, start, end float64) float64 {
	return start + (end-start)*t
}

// QuadraticEaseIn interpolates between a start and end value with a quadratic easing.
func QuadraticEaseIn(t, start, end float64) float64 {
	t = t * t
	return start + (end-start)*t
}

// QuadraticEaseOut interpolates between a start and end value with a quadratic easing.
func QuadraticEaseOut(t, start, end float64) float64 {
	t = -t * (t - 2)
	return start + (end-start)*t
}

// QuadraticEaseInOut interpolates between a start and end value with a quadratic easing.
func QuadraticEaseInOut(t, start, end float64) float64 {
	if t < 0.5 {
		t = 2 * t * t
	} else {
		t = -2*t*t + 4*t - 1
	}
	return start + (end-start)*t
}

// CubicEaseIn interpolates between a start and end value with a cubic easing.
func CubicEaseIn(t, start, end float64) float64 {
	t = t * t * t
	return start + (end-start)*t
}

// CubicEaseOut interpolates between a start and end value with a cubic easing.
func CubicEaseOut(t, start, end float64) float64 {
	t = math.Pow(t-1, 3) + 1
	return start + (end-start)*t
}

// CubicEaseInOut interpolates between a start and end value with a cubic easing.
func CubicEaseInOut(t, start, end float64) float64 {
	if t < 0.5 {
		t = 4 * t * t * t
	} else {
		t = math.Pow(2*t-2, 3)*0.5 + 1
	}
	return start + (end-start)*t
}

// QuarticEaseIn interpolates between a start and end value with a quartic easing.
func QuarticEaseIn(t, start, end float64) float64 {
	t = t * t * t * t
	return start + (end-start)*t
}

// QuarticEaseOut interpolates between a start and end value with a quartic easing.
func QuarticEaseOut(t, start, end float64) float64 {
	t = math.Pow(t-1, 3)*(1-t) + 1
	return start + (end-start)*t
}

// QuarticEaseInOut interpolates between a start and end value with a quartic easing.
func QuarticEaseInOut(t, start, end float64) float64 {
	if t < 0.5 {
		t = 8 * t * t * t * t
	} else {
		t = -math.Pow(t-1, 4)*8 + 1
	}
	return start + (end-start)*t
}

// QuinticEaseIn interpolates between a start and end value with a quintic easing.
func QuinticEaseIn(t, start, end float64) float64 {
	t = t * t * t * t * t
	return start + (end-start)*t
}

// QuinticEaseOut interpolates between a start and end value with a quintic easing.
func QuinticEaseOut(t, start, end float64) float64 {
	t = math.Pow(t-1, 5) + 1
	return start + (end-start)*t
}

// QuinticEaseInOut interpolates between a start and end value with a quintic easing.
func QuinticEaseInOut(t, start, end float64) float64 {
	if t < 0.5 {
		t = 16 * t * t * t * t * t
	} else {
		t = math.Pow(2*t-2, 5)*0.5 + 1
	}
	return start + (end-start)*t
}

// SineEaseIn interpolates between a start and end value with a sine easing.
func SineEaseIn(t, start, end float64) float64 {
	t = math.Sin((t-1)*math.Pi/2) + 1
	return start + (end-start)*t
}

// SineEaseOut interpolates between a start and end value with a sine easing.
func SineEaseOut(t, start, end float64) float64 {
	t = math.Sin(t * math.Pi / 2)
	return start + (end-start)*t
}

// SineEaseInOut interpolates between a start and end value with a sine easing.
func SineEaseInOut(t, start, end float64) float64 {
	t = 0.5 * (1 - math.Cos(t*math.Pi))
	return start + (end-start)*t
}

// CircularEaseIn interpolates between a start and end value with a circular easing.
func CircularEaseIn(t, start, end float64) float64 {
	t = 1 - math.Sqrt(1-t*t)
	return start + (end-start)*t
}

// CircularEaseOut interpolates between a start and end value with a circular easing.
func CircularEaseOut(t, start, end float64) float64 {
	t = math.Sqrt(t * (2 - t))
	return start + (end-start)*t
}

// CircularEaseInOut interpolates between a start and end value with a circular easing.
func CircularEaseInOut(t, start, end float64) float64 {
	if t < 0.5 {
		t = 0.5 * (1 - math.Sqrt(1-4*t*t))
	} else {
		t = 0.5 * (math.Sqrt((-2*t+3)*(2*t-1)) + 1)
	}
	return start + (end-start)*t
}

// ExponentialEaseIn interpolates between a start and end value with a exponential easing.
func ExponentialEaseIn(t, start, end float64) float64 {
	if t != 0.0 {
		t = math.Pow(2, 10*(t-1))
	}
	return start + (end-start)*t
}

// ExponentialEaseOut interpolates between a start and end value with a exponential easing.
func ExponentialEaseOut(t, start, end float64) float64 {
	if t != 1.0 {
		t = 1 - math.Pow(2, -10*t)
	}
	return start + (end-start)*t
}

// ExponentialEaseInOut interpolates between a start and end value with a exponential easing.
func ExponentialEaseInOut(t, start, end float64) float64 {
	if t != 0.0 && t != 1.0 {
		if t < 0.5 {
			t = 0.5 * math.Pow(2, 20*t-10)
		} else {
			t = -0.5*math.Pow(2, -20*t+10) + 1
		}
	}
	return start + (end-start)*t
}

// ElasticEaseIn interpolates between a start and end value with a elastic easing.
func ElasticEaseIn(t, start, end float64) float64 {
	t = math.Sin(13*math.Pi/2*t) * math.Pow(2, 10*(t-1))
	return start + (end-start)*t
}

// ElasticEaseOut interpolates between a start and end value with a elastic easing.
func ElasticEaseOut(t, start, end float64) float64 {
	t = math.Sin(-13*math.Pi/2*(t+1))*math.Pow(2, -10*t) + 1
	return start + (end-start)*t
}

// ElasticEaseInOut interpolates between a start and end value with a elastic easing.
func ElasticEaseInOut(t, start, end float64) float64 {
	if t < 0.5 {
		t = 0.5 * math.Sin(13*math.Pi*t) * math.Pow(2, 20*t-10)
	} else {
		t = 0.5 * (math.Sin(-13*math.Pi*t)*math.Pow(2, -20*t+10) + 2)
	}
	return start + (end-start)*t
}

// BackEaseIn interpolates between a start and end value with a back easing.
func BackEaseIn(t, start, end float64) float64 {
	t = t*t*t - t*math.Sin(t*math.Pi)
	return start + (end-start)*t
}

// BackEaseOut interpolates between a start and end value with a back easing.
func BackEaseOut(t, start, end float64) float64 {
	t = 1 - t
	t = 1 - (t*t*t - t*math.Sin(t*math.Pi))
	return start + (end-start)*t
}

// BackEaseInOut interpolates between a start and end value with a back easing.
func BackEaseInOut(t, start, end float64) float64 {
	if t < 0.5 {
		t = 2 * t
		t = 0.5 * (t*t*t - t*math.Sin(t*math.Pi))
	} else {
		t = 1 - 2*t + 1
		t = 0.5*(1-t*t*t+t*math.Sin(t*math.Pi)) + 0.5
	}
	return start + (end-start)*t
}

// BounceEaseIn interpolates between a start and end value with a bounce easing.
func BounceEaseIn(t, start, end float64) float64 {
	t = 1 - BounceEaseOut(1-t, 0, 1)
	return start + (end-start)*t
}

// BounceEaseOut interpolates between a start and end value with a bounce easing.
func BounceEaseOut(t, start, end float64) float64 {
	if t < 4/11.0 {
		t = 121 * t * t / 16.0
	} else if t < 8/11.0 {
		t = 363/40.0*t*t - 99/10.0*t + 17/5.0
	} else if t < 9/10.0 {
		t = 4356/361.0*t*t - 35442/1805.0*t + 16061/1805.0
	} else {
		t = 54/5.0*t*t - 513/25.0*t + 268/25.0
	}
	return start + (end-start)*t
}

// BounceEaseInOut interpolates between a start and end value with a bounce easing.
func BounceEaseInOut(t, start, end float64) float64 {
	if t < 0.5 {
		t = 0.5 * BounceEaseIn(t*2, 0, 1)
	} else {
		t = 0.5*BounceEaseOut(t*2-1, 0, 1) + 0.5
	}
	return start + (end-start)*t
}
