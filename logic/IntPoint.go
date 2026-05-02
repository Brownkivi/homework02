package logic

func ModifyInt(a *int) {
	*a += 10
}

func ModifySlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}
