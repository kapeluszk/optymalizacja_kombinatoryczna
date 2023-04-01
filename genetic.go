package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	cities   []City
	distance float64
}

type City struct {
	number int
	x      int
	y      int
}

func Distance(city1, city2 City) float64 {
	xDist := city1.x - city2.x
	yDist := city1.y - city2.y
	return math.Sqrt(math.Pow(float64(xDist), 2) + math.Pow(float64(yDist), 2))
}

func NearestNeighbor(cities []City) []City {
	visited := make(map[int]bool) // mapa odwiedzonych miast
	tour := make([]City, len(cities))

	// Wybieramy losowe miasto jako punkt początkowy
	currentCity := cities[rand.Intn(len(cities))]
	tour[0] = currentCity
	visited[currentCity.number] = true

	// Przeglądamy każde miasto w kolejności "najbliższego sąsiada"
	for i := 1; i < len(cities); i++ {
		nearestDistance := math.MaxFloat64
		var nearestCity City

		// Szukamy najbliższego miasta
		for _, city := range cities {
			if !visited[city.number] {
				dist := Distance(currentCity, city)
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

func ReadCitiesFromFile(filename string) ([]City, error) {

	//otwieramy plik
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//otwieramy skaner i tworzymy tablicę miast
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var cities []City
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
		cities = append(cities, City{
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
func ParseToStruct(inputMatrix [][]int) []City {
	ret := make([]City, len(inputMatrix))
	for i, row := range inputMatrix {
		ret[i] = City{
			number: row[0],
			x:      row[1],
			y:      row[2],
		}
	}
	return ret
}

func GenerateTxtInstance(citiesAmount int) (string, error) {
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

func runGeneticAlgorithm(cities []City, generations int, populationSize int) ([]City, float64) {

	cities = cities[:len(cities)-1]
	// initialize population
	population := make([][]City, populationSize)
	for i := range population {
		population[i] = shuffleCities(cities)
	}

	// evaluate fitness of initial population
	fitness := make([]float64, populationSize)
	for i, route := range population {
		fitness[i] = calculateFitness(route)
	}

	// run genetic algorithm
	for i := 0; i < generations; i++ {
		// select parents for crossover
		parents := selectParents(population, fitness)

		// perform crossover and mutation
		children := make([][]City, populationSize)
		for j := range children {
			parent1 := parents[rand.Intn(len(parents))]
			parent2 := parents[rand.Intn(len(parents))]
			child := crossover(parent1, parent2)
			child = mutate(child)
			children[j] = child
		}

		// replace population with children
		population = children

		// evaluate fitness of new population
		for j, route := range population {
			fitness[j] = calculateFitness(route)
		}
	}

	// find best route in final population
	bestRoute := population[0]
	bestFitness := fitness[0]
	for i := 1; i < len(population); i++ {
		if fitness[i] < bestFitness {
			bestRoute = population[i]
			bestFitness = fitness[i]
		}
	}

	bestRoute = append(bestRoute, bestRoute[0])
	bestDistance := 0.0
	for i := 1; i < len(bestRoute); i++ {
		bestDistance += distanceBetweenCities(bestRoute[i-1], bestRoute[i])
	}
	return bestRoute, bestDistance
}

func shuffleCities(cities []City) []City {
	shuffled := make([]City, len(cities))
	copy(shuffled, cities)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}

func calculateFitness(route []City) float64 {
	distance := 0.0
	for i := 1; i < len(route); i++ {
		distance += distanceBetweenCities(route[i-1], route[i])
	}
	distance += distanceBetweenCities(route[len(route)-1], route[0])
	return 1 / distance
}

func distanceBetweenCities(city1 City, city2 City) float64 {
	dx := city1.x - city2.x
	dy := city1.y - city2.y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func selectParents(population [][]City, fitness []float64) [][]City {
	parents := make([][]City, 0)
	for len(parents) < len(population) {
		index1 := rand.Intn(len(population))
		index2 := rand.Intn(len(population))
		if fitness[index1] > fitness[index2] {
			parents = append(parents, population[index1])
		} else {
			parents = append(parents, population[index2])
		}
	}
	return parents
}

func crossover(parent1 []City, parent2 []City) []City {
	child := make([]City, len(parent1))
	copy(child, parent1)

	start := rand.Intn(len(parent1))
	end := rand.Intn(len(parent1))
	if end < start {
		start, end = end, start
	}

	for i := start; i <= end; i++ {
		city := parent2[i]
		for j := range child {
			if child[j].number == city.number {
				child[j] = child[i]
				child[i] = city
				break
			}
		}
	}

	return child
}

func mutate(route []City) []City {
	if rand.Float64() < 0.01 {
		i := rand.Intn(len(route))
		j := rand.Intn(len(route))
		route[i], route[j] = route[j], route[i]
	}
	return route
}
