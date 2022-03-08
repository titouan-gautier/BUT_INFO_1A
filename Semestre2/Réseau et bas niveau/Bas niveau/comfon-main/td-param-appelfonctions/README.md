Cette séance est adaptée du chapitre _Go Assembly_ du livre _go-internals_  écrit par Clement Rey et [disponible en ligne](https://cmc.gitbook.io/go-internals) sous licence [BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/).


# Assembleur

Un assembleur est un langage textuel, qui permet de représenter le jeu d'instruction d'une machine.
C'est l'un des éléments de l'architecture d'une machine, avec les registres, et le système mémoire.

Un assembleur est spécifique à une machine donnée, mais l'ensemble des assembleurs utilisent des concepts communs.

En assembleur, une instruction est composée d'une opération et de zéro ou plusieurs opérandes.

## Opérations 

L'opération correspond à une fonction matérielle de la machine. On trouve :

  - des opérations arithmétiques et logiques, exécutées par l'UAL
  - des opérations de copie de données, entre la mémoire et les registres, ou entre deux registre
  - des opérations de branchement (parfois appelées opérations de saut), qui peuvent être conditionnelles ou non, et qui permettent de traduire les conditionnelles, les boucles, et les appels de sous-programme.
  
## Opérandes 

Les opérandes peuvent être :

  - des constantes (on parle également de valeur immédiate),
  - des adresses (c'est alors la valeur contenue en mémoire à cette adresse qui est l'opérande),
  - des registres (c'est alors la valeur contenue dans le registre qui est l'opérande),
  - ou des expressions qui permettent de calculer une adresse (c'est alors la valeur contenue en mémoire à l'adresse résultat du calcul qui est l'opérande).
  
# L'assembleur du langage Go 

Le langage Go propose un pseudo-assembleur qui est une forme d'assembleur portable qui abstrait certains mécanismes comme les copies de données ou l'appel de sous-programme, et propose également des registres virtuels.
Ce pseudo-assembleur est la cible du compilateur Go. 
Ensuite, une passe de génération de code permet de traduire le pseudo-assembleur vers le langage machine de la cible.

## Les registres virtuels SP, BP et SB 

Le registre virtuel `SP` est le pointeur de pile, il indique l'adresse du sommet de la pile.
Comme d'habitude, la pile croît vers les adresses basses.

Le registre virtuel `BP` est le pointeur de cadre, il indique l'adresse du début du cadre au sommet de la pile.

Enfin, le registre virtuel `SB` indique l'adresse de début de l'espace d'adressage du programme. Il est utilisé entre autre pour localiser les objets alloués dans les segments de texte et de données des programmes.


# ABI 

L'ABI (_Application Binary Interface_) d'un langage de programmation décrit l'implémentation des concepts du langage pour une architecture donnée.
Ce document précise ainsi, entre autres :

  - la représentation binaire des différents type de donnée : taille, alignement, codage ;
  - le protocole d'appel des sous-programmes.
  
Ces informations sont importantes car elles permettent d'assurer l'interopérabilité entre des composants binaires construits avec des chaînes de développement différentes.
En particulier, les chaînes doivent engendrer des codes conformes à l'ABI.

# L'appel de sous-programme en Go 

L'appel de sous-programme en Go repose sur l'utilisation de la pile. 
En particulier, l'appelant d'un sous-programme doit préparer la pile comme suit :

  - allouer de la place sur la pile pour la ou les valeurs de retour du sous-programme appelé ;
  - puis allouer de la place sur la pile pour les paramètres et y copier les valeurs des paramètres réels pour cet appel ;
  
  L'appel proprement dit peut alors être réalisé en utilisant la pseudo-opération `CALL`. Cela a pour effet d'empiler l'adresse de retour (c'est-à-dire l'adresse de l'instruction qui suit `CALL` dans le code de l'appelant), et de copier l'adresse de début du code de la fonction appelée dans le pointeur d'instruction.
  
  De son côté, le sous-programme appelé doit rendre la pile à l'appelant dans l'état qu'elle avait avant l'exécution de l'instruction `CALL`, à l'exception des mots alloués pour stocker les valeurs de retours, qui ont pu être modifiés par l'appelé.
  
  Le sous-programme appelé se termine par la pseudo-opération `RET`, qui dépile la valeur ai sommet de la pile et la copie dans le pointeur d'instruction.
  
# Dissection d'un programme simple.

Soit le programme ci-dessous.

``` go
package main

//go:noinline
func add(a, b int32) (int32, bool) {
	return a+b, true
}

func main() {
	add(10, 32)
}
```

Il est possible d'obtenir le pseudo-assembleur engendré par le compilateur en utilisant l'instruction suivante :

``` sh
GOOS=linux GOARCH=amd64 go tool compile -S add.go > compile.log
```

Le fichier `compile.log` contient alors le texte suivant 

``` text
"".add STEXT nosplit size=20 args=0x10 locals=0x0 funcid=0x0
	0x0000 00000 (add.go:4)	TEXT	"".add(SB), NOSPLIT|ABIInternal, $0-16
	0x0000 00000 (add.go:4)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (add.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (add.go:5)	MOVL	"".b+12(SP), AX
	0x0004 00004 (add.go:5)	MOVL	"".a+8(SP), CX
	0x0008 00008 (add.go:5)	ADDL	CX, AX
	0x000a 00010 (add.go:5)	MOVL	AX, "".~r2+16(SP)
	0x000e 00014 (add.go:5)	MOVB	$1, "".~r3+20(SP)
	0x0013 00019 (add.go:5)	RET
	0x0000 8b 44 24 0c 8b 4c 24 08 01 c8 89 44 24 10 c6 44  .D$..L$....D$..D
	0x0010 24 14 01 c3                                      $...
"".main STEXT size=66 args=0x0 locals=0x18 funcid=0x0
	0x0000 00000 (add.go:8)	TEXT	"".main(SB), ABIInternal, $24-0
	0x0000 00000 (add.go:8)	MOVQ	(TLS), CX
	0x0009 00009 (add.go:8)	CMPQ	SP, 16(CX)
	0x000d 00013 (add.go:8)	PCDATA	$0, $-2
	0x000d 00013 (add.go:8)	JLS	58
	0x000f 00015 (add.go:8)	PCDATA	$0, $-1
	0x000f 00015 (add.go:8)	SUBQ	$24, SP
	0x0013 00019 (add.go:8)	MOVQ	BP, 16(SP)
	0x0018 00024 (add.go:8)	LEAQ	16(SP), BP
	0x001d 00029 (add.go:8)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (add.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (add.go:9)	MOVQ	$137438953482, AX
	0x0027 00039 (add.go:9)	MOVQ	AX, (SP)
	0x002b 00043 (add.go:9)	PCDATA	$1, $0
	0x002b 00043 (add.go:9)	CALL	"".add(SB)
	0x0030 00048 (add.go:10)	MOVQ	16(SP), BP
	0x0035 00053 (add.go:10)	ADDQ	$24, SP
	0x0039 00057 (add.go:10)	RET
	0x003a 00058 (add.go:10)	NOP
	0x003a 00058 (add.go:8)	PCDATA	$1, $-1
	0x003a 00058 (add.go:8)	PCDATA	$0, $-2
	0x003a 00058 (add.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x003f 00063 (add.go:8)	PCDATA	$0, $-1
	0x003f 00063 (add.go:8)	NOP
	0x0040 00064 (add.go:8)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 2b 48  dH..%....H;a.v+H
	0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 b8 0a  ...H.l$.H.l$.H..
	0x0020 00 00 00 20 00 00 00 48 89 04 24 e8 00 00 00 00  ... ...H..$.....
	0x0030 48 8b 6c 24 10 48 83 c4 18 c3 e8 00 00 00 00 90  H.l$.H..........
	0x0040 eb be                                            ..
	rel 5+4 t=17 TLS+0
	rel 44+4 t=8 "".add+0
	rel 59+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
```

## Analyse du code de la fonction `main.add`

  1. Dessiner l'état que devrait avoir la pile en entrant dans la fonction.
  2. Exécuter à la main le code pseudo-assembleur de la fonction, instruction par instruction.
  
## Analyse du code de la fonction `main.main`

Dans un premier temps, s'intéresser au code entre les adresses `0x000f` et `0x0039`, et ne pas tenir compte des directives `FUNCDATA` et `PCDATA` qui sont à destination du ramasse-miettes.
Pour plus de facilité, voici le code correspondant.

``` text
	0x000f 00015 (add.go:8)	SUBQ	$24, SP
	0x0013 00019 (add.go:8)	MOVQ	BP, 16(SP)
	0x0018 00024 (add.go:8)	LEAQ	16(SP), BP
	0x001d 00029 (add.go:9)	MOVQ	$137438953482, AX
	0x0027 00039 (add.go:9)	MOVQ	AX, (SP)
	0x002b 00043 (add.go:9)	CALL	"".add(SB)
	0x0030 00048 (add.go:10)	MOVQ	16(SP), BP
	0x0035 00053 (add.go:10)	ADDQ	$24, SP
	0x0039 00057 (add.go:10)	RET
```

  3. Que font les trois premières instructions ?
  4. Que font les deux instructions suivantes ?
  5. Que font les trois dernières instructions ?
  
S'intéresser maintenant au code positionné en prologue et en épilogue de la fonction par le compilateur :

``` text
	0x0000 00000 (add.go:8)	MOVQ	(TLS), CX
	0x0009 00009 (add.go:8)	CMPQ	SP, 16(CX)
	0x000d 00013 (add.go:8)	PCDATA	$0, $-2
	0x000d 00013 (add.go:8)	JLS	58
    [...]
	0x003a 00058 (add.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x003f 00063 (add.go:8)	PCDATA	$0, $-1
	0x003f 00063 (add.go:8)	NOP
	0x0040 00064 (add.go:8)	JMP	0

```

  6. Que fait ce code ?
