#naked returns

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func concstr(a, b string) string{
	concat := a+b
	return concat
}


func main() {
	fmt.Println(split(20))
	fmt.Println(concstr("god", "zilla"))
}
