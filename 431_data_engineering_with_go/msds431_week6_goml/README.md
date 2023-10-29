# Linear Regression with Go + Concurrency

## Project Summary

For 6 we were tasked with creating two models, running them 100x each and doing this concurrently and sequentially.     

For this project we leveraged the linear regression model from GoML. We created a Batch Gradient Ascent version and a Stochastic Gradient Ascent version.

We then leveraged some python to handle a small bit of EDA & ETL work to get the data in the right format for GoML. 

## Findings


Running both models in a loop 100 times resulted in the following


![Alt text](results.png)

Doing a for loop, 1-100 the sequential runs took about 4 minutes.

Doing a for loop, 1-100 the concurrent version took about 1 minute 47 seconds.



## Go Setup


*./base/main.go:* \
This program located in the base folder ran both models sequentially 100 times.


*./concurrent/main.go:* \
This program located in the concurrent folder ran both models concurrently 100 times. To preserve calculations, we chunked out model 1 & model 2 sequences into separate go routines then used a waitgroup to make sure calculations were done on the right versions. This did not play a significant factor in this project, but for future deployments where parameters will be tweaked and results will be different, it'll be very important.


## Instructions

To recreate this, first run the python script to generate a simple linear model to give an estimate for where the model should be. It'll also spit out the two required files in the correct format and order. It'll also create a ./data folder for you and drop these files in there.

Then simply run the executables in each folder or run the batch file to execute both programs in sequence (batch file calls go run instead of using the created executable.

## References

https://github.com/cdipaolo/goml
