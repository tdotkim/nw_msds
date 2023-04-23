### Business Statistics for Contemporary Decision Making (Black)
### Chapter 11, Analysis of Variance and Design of Experiments

library(dplyr)
library(ggplot2)
library(tidyr)

# Chapter 11 introduces experimental design and, by way of this topic, introduces
# analysis of variance or ANOVA. Section 11.2 describes one-way ANOVA, wherein
# we test for a significant difference in the means of one variable across levels
# of a second factor or categorical variable. We will work Demonstration Problem
# 11.1 (p. 381) as our example.

# A company has three manufacturing plants, and company officials want to
# determine whether there is a difference in the average age of workers at the
# three locations. The following data are the ages of five randomly selected
# workers at each plant. Perform a one-way ANOVA to determine whether there is
# a significant difference in the mean ages of the workers at the three plants.
# Use alpha = 0.01 and note that the sample sizes are equal.

# So, we're going to break for a moment and discuss 'tidy data.' If you look at
# the table of ages in the Demonstration Problems, data are laid out like this:

# 1   2   3
# 29  32  25 
# 27  33  24  
# 30  31  24  
# 27  34  25
# 28  30  26

# This makes it very easy to read and grasp for humans and if you've worked in
# MS Excel or similar, you would likely be able to:  (1) select One-Way ANOVA
# from a menu, and (2) highlight the columns of data. In R, though, most functions
# will be written to accept what is called 'tidy data,' where each column is a
# variable, each row is a single observation and each element (or 'cell') is a
# single value. If you're familiar with database design, you may recognize this
# as Codd's 3rd normal form, but specific to a single table.

# The way our data is presented in the Demonstration Problem, we have three (3)
# columns with data, all representing plant employee ages and each row includes
# three (3) observations. Instead, we want to have one (1) column of ages and
# a second column indicating whether the age in a given row is of an employee
# from Plant 1, 2 or 3.

# Start from 'scratch,' we could just build our data frame in this way.

(plant_ages <- data.frame(age = c(29, 27, 30, 27, 28, 32, 33, 31, 34, 30, 25,
                                 24, 24, 25, 26),
                         plant = factor(rep(c(1, 2, 3), each = 5))))

# Then, using aov() to perform our ANOVA:
anova <- aov(age ~ plant, data = plant_ages)
summary(anova)

#             Df Sum Sq Mean Sq F value   Pr(>F)    
# plant        2  129.7   64.87   39.71 5.11e-06 ***
# Residuals   12   19.6    1.63                     
# ---
#   Signif. codes:  0 '***' 0.001 '**' 0.01 '*' 0.05 '.' 0.1 ' ' 1

# Alternatively, if we had a or imported a data object structured like the
# table in the Demo Problem, there are functions for getting us to a tidy object.

plant_age2 <- data.frame('1' = c(29, 27, 30, 27, 28),
                         '2' = c(32, 33, 31, 34, 30),
                         '3' = c(25, 24, 24, 25, 26))

tidy_plant <- plant_age2 %>%
  tidyr::pivot_longer(., cols = everything(), names_to = "Plant",
                      values_to = "Age")

summary(aov(Age ~ Plant, data = tidy_plant))

#             Df Sum Sq Mean Sq F value   Pr(>F)    
# Plant        2  129.7   64.87   39.71 5.11e-06 ***
# Residuals   12   19.6    1.63                     
# ---
#   Signif. codes:  0 '***' 0.001 '**' 0.01 '*' 0.05 '.' 0.1 ' ' 1

# We find a significant F statistic for Plant. This tells us that there is a
# significant difference between at least two (2) of the plants, but it doesn't
# tell us where that difference is. For that, we turn to post hoc testing.
# Specifically, Tukey's Honestly Significant Difference test, described in p. 386.

TukeyHSD(x = anova)


# 11.4 introduces randomized block design. In addition to our treatment variable,
# we will consider a 'blocking' variable; a variable we can control but is not
# our treatment variable, i.e. our primary interest. We will work Demonstration
# Problem 11.3 (p. 399) for our example.

