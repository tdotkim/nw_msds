
import pandas as pd
import numpy as np

import matplotlib.pyplot as plt
import seaborn as sns
sns.set()




pd.set_option('display.max_rows', None)
pd.set_option('display.max_columns', None)
pd.set_option('display.width', None)
pd.set_option('display.max_colwidth', None)


INFILE = "C:\\NWU422\\DATA\\Insurance.csv"

TARGET_F = "TARGET_CLM_FLAG"
TARGET_A = "TARGET_CLM_AMT"


df = pd.read_csv( INFILE )


dt = df.dtypes
objList = []
for i in dt.index :
    #print(" here is i .....", i , " ..... and here is the type", dt[i] )
    if i in ( [ TARGET_F, TARGET_A ] ) : continue
    if dt[i] in (["object"]) : objList.append( i )


##print(" OBJECTS ")
##print(" ------- ")
##for i in objList :
##    print( i )
##print(" ------- ")
##
##
##for i in objList :
##    print( i )
##    print( df[i].unique() )
##    g = df.groupby( i )
##    print( g[i].count() )
##    print( "MOST COMMON = ", df[i].mode()[0] )   
##    print( "MISSING = ", df[i].isna().sum() )
##    print( "\n\n")



"""
FILL IN MISSING WITH THE MODE
"""
##for i in objList :
##    if df[i].isna().sum() == 0 : continue
##    print( i ) 
##    print("HAS MISSING")
##    NAME = "IMP_"+i
##    print( NAME ) 
##    df[NAME] = df[i]
##    df[NAME] = df[NAME].fillna(df[NAME].mode()[0] )
##    print( "variable",i," has this many missing", df[i].isna().sum() )
##    print( "variable",NAME," has this many missing", df[NAME].isna().sum() )
##    g = df.groupby( NAME )
##    print( g[NAME].count() )
##    print( "\n\n")
##    df = df.drop( i, axis=1 )




"""
FILL IN MISSING WITH THE CATEGORY "MISSING"
"""
##for i in objList :
##    if df[i].isna().sum() == 0 : continue
##    print( i ) 
##    print("HAS MISSING")
##    NAME = "IMP_"+i
##    print( NAME ) 
##    df[NAME] = df[i]
##    df[NAME] = df[NAME].fillna("MISSING")
##    print( "variable",i," has this many missing", df[i].isna().sum() )
##    print( "variable",NAME," has this many missing", df[NAME].isna().sum() )
##    g = df.groupby( NAME )
##    print( g[NAME].count() )
##    print( "\n\n")
##    df = df.drop( i, axis=1 )



#print( df.head(3).T )




dt = df.dtypes
objList = []
for i in dt.index :
    #print(" here is i .....", i , " ..... and here is the type", dt[i] )
    if i in ( [ TARGET_F, TARGET_A ] ) : continue
    if dt[i] in (["object"]) : objList.append( i )


##print(" OBJECTS ")
##print(" ------- ")
##for i in objList :
##    print( i )
##print(" ------- ")


##for i in objList :
##    print(" Class = ", i )
##    print( df[i].unique() )
##    g = df.groupby( i )
##    x = g[ TARGET_F ].mean()
##    print( "Crash Prob", x )
##    print( " ................. ")
##    x = g[ TARGET_A ].mean()
##    #x = g[ TARGET_A ].median()
##    print( "Damage Amount", x )
##    print(" ===============\n\n\n ")


'''
EXPLORE THE CATEGORICAL / OBJECT VARIABLES
'''

#print( df["EDUCATION"].unique() )
#print( df.EDUCATION.unique() )


##df["y_EDU_4"] = (df.EDUCATION.isin( ["a_PhD"] ) + 0 )
##df["y_EDU_3"] = (df.EDUCATION.isin( ["a_PhD","b_Masters"] ) + 0)
##df["y_EDU_2"] = (df.EDUCATION.isin( ["a_PhD","b_Masters","c_Bachelors"] ) + 0)
##df["y_EDU_1"] = (df.EDUCATION.isin( ["a_PhD","b_Masters","c_Bachelors","d_High School"] ) + 0)
##df = df.drop( "EDUCATION", axis=1 )
##print( df.head().T )
##
##
##for i in ["y_EDU_4","y_EDU_3","y_EDU_2","y_EDU_1"] :
##    print(" Class = ", i )
##    g = df.groupby( i )
##    x = g[ TARGET_F ].mean()
##    print( "Crash Prob", x )
##    print( " ................. ")
##    x = g[ TARGET_A ].mean()
##    #x = g[ TARGET_A ].median()
##    print( "Damage Amount", x )
##    print(" ===============\n\n\n ")


    
##
##dt = df.dtypes
##objList = []
##for i in dt.index :
##    #print(" here is i .....", i , " ..... and here is the type", dt[i] )
##    if i in ( [ TARGET_F, TARGET_A ] ) : continue
##    if dt[i] in (["object"]) : objList.append( i )
##
##print(" OBJECTS ")
##print(" ------- ")
##for i in objList :
##    print( i )
##print(" ------- ")
##
##
##
##for i in objList :
##    #print(" Class = ", i )
##    thePrefix = "z_" + i
##    #print( thePrefix )
##    y = pd.get_dummies( df[i], prefix=thePrefix, dummy_na=False )   
##    #print( type(y) )
##    #print( y.head().T )
##    df = pd.concat( [df, y], axis=1 )
##    df = df.drop( i, axis=1 )
##    print( df.head(3).T )
##









