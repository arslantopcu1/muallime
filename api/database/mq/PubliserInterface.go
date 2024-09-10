package mq

type PubliserInterface interface {
	Publish() error
}

func Send(mq PubliserInterface) error {

	// puhlish message
	return mq.Publish()

}
