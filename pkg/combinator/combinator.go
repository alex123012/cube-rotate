package combinator

import (
	"fmt"
	"time"

	"github.com/alex123012/cube-rotate/pkg/common"
)

type CubePrinter interface {
	RotateAll(config *common.Config, framesPerSecond int)
}
type cubeCombinator struct {
	rotators []common.Rotator
}

func NewCubeCombinator(rotators []common.Rotator) CubePrinter {
	return &cubeCombinator{
		rotators: rotators,
	}
}
func (c *cubeCombinator) RotateAll(config *common.Config, framesPerSecond int) {
	cubeCount := len(c.rotators)
	cubesChan := make([]chan []rune, cubeCount)
	for i := 0; i < cubeCount; i++ {
		cubesChan[i] = make(chan []rune, 1)
		go c.rotators[i].Rotate(cubesChan[i])
	}

	resultBuffer := make([]rune, config.Screen.Height*config.Screen.Width)
	fmt.Printf("\x1b[2J")
	frame_delay := 1000 / framesPerSecond
	for range time.Tick(time.Duration(frame_delay) * time.Millisecond) {
		common.MemsetLoop(resultBuffer, config.BackgroundASCIICode)
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
