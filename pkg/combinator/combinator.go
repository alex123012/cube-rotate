package combinator

import (
	"fmt"
	"time"

	"github.com/alex123012/cube-rotate/pkg/cube"
)

type CubePrinter interface {
	RotateAll(config *cube.CubeConfig)
}
type cubeCombinator struct {
	cubes []*cube.Cube
}

func NewCubeCombinator(cubes []*cube.Cube) CubePrinter {
	return &cubeCombinator{
		cubes: cubes,
	}
}
func (c *cubeCombinator) RotateAll(config *cube.CubeConfig) {
	cubeCount := len(c.cubes)
	cubesChan := make([]chan []rune, cubeCount)
	for i := 0; i < cubeCount; i++ {
		cubesChan[i] = make(chan []rune, 1)
		go c.cubes[i].Rotate(cubesChan[i])
	}

	resultBuffer := make([]rune, config.Screen.Height*config.Screen.Width)
	fmt.Printf("\x1b[2J")
	for {
		cube.MemsetLoop(resultBuffer, config.BackgroundASCIICode)
		for _, ch := range cubesChan {
			CompareRunesSlices(resultBuffer, <-ch, config.BackgroundASCIICode)
		}
		fmt.Printf("\x1b[H")
		for k := 0; k < config.Screen.Width*config.Screen.Height; k++ {
			if k%config.Screen.Width != 0 {
				fmt.Print(string(resultBuffer[k]))
			} else {
				fmt.Println()
			}
		}
		time.Sleep(time.Microsecond * 5000 * 2)
	}
}

func CompareRunesSlices(orig, comp []rune, backgroundValue rune) {
	for i := range orig {
		if vc := comp[i]; vc != backgroundValue {
			orig[i] = vc
		}
	}
}
