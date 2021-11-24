import numpy as np
import matplotlib.pyplot as plt

a = 1
b = 0
c = -4

def resoudre(a,b,c) :
    d = b**2-4*a*c
    if d > 0 :
        x1 = (-b-np.sqrt(d)) / (2*a)
        x2 = (-b+np.sqrt(d)) / (2*a)
        f1 = (a*x1**2 + b*x1 + c)
        f2 = (a*x2**2 + b*x2 + c)
        return x1,x2
    elif d == 0 :
        return [(-b/(2*a))]
    else :
        return []

def parabole(a,b,c) :
    x = np.linspace(-3+(-b/(2*a)),3+(-b/(2*a)),100)
    y = a*x**2 + b*x + c
    plt.plot(x,y)
    xval = (b/(2*a))
    yval = a*xval**2 + b*xval + c
    plt.scatter(xval,yval)
    for i in resoudre(a,b,c) :
        plt.scatter(i,0)
    plt.show()

parabole(a,b,c)