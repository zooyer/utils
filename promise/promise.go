package promise

type Promise struct {
	done    chan struct{}
	state   int32
	resolve []interface{}
	reject  []interface{}
}

func New(fn func(resolve, reject func(v ...interface{}))) *Promise {
	var p = &Promise{
		done:    make(chan struct{}),
		state:   0,
		resolve: nil,
		reject:  nil,
	}
	go func() {
		defer close(p.done)
		fn(func(v ...interface{}) {
			p.state = 1
			p.resolve = v
		}, func(v ...interface{}) {
			p.state = 2
			p.reject = v
		})
	}()
	return p
}

func (p *Promise) Done() *Promise {
	<-p.done
	return p
}

func (p *Promise) Then(onFulfilled, onRejected func(v ...interface{})) *Promise {
	p.Done()
	if p.state == 1 && onFulfilled != nil {
		onFulfilled(p.resolve)
	}
	if p.state == 2 && onRejected != nil {
		onRejected(p.reject)
	}
	return p
}

func (p *Promise) Catch(onRejected func(v ...interface{})) *Promise {
	p.Done()
	if p.state == 2 && onRejected != nil {
		onRejected(p.reject)
	}
	return p
}
