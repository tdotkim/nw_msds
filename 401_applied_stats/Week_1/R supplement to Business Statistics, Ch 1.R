### Business Statistics for Contemporary Decision Making (Black)
### Chapter 1, Introduction to Statistics

# 1.2 Data Measurement describes 'four common levels of data measurement,'
# nominal, ordinal, interval and ratio.

# In contrast, R has data types. These do not correspond to the 'levels' defined
# in Black. However, we can represent data of any level using the available
# types in R.

### NOMINAL
# For example, nominal data is "used only to classify or categorize." In R,
# nominal data could be stored as a few different types.

# Black (p. 6) uses 'employment classifications' as an example, and encodes them
# as follows:

# 1. Educator
# 2. Construction worker
# 3. Manufacturing worker
# 4. Lawyer
# 5. Doctor
# 6. Other

# In R, we may store such data as a character string. Doing so preserves the 
# textual 'value,' and nothing would be lost because there's no meaningful order
# to these classifications. Below, we'll use the 'combine' function, c(), to
# create a character vector with our classifications.

our_char_vector <- c("Educator", "Construction worker", "Manufacturing worker",
                     "Lawyer", "Doctor", "Other")

# Executing the code above creates a vector of our employment classifications
# and assigns the vector to the name "our_char_vector." In the code below, we
# send our vector to the Console, then look at the structure and class. Note
# that there may be times when you need to explicitly print() and object, but
# here we just call our object and the output is the same.

our_char_vector # versus,
print(our_char_vector)

str(our_char_vector)
class(our_char_vector)  

# The class() and str() functions are quite useful fpr checking on objects
# you've created or changed; making sure whatever code you ran did what you'
# expected it do, especially when dealing with objects too large to view in
# their entirety.

# We may also store our employment classifications as a factor. The elements -
# i.e. the individual values will still be character - but R will recognize and
# treat the values as levels of factor or category.

our_factor_vector <- factor(x = our_char_vector)

str(our_factor_vector)
levels(our_factor_vector)

# Note that 'behind the scenes' R encodes the vector using integers 1 through 6
# and the character strings and attached to this integers. However, our vector
# is "unordered." 1 ("Construction worker") is neither greater nor less than 
# the other values.

# Lastly, we could've encoded our data numerically ourselves, separate from R.
# The only real risks are (1) making sure users know what the encoding means, 
# and (2) making sure users know what arithmetic or statistical operations will
# have meaningful results. What is the mean of our employment classification
# vector?

mean(our_factor_vector)


### ORDINAL
# Ordinal data is itself nominal data, where each element corresponds to some
# value, but with the added feature of having a rank order. In R, we would
# generally choose to store ordinal data as a factor like we did above, but
# specify an order to the values. Black (p. 7) uses a 'Likert-type' scale as an
# example. We'll create a vector or responses using this "not helpful" to
# "extremenly helpful" scale, then coerce this vector to an ordered factor.

our_vector <- c("very helpful", "extremely helpful", "extremely helpful",
                "not helpful", "extremely helpful", "moderately helpful",
                "extremely helpful", "moderately helpful", "very helpful",
                "very helpful")

# Now, to coerce to an ordered factor:

our_ord_factor <- factor(x = our_vector,
                         levels = c("not helpful", "somewhat helpful",
                                    "moderately helpful", "very helpful",
                                    "extremely helpful"),
                         ordered = TRUE)

str(our_ord_factor)
levels(our_ord_factor)


### INTERVAL AND RATIO
# Both interval- and ratio-level data will generally be stored as numeric data
# in R. It will ultimately be up to the user to understand and communicate if
# the data is interval or ratio; whether the ratio of two values is meaningful;
# whether zero represents a value on a scale or absence.

# It should be noted that R has two (2) types for numeric data:  double, short
# for double-precision floating-point, and integer. There is no 'single'
# precision mode in R.

# First, we'll create a numeric vector using only whole numbers and explore how
# R handles it.

our_num_vector <- c(1, 3, 5, 7, 9, 2, 4, 6, 8, 10)
typeof(our_num_vector)  # [1] "double"

# Okay, so R defaults to double. If we wanted, we could coerce it to integer.
# We won't lose any information and we'd save a few bytes.

our_int_vector <- as.integer(our_num_vector)
typeof(our_int_vector)

object.size(our_num_vector)
object.size(our_int_vector)

# Now, we can have vectors of any data type in R. The 'rule' is that every
# element of that vector - same is true of matrices - has to be of the same
# data type. We can have character, double or int vectors, but we can't have
# character and double or double and int vectors. The same is true for any of
# the data types available.

# So, what happens we try to add a double-precision or character element to
# our_int_vector?

(what_happens_double <- c(our_int_vector, sqrt(2)))
(what_happens_char <- c(our_int_vector, "ABCDEF"))

# R tip:  if you put assignment code in parentheses like the above, you'll
# assign the object to the name you have and output it to the Console. For
# 'assignment,' for creating objects in R, we'll generally use:
# name_of_our_object <- code_defining_the_object_of_interest

typeof(what_happens_double)
typeof(what_happens_char)


### LOGICAL

# A very common data type in R is logical, where every element in a vector
# (or matrix) takes the value TRUE or FALSE, encoded as 1 and 0. We'll create
# a logical vector below by evaluating whether the elements of our ordered
# factor vector above are "very" or "extremely helpful."

(our_logical_vector <- our_ord_factor >= "very helpful")

# We can do the same with our_num_vector; whether each element is less than 6:
(our_other_logical <- our_num_vector < 6)

# In addition to storing logical data, the 1|0 coding makes it easy to count
# the TRUEs. If every TRUE is 1, then:

sum(our_logical_vector)

# returns the number of "very helpful"s and "extremely helpful"s in our survey
# responses. To extend it just a bit further:

sum(our_logical_vector) / length(our_logical_vector)

# returns the proportion of "very helpful" or "extremely helpful" responses.


### COMPLEX AND RAW

# You will have long, wholly satisfying R-using lives and never see complex or
# raw data. However, they exist as data types, so...

# Complex numbers are of the form a + bi, where a and b are real numbers and i
# is the imaginary unit. The complex data type would be used to encode complex
# or imaginary numbers; imaginary numbers with zero for the complex component.

# For example, our complex vector:

(our_complex_vector <- c(2 + 1i, 3 + 2i, 3i))
class(our_complex_vector)
 
# Raw data encodes data as 'raw' bytes. For our example, we'll coerce
# our_int_vector to raw.

(our_raw_vector <- as.raw(our_int_vector))

# R tip:  for data types and objects in R, there are "as" and "is" functions.
# The "as" functions will attempt to coerce a data object to the form specified
# like we did above with as.raw(). The "is" functions will test whether an
# object is of a particular data type or object and returns a logical vector
# indicating whether something "is." Our concluding example:

is.complex(x = our_complex_vector)
