package exception

type NotFound struct {
	Err string
}

func NewNotFounErr(error string) NotFound {
	return NotFound{
		Err: error,
	}
}
