package main

func nest1(i int) int {
	one := 1
	i += 1
	nest2(1, 2)
	return one
}

func nest2(i1, i2 int) []int {
	ret := make([]int, 0)

	ret = append(ret, 1)
	ret = append(ret, 2)

	i1 += 1
	i2 += 1

	return ret
}

func main() {
	k := 100

	for i := 1; i <= 10; i++ {
		for j := i; j <= 10; j++ {
			k += i
		}
	}

	_ = nest1(1)
}
