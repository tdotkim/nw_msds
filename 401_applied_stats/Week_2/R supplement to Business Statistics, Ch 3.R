### Business Statistics for Contemporary Decision Making (Black)
### Chapter 3, Descriptive Statistics

library(gds)
library(ggplot2)
library(ie2misc)
library(moments)

# "The focus of Chapter 3 is the use of statistical techniques to describe
# data." The chapter introduces the basic population and sample parameters
# useful in describing data.

# 3.1 is concerned with Measures of Central Tendency; ways of describing data
# according to the value or values around which data is centered.

# The methods described in 3.1 are the mode, median, mean, percentiles and
# quartiles. We'll use the Shopping Centre and size data from Demonstration
# Problem 3.1 for exploring these in R.

# First, we'll create a data frame with two (2) columns:  Shopping Centre and
# Size.

(shop_centre <- data.frame(Shopping_Centre = c("MetroCentre", "Trafford Centre",
                                             "Westfield Stratford Centre",
                                             "Bluewater", "Liverpool One",
                                             "Westfield London",
                                             "Intu Merry Hill",
                                             "Manchester Arndale", "Meadowhall",
                                             "Lakeside", "St. David's",
                                             "Bullring", "Eldon Square"),
                         Size = c(190, 180.9, 175, 155.7, 154, 149.5, 140.8,
                                  139.4, 139.4, 133.8, 130.1, 127.1, 125.4)))

# The first measure introduced is the mode, "the most frequently occurring value
# in a set of data." There is not a built-in function for determining the mode
# of a dataset. Please do not be lead astray by the mode() function which is
# used to get or set the storage mode (think:  "type") of an R object.

mode(shop_centre)  # [1] "list"

# However, we readily find the mode by (1) identifying the counts of values, and
# (2) identifying the largest count.

# We can 'eyeball' our data frame and recognize that 139.4 is the mode. But, we
# can do a bit better:
table(shop_centre$Size)[which(table(shop_centre$Size) ==
                                max(table(shop_centre$Size)))]

sort(x = table(shop_centre$Size), decreasing = TRUE)

# Just to make sure our approach would work for character data, too.
set.seed(9999)
our_character_vector <- sample(x = LETTERS, size = 30, replace = TRUE)

sort(x = table(our_character_vector), decreasing = TRUE)

table(our_character_vector)[which.max(table(our_character_vector))]
# versus
table(our_character_vector)[which(table(our_character_vector) ==
                                    max(table(our_character_vector)))]

# For the mean and median, however, we do have built-in R functions.
?mean
?median

mean(x = shop_centre$Size, trim = 0, na.rm = TRUE)  # trim and na.rm args

# But, for fun, let's not trust mean(). Instead, we'll sum our Size vector and
# divide by its length.

sum(x = shop_centre$Size) / length(shop_centre$Size)

# For our example use, we'll remove 10% of observations from each 'end.'
mean(x = shop_centre$Size, trim = 0.1)

# And, to implement manually,
length(shop_centre$Size) * 0.1  # [1] 1.3
trim_vect <- sort(shop_centre$Size,
                  decreasing = FALSE)[2:(length(shop_centre$Size) - 1)]
sum(trim_vect) / length(trim_vect)


# The median() function takes the same 'na.rm' argument as mean().
median(x = shop_centre$Size, na.rm = TRUE)
median.default

# Percentiles consider our data in 100 equal parts. Each percentile is the value
# at which at least n percent of the data is below it. Percentiles will be
# either a value in your dataset or a value the 'makes sense' for your data.

# Quartiles are similar, but consider your data in four (4) equal parts. The
# 1st quartile is equal to the 25th percentile.

# Both percentiles and quartiles are 'quantiles'; specifying a number of cuts
# or breaks in a numeric vector and allowing us to identify values at which
# some quantity of our data are below (or above). Percentiles consider our
# data in 100 parts; 99 cuts. Quartiles in four (4) parts; three (3) cuts.
# Quantiles in x parts; x - 1 cuts.

quantile(x = shop_centre$Size, probs = 0.5, type = 7)

quantile(x = shop_centre$Size, probs = c(0.25, 0.5, 0.75), type = 7)
boxplot(x = shop_centre$Size, horizontal = TRUE,
        xlab = expression(Size ~ (1000 ~ m^2)),
        ylim = c(100, max(shop_centre$Size) + 10))

quantile(x = shop_centre$Size, probs = c(0.1, 0.80, 0.99), type = 7)


# 3.2 is concerned with Measures of Variability; ways of describing "the spread
# or dispersion of a set of data."

# The methods described in 3.2 are the range, the interquartile range, mean
# absolute deviation, variance, and standard deviation. We'll use the top 15
# trading partners of the United States in exports, 2014, data on p. 55.

top_15 <- data.frame(county = c("Canada", "Mexico", "China", "Japan",
                                "United Kingdom", "Germany", "South Korea",
                                "Netherlands", "Brazil", "Hong Kong", "Belgium",
                                "France", "Singapore", "Taiwan", "Switzerland"),
                     exports = c(312.4, 240.2, 123.7, 66.8, 53.8, 49.4, 44.5,
                                 43.1, 42.4, 40.9, 34.8, 31.3, 30.2, 26.7,
                                 22.2))

