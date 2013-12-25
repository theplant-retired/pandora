package pandora

import (
	"encoding/json"

	zmq "github.com/alecthomas/gozmq"
)

var Address = "tcp://127.0.0.1:9999"
var socket *zmq.Socket

const (
	FLAG_HTML     = "html"
	FLAG_MARKDOWN = "markdown"
)

type Input struct {
	Text string `json:"text"`
	From string `json:"from"`
	To   string `json:"to"`
}

func init() {
	context, _ := zmq.NewContext()
	socket, _ = context.NewSocket(zmq.REQ)
	err := socket.Connect(Address)
	if err != nil {
		panic("Please make sure the pandora daemon is running")
	}
}

func ToMarkdown(html string) (r string, err error) {
	data := &Input{
		Text: html,
		From: FLAG_HTML,
		To:   FLAG_MARKDOWN,
	}

	return send(data)
}

func ToHTML(md string) (r string, err error) {
	data := &Input{
		Text: md,
		From: FLAG_MARKDOWN,
		To:   FLAG_HTML,
	}

	return send(data)
}

func send(data interface{}) (r string, err error) {
	var sendBytes []byte
	sendBytes, err = json.Marshal(data)
	if err != nil {
		return
	}

	err = socket.Send(sendBytes, 0)
	if err != nil {
		return
	}

	var b []byte
	b, err = socket.Recv(0)
	if err != nil {
		return
	}
	r = string(b)
	return
}
