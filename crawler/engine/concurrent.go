package engine

type ConcurrentEngine struct {

}

type Scheduler interface {
	Submit(Request)
}

func (ConcurrentEngine) run (seeds ... Request)  {
	for _, r := range seeds {
		schedule.Submit(r)
	}
}