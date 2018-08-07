package util

//public method: uppercase first letter
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
