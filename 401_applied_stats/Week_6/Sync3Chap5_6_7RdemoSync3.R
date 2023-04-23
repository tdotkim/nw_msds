

# Binomial - chap 5

#dbinom(x, size, prob, log = FALSE)
#pbinom(q, size, prob, lower.tail = TRUE, log.p = FALSE)
#qbinom(p, size, prob, lower.tail = TRUE, log.p = FALSE)
#rbinom(n, size, prob)
## 'd' density 'p' cumulative probability 'q' quantile 'r' random


### 10 coins are tossed. p = probability(H) = 0.5, x= Number of heads

n=10
p=0.5
x = 3

dbinom(x, n, p) ## 3 heads

pbinom(x, n, p, lower.tail = TRUE) ## <= 3 heads

dbinom(0, n, p) + dbinom(1, n, p) + dbinom(2, n, p) + dbinom(3, n, p)

pbinom(x, n, p, lower.tail = FALSE) ## > 3 heads

1 - (dbinom(0, n, p) + dbinom(1, n, p) + dbinom(2, n, p) + dbinom(3, n, p))

qbinom(0.171875, n, p, lower.tail = TRUE, log.p = FALSE)

qbinom(0.171875, n, p, lower.tail = FALSE, log.p = FALSE)

qbinom(0.83, n, p, lower.tail = TRUE, log.p = FALSE)

qbinom(0.85, n, p, lower.tail = TRUE, log.p = FALSE)

rbinom(20, n, p)

## Binomial Distribution

# Create a sample of 10 numbers which are incremented by 1.
x <- seq(0,10,by = 1)
# Create the binomial distribution.
y <- dbinom(x,10,0.5)
# Plot the graph for this sample.
plot(x,y)

df <- data.frame(x,y)
df$prd <- df$x*df$y
mean_x <- sum(df$prd)
mean_x



#########################################################################################

# Poisson - chap 5

#dpois(x, lambda, log = FALSE)
#ppois(q, lambda, lower.tail = TRUE, log.p = FALSE)
#qpois(p, lambda, lower.tail = TRUE, log.p = FALSE)
#rpois(n, lambda)

#If there are 10 cars crossing a bridge per minute on average, find the probability of having 
#10 or more cars crossing the bridge in a particular minute.

lambda <- 10  ## mean number of cars

dpois(10, lambda, log = FALSE)

ppois(10, lambda, lower.tail = TRUE, log.p = FALSE)

ppois(10, lambda, lower.tail = FALSE, log.p = FALSE)

ppois(9, lambda, lower.tail = FALSE, log.p = FALSE)

ppois(9, lambda, lower.tail = TRUE, log.p = FALSE)

qpois(0.457, lambda, lower.tail = TRUE, log.p = FALSE)

qpois(0.99999999, lambda, lower.tail = TRUE, log.p = FALSE)

rpois(20, lambda)

############################################################################################

# hypergrometric - chap 5

#dhyper(x, m, n, k, log = FALSE)
#phyper(q, m, n, k, lower.tail = TRUE, log.p = FALSE)
#qhyper(p, m, n, k, lower.tail = TRUE, log.p = FALSE)
#rhyper(nn, m, n, k)
#x,q, and m are white balls, n black balls, m+n total balls, k sample size

#A small voting district has 101 female voters and 95 male voters. A random sample of 10 voters is drawn. 
#What is the probability exactly 7 of the voters will be female?

x = 7                 ##  101 (m)  &  M: 95 (n)
m = 101               ## sample of  10  (k)
n = 95                ##  7 (x)  & M: 3  
k =10

dhyper(x, m, n, k, log = FALSE)

#################################################################################################

# Normal - chap 6

#dnorm(x, mean = 0, sd = 1, log = FALSE)
#pnorm(q, mean = 0, sd = 1, lower.tail = TRUE, log.p = FALSE)
#qnorm(p, mean = 0, sd = 1, lower.tail = TRUE, log.p = FALSE)
#rnorm(n, mean = 0, sd = 1)

# avg monthly cell phone bill has mean =$100, sd=$10

