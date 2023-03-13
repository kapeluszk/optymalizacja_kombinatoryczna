package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type city struct {
	number int
	x      int
	y      int
}

func parseToStruct(inputMatrix [][]int) []city {
	ret := make([]city, len(inputMatrix))
	for i, row := range inputMatrix {
		ret[i] = city{
			number: row[0],
			x:      row[1],
			y:      row[2],
		}
	}
	return ret
}

func main() {
	filename := flag.String("file", "optymalizacja_kombinatoryczna/city.txt", "podaj ścieżkę do pliku txt z miastami")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	cities := make([][]int, n)
	rows := make([]int, 3)

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			scanner.Scan()
			k, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			rows[j] = k

		}
		cities[i] = rows
	}

	xyz := parseToStruct(cities)
	fmt.Println(xyz)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