range(top_15$exports)  # alternatively, top_15[, "exports"] OR top_15[, 2]
# hmmmmm, okay. how 'bout:
diff(range(top_15$exports))  # alternatively,
max(top_15$exports) - min(top_15$exports)

# Interquartile range is the distance between the 1st and 3rd quartile. There is
# built-in function, IQR(), for returning the interquartile range.

IQR(x = top_15$exports)  # alternatively,
as.numeric(quantile(x = top_15$exports, probs = 0.75) -
  quantile(x = top_15$exports, probs = 0.25))

# There isn't a base R function for calculating the mean absolute deviation
# (MAD), but there are packages that define such a function. We will calculate
# it for ourselves for our export data, then test the "ie2misc" function.

mad_top_15 <- sum(abs(top_15$exports - mean(top_15$exports))) / nrow(top_15)
mad_top_15

madstat(observed = top_15$exports)

# Please note that there is a base R mad() function, but it returns the MEDIAN
# absolute deviation, not the MEAN.

var(top_15$exports)

sum((top_15$exports - mean(top_15$exports))^2) / nrow(top_15)

# Wait, what?
# The var() and sd() functions return the SAMPLE variance and standard
# deviation. These use n - 1, instead of N, as the divisor.

sum((top_15$exports - mean(top_15$exports))^2) / (nrow(top_15) - 1)

# And,
sd(top_15$exports)

sqrt(sum((top_15$exports - mean(top_15$exports))^2) / (nrow(top_15) - 1))

# 3.2 also discusses z-scores for the purpose of representing the distance
# from the mean in units of standard deviation for normallly distributed data.

# First, we'll manually calculate the z-scores for our export data. Then, we'll
# use the base scale() function.

(our_z <- (top_15$exports - mean(top_15$exports)) / sd(top_15$exports))
# versus
as.vector( scale(x = top_15$exports) )

# The last statistic introduced in 3.2 is the coefficient of variation; "the
# ratio of the standard deviation to the mean expressed in percentage." There
# is not an existing base R function, but we could (1) easily calculate the
# sample CV using the sd() and mean() functions, and (2) write our own function
# for later use | reuse.

sd(top_15$exports) / mean(top_15$exports) * 100

our_CV_function <- function(x) {
  return(list(CV = sd(x) / mean(x) * 100))
}

our_CV_function(x = top_15$exports)


# 3.3 discusses many of these same measures of central tendency and variability,
# but for grouped data. There are not base R functions for grouped data. You'll
# find in R that, as you might expect, many functions - including statistical
# test methods - are written expecting raw data. So, you may have functions
# to help us create grouped or grouped-by data, e.g. cut(), but not have
# functions that are written "expecting" grouped data.

# However, to not let the section go unaddressed here, there is a package 'gds'
# that defines a single function, gds(), that returns summary statistics for
# grouped data. I tested it on the grouped Canadian unemployment
# data (Black, p. 70) and the statistics output matched those in Black.

# https://cran.r-project.org/web/packages/gds/gds.pdf

gds(ll = c(1, 3, 5, 7, 9, 11), ul = c(3, 5, 7, 9, 11, 13),
    freq = c(4, 12, 13, 19, 7, 5))


# 3.4 discusses Measures of Shape, "tools that can be used to describe the
# shape of a distribution of data."

# Measures discussed in skewness, kurtosis, boxplot-and-whisker plots and
# five-number summary.

# We will use the 'moments' package skewness() and kurtosis() functions.
?skewness
?kurtosis

# First, we'll create density plot of our shopping centre data. Then, we'll
# calculate and discuss skewness and kurtosis.

ggplot(data = top_15, aes(x = exports)) + geom_density()

ggplot(data = shop_centre, aes(x = Size)) +
  geom_density(fill = "cornflowerblue", alpha = 0.75) +
  xlim(75, 225) + geom_vline(xintercept = mean(shop_centre$Size),
                             col = "red4") +
  geom_vline(xintercept = median(shop_centre$Size), col = "gray30",
             linetype = "dashed")

skewness(x = shop_centre$Size)  # think:  symmetry
kurtosis(x = shop_centre$Size)  # think:  peakedness; discuss 'excess'

# Boxplots are very common for evaluating the distribution of data, along with
# histograms and density curves. The boxplot includes a 'box' that shows the 
# first quartile, median and third quartile, and often includes 'whiskers' that
# (usually) extend upward to 1.5 * the interquartile range above the 3rd
# quartile and downward to 1.5 * the interquartile range less the 1st quartile.

ggplot(data = shop_centre, aes(y = Size)) + geom_boxplot() +
  theme(axis.line.x = element_blank(),
        axis.text.x = element_blank(),
        axis.ticks.x = element_blank())

# Lastly, we have a few base R options for returning five (5) or six (6) number
# summaries.

fivenum(x = shop_centre$Size)  # min, lower-hinge, median, upper-hinge, max
summary(object = shop_centre$Size)  # min, Q1, median, mean, Q3, max
boxplot.stats(x = shop_centre$Size)$stats  # lower-whisk, lower-hinge, median,
                                           # upper-hinge, upper-whisk
