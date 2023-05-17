package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	filename := flag.String("file", "city.txt", "podaj ścieżkę do pliku txt z miastami")
	flag.Parse()

	var choose int
	fmt.Println("1 - wygeneruj plik, 2 - użyj pliku podanego we fladze, 3 - algorytm genetyczny dla pliku we fladze")
	fmt.Scanln(&choose)

	switch choose {
	case 1:
		var amount int
		fmt.Println("Podaj ile miast ma być wygenerowane")
		fmt.Scanln(&amount)
		file, err := base.GenerateTxtInstance(amount)
		if err != nil {
			log.Fatal()
		}
		fmt.Printf("utworzono plik %s", file)

	case 2:
		start := time.Now()
		xyz, err := base.ReadCitiesFromFile(*filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(base.NearestNeighbor(xyz))

		elapsed := time.Since(start)
		println("czas wykonania algorytmu w sekundach to: ", elapsed.Seconds())

	case 3:
		xyz, err := base.ReadCitiesFromFile(*filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("podaj liczbę generacji")
		var generation int
		fmt.Scanln(&generation)

		fmt.Println("podaj wielkość populacji")
		var pop int
		fmt.Scanln(&pop)

		greedy := base.NearestNeighbor(xyz)

		answer, bestRoute := base.runGeneticAlgorithm(greedy, generation, pop)

		fmt.Println(answer)
		fmt.Println(bestRoute)

	default:
		fmt.Println("wprowadzono niepoprawną liczbę")
		break
	}

}
