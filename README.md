# Cube Rotate
This is simple CLI app that will print rotating cubes in terminal.
![Screencast](.github/screencast.gif)
## Installation
Open [release](https://github.com/alex123012/cube-rotate/releases) page and find appropriate release for your system and arch, than download and unpack binary
```bash
mkdir -p bin
wget https://github.com/alex123012/cube-rotate/releases/download/v0.0.1/cube-rotate-${VERSION}-${OS}-${ARCH}.tar.gz
tar -C bin/ -xf cube-rotate-${VERSION}-${OS}-${ARCH}.tar.gz cube-rotate
```

## Usage
```text
Usage of bin/cube-rotate:
  -cubes.count int
        your terminal symbol counts in one line (default 287) (default 1)
  -cubes.horizontal-offsets value
        horizontal positions of cubes, count must be equal to cubes.count
  -cubes.widths value
        widths of cubes, count must be equal to cubes.count
  -screen.height int
        your terminal symbol counts in one line (default 287) (default 67)
  -screen.width int
        your terminal symbol counts in one column (default 67) (default 287)
```