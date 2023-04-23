### Business Statistics for Contemporary Decision Making (Black)
### Chapter 10, Statistical Inferences About Two Populations


# Chapter 10 extends the learning we did in chapter 9 and introduces us to
# two-sample tests. In chapter 9, we looked at one-sample tests:  comparing
# a sample parameter to an established or population estimate. In chapter 10,
# we look at structuring tests to compare  two (2) samples and consider the
# probability that the samples are from the same or different populations.

# Again, we will specifically work the Demonstration Problems that give us
# 'raw' data to work with. Any of the problems that start with sample parameters
# can be approached 'manually,' implementing the formulae like we have for
# previous chapters. As we did with chapter 9, we'll focus on the problems for
# which there are R statistical test functions.

# Demonstration Problem 10.5 works through a paired, or "related measures"
# t-test. Here, our data are consumer ratings of a company before and after
# viewing video content.

# Demonstration Problem 10.5 (p. 341)
# Let us revisit  the hypothetical study discussed earlier in the section in
# which consumers are asked to rate a company both before and after viewing a
# video on the company twice a day for a week. That data from Table 10.4 are
# displayed again here. Use an alpha of 0.05 to test to determine whether there
# is a significant increase in the ratings of the company after the one-week
# video treatment. Assume that differences in ratings are normally distributed
# in the population.

before <- c(32, 11, 21, 17, 30, 38, 14)
after <- c(39, 15, 35, 13, 41, 39, 22)

t.test(x = after, y = before, paired = TRUE, alternative = "greater") # discuss 'direction'
# Paired t-test

# data:  after and before
# t = 2.5427, df = 6, p-value = 0.02196
# alternative hypothesis: true difference in means is greater than 0
# 95 percent confidence interval:
#   1.381023      Inf
# sample estimates:
#   mean of the differences 
#   5.857143

# We should note that our paired test is equivalent to testing whether the mean
# difference between the before and after vectors is greater than zero.
t.test(x = after - before, mu = 0, alternative = "greater")

# One Sample t-test

# data:  after - before
# t = 2.5427, df = 6, p-value = 0.02196
# alternative hypothesis: true mean is greater than 0
# 95 percent confidence interval:
#   1.381023      Inf
# sample estimates:
#   mean of x 
#   5.857143 


# We will next work Demonstration Problem 10.7 (p. 359) which asks us to compare
# two sample variances to evaluate the probability that our samples are drawn
# from the same population.

# According to Runzheimer International, a family of four in Manhattan with
# $60,000 annual income spends more than $22,000 a year on basic goods and 
# services. In contrast, a family of four in San Antonio with the same annual
# income spends only $15,460 on the same items. Suppose we want to test to
# determine whether the variance of money spent per year on the basics by
# families across the United States is greater than the variance of money spent
# on the basics by families of four in Manhattan - that is, whether the amounts
# spent by families of four in Manhattan are more homogeneous than the amounts
# spent by such families nationally. Suppose a random sample of eight Manhattan
# families produces the figures in the table, which are given along with those
# reported from a random sample of seven families across the United States.
# Complete a hypothesis-testing procedure to determine whether the variance
# of values taken from across the United States can be shown to be greater than
# the variance of values obtained from families in Manhattan. Let alpha = 0.01.
# Assume the amount spent on the basics is normally distributed in the
# population.

alpha <- 0.01

across_us <- c(18500, 19250, 16400, 20750, 17600, 21800, 14750)
manhattan <- c(23000, 21900, 22500, 21200, 21000, 22800, 23100, 21300)

# We can use the base R stats::var.test() function because we're comparing two
# samples from (one or more) normal populations.
var.test(x = across_us, y = manhattan, ratio = 1, alternative = "greater",
         conf.level = 1 - alpha)
# F test to compare two variances

# data:  across_us and manhattan
# F = 8.0872, num df = 6, denom df = 7, p-value = 0.007167
# alternative hypothesis: true ratio of variances is greater than 1
# 99 percent confidence interval:
#   1.124566      Inf
# sample estimates:
#   ratio of variances 
# 8.087209

# And, for ourselves:
(F <- var(across_us) / var(manhattan))
pf(q = F, df1 = length(across_us) - 1, df2 = length(manhattan) - 1,
   lower.tail = FALSE)  # [1] 0.007167222