mean = 100
sd = 10

dnorm(100, mean = mean, sd = sd, log = FALSE)

set.seed(3000)
xseq<-seq(60,140,1)
densities<-dnorm(xseq, 100,10)
plot(xseq, densities, col="darkgreen",xlab="", ylab="Density", type="l",lwd=2, cex=2, 
     main="Probability Density Function of Normal", cex.axis=.8)

pnorm(100, mean = mean, sd = sd, lower.tail = TRUE, log.p = FALSE)
abline(v = 100, col = "black", lwd = 2)
pnorm(90, mean = mean, sd = sd, lower.tail = TRUE, log.p = FALSE)
abline(v = 90, col = "green", lwd = 2)
pnorm(80, mean = mean, sd = sd, lower.tail = TRUE, log.p = FALSE)
abline(v = 80, col = "blue", lwd = 2)
pnorm(70, mean = mean, sd = sd, lower.tail = TRUE, log.p = FALSE)
abline(v = 70, col = "red", lwd = 2)

pnorm(90, mean = mean, sd = sd, lower.tail = TRUE, log.p = FALSE)
abline(v = 110, col = "green", lwd = 2)
pnorm(80, mean = mean, sd = sd, lower.tail = TRUE, log.p = FALSE)
abline(v = 120, col = "blue", lwd = 2)
pnorm(70, mean = mean, sd = sd, lower.tail = TRUE, log.p = FALSE)
abline(v = 130, col = "red", lwd = 2)

qnorm(0.95, mean = mean, sd = sd, lower.tail = TRUE, log.p = FALSE)
abline(v = 116.4, col = "cyan", lwd = 2, lty =2)
qnorm(0.975, mean = mean, sd = sd, lower.tail = TRUE, log.p = FALSE)
abline(v = 119.6, col = "cyan", lwd = 2, lty =2)
qnorm(0.99, mean = mean, sd = sd, lower.tail = TRUE, log.p = FALSE)
abline(v = 123.3, col = "cyan", lwd = 2, lty =2)



#############################################

#qnorm(0.99, mean = 0, sd = 1, lower.tail = TRUE, log.p = FALSE)
#pnorm(-1.63, mean = 0, sd = 1, lower.tail = TRUE, log.p = FALSE)

## 6 sigma program shift of 1.5 st dev
pnorm(-4.5, mean = 0, sd = 1, lower.tail = TRUE, log.p = FALSE) ## 6 sigma means 3.4 ppm


#qchisq(0.975,df=19)
#qchisq(0.025,df=19)

#qt(0.9,27, lower.tail = TRUE)
##############################################################################################

# exponential - chap 6

#dexp(x, rate = 1, log = FALSE)
#pexp(q, rate = 1, lower.tail = TRUE, log.p = FALSE)
#qexp(p, rate = 1, lower.tail = TRUE, log.p = FALSE)
#rexp(n, rate = 1)

#Suppose the mean checkout time of a supermarket cashier is 3 minutes. Find the probability of a 
#customer checkout being completed by the cashier in less than 2 minutes.

pexp(2, rate=1/3) 



#######################################################################################

# sampling distribution for the mean & p - chap 7

# In a given hour, avg # of shoppers = 448 & sd = 21; sample of size 49 taken & sample mean computed.
# What is the probability that the sample mean < 446

z <- (446-448)/(21/7)
pnorm(z, 0, 1, lower.tail = TRUE)

## same question but NOT mean; What is the probability of 446 or less shoppers?

z <- (446-448)/(21)
pnorm(z, 0, 1, lower.tail = TRUE)


# 10% defective population. Select 80 parts. What is the probability of 12 or more defective? - chap 7

p = 0.1
n = 80
phat <- 12/80 ## 0.15

sd = sqrt((p*(1-p)/n))
z <- (phat - p)/sd

pnorm(z, mean = 0, sd = 1, lower.tail = FALSE, log.p = FALSE)
#######################################################################

###################### stop here #######################################









































pnorm(1, mean=0, sd=1, lower.tail = FALSE)