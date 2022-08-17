package slack

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"notifications/errors"
	"os"
)

type JsonInfo struct {
	HookUrl string
}

// Load template information fr
func (ji *JsonInfo) setJsonInfo() error {
	var present bool

	if ji.HookUrl, present = os.LookupEnv("HOOK_URL"); !present {
		return &errors.EnvVarNotDefined{Name: "HOOK_URL"}
	}

	return nil
}

// Send Slack notification
func (ji *JsonInfo) SendJsonNotification() error {

	var tplBuffer bytes.Buffer

	//parse json tempalte
	var jmt JsonMsgTemplate

	err := jmt.ParseJsonMsgTemplate(&tplBuffer)

	if err != nil {
		log.Fatalln(err)
	}

	// set json info
	err = ji.setJsonInfo()
	if err != nil {
		return err
	}

	//Slack request
	request, err := http.NewRequest("POST", ji.HookUrl, &tplBuffer)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	log.Println("response Status:", response.Status)
	body, _ := ioutil.ReadAll(response.Body)
	log.Println("response Body:", string(body))

	log.Println("Slack message delivered.")

	return nil
}
