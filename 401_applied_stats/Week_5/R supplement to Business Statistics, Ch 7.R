### Business Statistics for Contemporary Decision Making (Black)
### Chapter 7, Sampling and Sampling Distributions

library(dplyr)

# Chapter 7 is largely conceptual, describing sampling methods, distributions
# and the Central Limit Theorem. As such, there is little for us to 'do' in
# R. However, there are sampling functions we should give some coverage to.

# Base R defines two (2) related functions:  sample() and sample.int().
# Each function, when run, looks to a 'hidden' object, .Random.seed.
# .Random.seed is a 626 element integer vector; the state of R's random number
# generator at the time. Calling a function that looks to this vector results
# in changes to it.

# sample() is used to take a sample of size 'size' from the elements of 'x' with
# or without replacement. The default weighting of the elements is uniformity;
# each element is equally likely to be sampled. However, if necessary, we can
# pass a vector of probability weights to the argument 'prob.'

# We will create a vector of 1000 values generated from a standard normal
# distribution. Then, we will take a sample of size ten (10) from that vector.

our_vector <- rnorm(n = 1000, mean = 0, sd = 1)
hist(our_vector, xlab = "Value", main = "")

(our_sample <- sample(x = our_vector, size = 10, replace = FALSE))

# We can use set.seed() when  we need to get repeatable behavior from a random
# or pseudo-random function.

set.seed(999)
(our_sample <- sample(x = our_vector, size = 10, replace = FALSE))

set.seed(999)
(our_sample <- sample(x = our_vector, size = 10, replace = FALSE))

set.seed(999)
(our_sample <- sample(x = our_vector, size = 10, replace = FALSE))

# The value used is effectively arbitrary.

# Now, often we'll use sample() to sample rows from a data frame or matrix. But,
# the sample() function, as we said, samples from the 'elements of 'x' ' and in
# R, the elements of a data frame are the columns. So,

data(iris)

set.seed(8888)
(iris_sample <- sample(x = 1:nrow(iris), size = 10, replace = FALSE))

# Then, use this numeric vector to get our sample rows:
iris[iris_sample, ]

# There is also a sample.int() 'version' of the function that can save some
# keystrokes in some cases. Rather than specifying an 'x' to sample from,
# sample.int() has us specify an 'n' to represent 'the number of items to
# choose from.' For example:

set.seed(8888)
iris_sample_2 <- sample.int(n = nrow(iris), size = 10, replace = FALSE)

iris[iris_sample_2, ]

# Lastly, the 'tidyverse' package defines a function sample_n() that will
# directly sample rows from a data frame (or 'tibble' in the tidyverse).

set.seed(8888)
sample_n(tbl = iris, size = 10, replace = FALSE)


# We can 
set.seed(777)

iris %>%
  group_by(., Species) %>%
  sample_n(., size = 10, replace = FALSE)

# versus

set.seed(777)

split_iris <- split(x = iris, f = iris$Species)
iris_samples <- lapply(X = split_iris, function(x) x[sample(x = 1:nrow(x),
                                                            size = 10,
                                                            replace = FALSE), ])
do.call(rbind, iris_samples)
