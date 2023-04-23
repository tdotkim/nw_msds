### Business Statistics for Contemporary Decision Making (Black)
### Chapter 2, Charts and Graphs

library(ggplot2)

# Ch. 2 is concerned with "techniques for summarizing and depicting data."
# We'll use the Canadian unemployment rate data introduced in 2.1 and used
# throughout to show

# Our vector of unemployment rate data:
can_unemp <- c(2.3, 2.8, 3.6, 2.4, 2.9, 3, 4.6, 4.4, 3.4, 4.6, 6.9, 6, 7,
               7.1, 5.9, 5.5, 4.7, 3.9, 3.6, 4.1, 4.8, 4.7, 5.9, 6.4, 6.3, 5.6,
               5.4, 7.1, 7.1, 8, 8.4, 7.5, 7.5, 7.6, 11, 12, 11.3, 10.6,
               9.7, 8.8, 7.8, 7.5, 8.1, 10.3, 11.2, 11.4, 10.4, 9.5, 9.6, 9.1,
               8.3, 7.6, 6.8, 7.2, 7.7, 7.6, 7.2, 6.8, 6.3, 6)

length(can_unemp)  # just to check. we should have 60 values

# In Table 2.2 (p. 17), this data has been grouped to give us frequencies; for
# how many years was the rate 1 to less than 3, 3 to less than 5, etc. We'll do
# this ourselves in R.

# 
our_breaks <- c(1, 3, 5, 7, 9, 11, 13)  # or, seq(from = 1, to = 13, by = 2)
(our_grouped_data <- cut(x = can_unemp, breaks = our_breaks,
                         right = FALSE))  # 'right = FALSE' gives us x to < y

# And, to see how we did:
(freqs <- table(our_grouped_data))  # hooray for us.

# To Table 2.3 (p. 18), we can also calculate our relative and cumulative
# frequencies. Then, we'll create a matrix with our intervals, class midpoints,
# relative and cumulative frequences.

(relative_freqs <- freqs / length(can_unemp))

(cumulative_freqs <- cumsum(freqs))

(class_midpoints <- seq(from = 2, to = 12, by = 2))

# Lastly,
our_matrix <- cbind(freqs,
                    class_midpoints,
                    relative_freqs,
                    cumulative_freqs)

our_matrix


# In 2.2, Black describes a few types of graph for presenting quantitative data.
# We'll go over R code for building those same graphs in R. We will do so using
# the base R and "ggplot2" syntaxes.

# Histograms
# To create a histogram using base R, we'll use the hist() function.
?hist  # or, help(hist)

hist(x = can_unemp,
     breaks = "Sturges",  # we'll use R's default break-determining method
     main = "",  # suppress title
     xlab = "Unemployment Rates for Canada",  # x-axis label
     ylab = "Frequency"  # y-axis label
     )

# Now, simply to demonstrate how, we'll try and match the bins and breaks for 
# Figure 2.1 (p. 21).

hist(x = can_unemp,
     breaks = c(1, 3, 5, 7, 9, 11, 13),
     main = "",  # suppress title
     xaxt = "n",  # suppress x-axis so we can specify exactly what we want
     xlab = "Unemployment Rates for Canada",
     xlim = c(0, 14),
     ylim = c(0, 20)  # y-axis range
     )
axis(side = 1, at = c(0, 14), labels=c("", ""), lwd.ticks = 0)
axis(side = 1, at = c(1, 3, 5, 7, 9, 11, 13))

# The above is a bit awkward; two axis() calls, but it seemed worthwhile to
# demonstrate that R - both base R and ggplot2 - gives you a great deal of
# control over graphics.

# And now, we'll create the same histogram in ggplot2.
ggplot(data = NULL, aes(x = can_unemp)) +
  geom_histogram(breaks = c(1, 3, 5, 7, 9, 11, 13)) +
  labs(x = "Unemployment Rates of Canada", y = "Frequency") +
  scale_x_continuous(breaks = c(1, 3, 5, 7, 9, 11, 13),
                   limits = c(0, 14))


# On p. 23, we a frequency polygon of the unemployment data. In base R,
plot(x = class_midpoints,
     y = as.vector(freqs),  # recall that our "freqs" object is a table
     type = "l",  # to create a line plot
     xlab = "Class Midpoints",
     ylab = "Frequency")

# And, in ggplot2,
ggplot(data = NULL, aes(x = class_midpoints, y = as.vector(freqs))) +
  geom_line() + labs(x = "Class Midpoints", y = "Frequency") +
  scale_x_continuous(breaks = class_midpoints)


# P. 23 also includes an an ogive or cumulative frequency polygon.
plot(x = c(1, 3, 5, 7, 9, 11, 13),
     y = c(0, cumsum(freqs)),
     type = "b",  # b for "both" a point and line plot
     xaxt = "n",
     xlab = "Class Endpoints",
     xlim = c(0, 14),
     ylab = "Cumulative Frequency")
axis(side = 1, at = c(0, 14), labels=c("", ""), lwd.ticks = 0)
axis(side = 1, at = c(1, 3, 5, 7, 9, 11, 13))

ggplot(data = NULL, aes(x = c(1, 3, 5, 7, 9, 11, 13),
                        y = c(0, cumsum(freqs)))) +
  geom_line() +
  geom_point() +
  labs(x = "Class Midpoints", y = "Frequency") +
  scale_x_continuous(breaks = c(1, 3, 5, 7, 9, 11, 13))


# Dot plots, p. 24
stripchart(x = can_unemp, method = "stack", pch = 20, at = 0,
           xlab = "Annual Unemployment Rates for Canada")

