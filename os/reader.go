package os

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(fpath string, lineNum uint) (lines []string, err error) {
	f, err := os.Open(fpath)
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var count uint
	for scanner.Scan() {
		count++
		if count > lineNum {
			return
		}
		line := scanner.Text()
		lines = append(lines, line)
	}
	return
}

func ReadFirstLine(fpath string) (line string, err error) {
	lines, err := ReadLines(fpath, 1)
	if err != nil {
		return
	}
	return lines[0], nil
}

func ReadFirstLineAsInt(fpath string) (i int, err error) {
	data, err := ReadFirstLine(fpath)
	if err != nil {
		return
	}
	i, err = strconv.Atoi(data)
	if err != nil {
		return
	}
	return
}

func ReadInt(fpath string) (i int, err error) {
	data, err := os.ReadFile(fpath)
	if err != nil {
		return
	}
	i, err = strconv.Atoi(string(data))
	if err != nil {
		return
	}
	return
}
