package cube

import (
	"math"
)

type Screen struct {
	Width  int
	Height int
}
type CubeConfig struct {
	CubeWidth           int
	BackgroundASCIICode rune
	DistanceFromCam     int
	HorizontalOffset    float64
	K1                  float64
	IncrementSpeed      float64
	Screen              Screen
}

func (c *CubeConfig) CopyWithDefault() *CubeConfig {
	config := new(CubeConfig)
	*config = *c
	if config.CubeWidth == 0 {
		config.CubeWidth = 20
	}
	if config.BackgroundASCIICode == 0 {
		config.BackgroundASCIICode = '.'
	}
	if config.DistanceFromCam == 0 {
		config.DistanceFromCam = 100
	}
	if config.K1 == 0 {
		config.K1 = 40
	}
	if config.IncrementSpeed == 0 {
		config.IncrementSpeed = 0.6
	}
	if config.Screen.Height == 0 {
		config.Screen.Height = 44
	}
	if config.Screen.Width == 0 {
		config.Screen.Width = 160
	}

	return config
}

type Cube struct {
	cubeWidth           float64
	zBuffer             []float64
	buffer              []rune
	backgroundASCIICode rune
	distanceFromCam     float64
	horizontalOffset    float64
	k1                  float64
	incrementSpeed      float64
	heightWidthMultiply int
	screen              Screen
	coefA, coefB, coefC float64
}

func NewCube(cubeConfig *CubeConfig) *Cube {
	config := cubeConfig.CopyWithDefault()
	heightWidthMultiply := config.Screen.Width * config.Screen.Height
	return &Cube{
		cubeWidth:           (float64)(config.CubeWidth),
		zBuffer:             make([]float64, heightWidthMultiply),
		buffer:              make([]rune, heightWidthMultiply),
		backgroundASCIICode: config.BackgroundASCIICode,
		distanceFromCam:     (float64)(config.DistanceFromCam),
		horizontalOffset:    config.HorizontalOffset,
		k1:                  config.K1,
		incrementSpeed:      config.IncrementSpeed,
		screen:              config.Screen,
		heightWidthMultiply: heightWidthMultiply,
	}
}
func (c *Cube) Rotate(bufferChan chan []rune) {
	for {
		MemsetLoop(c.buffer, c.backgroundASCIICode)
		MemsetLoop(c.zBuffer, 0)
		for cubeX := -c.cubeWidth; cubeX < c.cubeWidth; cubeX += c.incrementSpeed {
			for cubeY := -c.cubeWidth; cubeY < c.cubeWidth; cubeY += c.incrementSpeed {
				c.calculateForSurface(cubeX, cubeY, -c.cubeWidth, '@')
				c.calculateForSurface(c.cubeWidth, cubeY, cubeX, '$')
				c.calculateForSurface(-c.cubeWidth, cubeY, -cubeX, '~')
				c.calculateForSurface(-cubeX, cubeY, c.cubeWidth, '#')
				c.calculateForSurface(cubeX, -c.cubeWidth, -cubeY, ';')
				c.calculateForSurface(cubeX, c.cubeWidth, cubeY, '+')
			}
		}
		bufferChan <- c.buffer

		c.coefA += 0.05
		c.coefB += 0.05
		c.coefC += 0.01
	}
}
func (c *Cube) calculateForSurface(cubeX, cubeY, cubeZ float64, ch rune) {
	x := c.calculateX(cubeX, cubeY, cubeZ)
	y := c.calculateY(cubeX, cubeY, cubeZ)
	z := c.calculateZ(cubeX, cubeY, cubeZ) + c.distanceFromCam

	ooz := 1 / z

	xp := (int)(float64(c.screen.Width)/2 + c.horizontalOffset + c.k1*ooz*x*2)
	yp := (int)(float64(c.screen.Height)/2 + c.k1*ooz*y)

	idx := xp + yp*c.screen.Width
	if idx >= 0 && idx < c.heightWidthMultiply && ooz > c.zBuffer[idx] {
		c.zBuffer[idx] = ooz
		c.buffer[idx] = ch
	}
}

func (c *Cube) calculateX(i, j, k float64) float64 {
	return j*math.Sin(c.coefA)*math.Sin(c.coefB)*math.Cos(c.coefC) - k*math.Cos(c.coefA)*math.Sin(c.coefB)*math.Cos(c.coefC) +
		j*math.Cos(c.coefA)*math.Sin(c.coefC) + k*math.Sin(c.coefA)*math.Sin(c.coefC) + i*math.Cos(c.coefB)*math.Cos(c.coefC)
}

func (c *Cube) calculateY(i, j, k float64) float64 {
	return j*math.Cos(c.coefA)*math.Cos(c.coefC) + k*math.Sin(c.coefA)*math.Cos(c.coefC) -
		j*math.Sin(c.coefA)*math.Sin(c.coefB)*math.Sin(c.coefC) + k*math.Cos(c.coefA)*math.Sin(c.coefB)*math.Sin(c.coefC) -
		i*math.Cos(c.coefB)*math.Sin(c.coefC)
}

func (c *Cube) calculateZ(i, j, k float64) float64 {
	return k*math.Cos(c.coefA)*math.Cos(c.coefB) - j*math.Sin(c.coefA)*math.Cos(c.coefB) + i*math.Sin(c.coefB)
}

func MemsetLoop[T any](a []T, v T) {
	for i := range a {
		a[i] = v
	}
}
