library(readr)

library(randomForest)

train_set <- read_csv("data/mnist_train.csv", col_names = FALSE, show_col_types = FALSE)
test_set <- read_csv("data/mnist_test.csv", col_names = FALSE, show_col_types = FALSE)
train_labels <- as.factor(train_set[, 1]$X1)
test_labels <- as.factor(test_set[, 1]$X1)
#print(head(train_labels, 20))




rf <- randomForest(x = train_set, y = train_labels, xtest = test_set, ntree = 20, nodesize=50)

print(c("Accuracy", 1 - mean(rf$err.rate)))