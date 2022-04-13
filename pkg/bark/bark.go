package bark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BarkMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// push msg to bark server
func PushToBark(msg *BarkMessage) {
	b, _ := json.Marshal(msg)
	resp, _ := http.Post(BarkPushUrl, "application/json", bytes.NewBuffer(b))
	//
	defer resp.Body.Close()
	//io.Reader

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
