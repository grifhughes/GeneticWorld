#GeneticWorld
A learning exercise in Go. Uses a genetic algorithm to evolve a 
random string into the user's input. 

##Algorithm outline
1. Generate random initial population
2. Calculate each fitness compared to user input and store the fittest
3. Randomly select (population size times) 2 candidates; fitter becomes a parent
4. Breen consecutive parents to produce new generation
5. Recalculate fitnesses
6. Repeat 3-5 until fittest is perfect
