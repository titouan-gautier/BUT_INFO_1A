import numpy as np
import matplotlib.pyplot as plt

print(np.array([3,7,-1,2]))
#créer le tableau
print(np.ones(5))
# écrit n fois 1
print(np.array([3,7,-1,2])+np.ones(4))
#ajoute 1 à chaque élément du tableau(avec virgule)
print(np.array([3,7,-1,2])+1)
#ajoute 1 à chaque élément du tableau(sans virgule)
print(np.array([[3,7],[-1,2]]))
#créer un tableau avec deux tableau
print(np.array([1,2,3,4])**2)
#calcul le carré de chaque élément du tableau
print(np.arange(1,5,1))
#créer le tableau de n à  y exclus avec une marche de z (limite gauche, limite droite,marche)
print(np.arange(10,30,5))
#ici, créer le tableau de 10 à 30 exclus avec une marche de 5
print(np.linspace(0,2,9))
#Créer un tableau de n à y inclus de z valeur d'écart identique
print(np.sin(np.linspace(0,2*np.pi,20)))
#Ici, Créer le tableau de valeur de O à 2*pi du sinus de 20 valeur d'écart identique

x=np.array([1,3,5,7,10,13])
y=np.array([3,2,0,1,-4,6])

print(plt.plot(x,y))
print(plt.show())