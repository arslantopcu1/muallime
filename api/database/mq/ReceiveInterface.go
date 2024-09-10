package mq

type ReceiveInterface interface {
	Consumer()
}

func BeginListenMq(mq ReceiveInterface) {
	// start a thread for listen
	go mq.Consumer()
}
