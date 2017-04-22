# GeneticWorld
A learning exercise in Go, first ever exposure to it. Uses a genetic 
algorithm to evolve a random string into the user's input.

Parent selection algorithm: [Binary
tournament](https://en.wikipedia.org/wiki/Tournament_selection)

Breeding algorithm: [Single point
crossover](https://en.wikipedia.org/wiki/Crossover_(genetic_algorithm))

[Here](https://www.reddit.com/r/dailyprogrammer/comments/40rs67/20160113_challenge_249_intermediate_hello_world/?st=ium6p9rn&sh=4cabed37) is the challenge source. 

## Algorithm outline
1. Generate random initial population
2. Calculate each fitness compared to user input and store the fittest
3. Randomly select (population size times) 2 candidates; fitter becomes a parent
4. Breed consecutive parents to produce new generation
5. Recalculate fitnesses
6. Repeat 3-5 until fittest is perfect