# Suppose a national travel association studied the cost of premium unleaded
# gasoline in the United States during the summer of 2016. From experience,
# association directors believed there was a significant difference in the
# average cost of a gallon of premium gasoline among urban areas in different
# parts of the country. To test this belief, they placed random calls to gasoline
# stations in five different cities. In addition, the researchers realized that
# the brand of gasoline might make a difference. They were mostly interested in
# the differences between cities, so they made city their treatment variable. To
# control for the fact that pricing varies with brand, the researchers included
# brand as blocking variable and selected six different brands to participate.
# The researchers randomly telephoned one gasoline station for each brand in
# each city, resulting in 30 measurements (five cities and six brands). Each
# station operator was asked to report the current cost of a gallon of premium
# unleaded gasoline at the station. The data are shown here. Test these data
# by using a randomized block design analysis to determine whether there is a
# significant difference in the average cost of premium unleaded gasoline by
# city. Let alpha = 0.01.

A <- c(3.47, 3.4, 3.38, 3.32, 3.5)
B <- c(3.43, 3.41, 3.42, 3.35, 3.44)
C <- c(3.44, 3.41, 3.43, 3.36, 3.45)
D <- c(3.46, 3.45, 3.4, 3.3, 3.45)
E <- c(3.46, 3.4, 3.39, 3.39, 3.48)
F <- c(3.44, 3.43, 3.42, 3.39, 3.49)

price <- c(A, B, C, D, E, F)
brand <- rep(x = c("A", "B", "C", "D", "E", "F"), each = 5)
city <- rep(x = c("Miami", "Phila", "Minn", "San Ant", "Oak"), times = 6)

gasoline <- data.frame(price, brand, city)
head(gasoline, n = 10)

gas_anova <- aov(price ~ city + brand, data = gasoline)
summary(gas_anova)


# 11.5 introduces two-way ANOVA, where two (or more) treatment variables are
# explored concurrently, and the interaction between them may be considered.

# We will work Demonstration Problem 11.4 (p. 411) as our example.
# Some theorists believe that training warehouse workers can reduce absenteeism.
# Suppose an experimental design is structured to test this belief. Warehouses
# in which training sessions have been held for workers are selected for the
# study. The four types of warehouse are (1) general merchandise, (2) commodity,
# (3) bulk storage, and (4) cold storage. The training sessions are
# differentiated by length. Researchers identify three levels of training
# sessions according to the length of sessions: (1) 1 - 20 days, (2) 21 - 50 days,
# and (3) more than 50 days. Three warehouse workers are selected randomly for
# each particular combination of type of warehouse and session length. The workers
# are monitored for the next year to determine how many days they are absent.
# The resulting data are in the following 4 x 3 design (4 rows, 3 columns)
# structure. Using this information, calculate a two-way ANOVA to determine
# whether there are any significant differences in effects. Use alpha = 0.05.

# The below code builds a 'tidy' data frame from the data provided in the 4 x 3
# table on p. 412.

days_absent <- c(3, 4.5, 4, 5, 4.5, 4, 2.5, 3, 3.5, 2, 2, 3, 2, 2.5, 2, 1,
                      3, 2.5, 1, 3, 1.5, 5, 4.5, 2.5, 2.5, 1, 1.5, 0, 1.5, 2,
                      3.5, 3.5, 4, 4, 4.5, 5)
days_training <- rep(x = c("1-20", "21-50", "More than 50"), each = 12)
warehouse_type <- rep(x = c("Gen. Merch", "Commodity", "Bulk Storage",
                            "Cold Storage"), times = 3, each = 3)

absenteeism <- data.frame(absent = days_absent, training = days_training,
                          warehouse = warehouse_type)

absent_anova <- aov(absent ~ warehouse * training,
                    data = absenteeism)
summary(absent_anova)

mean_absent <- absenteeism %>%
  group_by(., warehouse, training) %>%
  summarise(absent = mean(absent))

# And, we'll go ahead and recreate the Minitab plot from the bottom of page 413.
ggplot(data = mean_absent, aes(x = factor(training), y = absent, group = warehouse,
                               col = warehouse)) +
  geom_line(size = 1.5) +
  labs(x = "Length of Training", y = "Absenteeism") +
  scale_x_discrete(limits = c("1-20", "21-50", "More than 50")) +
  theme(legend.title = element_blank(), legend.position = "bottom")
