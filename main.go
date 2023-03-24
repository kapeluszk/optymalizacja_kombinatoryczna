package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type city struct {
	number int
	x      int
	y      int
}

func distance(city1, city2 city) float64 {
	xDist := city1.x - city2.x
	yDist := city1.y - city2.y
	return math.Sqrt(math.Pow(float64(xDist), 2) + math.Pow(float64(yDist), 2))
}

func nearestNeighbor(cities []city) []city {
	visited := make(map[int]bool) // mapa odwiedzonych miast
	tour := make([]city, len(cities))

	// Wybieramy losowe miasto jako punkt początkowy
	currentCity := cities[rand.Intn(len(cities))]
	tour[0] = currentCity
	visited[currentCity.number] = true

	// Przeglądamy każde miasto w kolejności "najbliższego sąsiada"
	for i := 1; i < len(cities); i++ {
		nearestDistance := math.MaxFloat64
		var nearestCity city

		// Szukamy najbliższego miasta
		for _, city := range cities {
			if !visited[city.number] {
				dist := distance(currentCity, city)
				if dist < nearestDistance {
					nearestDistance = dist
					nearestCity = city
				}
			}
		}

		// Dodajemy najbliższe miasto do trasy i oznaczamy jako odwiedzone
		tour[i] = nearestCity
		visited[nearestCity.number] = true
		currentCity = nearestCity // aktualizujemy aktualne miasto
	}

	// Dodajemy punkt początkowy na końcu trasy
	tour = append(tour, tour[0])

	return tour
}

func readCitiesFromFile(filename string) ([]city, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var cities []city
	//n, err := strconv.Atoi(scanner.Text())
	//if err != nil {
	//	fmt.Println("first line is not an amount of cities")
	//}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 3 {
			return nil, fmt.Errorf("incorrect format of input file")
		}

		// Konwersja wartości na odpowiednie typy
		num, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		xArg, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		yArg, err := strconv.Atoi(fields[2])
		if err != nil {
			return nil, err
		}

		// Dodanie miasta do tablicy
		cities = append(cities, city{
			number: num,
			x:      xArg,
			y:      yArg,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cities, nil
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

	xyz, err := readCitiesFromFile(*filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nearestNeighbor(xyz))

}
