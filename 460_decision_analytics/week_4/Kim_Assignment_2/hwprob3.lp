\* problem *\
Minimize
OBJ: 0.1 Fast_Clean_FROM_Day1 + 0.1 Fast_Clean_FROM_Day2
 + 0.1 Fast_Clean_FROM_Day3 + 0.1 Fast_Clean_FROM_Day4 + 0.2 New_Buy_Day1
 + 0.2 New_Buy_Day2 + 0.2 New_Buy_Day3 + 0.2 New_Buy_Day4
 + 0.06 Slow_Clean_Day1_Delivered_Day3 + 0.06 Slow_Clean_Day2_Delivered_Day4
 + 0.06 Slow_Clean_Day3_Delivered_Day5 + 0.06 Slow_Clean_Day4_Delivered_Day6
Subject To
_C1: New_Buy_Day1 = 15
_C2: Fast_Clean_FROM_Day1 + New_Buy_Day2 = 12
_C3: Fast_Clean_FROM_Day2 + New_Buy_Day3 + Slow_Clean_Day1_Delivered_Day3 = 18
_C4: Fast_Clean_FROM_Day3 + New_Buy_Day4 + Slow_Clean_Day2_Delivered_Day4 = 6
_C5: Fast_Clean_FROM_Day1 + Slow_Clean_Day1_Delivered_Day3 <= 15
_C6: Fast_Clean_FROM_Day2 + Slow_Clean_Day2_Delivered_Day4 <= 12
_C7: Fast_Clean_FROM_Day3 <= 18
End
