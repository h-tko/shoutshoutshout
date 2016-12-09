package libraries

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BLUEMIX_CONVERSATION_URL  string = "https://gateway.watsonplatform.net/conversation/api/v1/workspaces/%s/message?version=%s"
	CONVERSATION_WORKSPACE_ID string = "xxxxxxxxxxxxx"
	CONVERSATION_USERNAME     string = "xxxxxxxxxxxxx"
	CONVERSATION_PASSWORD     string = "xxxxxxxxxxxx"
	VERSION                   string = "yyyy-mm-dd"
)

type Watson struct {
}

type resdata struct {
	Output struct {
		Text []string `json:"text"`
	} `json:"output"`
}

func NewWatson() *Watson {
	return new(Watson)
}

func (this *Watson) ConversationApi(send_message string) string {

	jsonParam := fmt.Sprintf("{\"input\": {\"text\": \"%s\"}}", send_message)

	client := &http.Client{}

	req, _ := http.NewRequest("POST", fmt.Sprintf(BLUEMIX_CONVERSATION_URL, CONVERSATION_WORKSPACE_ID, VERSION), bytes.NewBuffer([]byte(jsonParam)))

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(CONVERSATION_USERNAME, CONVERSATION_PASSWORD)

	res, err := client.Do(req)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		print("%v", res)
	}

	output := resdata{}
	json.Unmarshal(body, &output)

	return output.Output.Text[0]
}
