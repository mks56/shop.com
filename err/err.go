package err

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintError(err error) {
	if err != nil {
		panic(err)
	}
}
