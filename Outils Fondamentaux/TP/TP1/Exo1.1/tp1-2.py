import numpy as np
import matplotlib.pyplot as plt

#x=np.array([1,3,5,7,10,13])
#y=np.array([3,2,0,1,-4,6])

#print(plt.plot(x,y, linestyle= 'dashed'))
#print(plt.show())

#Créer le graphique des points avec x abscisse et y ordonné

#option
#— color = ’r’ : change la couleur de la ligne en rouge
#— linestyle = ’dashed’ : affche la ligne en pointillés
#— marker = ’o’ : place un cercle sur chaque point dessiné

#Legendes
#— plt.legend(["fonction 1","fonction 2"]) : pour préciser les fonctions tracées
#— plt.ylabel("unit ́e ord.") : pour un label sur l’axe des ordonnées
#— plt.xlabel("unit ́e abs.") : pour un label sur l’axe des abscisses
#— plt.title("titre graphique") : pour un titre au graphique

x = np.linspace(0,3*np.pi,100)
f = x**2*np.sin(x)+4
g = 30 / (x**2+1)

plt.plot(x,f)
plt.plot(x,g)
plt.show()
plt.clf()
#supprime la fenetre graphique courante
plt.scatter(x,f)
#Met en forme la foction sous forme de points
plt.show()