ggplot(data = NULL, aes(x = can_unemp)) + geom_dotplot(binwidth = 0.2) +
  labs(x = "Annual Unemployment Rates for Canada") +
  theme(axis.text.y = element_blank(),
        axis.ticks.y = element_blank(),
        axis.title.y = element_blank()) +
  scale_x_continuous(breaks = c(1, 3, 5, 7, 9, 11, 13),
                     limits = c(1, 13))

# Stem-and-Leaf Plots
stem(can_unemp)


# In 2.3, a few types of qualitative graphs are introduced for presenting non-
# numeric, categorical data; specifically, frequencies and proportions of class
# memberships.

# Personally, i think pie charts deserve the criticism advanced against them.
# Humans, generally (and me, specifically), don't seem to be as good at
# comparing relative area as we are at comparing height | length.

# We do need some non-numeric data, so we're going to put aside the Canadian
# unemployment rate data. Instead, we'll use the U.S. petroleum company capacity
# data from Table 2.6 (p. 28). Code below will create a data frame with two (2)
# columns - i.e. two (2) vectors - with the company name and capacity.

petro_cap <- data.frame(company = c("Exxon Mobil", "Valero Energy", "Chevron",
                                    "ConocoPhilips", "Marathon Oil"),
                        capacity_1000 = c(5589, 2777, 2540, 2514, 1714))
str(petro_cap)

pie(x = petro_cap$capacity_1000, labels = petro_cap$company)

# There isn't a built-in pie "geom" in ggplot2. However, we can get to one
ggplot(data = petro_cap, aes(x = "", y = capacity_1000, fill = company)) +
  geom_bar(stat = "identity") +
  scale_fill_discrete(name = "Company") +  # just to capitalize our legend title
  coord_polar("y", start = 0) +
  theme(axis.ticks = element_blank(),
        axis.text = element_blank(),
        panel.grid = element_blank(),
        axis.title = element_blank())

# You will see barplots far more often. Instead of comparing
# areas, we're comparing heights or lengths, and we can do a secondary grouping
# when called for by having 'stacked' or 'beside' ('dodge' in ggplot2) barplots.
barplot(height = petro_cap$capacity_1000,
        names.arg = petro_cap$company,
        ylab = "Capacity (in 1,000s barrels/day)",
        col = c("antiquewhite", "darksalmon", "darkseagreen", "cornflowerblue",
                "darkkhaki"))  # if we wanted color

ggplot(data = petro_cap, aes(y = capacity_1000,
                             x = reorder(company, -capacity_1000),
                             fill = reorder(company, -capacity_1000))) +
         geom_bar(stat = "identity", col = "gray80") +
         labs(x = "", y = "Capacity (in 1,000s barrels/day)", fill = "Company") +
         scale_fill_discrete(type = c("antiquewhite", "darksalmon",
                                      "darkseagreen", "cornflowerblue",
                                      "darkkhaki")) 

# Pareto chart; think rank-ordered bar chart with a cumulative count or
# proportion, usually. For our purposes, we'll continue with the petro_cap
# data frame.

# There is not a built-in Pareto chart function in base R. However, we can
# construct one.

# First, we'll create a vector with the cumulative proportions:
petro_cap$cum_prop <- cumsum(petro_cap$capacity_1000) /
  sum(petro_cap$capacity_1000)

# Now, we'll start with a barplot and overlay a line plot representing
# our cumulative capacity.
pareto_chart <- barplot(height = petro_cap$capacity_1000,
        names.arg = petro_cap$company,
        ylab = "Capacity (in 1,000s barrels/day)",
        ylim = c(0, 1.05 * max(cumsum(petro_cap$capacity_1000)))
        )
lines(x = pareto_chart, y = cumsum(petro_cap$capacity_1000),
      type = "b", col = "deepskyblue4")

# Now, we're going to 'cheat' just a bit and create a right-hand side axis
# scaled to treat our cumulative 1,000 gallons as percentages. We're also
# going to increase the amount of space available for our secondary y-axis.
par(mar = c(5.1, 4.1, 4.1, 4.1))  # default is [1] 5.1 4.1 4.1 2.1
axis(side = 4, at = c(0, cumsum(petro_cap$capacity_1000)),
     labels = paste(c(0, round(petro_cap$cum_prop * 100)), "%", sep = ""),
     las = 1, col.axis = "deepskyblue4", col = "deepskyblue4")

# and, again, for civility:
par(mar = c(5.1, 4.1, 4.1, 2.1))

# ggplot2
ggplot(data = petro_cap, aes(y = capacity_1000,
                             x = reorder(company, -capacity_1000))) +
  geom_bar(stat = "identity") +
  geom_point(aes(y = cumsum(capacity_1000)), color = "deepskyblue4", pch = 16) +
  geom_path(aes(y = cumsum(capacity_1000), group = 1),
            color = "deepskyblue4", lty = 2) +
  labs(x = "Company", y = "Capacity (in 1,000s barrels/day)") +
  theme(axis.text.x = element_text(angle = 90, vjust = 0.5))

# Tl;dr there aren't existing Pareto chart functions in base R or ggplot2. This
# might be one of the times we look to an outside package to try and find an
# option. We might just find the "qcc" package; initialism for Quality Control
# Charts.

library(qcc)

# We need a vector of capacities and we want to 'name' the
caps <- petro_cap$capacity_1000
names(caps) <- petro_cap$company

# And, finally,
pareto.chart(data = caps, main = "Pareto Chart for 1,000 barrel/day capacities")











