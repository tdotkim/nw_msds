### Business Statistics for Contemporary Decision Making (Black)
### Chapter 12, Simple Regression Analysis and Correlation

library(ggplot2)

# Chapter 12 introduces concepts and methods that help us explore bivariate
# relationships; i.e. how two (2) variables may or may not be related.
# Specifically, correlation and simple linear regression are considered as well
# as methods for considering model fit and specification.

# 12.1 defines correlation as "a measure of the degree of relatedness of
# variables." The first measure introduced is Pearson's product-moment
# correlation coefficient. We'll look at the Interest Rate and Futures Index
# data from Table 12.1 on page 426.

interest_rate <- c(7.43, 7.48, 8, 7.75, 7.6, 7.63, 7.68, 7.67, 7.59, 8.07,
                   8.03, 8)
futures <- c(221, 222, 226, 225, 224, 223, 223, 226, 226, 235, 233, 241)

# We will first look to the R function, cor.test().

cor.test(x = interest_rate, y = futures, alternative = "two.sided",
         method = "pearson")

# Pearson's product-moment correlation

# data:  interest_rate and futures
# t = 4.4518, df = 10, p-value = 0.001232
# alternative hypothesis: true correlation is not equal to 0
# 95 percent confidence interval:
#   0.4535630 0.9463714
# sample estimates:
#       cor 
# 0.8152537

# There is a second function, cor(), that does the same, but instead of providing
# the detailed statistical test ouput that cor.test() does, it outputs a one-
# element numeric vector with only the correlation coefficient. This can be
# very useful when we need to calculate a correlation coefficient and pass it
# to another operation | calculation.

cor(x = interest_rate, y = futures)  # [1] 0.8152537


# 12.2 introduces simple regression analysis. We will start with the Airline
# Cost Data from Table 12.3 (p. 429).

no_of_pass <- c(61, 63, 67, 69, 70, 74, 76, 81, 86, 91, 95, 97)
cost_1000 <- c(4.28, 4.08, 4.42, 4.17, 4.48, 4.3, 4.82, 4.7, 5.11, 5.13, 5.64,
               5.56)

# First, we'll create the scatterplot on page 430.
ggplot(data = NULL, aes(x = no_of_pass, y = cost_1000)) +
  geom_point(col = "red2") +
  labs(x = "Number of Passengers", y = "Cost ($1,000)") +
  theme_minimal()

# 12.3 provides us the math for fitting our regression line. In R, we can use
# the lm() function.

our_model <- lm(cost_1000 ~ no_of_pass)
summary(our_model)

# y = 1.57 + 0.0407x

# We can also easily add our regression fit to our existing scatterplot:
ggplot(data = NULL, aes(x = no_of_pass, y = cost_1000)) +
  geom_point(col = "red2") +
  geom_smooth(method = "lm", se = FALSE) +  # Huzzah!
  labs(x = "Number of Passengers", y = "Cost ($1,000)") +
  theme_minimal()

# Now, we have our model. Next step is to determine whether it's a 'good' model.
# We do that, primarily, by analyzing the residuals; i.e. the errors, the
# difference between each fitted value and observed value.

# Fortunately, our fitted model object - "our_model" - is a list of thirteen (13)
# elements, two (2) of which are vectors that include the predicted (i.e. fitted)
# values and the residuals.

str(our_model)

# We can get the residuals via residuals() or via our_model$residuals. Pages
# 438 to 441 give us examples of the plots normally used perform residual
# analysis. The base R plot() function will actually accept a fitted model
# object and return these figures for us.

plot(our_model)

# Also, just for fun, let's recreate the scatterplot-w/-residuals from the top
# of page 438.

ggplot(data = NULL, aes(x = no_of_pass, y = cost_1000)) +
  geom_point(col = "red2") +
  geom_smooth(method = "lm", se = FALSE) +
  geom_segment(aes(x = no_of_pass, y = cost_1000, xend = no_of_pass,
                   yend = fitted(our_model))) +
  labs(x = "Number of Passengers", y = "Cost ($1,000)") +
  theme_minimal()


# 12.5 The standard error of the estimate is our estimate of a standard deviation
# our residuals; i.e. a way to consider the accuracy of our predictions/fitted
# values.
(sse <- sum(our_model$residuals^2))  # [1] 0.3140603

# In our summary(lm(...)) output, the standard error is referred to as the
# residual standard error
(rse <- sqrt(sse / (length(our_model$residuals) - 2)))  # [1] 0.1772175
 

# 12.6 introduces the coefficient of determination or R^2. This a common metric
# for understanding model-fit. It is, in short, "the proportion of variability
# of the dependent variable (y) accounted for or explained by the independent
# variable (x).

# Again, a summary() of our our_model object returns the coefficient of
# determination.

summary(our_model)$r.squared  # [1] 0.8990839

# for discussion,
str(our_model)
str(summary(our_model))

# 12.7 walks through conducting a t-test on our slope. Here, our null hypothesis
# is that slope is equal to zero; i.e. that our coefficient estimate is not
# statistically significantly not zero.

# If we wanted, we perform the t-test ourselves:
sum_x_sq <- sum(no_of_pass^2)
sum_x <- sum(no_of_pass)
n <- length(no_of_pass)

(our_t <- (0.040702 - 0) / (rse / sqrt(sum_x_sq - (sum_x^2 / n))))  # [1] 9.438958
# And,
pt(q = our_t, df = n - 2, lower.tail = FALSE) * 2

# Or, if we'd rather:
coef(summary(our_model))["no_of_pass", c("t value", "Pr(>|t|)")]

# 12.7 also introduces testing the overall significance of the model. This is
# an F-test to identify whether at least one of the regression coefficients is
# not zero. We can do the calculation ourselves:

anova(our_model)

(F <- (2.79803 / 1) / (0.31406 / 10))  # [1] 89.09221
pf(q = F, df1 = 1, df2 = 10, lower.tail = FALSE)  # [1] 2.691639e-06

# Or,
(f_stat <- summary(our_model)$fstatistic)
pf(q = f_stat[1], df1 = f_stat[2], df2 = f_stat[3], lower.tail = FALSE)  # 2.691644e-06

# 12.8 goes over estimation | prediction using our linear model fit. R makes
# this relatively straightforward via the predict() function. On page 453, Black
# uses the cost_in_1000s ~ no_of_passengers model fit to estimate the cost if
# the number of passengers is 73.
?predict

predict(object = our_model, newdata = data.frame(no_of_pass = 73))  # 4.541009

# predict() can also provide confidence intervals:
predict(object = our_model, newdata = data.frame(no_of_pass = 73),
        interval = "confidence")
#        fit      lwr      upr
# 1 4.541009 4.419097 4.662922

# Or prediction intervals (p. 454):
predict(object = our_model, newdata = data.frame(no_of_pass = 73),
        interval = "prediction")
#        fit      lwr      upr
# 1 4.541009 4.127753 4.954266


# 12.9 describes using regression to develop a forecasting trend line; i.e.
# using lm() to create forecast time-series data. We'll work Demonstration
# Problem 12.7.

month <- factor(x = c("January", "February", "March", "April", "May", "June",
                      "July", "August"), levels = month.name)
sales <- c(32569, 32274, 32583, 32304, 32149, 32077, 31989, 31977)

ggplot(data = NULL, aes(x = month, y = sales)) + geom_point()

ts_regression <- lm(sales ~ as.numeric(month))
summary(ts_regression)

# And, to our October prediction:
predict(object = ts_regression, newdata = data.frame(month = 10))  # 31766.07


# And, it turns out, we've been living 12.10 all along.
