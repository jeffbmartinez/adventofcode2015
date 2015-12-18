// puzzle at http://adventofcode.com/day/9

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var locations []string
var distanceMap = make(map[string]map[string]int)

func main() {
	distances := strings.Split(Input, "\n")

	for _, distanceString := range distances {
		location1, location2, distance := ParseDistance(distanceString)

		if _, ok := distanceMap[location1]; !ok {
			distanceMap[location1] = make(map[string]int)
		}

		if _, ok := distanceMap[location2]; !ok {
			distanceMap[location2] = make(map[string]int)
		}

		distanceMap[location1][location2] = distance
		distanceMap[location2][location1] = distance
	}

	for location := range distanceMap {
		locations = append(locations, location)
	}

	sort.Strings(locations)

	paths := Permutations(locations)

	shortestPath := paths[0]
	shortestDistance := PathDistance(shortestPath)

	for _, path := range paths {
		pathDistance := PathDistance(path)

		if pathDistance < shortestDistance {
			shortestPath = path
			shortestDistance = pathDistance

			fmt.Printf("Shortest path so far is %v long:\n\t%v\n", shortestDistance, shortestPath)
		}
	}
}

func PathDistance(path []string) int {
	distance := 0

	for i := 0; i < len(path)-1; i++ {
		distance += Distance(path[i], path[i+1])
	}

	return distance
}

func Distance(city1, city2 string) int {
	return distanceMap[city1][city2]
}

func ParseDistance(s string) (string, string, int) {
	tokens := strings.Split(s, " ")
	if len(tokens) != 5 || tokens[1] != "to" || tokens[3] != "=" {
		fmt.Printf("Warning: Unexpected input line format (%v)\n", s)
		return "", "", 0
	}

	distance, err := strconv.Atoi(tokens[4])
	if err != nil {
		fmt.Printf("Warning: Expected distance to be an integer, found (%v)\n", s)
		return "", "", 0
	}

	return tokens[0], tokens[2], distance
}

func Permutations(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	if len(strs) == 1 {
		return [][]string{strs}
	}

	var permutations [][]string

	for _, permutation := range Permutations(strs[1:]) {
		permutations = append(permutations, permutateIn(strs[0], permutation)...)
	}

	return permutations
}

func permutateIn(s string, strs []string) [][]string {
	permutations := make([][]string, 0)
	for i := 0; i < len(strs)+1; i++ {
		tempStrs := make([]string, len(strs))
		copy(tempStrs, strs[:])

		newPermutation := append(tempStrs, "")
		copy(newPermutation[i+1:], newPermutation[i:])
		newPermutation[i] = s

		permutations = append(permutations, newPermutation)
	}

	return permutations
}

var Input = `Tristram to AlphaCentauri = 34
Tristram to Snowdin = 100
Tristram to Tambi = 63
Tristram to Faerun = 108
Tristram to Norrath = 111
Tristram to Straylight = 89
Tristram to Arbre = 132
AlphaCentauri to Snowdin = 4
AlphaCentauri to Tambi = 79
AlphaCentauri to Faerun = 44
AlphaCentauri to Norrath = 147
AlphaCentauri to Straylight = 133
AlphaCentauri to Arbre = 74
Snowdin to Tambi = 105
Snowdin to Faerun = 95
Snowdin to Norrath = 48
Snowdin to Straylight = 88
Snowdin to Arbre = 7
Tambi to Faerun = 68
Tambi to Norrath = 134
Tambi to Straylight = 107
Tambi to Arbre = 40
Faerun to Norrath = 11
Faerun to Straylight = 66
Faerun to Arbre = 144
Norrath to Straylight = 115
Norrath to Arbre = 135
Straylight to Arbre = 127`
