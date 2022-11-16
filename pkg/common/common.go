package common

type Rotator interface {
	Rotate(chan<- []rune)
}

type Screen struct {
	Width  int
	Height int
}
type Config struct {
	CubeWidth           int
	BackgroundASCIICode rune
	DistanceFromCam     int
	HorizontalOffset    float64
	K1                  float64
	K2                  float64
	R1                  float64
	R2                  float64
	IncrementSpeed      float64
	ThetaIncrementSpeed float64
	PhiIncrementSpeed   float64
	Screen              Screen
}

func (c *Config) CopyWithDefaultForCube() *Config {
	config := new(Config)
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
func (c *Config) CopyWithDefaultForTorus() *Config {
	config := new(Config)
	*config = *c
	if config.BackgroundASCIICode == 0 {
		config.BackgroundASCIICode = '.'
	}
	if config.DistanceFromCam == 0 {
		config.DistanceFromCam = 25
	}
	if config.K2 == 0 {
		config.K2 = 5
	}
	if config.R1 == 0 {
		config.R1 = 1
	}
	if config.R2 == 0 {
		config.R2 = 2
	}
	if config.ThetaIncrementSpeed == 0 {
		config.ThetaIncrementSpeed = 0.03
	}
	if config.PhiIncrementSpeed == 0 {
		config.PhiIncrementSpeed = 0.01
	}
	if config.Screen.Height == 0 {
		config.Screen.Height = 44
	}
	if config.Screen.Width == 0 {
		config.Screen.Width = 160
	}

	return config
}
func MemsetLoop[T any](a []T, v T) {
	for i := range a {
		a[i] = v
	}
}
