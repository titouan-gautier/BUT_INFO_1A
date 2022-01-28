import graphviz

graphe1 = [[0,1,1],
          [1,0,1],
          [1,0,0]]

def creation_graphe_graphviz(g):
    # création d'un graphe orienté
    dot = graphviz.Digraph()
    
    # ajout des sommets
    for i in range(len(g)):
        dot.node(str(i))
    
    # ajout des arcs
    for i in range(len(g)):
        for j in range(len(g)):
            if g[i][j] == 1:
                dot.edge(str(i), str(j))
    
    return dot

gdot1 = creation_graphe_graphviz(graphe1)
print(gdot1.source)