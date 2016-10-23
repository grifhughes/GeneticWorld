package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

//helpers
func r_int_range(min, max int) int {
	return min + rand.Intn(max-min)
}

func r_string(s int) string {
	b := make([]byte, s)
	for i := 0; i < s; i++ {
		b[i] = byte(r_int_range(32, 126))
	}
	return string(b)
}

func replace_at_index(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

//creates initial random generation
func init_generation(c []string, f []int, t string) {
	for i := 0; i < len(f); i++ {
		c[i] = r_string(len(t) - 1)
		f[i] = test_fitness(t, c[i])
	}
}

//tests candidate strings against target by Hamming distance
func test_fitness(s1, s2 string) int {
	count := len(s1) - 1
	for i := 0; i < len(s1)-1; i++ {
		if s1[i] == s2[i] {
			count--
		}
	}
	return count
}

//finds fittest candidate in a generation
func lowest_fitness(f []int) int {
	min := math.MaxInt32
	for i := 0; i < len(f); i++ {
		if f[i] < min {
			min = f[i]
			break
		}
	}
	return min
}

//calculates fitnesses for new candidate generation
func calc_new_fitnesses(c []string, f []int, t string) {
	for i := 0; i < len(f); i++ {
		f[i] = test_fitness(t, c[i])
	}
}

//picks 2 candidates at random and selects the fitter to become a parent
func binary_tournament(c []string, f []int, p []string) {
	size := len(f)
	for i := 0; i < size; i++ {
		//get two random participants
		i_one := rand.Intn(size)
		i_two := rand.Intn(size)

		if f[i_one] < f[i_two] {
			p[i] = c[i_one]
		} else {
			p[i] = c[i_two]
		}
	}
}

//breeds new candidate generation from parents
func breed(c []string, p []string, mf, cp, tl int) {
	for i := 0; i < len(c); i += 2 {
		//single point crossover
		c[i] = p[i][:cp] + p[i+1][cp:tl]
		c[i+1] = p[i+1][:cp] + p[i][cp:tl]

		//(terrible) scaling mutation rate adds diversity to break local optimum
		chance := 240
		if mf < 7 && mf >= 4 {
			chance /= 6
		} else if mf < 4 && mf > 0 {
			chance /= 12
		}

		//mutation possibilities
		//entire candidate is replaced by random string
		//one character in candidate is replaced by random character
		if rand.Intn(chance) == rand.Intn(chance) {
			c[i] = r_string(tl)
		}

		if rand.Intn(chance) == rand.Intn(chance) {
			c[i] = replace_at_index(c[i], rune(r_int_range(32, 126)), rand.Intn(tl))
		}
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	const population_size = 200
	generations := 1

	scanner := bufio.NewReader(os.Stdin)
	target, _ := scanner.ReadString('\n')
	target_len := len(target) - 1

	candidates := make([]string, population_size)
	fitnesses := make([]int, population_size)
	parents := make([]string, population_size)

	init_generation(candidates, fitnesses, target)
	min_fitness := lowest_fitness(fitnesses)

	for ok := true; ok; ok = (min_fitness != 0) {
		binary_tournament(candidates, fitnesses, parents)
		breed(candidates, parents, min_fitness, rand.Intn(target_len), target_len)
		calc_new_fitnesses(candidates, fitnesses, target)

		min_fitness = lowest_fitness(fitnesses)
		generations++
	}
	fmt.Println("Generation", generations, "success!")
}
