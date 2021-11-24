P = [4,3,2,1]
e = 1
def affiche(P) :
    a = ""
    for i in range(len(P)) :
        a = str(P[i]) + "x^" + str(i) + " +" + a
    return a

print(affiche(P))