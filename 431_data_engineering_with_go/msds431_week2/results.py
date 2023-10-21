# %%
import pandas as pd


# %%
r_runs = pd. read_csv('./runs/r_runs.csv',header=None,names=['R_Runtime'])
go_runs = pd. read_csv('./runs/go_runs.csv',header=None,names=['Go_Runtime'])
py_runs = pd. read_csv('./runs/python_runs.csv',header=None,names=['Py_Runtime'])

# %%
merged_df = pd.merge(r_runs,(pd.merge(go_runs, py_runs, left_index=True, right_index=True)), left_index=True, right_index=True)

# %%
print(merged_df.describe())


