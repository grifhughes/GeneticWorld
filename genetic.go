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
func init_generation(t string, population_size int) (c []string, f []int, m int) {
	c = make([]string, population_size)
	f = make([]int, population_size)
	for i := 0; i < population_size; i++ {
		c[i] = r_string(len(t) - 1)
		f[i] = test_fitness(t, c[i])
	}
	m = lowest_fitness(f)
	return c, f, m
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
		}
	}
	return min
}

//picks 2 candidates at random and selects the fitter to become a parent
func binary_tournament(c []string, f []int, population_size int) (p []string) {
	p = make([]string, population_size)
	for i := 0; i < population_size; i++ {
		//get two random participants
		i_one := rand.Intn(population_size)
		i_two := rand.Intn(population_size)

		if f[i_one] < f[i_two] {
			p[i] = c[i_one]
		} else {
			p[i] = c[i_two]
		}
	}
	return p
}

//breeds new candidate generation from parents
func breed(p []string, target string, mf, tl int) (nc []string, nf []int, nmf int) {
	//(terrible) scaling mutation rate adds diversity to break local optimum
	chance := 240
	if tl > 10 && tl < 30 {
		chance /= 3
	} else if tl >= 30 {
		chance /= 6
	}

	nc = make([]string, len(p))
	nf = make([]int, len(p))
	crossover_point := rand.Intn(tl)

	for i := 0; i < len(p); i += 2 {
		//single point crossover
		nc[i] = p[i][:crossover_point] + p[i+1][crossover_point:tl]
		nc[i+1] = p[i+1][:crossover_point] + p[i][crossover_point:tl]

		nf[i] = test_fitness(nc[i], target)
		nf[i+1] = test_fitness(nc[i+1], target)

		//mutation possibilities
		//entire candidate is replaced by random string
		//one character in candidate is replaced by random character
		if rand.Intn(chance) == rand.Intn(chance) {
			nc[i] = r_string(tl)
		}

		if rand.Intn(chance) == rand.Intn(chance) {
			nc[i] = replace_at_index(nc[i], rune(r_int_range(32, 126)), rand.Intn(tl))
		}
	}
	nmf = lowest_fitness(nf)
	return nc, nf, nmf
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	const population_size = 200
	generations := 1
	scanner := bufio.NewReader(os.Stdin)
	target, _ := scanner.ReadString('\n')
	target_len := len(target) - 1

	candidates, fitnesses, min_fitness := init_generation(target, population_size)

	for ok := true; ok; ok = (min_fitness != 0) {
		parents := binary_tournament(candidates, fitnesses, population_size)
		candidates, fitnesses, min_fitness = breed(parents, target, min_fitness, target_len)
		generations++
	}
	fmt.Println("Generation", generations, "success!")
}
