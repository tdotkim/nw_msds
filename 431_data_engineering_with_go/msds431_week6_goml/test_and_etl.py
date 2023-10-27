# %%
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt
from matplotlib.colors import LinearSegmentedColormap
import numpy as np
from sklearn.model_selection import train_test_split 
from sklearn.linear_model import LinearRegression 
from sklearn.metrics import mean_squared_error, mean_absolute_error 
from pathlib import Path


# %%
df = pd.read_csv('boston.csv')

# %%
plt.figure(figsize=(20, 10))
cmap = LinearSegmentedColormap.from_list('', ['crimson', 'gold', 'lime'])

sns.heatmap(df.corr().abs(), cmap=cmap, annot=True)

# %%
train, test = train_test_split(df, test_size=0.2)

# %% [markdown]
# # simple test

# %%
X = df.drop('mv', axis=1)

# %%
y = df['mv']

# %%
X_train, X_test, y_train, y_test = train_test_split( 
    X, y, test_size=0.3, random_state=101) 

# %%
reg = LinearRegression().fit(X_train, y_train)


# %%
predictions = reg.predict(X_test) 

# %%
reg.coef_

# %%
print('mean_squared_error : ', mean_squared_error(y_test, predictions)) 
print('mean_absolute_error : ', mean_absolute_error(y_test, predictions)) 

# %%
train_df = pd.concat([X_train,y_train],axis=1)
test_df = pd.concat([X_test,y_test],axis=1)

Path("./data").mkdir(parents=True, exist_ok=True)

train_df.to_csv('./data/training_set_noheader.csv', index=False, header=False)
test_df.to_csv('./data/testing_set_noheader.csv', index=False, header=False)


