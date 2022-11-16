package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alex123012/cube-rotate/pkg/combinator"
	"github.com/alex123012/cube-rotate/pkg/common"
	"github.com/alex123012/cube-rotate/pkg/cube"
	"github.com/alex123012/cube-rotate/pkg/torus"
)

var (
	screenWidth  int = 287
	screenHeight int = 67
	cubesCount   int = 1
	torusesCount int = 0

	framesPerSecond = 30

	backgroudnSymbol = "."

	cubesHorizontalOffsetsCMD   = make(SlicaVar, 0)
	torusesHorizontalOffsetsCMD = make(SlicaVar, 0)

	cubesWidthsCMD      = make(SlicaVar, 0)
	torusesDistancesCMD = make(SlicaVar, 0)
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

	parseFlags()
	cubesHorizontalOffsetValues, torusesHorizontalOffsetValues, cubesWidthValues, torusesDistanceValues := prepareValues()
	config := &common.Config{
		Screen: common.Screen{
			Width:  screenWidth,
			Height: screenHeight,
		},
		BackgroundASCIICode: []rune(backgroudnSymbol)[0],
	}
	cubeConfig := config.CopyWithDefaultForCube()
	torusConfig := config.CopyWithDefaultForTorus()

	cubes := make([]common.Rotator, cubesCount)
	for i := range cubes {
		cubeConfig.HorizontalOffset = cubesHorizontalOffsetValues[i]
		cubeConfig.CubeWidth = cubesWidthValues[i]
		cubes[i] = cube.NewCube(cubeConfig)
	}

	toruses := make([]common.Rotator, torusesCount)
	for i := range toruses {
		torusConfig.HorizontalOffset = torusesHorizontalOffsetValues[i]
		torusConfig.DistanceFromCam = torusesDistanceValues[i]
		toruses[i] = torus.NewTorus(torusConfig)
	}
	comb := combinator.NewCubeCombinator(append(cubes, toruses...))
	comb.RotateAll(config, framesPerSecond)
}

func parseFlags() {
	flag.IntVar(&screenHeight,
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
		"count of cubes to print",
	)
	flag.IntVar(
		&torusesCount,
		"toruses.count",
		cubesCount,
		"count of toruses to print",
	)
	flag.Var(
		&cubesHorizontalOffsetsCMD,
		"cubes.horizontal-offsets",
		"horizontal positions of cubes, count must be equal to cubes.count",
	)
	flag.Var(
		&torusesHorizontalOffsetsCMD,
		"toruses.horizontal-offsets",
		"horizontal positions of toruses, count must be equal to toruses.count",
	)
	flag.Var(
		&cubesWidthsCMD,
		"cubes.widths",
		"widths of cubes, count must be equal to cubes.count",
	)
	flag.Var(
		&torusesDistancesCMD,
		"toruses.distances",
		"distances for which toruses will be print",
	)
	flag.StringVar(
		&backgroudnSymbol,
		"screen.background-symbol",
		backgroudnSymbol,
		"background symbol to use",
	)
	flag.IntVar(
		&framesPerSecond,
		"screen.fps",
		framesPerSecond,
		"fps",
	)

	flag.Parse()

}

func prepareValues() ([]float64, []float64, []int, []int) {

	var (
		cubesHorizontalOffsetValues   = make([]float64, 0)
		cubesWidthValues              = make([]int, 0)
		torusesHorizontalOffsetValues = make([]float64, 0)
		torusesDistanceValues         = make([]int, 0)
	)
	if len(cubesHorizontalOffsetsCMD) > 0 {
		for _, v := range cubesHorizontalOffsetsCMD {
			floatV, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println("Error parsing CMD horisontal offset value for cube")
				os.Exit(1)
			}
			cubesHorizontalOffsetValues = append(cubesHorizontalOffsetValues, floatV)
		}
	} else {
		tmp := 20.0
		cubesHorizontalOffsetValues = []float64{tmp * -3}
	}
	if len(torusesHorizontalOffsetsCMD) > 0 {
		for _, v := range torusesHorizontalOffsetsCMD {
			floatV, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println("Error parsing CMD horisontal offset value for cube")
				os.Exit(1)
			}
			torusesHorizontalOffsetValues = append(torusesHorizontalOffsetValues, floatV)
		}
	} else {
		tmp := 20.0
		torusesHorizontalOffsetValues = []float64{tmp * 3}
	}

	if len(cubesWidthsCMD) > 0 {
		for _, v := range cubesWidthsCMD {
			intV, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error parsing CMD cube width value")
				os.Exit(1)
			}
			cubesWidthValues = append(cubesWidthValues, intV)
		}
	} else {
		cubesWidthValues = []int{25}
	}

	if len(torusesDistancesCMD) > 0 {
		for _, v := range torusesDistancesCMD {
			intV, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Error parsing CMD cube width value")
				os.Exit(1)
			}
			torusesDistanceValues = append(torusesDistanceValues, intV)
		}
	} else {
		torusesDistanceValues = []int{25}
	}

	if len(cubesHorizontalOffsetValues) < cubesCount || len(cubesWidthValues) < cubesCount {
		fmt.Println(
			"Provide horisontal offsets or cube widths for all cubes (count must be equal to cubes.count)",
		)
		os.Exit(1)
	}

	if len(torusesHorizontalOffsetValues) < torusesCount || len(torusesDistanceValues) < torusesCount {
		fmt.Println(
			"Provide horisontal offsets or cube widths for all cubes (count must be equal to cubes.count)",
		)
		os.Exit(1)
	}
	return cubesHorizontalOffsetValues, torusesHorizontalOffsetValues, cubesWidthValues, torusesDistanceValues
}
