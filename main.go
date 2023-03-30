package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

	//otwieramy plik
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//otwieramy skaner i tworzymy tablicę miast
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

// to jest niepotrzebne ale może się przydać kiedyś
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

func generateTxtInstance(citiesAmount int) (string, error) {
	var newFileName string
	fmt.Println("Podaj nazwę nowego pliku")
	fmt.Scanln(&newFileName)

	fileNew, err := os.Create(newFileName)
	if err != nil {
		log.Fatal()
	}
	defer fileNew.Close()

	writer := bufio.NewWriter(fileNew)

	fmt.Fprintf(fileNew, "%d\n", citiesAmount)
	writer.Flush()

	for i := 0; i < citiesAmount; i++ {
		x := rand.Intn(2000)
		y := rand.Intn(2000)
		fmt.Fprintf(writer, "%d %d %d\n", i+1, x, y)
		writer.Flush()
	}
	return newFileName, err
}

func main() {
	filename := flag.String("file", "optymalizacja_kombinatoryczna/city.txt", "podaj ścieżkę do pliku txt z miastami")
	flag.Parse()

	var choose int
	fmt.Println("1 - wygeneruj plik, 2 - użyj pliku podanego we fladze")
	fmt.Scanln(&choose)

	switch choose {
	case 1:
		var amount int
		fmt.Println("Podaj ile miast ma być wygenerowane")
		fmt.Scanln(&amount)
		file, err := generateTxtInstance(amount)
		if err != nil {
			log.Fatal()
		}
		fmt.Printf("utworzono plik %s", file)

	case 2:
		start := time.Now()
		xyz, err := readCitiesFromFile(*filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(nearestNeighbor(xyz))

		elapsed := time.Since(start)
		println("czas wykonania algorytmu w sekundach to: ", elapsed.Seconds())

	default:
		fmt.Println("wprowadzono niepoprawną liczbę")
		break
	}

}
