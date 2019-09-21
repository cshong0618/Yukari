package yukari

import "errors"

type YukariEvent struct {
	Topic string
	Data []byte
	ContentType string
	Receivers []string

	HttpCaller HttpCaller
}

func (ye *YukariEvent) Announce() error {
	ints := make(chan int, len(ye.Receivers))
	for _, receiver := range ye.Receivers {
		go ye.call(receiver, ye.ContentType, ye.Data, ints)
	}

	var err error

	for i := 0; i < len(ye.Receivers); i++ {
		responseCode := <- ints
		if responseCode / 100 != 2 {
			err = errors.New("")
		}
	}

	return err
}

func (ye YukariEvent) call(url string, contentType string, data []byte, ch chan int) {
	responseCode := ye.HttpCaller.POST(url, contentType, data)
	ch <- responseCode
}

func CreateYukariEvent(topic string,
	data []byte, contentType string,
	receivers []string, caller HttpCaller) Event {
		return &YukariEvent{
			Topic:       topic,
			Data:        data,
			ContentType: contentType,
			Receivers:   receivers,
			HttpCaller:  caller,
		}
}