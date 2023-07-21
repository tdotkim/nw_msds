\* problem *\
Minimize
OBJ: 40 Atlanta_to_Columbus + 35 Atlanta_to_Mobile + 30 Atlanta_to_Richmond
 + 50 Boston_to_Columbus + 35 Columbus_to_Atlanta + 45 Jacksonville_to_Atlanta
 + 50 Jacksonville_to_Mobile + 50 Jacksonville_to_Richmond
 + 25 Mobile_to_Atlanta + 30 Newark_to_Boston + 40 Newark_to_Richmond
Subject To
_C1: - Newark_to_Boston - Newark_to_Richmond >= -200
_C2: - Boston_to_Columbus + Newark_to_Boston >= 100
_C3: Atlanta_to_Columbus + Boston_to_Columbus - Columbus_to_Atlanta >= 60
_C4: - Atlanta_to_Columbus + Atlanta_to_Mobile + Atlanta_to_Richmond
 + Columbus_to_Atlanta + Jacksonville_to_Atlanta + Mobile_to_Atlanta >= 170
_C5: Atlanta_to_Richmond + Jacksonville_to_Richmond + Newark_to_Richmond >= 80
_C6: Atlanta_to_Mobile + Jacksonville_to_Mobile - Mobile_to_Atlanta >= 70
_C7: - Jacksonville_to_Atlanta - Jacksonville_to_Mobile
 - Jacksonville_to_Richmond >= -300
Bounds
 Atlanta_to_Columbus free
 Atlanta_to_Mobile free
 Atlanta_to_Richmond free
 Boston_to_Columbus free
 Columbus_to_Atlanta free
 -inf <= Jacksonville_to_Atlanta <= 300
 -inf <= Jacksonville_to_Mobile <= 300
 -inf <= Jacksonville_to_Richmond <= 300
 Mobile_to_Atlanta free
 -inf <= Newark_to_Boston <= 200
 Newark_to_Richmond free
Generals
Atlanta_to_Columbus
Atlanta_to_Mobile
Atlanta_to_Richmond
Boston_to_Columbus
Columbus_to_Atlanta
Jacksonville_to_Atlanta
Jacksonville_to_Mobile
Jacksonville_to_Richmond
Mobile_to_Atlanta
Newark_to_Boston
Newark_to_Richmond
End
