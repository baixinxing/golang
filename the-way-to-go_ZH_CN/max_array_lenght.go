package main

//100M
func main() {
	var arr [63 * 1024 * 1024]int
	for i := 0; i < 63*1024*1024; i++ {
		arr[i] = i
	}

	if arr[0] != nil {

	}
}
