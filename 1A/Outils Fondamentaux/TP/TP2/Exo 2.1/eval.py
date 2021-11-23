P = [2,1,1]
x = 3
def eval(P,x) :
    r = 0
    for i in range(len(P)) :
        r = r + P[i]*x**i
    return r

print(eval(P,x))
