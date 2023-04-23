# Illustration of the differences in how quartiles are computed.

# Options available in the quantile() function of R.

# The three quartiles are supposed to divide the data into four equal parts.
# The question is how this should best be done. There are various methods.

# Example using Black page 55 - Chapter 3 with an odd number of observations. #n=15
# Page 55 of Black; top 15 trading partners of US. exports in billions of dollars
## Canada, Mexico, China, Japa, UK, ...Taiwan, Switzerland

x <- c(312,240,124,67,54,49,45,43,42,41,35,31,30,27,22)

p <- c(0.25, 0.5, 0.75)

length(x)
length(p)


sort(x)
sort(x, decreasing = FALSE)
sort(x, decreasing = TRUE)

?sort

# The summary() function uses the type = 7 option.
summary(x)

hist(x)

?quantile()

# Here are some results using each of the different options in R.

quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 1)
quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 2)
quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 3)
quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 4) ## exact interpolation 7.5th which is 42.5; 
## 3.75th 30.75 ##11.25th
quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 5)
quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 6)
quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 7) ## summary option: median mid value; 
## Q1 interpolate 31 & 35
quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 8)
quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 9)


boxplot(x)
boxplot.stats(x)
?boxplot.stats
?boxplot

## Note the whisker = 'nearest' data point from 1.5*IQR
## conf represents confidence interval for the median

boxplot(x, range = 1.5,  col = "red", main = "Boxplot with 1.5*IQR to identify outliers")
boxplot.stats(x, coef = 1.5, do.conf = TRUE, do.out = TRUE)

boxplot(x, range = 3.0,  col = "blue", main = "Boxplot with 3.0*IQR to identify extreme outliers") ## extreme outliers
boxplot.stats(x, coef = 3.0, do.conf = TRUE, do.out = TRUE)

## stop here ##############

#-------------------------------------------------------------------------------
# Boxplot uses type 2 if n is even and type 7 if n is odd as verified using
# boxplot.stats
#-------------------------------------------------------------------------------

x <- c(27 , 65 , 28 , 61,  34 , 91 , 61  ,37 , 58 , 31,
       43 , 47 , 44 , 20,  48 , 50 , 49 , 43 , 19,  52)

summary(x)

hist(x)

mean(x)
median(x)
mode(x)
sd(x)


library(moments)


?skewness(x) 
skewness(x) 

x <- c(1750.0, 2347.5, 3250.0, 3997.5, 3125.0, 2950.0, 4062.5, 5250.0, 2822.5, 3325.0, 5375.0, 3900.0)
summary(x)

quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 4)
quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 7)
quantile(x, probs = p, na.rm = FALSE, names = TRUE, type = 5)


