package error

func VerifyError(err error) {
	if err != nil {
		panic(err)
	}
}
