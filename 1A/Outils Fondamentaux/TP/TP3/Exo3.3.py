import numpy as np

# 1 = colonne
# 0 = ligne

A = np.array([[1,1,-6],[1,2,0],[-6,1,1]])


def recherche_pivot(A) :
    print(A)
    ligne = 0
    comp = 0
    pivot = 0
    for i in range(A.shape[0]) :
        pivot = i
        j = i
        for j in range(A.shape[1]) :
            if A[i][j] < 0 :
                comp = A[i][j]*-1
            if comp > A[i][i] :
                pivot = comp
                ligne = j
    A[j],A[i] = A[i],A[j]
    print(A)

print(recherche_pivot(A))