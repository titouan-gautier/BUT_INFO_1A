package stevej

import (
	"errors"
)

var errPasUneDate error = errors.New("la date indiquée n'est pas valide")

/*
Steve Jobs est né le 24 février 1955. La fonction ecart devra indiquer le nombre
de jours de différence entre une date (à partir du 1er janvier 1900 jusqu'au 31
décembre 2021) et la date de naissance de Steve Jobs. L'écart en jours sera toujours
un nombre positif ou nul (on ne se préoccupera pas de savoir si la date est antérieure
ou postérieure à la naissance de steve Jobs).

# Entrées
- j : un numéro de jour
- m : un numéro de mois
- a : un numéro d'année

# Sorties
- ej : l'écart en jours entre j/m/a et le 24/2/1955
- err : une erreur qui vaudra errPasUneDate quand j/m/a n'est pas une date valide et nil sinon

# Exemple
ecart(17, 4, 1986) = 11375
*/
func ecart(j, m, a uint) (ej uint, err error) {
	if a < 1900 || a > 2021 || m > 12 || m == 0 || j > 31 {
		return 0, errPasUneDate
	}
	if (m == 2 || m == 4 || m == 6 || m == 9 || m == 11) && j > 30 {
		return 0, errPasUneDate
	}
	if m == 2 {
		if j > 28 && a%4 != 0 {
			return 0, errPasUneDate
		}
		if j > 29 && a%4 == 0 {
			if a%100 != 0 {
				return 0, errPasUneDate
			}
		}
	} else {
		ej = calc(j, m, a)
	}
	return ej, nil
}

func calc(j, m, a uint) (somme uint) {
	var somme1 uint
	var somme2 uint
	var i uint
	var k uint

	//Calcul mois
	if m < 2 {
		for i = m; i <= 2; i++ {
			if i == 2 {
				if a%4 == 0 {
					somme1 = somme1 + 29
				} else {
					somme1 = somme1 + 28
				}
			} else {
				if i == 1 || i == 3 || i == 5 || i == 7 || i == 8 || i == 10 || i == 12 {
					somme1 = somme1 + 31
				} else {
					somme1 = somme1 + 30
				}
			}
		}
	} else {
		for i = 1; i < m; i++ {
			if i == 2 {
				if a%4 == 0 {
					somme1 = somme1 + 29
				} else {
					somme1 = somme1 + 28
				}
			} else {
				if i == 1 || i == 3 || i == 5 || i == 7 || i == 8 || i == 10 || i == 12 {
					somme1 = somme1 + 31
				} else {
					somme1 = somme1 + 30
				}

			}
		}
	}

	//Calcul année
	if a != 1955 {
		if a < 1955 {
			for k = a; k <= 1955; k++ {
				if k%4 == 0 {
					somme2 = somme2 + 366
				} else {
					somme2 = somme2 + 365
				}
			}
		} else {
			for k = 1955; k < a; k++ {
				if k%4 == 0 {
					somme2 = somme2 + 366
				} else {
					somme2 = somme2 + 365
				}
			}
		}
	} else {
		somme2 = 0
	}

	somme = somme2 + somme1 - j
	somme = somme - 7
	return somme
}
