# %%
import numpy as np
import pandas as pd
import random

# %%
countries = {
    'Argentina': ['Bolivia', 'Brazil', 'Chile', 'Paraguay', 'Uruguay'],
    'Bolivia': ['Argentina', 'Brazil', 'Chile', 'Paraguay', 'Peru'],
    'Brazil': ['Argentina', 'Bolivia', 'Colombia', 'Guyana', 'Paraguay', 'Peru', 'Suriname', 'Uruguay', 'Venezuela', 'French Guiana'],
    'Chile': ['Argentina', 'Bolivia', 'Peru'],
    'Colombia': ['Brazil', 'Ecuador', 'Peru', 'Venezuela'],
    'Ecuador': ['Colombia', 'Peru'],
    'French Guiana': ['Brazil', 'Suriname'],
    'Guyana': ['Brazil', 'Suriname', 'Venezuela'],
    'Paraguay': ['Argentina', 'Bolivia', 'Brazil'],
    'Peru': ['Bolivia', 'Brazil', 'Chile', 'Colombia', 'Ecuador'],
    'Suriname': ['Brazil', 'Guyana', 'French Guiana'],
    'Uruguay': ['Argentina', 'Brazil'],
    'Venezuela': ['Brazil', 'Colombia', 'Guyana']
}


# %%
country_list = list(countries.keys())
country_matrix = np.zeros(shape=(len(country_list),len(country_list)))

# %%
# loop through for each country (13x)
for i, country in enumerate(country_list):
    # loop through each country again to check if any of them are neighbors
    for j, neighbor in enumerate(country_list):
        # if the country is in the list of neighbors in the dict, flag 1
        if neighbor in countries[country]:
            country_matrix[i][j] = 1

country_matrix

# %%
# double check, returns True if transposed = regular
# argentina first column going down == argentina first row across 
np.all(country_matrix.T == country_matrix)

# %%
# loop through array
# value/sum of row
for i in range(len(country_list)):
    country_matrix[i,:] /= country_matrix[i,:].sum()

# %%
# visualize it
# check uruguay only borders Argentina and Brazil
# row should total 1 and argentina/brazil should be .5/.5
df = pd.DataFrame(country_matrix)
df.columns = country_list
df.index = country_list
df.style.format(precision=2).background_gradient(cmap="Purples", axis=None)


# %%
# what are the probabilities for the spy being in each country
# given we start in Chile (get Chiles neighbors)
df.loc['Chile']

# %%
def get_probs(days):
    # each successive day is probability ^ day
    probs = np.dot(df.loc['Chile'],np.linalg.matrix_power(country_matrix, days))
    return probs

# %%
biglist = []
for i in range(0,1000):
    smalldict = {}
    row = get_probs(i)
    smalldict['day'] = i+1
    for j,thing in enumerate(row):
        smalldict[country_list[j]] = thing
    biglist.append(smalldict)



# %%
newdf.head()

# %%
newdf = pd.DataFrame(biglist)
newdf.set_index('day', inplace=True)
newdf

# check our math
# day 1 countries can be Argentina, Bolivia, Peru
# day 2 countries must touch any of the above three
temp = newdf.loc[2]
holder = temp[temp>0]
d2countries = countries['Argentina'] + countries['Bolivia'] + countries['Peru']
for country in list(holder.index):
    if country in d2countries:
        print(f'{country} checks out')
    else:
        print(f'{country} error')


