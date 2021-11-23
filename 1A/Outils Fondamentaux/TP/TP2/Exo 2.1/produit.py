P = [1,2,3,4]
Q = [1,2,3]
R = []
i = 0

def produit(P,Q) :
   #Créer un tableau R de la bonne taille selon le degré de PxQ
    for i in range(len(P)+len(Q)-1) :
       R.append(0)
    for j in range(len(P)) :
        for k in range(len(Q)) :
            R[j+k] += P[j] + Q[k]
    return R

print(produit(P,Q))