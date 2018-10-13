package tools

func CheeckErr(err error) {
	if err != nil {
		panic(err)
	}
}
