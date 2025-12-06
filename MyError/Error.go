package err

type ErrorInfo struct {
	Code    int
	Sussecc bool
	Result  string
}

func(e *ErrorInfo) Error() string {
	return e.Result
}
