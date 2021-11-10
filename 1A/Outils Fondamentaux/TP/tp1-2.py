import numpy as np
import matplotlib.pyplot as plt

x=np.array([1,3,5,7,10,13])
y=np.array([3,2,0,1,-4,6])

print(plt.plot(x,y, linestyle= 'dashed'))
print(plt.show())
#Créer le graphique des points avec x abscisse et y ordonné
#— color = ’r’ : change la couleur de la ligne en rouge
#— linestyle = ’dashed’ : ache la ligne en pointillés
#— marker = ’o’ : place un cercle sur chaque point dessiné