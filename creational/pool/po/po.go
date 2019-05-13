package po

type Pool chan *interface{}

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(interface{})
	}

	return &p
}
