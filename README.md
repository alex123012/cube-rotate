# Cube Rotate
This is simple CLI app that will print rotating cubes and toruses in terminal.
![Screencast](.github/screencast.gif)
## Installation
Open [release](https://github.com/alex123012/cube-rotate/releases) page and find appropriate release for your system and arch, than download and unpack binary
```bash
mkdir -p bin
wget https://github.com/alex123012/cube-rotate/releases/download/v0.0.1/cube-rotate-${VERSION}-${OS}-${ARCH}.tar.gz
tar -C bin/ -xf cube-rotate-${VERSION}-${OS}-${ARCH}.tar.gz cube-rotate
```

## Usage
```bash
bin/cube-rotate --screen.width 205 --screen.height 44 --cubes.horizontal-offsets -50 --toruses.horizontal-offsets 50 --screen.background-symbol '.' --cubes.count 1 --toruses.count 1 --toruses.distances 10 --screen.fps 100
```
```text
Usage of cube-rotate:
  -cubes.count int
    	count of cubes to print (default 1)
  -cubes.horizontal-offsets value
    	horizontal positions of cubes, count must be equal to cubes.count
  -cubes.widths value
    	widths of cubes, count must be equal to cubes.count
  -screen.background-symbol string
    	background symbol to use (default ".")
  -screen.fps int
    	fps (default 30)
  -screen.height int
    	your terminal symbol counts in one line (default 67)
  -screen.width int
    	your terminal symbol counts in one column (default 287)
  -toruses.count int
    	count of toruses to print (default 1)
  -toruses.distances value
    	distances for which toruses will be print
  -toruses.horizontal-offsets value
    	horizontal positions of toruses, count must be equal to toruses.count
```