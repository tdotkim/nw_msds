\* problem *\
Minimize
OBJ: 30 Route_(1,_2) + 40 Route_(1,_4) + 50 Route_(2,_3) + 35 Route_(3,_5)
 + 40 Route_(5,_3) + 30 Route_(5,_4) + 35 Route_(5,_6) + 25 Route_(6,_5)
 + 50 Route_(7,_4) + 45 Route_(7,_5) + 50 Route_(7,_6)
Subject To
Constraint_1: - Route_(1,_2) - Route_(1,_4) >= -200
Constraint_2: Route_(1,_2) - Route_(2,_3) >= 100
Constraint_3: Route_(2,_3) - Route_(3,_5) + Route_(5,_3) >= 60
Constraint_4: Route_(1,_4) + Route_(5,_4) + Route_(7,_4) >= 80
Constraint_5: Route_(3,_5) - Route_(5,_3) - Route_(5,_4) - Route_(5,_6)
 + Route_(6,_5) + Route_(7,_5) >= 170
Constraint_6: Route_(5,_6) - Route_(6,_5) + Route_(7,_6) >= 70
Constraint_7: - Route_(7,_4) - Route_(7,_5) - Route_(7,_6) >= -300
Bounds
 Route_(1,_2) <= 200
 Route_(1,_4) <= 200
 Route_(2,_3) <= 200
 Route_(3,_5) <= 200
 Route_(5,_3) <= 200
 Route_(5,_4) <= 200
 Route_(5,_6) <= 200
 Route_(6,_5) <= 200
 Route_(7,_4) <= 300
 Route_(7,_5) <= 300
 Route_(7,_6) <= 300
End
