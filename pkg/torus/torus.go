package torus

import (
	"math"

	"github.com/alex123012/cube-rotate/pkg/common"
)

var luminanceSymbols = []rune("',-~:;=!*#$@")

const (
	theta_spacing = 0.07
	phi_spacing   = 0.02
)

type Torus struct {
	coefA  float64
	coefB  float64
	coefR1 float64
	coefR2 float64
	coefK1 float64
	coefK2 float64

	heightWidthMultiply int
	horizontalOffset    float64
	distanceFromCam     float64

	screen              common.Screen
	backgroundASCIICode rune
	zBuffer             []float64
	buffer              []rune
}

func NewTorus(cfg *common.Config) common.Rotator {
	config := cfg.CopyWithDefaultForTorus()
	heightWidthMultiply := config.Screen.Width * config.Screen.Height
	return &Torus{
		coefR1: config.R1, coefR2: config.R2,
		coefA: 1, coefB: 1,
		coefK2:              config.K2,
		coefK1:              float64(config.Screen.Width) * config.K2 * 3.0 / (8.0 * (config.R1 + config.R2)),
		zBuffer:             make([]float64, heightWidthMultiply),
		buffer:              make([]rune, heightWidthMultiply),
		distanceFromCam:     (float64)(config.DistanceFromCam),
		backgroundASCIICode: config.BackgroundASCIICode,
		horizontalOffset:    config.HorizontalOffset,
		screen:              config.Screen,
		heightWidthMultiply: heightWidthMultiply,
	}
}

func (t *Torus) Rotate(bufferChan chan<- []rune) {
	for {
		common.MemsetLoop(t.buffer, t.backgroundASCIICode)
		common.MemsetLoop(t.zBuffer, 0)

		// theta goes around the cross-sectional circle of a torus
		for theta := 0.0; theta < 2.0*math.Pi; theta += theta_spacing {
			// phi goes around the center of revolution of a torus
			for phi := 0.0; phi < 2.0*math.Pi; phi += phi_spacing {
				t.calculateSurface(phi, theta)
			}
		}
		bufferChan <- t.buffer
		t.coefA += 0.07
		t.coefB += 0.03
	}
}
func (t *Torus) calculateSurface(phi, theta float64) {
	// final 3D (x,y,z) coordinate after rotations, directly from our math above
	x := t.calculateX(phi, theta)
	y := t.calculateY(phi, theta)
	z := t.calculateZ(phi, theta) + t.distanceFromCam

	ooz := 1 / z // "one over z"

	// x and y projection.  note that y is negated here, because y goes up in
	// 3D space but down on 2D displays.
	xp := (int)(float64(t.screen.Width)/2 + t.horizontalOffset + t.coefK1*ooz*x*2)
	yp := (int)(float64(t.screen.Height)/2 + t.coefK1*ooz*y)

	// calculate luminance.  ugly, but correct.
	// L ranges from -sqrt(2) to +sqrt(2).  If it's < 0, the surface is
	// pointing away from us, so we won't bother trying to plot it.
	if L := t.calculateLum(phi, theta); L > 0 {
		// test against the z-buffer.  larger 1/z means the pixel is closer to
		// the viewer than what's already plotted.
		if idx := xp + yp*t.screen.Width; idx >= 0 && idx < t.heightWidthMultiply && ooz > t.zBuffer[idx] {
			t.zBuffer[idx] = ooz
			luminance_index := int(L * 8.0) // this brings L into the range 0..11 (8*sqrt(2) = 11.3)
			// now we lookup the character corresponding to the luminance and plot it in our output:
			t.buffer[idx] = luminanceSymbols[luminance_index]
		}
	}
}
func (t *Torus) calculateX(phi, theta float64) float64 {
	return (t.coefR2+t.coefR1*math.Cos(theta))*
		(math.Cos(t.coefB)*math.Cos(phi)+
			math.Sin(t.coefA)*math.Sin(t.coefB)*math.Sin(phi)) -
		t.coefR1*math.Sin(theta)*math.Cos(t.coefA)*math.Sin(t.coefB)
}

func (t *Torus) calculateY(phi, theta float64) float64 {
	return (t.coefR2+t.coefR1*math.Cos(theta))*
		(math.Sin(t.coefB)*math.Cos(phi)-
			math.Sin(t.coefA)*math.Cos(t.coefB)*math.Sin(phi)) +
		t.coefR1*math.Sin(theta)*math.Cos(t.coefA)*math.Cos(t.coefB)
}

func (t *Torus) calculateZ(phi, theta float64) float64 {
	return t.coefK2 +
		math.Cos(t.coefA)*(t.coefR2+t.coefR1*math.Cos(theta))*math.Sin(phi) +
		t.coefR1*math.Sin(theta)*math.Sin(t.coefA)
}

func (t *Torus) calculateLum(phi, theta float64) float64 {
	return math.Cos(phi)*math.Cos(theta)*math.Sin(t.coefB) -
		math.Cos(t.coefA)*math.Cos(theta)*math.Sin(phi) -
		math.Sin(t.coefA)*math.Sin(theta) +
		math.Cos(t.coefB)*(math.Cos(t.coefA)*math.Sin(theta)-math.Cos(theta)*math.Sin(t.coefA)*math.Sin(phi))
}
