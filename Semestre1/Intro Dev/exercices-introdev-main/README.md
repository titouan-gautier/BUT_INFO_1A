# Exercices introdev

Collection d'exercices pour le cours d'introduction au développement au département informatique de l'IUT de Nantes.

## Instructions pour récupérer les exercices

Vous pouvez récupérer tout le contenu de ce dépôt avec la commande suivante : git clone https://gitlab.univ-nantes.fr/jezequel-l/exercices-introdev.git

Les semaines suivantes vous pourrez le mettre à jour avec les nouveaux exercices simplement en vous plaçant dans le dossier et en utilisant la commande suivante : git pull

**Attention, si vous modifiez le contenu du dossier, la commande git pull va vous signaler des erreurs, vous verrez à quoi cela correspond un peu plus tard. Ne modifiez donc pas le contenu du dossier.**

## Instructions pour faire les exercices

Le dossier que vous aurez récupéré contient plusieurs dossiers (basiques, recherche, etc) qui eux-même contiennent des sous-dossiers (dans basiques il y a par exemple facteurspremiers, monnaie, etc). Chacun de ces sous-dossiers est un exercice.

Chaque exercice contient trois fichiers :
- go.mod, décrivant un module Go
- xxx_test.go, décrivant un jeu de tests
- xxx.go, l'exercice en lui même

Pour résoudre un exercice vous devez coder la fonction Go qui est décrite dans le fichier de l'exercice et vérifier qu'elle passe bien tous les tests.

**Ne faites pas cela dans le dossier récupéré par git clone, copiez l'exercice ailleurs, sinon ça vous causera des problèmes pour récupérer les nouveaux exercices durant les prochaines semaines.**

Vous pouvez faire les exercices dans l'ordre que vous voulez. Dans chaque dossier vous trouverez des exercices faciles et d'autres plus difficiles, ne restez pas bloqués sur un exercice : passez à un autre et revenez-y plus tard.

## Historique

### 17 septembre 2021

#### basiques
- facteurspremiers
- factorielle
- monnaie
- monnaie2
- monnaie3
- nombreparfait
- nombrepremier
- nombrespairs
- nombrespremiers
- stevej

#### recherche
- matrice
- maxoccurrences
- occurrences
- occurrencesmax
- prefixes
- tabtab

### 27 septembre 2021

#### basiques
- ensemble

#### recherche
- nombrespremiers
- nombrespremiers2

#### recursivite
- chainesbinaires
- chocolats
- palindrome
- sousensembles
- sousensembles2

### 1 octobre 2021

#### programmation dynamique
- piecesjaunes
- souschaine

#### recherche
- classement

#### recursivite
- huitreines
- huitreines2
- piecesjaunes
- racaman
- souschaine

### 18 octobre 2021

#### recherche
- doublons1
- doublons3
- doublons4
- doublons5

#### tri
- alphabetique
- bienrange
- decroissant
- valabs

### 2021-2022 Test machine 1

#### basiques
- alphabet
- calcul
- somme

#### pointeurs
- double
- varswitch

#### recherche
- compte (remplace occurrences, qui est similaire)
- doublons2
- puissant
- recherche

#### recursivite
- suite

### 2021-2022 Test machine 2

basiques devient divers

#### divers
- chiffres

#### pointeurs
- add
- copy

#### recherche
- egalite
- recherche2

#### recursivite
- conway
- ppcm
- syracuse

#### tri
- tri
- tri2

### 23 novembre 2021

#### fichiers
- acrostiche
- existe
- lignes
- lignes2

#### recherche
- egalite2

#### recursivite
- pgcd
- pgcd2
- ppcm2
