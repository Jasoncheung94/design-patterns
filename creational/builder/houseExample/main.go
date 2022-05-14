package main

func main() {
	// go run *.go to build the example due to all files in the main package
	standardBuilder := getBuilder("house")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(standardBuilder)
	normalHouse := director.construct()
	normalHouse.printDetails()

	director.setBuilder(iglooBuilder)
	iglooHouse := director.construct()
	iglooHouse.printDetails()
}
