package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alex123012/cube-rotate/pkg/combinator"
	"github.com/alex123012/cube-rotate/pkg/cube"
)

type SlicaVar []string

func (h *SlicaVar) String() string {
	return strings.Join(*h, ", ")
}

func (h *SlicaVar) Set(value string) error {
	if strings.Contains(value, ",") {
		values := strings.Split(value, ",")
		*h = append(*h, values...)
	} else {
		*h = append(*h, value)
	}
	return nil
}

func main() {
	var (
		screenWidth  int = 287
		screenHeight int = 67
		cubesCount   int = 1

		backgroudnSymbol = "."

		horizontalOffsetsCMD   = make(SlicaVar, 0)
		horizontalOffsetValues = make([]float64, 0)

		cubesWidthsCMD  = make(SlicaVar, 0)
		cubeWidthValues = make([]int, 0)
	)

	flag.IntVar(
		&screenHeight,
		"screen.height",
		screenHeight,
		"your terminal symbol counts in one line",
	)
	flag.IntVar(
		&screenWidth,
		"screen.width",
		screenWidth,
		"your terminal symbol counts in one column",
	)
	flag.IntVar(
		&cubesCount,
		"cubes.count",
		cubesCount,
		"your terminal symbol counts in one line",
	)
	flag.Var(
		&horizontalOffsetsCMD,
		"cubes.horizontal-offsets",
		"horizontal positions of cubes, count must be equal to cubes.count",
	)
	flag.StringVar(
		&backgroudnSymbol,
		"cubes.background-symbol",
		backgroudnSymbol,
		"background symbol to use",
	)
	flag.Var(&cubesWidthsCMD, "cubes.widths", "widths of cubes, count must be equal to cubes.count")

	flag.Parse()

	config := (&cube.CubeConfig{
		Screen: cube.Screen{
			Width:  screenWidth,
			Height: screenHeight,
		},
		BackgroundASCIICode: []rune(backgroudnSymbol)[0],
	}).CopyWithDefault()

	cubes := make([]*cube.Cube, cubesCount)

	if len(horizontalOffsetsCMD) > 0 {
		for _, v := range horizontalOffsetsCMD {
			floatV, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println("Error parsing CMD horisontal offset value")
				os.Exit(1)
			}
			horizontalOffsetValues = append(horizontalOffsetValues, floatV)
		}
	} else {
		tmp := 20.0
		horizontalOffsetValues = []float64{tmp * -5, tmp * -1, tmp * 2, tmp * 4.5, tmp * 6}
	}

	if len(cubesWidthsCMD) > 0 {
		for _, v := range cubesWidthsCMD {
			floatV, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error parsing CMD cube width value")
				os.Exit(1)
			}
			cubeWidthValues = append(cubeWidthValues, floatV)
		}
	} else {
		cubeWidthValues = []int{25, 20, 15, 10, 5}
	}

	if len(horizontalOffsetValues) < cubesCount || len(cubeWidthValues) < cubesCount {
		fmt.Println(
			"Provide horisontal offsets or cube widths for all cubes (count must be equal to cubes.count)",
		)
		os.Exit(1)
	}
	for i := range cubes {
		config.HorizontalOffset = horizontalOffsetValues[i]
		config.CubeWidth = cubeWidthValues[i]
		cubes[i] = cube.NewCube(config)
	}
	comb := combinator.NewCubeCombinator(cubes)
	comb.RotateAll(config)
}
