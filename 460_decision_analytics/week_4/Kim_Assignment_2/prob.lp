\* problem *\
Maximize
OBJ: 8 Prof_1_Finance_Fall + 8 Prof_1_Finance_Spring + 9 Prof_1_Marketing_Fall
 + 9 Prof_1_Marketing_Spring + 7 Prof_1_Production_Fall
 + 7 Prof_1_Production_Spring + 11 Prof_2_Finance_Fall
 + 11 Prof_2_Finance_Spring + 9 Prof_2_Marketing_Fall
 + 9 Prof_2_Marketing_Spring + 10 Prof_2_Production_Fall
 + 10 Prof_2_Production_Spring + 8 Prof_3_Finance_Fall
 + 8 Prof_3_Finance_Spring + 9 Prof_3_Marketing_Fall
 + 9 Prof_3_Marketing_Spring + 10 Prof_3_Production_Fall
 + 10 Prof_3_Production_Spring
Subject To
_C1: Prof_1_Finance_Fall + Prof_1_Finance_Spring + Prof_1_Marketing_Fall
 + Prof_1_Marketing_Spring + Prof_1_Production_Fall + Prof_1_Production_Spring
 = 4
_C10: Prof_1_Marketing_Fall + Prof_1_Marketing_Spring + Prof_2_Marketing_Fall
 + Prof_2_Marketing_Spring + Prof_3_Marketing_Fall + Prof_3_Marketing_Spring
 <= 4
_C11: Prof_1_Finance_Fall + Prof_1_Finance_Spring + Prof_2_Finance_Fall
 + Prof_2_Finance_Spring + Prof_3_Finance_Fall + Prof_3_Finance_Spring <= 4
_C12: Prof_1_Production_Fall + Prof_1_Production_Spring
 + Prof_2_Production_Fall + Prof_2_Production_Spring + Prof_3_Production_Fall
 + Prof_3_Production_Spring <= 4
_C2: Prof_2_Finance_Fall + Prof_2_Finance_Spring + Prof_2_Marketing_Fall
 + Prof_2_Marketing_Spring + Prof_2_Production_Fall + Prof_2_Production_Spring
 = 4
_C3: Prof_3_Finance_Fall + Prof_3_Finance_Spring + Prof_3_Marketing_Fall
 + Prof_3_Marketing_Spring + Prof_3_Production_Fall + Prof_3_Production_Spring
 = 4
_C4: Prof_1_Marketing_Spring + Prof_2_Marketing_Spring
 + Prof_3_Marketing_Spring >= 1
_C5: Prof_1_Finance_Spring + Prof_2_Finance_Spring + Prof_3_Finance_Spring
 >= 1
_C6: Prof_1_Production_Spring + Prof_2_Production_Spring
 + Prof_3_Production_Spring >= 1
_C7: Prof_1_Marketing_Fall + Prof_2_Marketing_Fall + Prof_3_Marketing_Fall
 >= 1
_C8: Prof_1_Finance_Fall + Prof_2_Finance_Fall + Prof_3_Finance_Fall >= 1
_C9: Prof_1_Production_Fall + Prof_2_Production_Fall + Prof_3_Production_Fall
 >= 1
End
