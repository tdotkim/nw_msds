## Testing Go for Statistics

### Project Summary

The company's data scientists are concerned about the prospect of having to use Go for their work. At the very least, the data scientists want to ensure that the proposed Go statistics package will provide correct answers. Tests could examine Go linear regression results against results from Python and R. It is suggested that an initial test be run on four small data sets: The Anscombe Quartet as described by Anscombe (1973) and Miller (2015). 

To reassure the company's data scientists we tested Ordinary Least Squares in Go, Python, and R. Our aim was to provide two levels of cetainty, one that the results are the same to four decimal places and two, that the time to run similar analysis is close enough that the runtime differences are negligible. 

For the first level of certainty that Go will be usable for the team we looked for the calculated slope, intercept, and R-Squared value. These three metrics when equal across all thre packages should ensure a high level of confidence in the correctness across languages/packages.

For the second level of certainty, we timed the executions. In a batch file we ran all three (Go, Python, R) scripts with timers added to the main chunks of each code (excluding the write to csv part). We ran each script 100 times.

In the results.py file we ran summary statistics against the runs.

R had a mean run time of 0.019205 seconds, Go had a mean run time of 0.000099 seconds, and Python had a mean run time of 1.383823 seconds. 

Go is significantly faster but we must note we're nowhere near as comprehensive as the packages being run for R or Python to generate the models and its associated metrics.


### Files

*main.go:* \
Similar to the python OLS function in statsmodels.api, we use Flynn's [LinearRegression](https://github.com/montanaflynn/stats/blob/master/regression.go) in the stats package to output predicted values and gather a few key datapoints into an object. We then use these results to compare against the same points in both R and Python. For the purposes of this exercise, only a few were chosen instead of refactoring the entire package and function to Go.

*main_test.go:* \
Table testing to verify outputs match Python and R results. Some finangling done to get rounding to match.

*miller_py.py:* \
Professor Miller's original file with graph outputs removed, timers added, warnings suppressed, and runtime results export to csv added.

*miller_r.R:* \
Professor Miller's original file with graph outputs removed, timers added, warnings suppressed, and runtime results export to csv added.

*results.py:* \
Compiles the three different CSVs of run times then does a pd.describe() to get some basic stats about the run times.

*run_all.bat:* \
Clears the runs folder. Then executes the main.go, miller_py, and miller_r files 100 times each. At the end it runs the results.py file.

### References

Anscombe, F. J. 1973. “Graphs in Statistical Analysis.” The American Statistician 27 (1): 17–21. https://doi.org/10.2307/2682899. \
Flynn, Montana. 2023. “Stats - Golang Statistics Package.” 2023. https://github.com/montanaflynn/stats. \
Miller, Thomas. 2015. “Modeling Techniques in Predictive Analytics Chapter 1.” 2015. https://github.com/mtpa/mtpa/tree/master/MTPA_Chapter_1.