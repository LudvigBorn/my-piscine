package test

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}
