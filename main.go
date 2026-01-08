package main

func main() {
	_, err := linkXG()
	if err != nil {
		panic(err)
	}
	_, err = linkMG()
	if err != nil {
		panic(err)
	}

}
