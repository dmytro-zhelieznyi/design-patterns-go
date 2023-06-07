package reader

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadData(path string) (map[string]int, error) {
	capitals := make(map[string]int)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		keyValue := strings.Split(line, ",")
		value, err := strconv.Atoi(keyValue[1])
		if err != nil {
			log.Fatal(err)
		}
		capitals[keyValue[0]] = value
	}

	return capitals, nil
}
