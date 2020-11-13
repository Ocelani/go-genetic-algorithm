# Genetic algorithm

This repo is inspired by the article:
"A new approach to the software release planning", deColares et al.

The following texts was retired from the article.

## Metaheuristic Settings

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

### Variables

- Requirements: There are several requirements to be implemented, which may be allocated to a specific release or may even not be allocated, that is, not implemented. Associated to each requirement there is a risk value. For the sake of simplicity the risk levels may vary from 0 (lowest) up to 5 (highest).

- Stakeholders: There are multiple stakeholders with different needs and possibly conflicting interests. These stakeholders have different importance levels to the company. These importance levels vary from 1 (lowest) up to 5 (highest).

- Releases: All releases that will be developed. Handles all those requirements that due to the restrictions of resources will be left for some other following releases

- Resources: Are all sort of goods available for the completion of the project activities. Naturally, release planning must respect the availability of resources.

#### Relations

- Stakeholders vs. Requirements: Associated to every requirement there are importance (priority) values, which measure its importance to each stakeholders.

- Resources vs. Releases vs. Requirements: To implement each requirement there is a cost related to each resource.

- Risks vs. Releases: The requirements with higher risk level are supposed to be implemented in earlier releases.

### The problem

It was considered the following problem instance: 19 requirements that should be implemented in 5 releases; 5 stakeholders, each one with a weight (importance) and his/her requirements priorities; 3 different resources were considered, where each release had a limited amount of them, and the total amount was less than the effort needed to implement all the 19 requirements.

Given the aspects above mentioned, the problem of soft-ware release planning can be described as follows:

##### We want to both:

– maximize stakeholders satisfaction and

– minimize project risks.

##### Respecting:

– the available resources and

– the requirements interdependencies.
