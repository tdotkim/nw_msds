### Business Statistics for Contemporary Decision Making (Black)
### Chapter 4, Probability

library(gtools)

# "The main objective of Chapter 4 is to help you understand the basic
# principles of probability." Like Chapter 1, Chapter 4 is largely
# conceptual, but we will introduce and explore some of the R functions that
# may prove useful in approaching probability problems.

# 4.2 discusses a number of concepts using dice rolls. We'll do the same.

# We may, at some point, need to sample from a data object. A base R sample()
# function exist for this purpose.

?sample

# Now, how do we represent our die rolls? Each roll is an event. We will 
# simulate ten (10) rolls using sample().
sample(x = 1:6, size = 10, replace = TRUE, prob = NULL)

# In the above, we're sampling from the 'sample space' of a die. Any given
# standard die roll must have one of six outcomes: 1, 2, 3, 4, 5 or 6.

# So, we'll sample from the integer sequence 1 through 6. We'll sample from
# that sequence ten (10) times, with replacement. And, there is an equal
# probability of each outcome.

sample(x = 1:6, size = 10, replace = TRUE, prob = NULL)
sample(x = 1:6, size = 10, replace = TRUE, prob = NULL)
sample(x = 1:6, size = 10, replace = TRUE, prob = NULL)

set.seed(1234)
sample(x = 1:6, size = 10, replace = TRUE, prob = NULL)
set.seed(1234)
sample(x = 1:6, size = 10, replace = TRUE, prob = NULL)
set.seed(1234)
sample(x = 1:6, size = 10, replace = TRUE, prob = NULL)

# Obviously, rolling one (1) die once will have one (1) of six (6) outcomes,
# each with a probability of:
1 / 6  # [1] 0.1666667

# So, what if we add a second die? What is the probability of rolling two (2)
# sixes?
(1 / 6) * (1 / 6)  # [1] 0.02777778

# Equivalently, we may've realized that two (2) sixes is one (1) possible
# outcome, that there is only one (1) way it can be achieved and that there are
# 6^2 possible outcomes when rolling two (2) dice.

1 / 6^2  # [1] 0.02777778

# For some problems, we may well be able to think out the universe of outcomes,
# the outcomes of interest and 'manually' calculate the probability. However,
# in many cases, the number of total combinations or the number of 'successful'
# outcomes satisfying some criterion/criteria. For this,
(our_grid <- expand.grid(x = rep(list(1:6), 2)))

# our_grid has our thirty-six (36) possible outcomes; our 'universe' of
# outcomes.

# Now, how many of those outcomes are two (2) sixes?
sum(apply(X = our_grid, MARGIN = 1, function(x) sum(x) == 12)) / nrow(our_grid)

# Now, what proportion of rolls have a sum of less than 6? We'll use our_grid,
# again, but specify a different function:

sum(apply(X = our_grid, MARGIN = 1, function(x) sum(x) < 6)) / 
  nrow(our_grid)

# Other packages exist that can help us explore permutations and combinations.
# For example,
gtools::permutations(n = 6, r = 2, v = 1:6, repeats.allowed = TRUE)


# R provides functions for identifying the 'union' and 'intersection' of two (2)
# vectors (Black, p. 94).
our_first_vector <- c(1, 3, 5, 2, 4, 6)
our_second_vector <- c(2, 4, 6, 8, 10)

# What is the union of these two (2) vectors? i.e. what elements are present
# in either or both vectors?
union(x = our_first_vector, y = our_second_vector)

# What is the intersection of these two (2) vectors? i.e. what elements are 
# present in both?
intersect(x = our_first_vector, y = our_second_vector)


# In our discussion of dice rolling, we were sampling 'with replacement.' We
# wanted to simulate the outcomes of ten (10) dice rolls and rolling a one (1)
# on the first roll or with the first die doesn't preclude us from rolling 
# another one (1). However, there are many situations where we would need
# to approach a problem or sampling task 'without replacement,' where each
# draw or sample does meaningfully remove an outcome from any draws to follow.

# For an example, we'll look at the 'small law firm' example; 4.2, p. 97.

# "Suppose a small law firm has 16 employees and 3 are to be selected randomly
# to represent the company at the annual meeting of the American Bar
# Association. How many different combinations of lawyers could be sent to the
# meeting? This situation does not allow sampling with replacement because
# three different lawyers will be selected to go. This problem is solved by
# using combinations. N = 16 and n = 3, so..."

factorial(16) / (factorial(3) * factorial(16 - 3))  # [1] 560

# OR,
choose(16, 3)  # [1] 560


# 4.3, 4.4, 4.5, 4.6 and 4.7 describe marginal, union, join and conditional
# probabilities and the laws and rules for calculating them. For problems like
# those explored in these sections, you'll ultimately be using R as a calculator
# and 'manually' implementing the appropriate formulae. We can, if we like,
# create functions from | for these implementations.
