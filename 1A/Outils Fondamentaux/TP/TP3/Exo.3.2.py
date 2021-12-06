import numpy as np
import time
# print(np.zeros(7))
#Ecrit 7 zeros
# print(np.ones(6))
#Ecrit 6 un
# print(np.identity(3))
#Créer une matrice de 3 par 3 avec des 1 sur la diagonale
# print(np.random.rand(3,4))
#Créer une matrice avec des nombres aléatoire entre 0 et 1 de hauteur 3 et largeur 4
# print(np.random.randint(-4,3,size=(3,5),dtype=int))
#Créer une matrice de nombres entier aléatoire entre -4 et 3 de hauteur 3 et largeur 5

#print(a+b)
#A un sens, on peut les aditionner car elles sont de meme taille.
#print(a+4)
#A du sens, on ajoue 4 a chaque terme de la matrice.
#print(a*b)
#N'a pas de sens, multiplie chaque terme un a un(pas bonne methode)
#print(3*a)
#A du sens, on mltiplie par 3 chaque terme de la matrice
#print(np.add(a,b))
#A du sens, additionne la matrice a et b : a+b
#print(a.dot(b))
#A du sens, fait un réel produit de matrice
#print(a@b)
#A du sens, fait un réel produit de matrice
#print(np.linalg.matrix_power(a,2))
#A du sens, fait le réel produit de matrice par lui même, ici a@a
#print(np.linalg.inv(a))
#A du sens, nous donne l'inverse de la matrice a

#print(a.shape)
#A du sens, donne la hauteur et la largeur.

#print(a.sum())
#Fait la somme des terme de la matrice
#print(a.sum(axis=0))
#Fait la somme des terme en vertical
#print(a.sum(axis=1))
#Fait la somme des terme en horizontal
#print(a.min())
#Donne le plus petit terme de la matrice
#print(a.max())
#Donne le plus grand terme de la matrice
#print(a[0,1])
#Donne le terme se trouvant à la ligne 0 (première ligne en partant du haut) et à la case 1 
#print(a[0][1])
#Donne le terme se trouvant à la ligne 0 (première ligne en partant du haut) et à la case 1


# def somme(a) :
#     t1 = time.time()
#     res = 0
#     for i in a :
#         for j in i :
#             res = res + j
#     t2 = time.time() - t1
#     return res,t2

# def somme2(a) :
#     t1 = time.time()
#     res = a.sum()
#     t2 = time.time() - t1
#     return res, t2

# print(somme(a))
# print(somme2(a))

# = np.array([[1,2],[3,4],[0,2],[2,5]])
# = np.array([[1,2,3],[3,4,5]])

a = np.random.rand(1000,1000)
b = np.random.rand(1000,1000)

# 1 = colonne
# 0 = ligne


def produit(a,b) :
    t1 = time.time()
    if a.shape[1] == b.shape[0] :
        prod = np.zeros((a.shape[0],b.shape[1]))
        for i in range(prod.shape[0]) :
            for j in range(prod.shape[1]) :
                for k in range(a.shape[1]) :
                    prod[i][j] += a[i,k]*b[k,j]
        t2 = time.time() - t1
        return t2
    else :
        return "Produit impossible"

def produit2(a,b) :
    t1 = time.time()
    res= a@b
    t2 = time.time() - t1
    return t2


print(produit(a,b) / produit2(a,b))
