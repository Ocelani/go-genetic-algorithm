# Genetic algorithm

This repo is inspired by the article:
"A new approach to the software release planning", deColares et al.

## Metaheuristic Settings

The following text was retired from the article:

``
This subsection will discuss specific aspects related to problem and settings of the algorithm. Below we show these aspects and settings.

1. Solution Encoding: Each solution is encoded as binary strings. For each requirement there are |K| possibilities (releases). To represent these |K| possibilities in binary systems, it is necessary ⌈log2(|K|)⌉ bits. So, to represent the n = |R| requirements, it is necessary ⌈log2(|K|)⌉ ∗ n bits.

2. Initial Population: Initial population is the initial set of solutions from which new solutions will be continuously generated, until is reached a termination condition. In our approach, the initial population is obtained randomly.

3. Termination Condition: Termination condition is a condition that, when reached, stops the algorithm running. We used the number of generations, which is the number of iterations of the algorithm.

4. Algorithms Parameters: There are some parameters that should be set when using the NSGA-II algorithm. After some calibration tests, we defined the following values:
   • cross-over tax: 0.9 or 90%
   • mutation tax: 0.1 or 10%
   • population size (solutions per generation): 400
   • number of generations: 5000
   ``
