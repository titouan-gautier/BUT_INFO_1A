package nombrepremier

func estPremier(n int) (prem bool) {
	if n > 1 {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
