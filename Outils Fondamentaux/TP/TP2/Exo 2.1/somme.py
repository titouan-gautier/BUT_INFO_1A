P = [0,2,1,1]
Q = [0,4,-1,-1]
R = []
i = 0
def somme(P,Q) :
    if len(P) >= len(Q) :
        max = P
        min = Q
    else :
        max = Q
        min = P
    for i in range(len(min)) :
        R.append(max[i]+min[i])
    R.extend(max[i+1:])
    print(R)
    for i in range(len(R)) :
        if R[-1] == 0 :
            R.pop(-1)
    return R

print(somme(P,Q))