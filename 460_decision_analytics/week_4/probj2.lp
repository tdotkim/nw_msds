\* with_jeep *\
Maximize
OBJ: 300 Cars + 600 Jeeps - 50 Machine_1s + 400 Trucks
Subject To
_C1: Cars >= 88
_C2: Trucks >= 26
_C3: Jeeps >= 0
_C4: 2 Cars + 4 Jeeps + 3 Trucks <= 260
_C5: 0.8 Cars + 2 Jeeps - Machine_1s + Trucks <= 0
_C6: Machine_1s <= 98
_C7: 0.6 Cars + 2 Jeeps + 0.7 Trucks <= 73
End
