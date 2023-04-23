 

setwd("C:/NorthWestern/LectureMaterialSummer2018MSDS401/RProjects")

require(ggplot2)  
require(lattice)
library(car)

#################### Chapter 1 demonstrate variable types, sample stats #############

mydata <- read.csv("ConsumerData.csv",head=TRUE, sep = ",", stringsAsFactors = TRUE)

str(mydata)

mean(mydata$Region)

summary(mydata)

mydata$RegionLab <- ifelse(mydata$Region  == 1,"NE",
                      ifelse(mydata$Region  == 2,"MW",
                             ifelse(mydata$Region  == 3,"S",
                                    "W")))

mydata$LocationLab <- ifelse(mydata$Location  == 1,"Metro", "Outside")

str(mydata)

summary(mydata)

mean(mydata$RegionLab)

newdata <- mydata[complete.cases(mydata), ]

newdata$Region <- NULL
newdata$Location <- NULL

str(newdata)

names(newdata)

summary(newdata)

mean(newdata$AnnualHouseholdIncome)  ## sample stats
sd(newdata$AnnualHouseholdIncome)    ## sample stats

################### Chapter 2  #######################

hist(newdata$AnnualHouseholdIncome, main="Distribution of AnnualHouseholdIncome")

ggplot(newdata, aes(x=AnnualHouseholdIncome)) + 
  geom_histogram(color="blue", fill = "blue", binwidth= 10000) +
  labs(title="Distribution of AnnualHouseholdIncome") +
  theme(plot.title=element_text(lineheight=0.8, face="bold", hjust=0.5))

counts <- table(newdata$RegionLab)
counts
barplot(counts, main="Number of samples by Region")

counts <- table(newdata$LocationLab)
counts
barplot(counts, main="Number of samples by Location")

counts <- table(newdata$RegionLab, newdata$LocationLab)
counts
barplot(counts, main="Number of samples by Region & Location")

library(plyr)

dfsummary <- ddply(newdata, .(LocationLab,RegionLab), summarize, MeanIncome=mean(AnnualHouseholdIncome))

dfsummary

pairs(newdata[1:3], pch = 21)

scatterplotMatrix(~AnnualFoodSpending+AnnualHouseholdIncome+NonMortgageDebt|LocationLab, 
                   data=newdata, main="Pairwise scatter plots")

scatterplotMatrix(~AnnualFoodSpending+AnnualHouseholdIncome+NonMortgageDebt|RegionLab, 
                   data=newdata, main="Pairwise scatter plots")



### using ggplot ####

library(ggplot2)
library(GGally)
ggpairs(newdata, columns=1:3, aes(color=RegionLab)) + 
  ggtitle("Pairwise scatter plots")

library(lattice)
splom(newdata[1:3], 
      groups=newdata$RegionLab, 
      main="Pairwise scatter plots")

library(ggplot2)
library(GGally)
ggpairs(newdata, columns=1:3, aes(color=LocationLab)) + 
  ggtitle("Pairwise scatter plots")

library(lattice)
splom(newdata[1:3], 
      groups=newdata$LocationLab, 
      main="Pairwise scatter plots")
################################################################################

##############  Chapter 3 ###################################################

summary(newdata)

library(pastecs)
stat.desc(newdata[1:3])

library(psych)
describe(newdata[1:3])
?describe

library(psych)
describeBy(newdata[1:3], newdata$LocationLab)

library(doBy)
summaryBy(AnnualHouseholdIncome + AnnualFoodSpending ~ RegionLab + LocationLab, data = mydata, 
          FUN = function(x) { c(m = mean(x), s = sd(x)) } )

CV1 <- sd(newdata$AnnualFoodSpending)/mean(newdata$AnnualFoodSpending)
CV2 <- sd(newdata$AnnualHouseholdIncome)/mean(newdata$AnnualHouseholdIncome) 
          

ggplot(newdata, aes(x=factor(RegionLab), y=AnnualHouseholdIncome)) + 
  geom_boxplot(fill="gray")+
  labs(title="Distribution of AnnualHouseholdIncome by RegionLab")+
  theme(plot.title=element_text(lineheight=0.8, face="bold", hjust=0.5))

ggplot(newdata, aes(x=factor(RegionLab), y=AnnualHouseholdIncome, fill=factor(LocationLab))) + 
  geom_boxplot()+
  labs(title="Distribution of AnnualHouseholdIncome")+
  theme(plot.title=element_text(lineheight=0.8, face="bold", hjust=0.5))

ggplot(newdata, aes(x= AnnualHouseholdIncome, y=AnnualFoodSpending)) + geom_point(size=3) +
  ggtitle("AnnualHouseholdIncome vs AnnualFoodSpending") +
  theme(plot.title=element_text(lineheight=0.8, face="bold", hjust=0.5)) +
  geom_smooth(method=lm, se=FALSE)

ggplot(newdata, aes(x= AnnualHouseholdIncome, y=AnnualFoodSpending, shape=RegionLab, color=LocationLab)) + 
  geom_point(size=3) +
  ggtitle("AnnualHouseholdIncome vs AnnualFoodSpending") +
  theme(plot.title=element_text(lineheight=0.8, face="bold", hjust=0.5))

###################################  box plot ################################################

boxplot(newdata$AnnualHouseholdIncome, main = "Distribution of Annual household income", col = "red", 
        horizontal = TRUE, range = 1.5, notch = TRUE, xlab = "Annual household income ($)")
boxplot.stats(newdata$AnnualHouseholdIncome, coef = 1.5, do.conf = TRUE, do.out = TRUE)


boxplot(newdata$AnnualHouseholdIncome, main = "Distribution of Annual household income", col = "red", 
        horizontal = TRUE, range = 3.0, notch = TRUE, xlab = "Annual household income ($)")
boxplot.stats(newdata$AnnualHouseholdIncome, coef = 3.0, do.conf = TRUE, do.out = TRUE)
######################################################################################################
fivenum(newdata$AnnualHouseholdIncome)
IQR(newdata$AnnualHouseholdIncome)

