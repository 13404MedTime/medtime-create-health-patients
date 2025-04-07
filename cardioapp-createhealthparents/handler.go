package function

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Datas struct {
	Data struct {
		Data struct {
			Data map[string]interface{} `json:"data"`
		} `json:"data"`
	} `json:"data"`
}

type ClientApiResponse struct {
	Data ClientApiData `json:"data"`
}

type ClientApiData struct {
	Data ClientApiResp `json:"data"`
}

type ClientApiResp struct {
	Response map[string]interface{} `json:"response"`
}

type Response struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

type HttpRequest struct {
	Method  string      `json:"method"`
	Path    string      `json:"path"`
	Headers http.Header `json:"headers"`
	Params  url.Values  `json:"params"`
	Body    []byte      `json:"body"`
}

type AuthData struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type NewRequestBody struct {
	RequestData HttpRequest            `json:"request_data"`
	Auth        AuthData               `json:"auth"`
	Data        map[string]interface{} `json:"data"`
}

type Request struct {
	Data map[string]interface{} `json:"data"`
}

type GetListClientApiResponse struct {
	Data GetListClientApiData `json:"data"`
}

type GetListClientApiData struct {
	Data GetListClientApiResp `json:"data"`
}

type GetListClientApiResp struct {
	Response []map[string]interface{} `json:"response"`
}

func DoRequest(url string, method string, body interface{}, appId string) ([]byte, error) {
	data, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Add("authorization", "API-KEY")
	request.Header.Add("X-API-KEY", "P-JV2nVIRUtgyPO5xRNeYll2mT4F5QG4bS")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func Send(text string) {
	bot, _ := tgbotapi.NewBotAPI("6241555505:AAHPpkXj-oHBGblWd_7O9kxc9a05tJUIFRw")
	msg := tgbotapi.NewMessage(1194897882, text)
	bot.Send(msg)
}

func Handle(req []byte) string {
	var response Response
	var request NewRequestBody
	const urlConst = "https://api.admin.u-code.io"

	if err := json.Unmarshal(req, &request); err != nil {
		return errorResponse("Error while unmarshalling request")
	}
	if request.Data["app_id"] == nil {
		return errorResponse("App id required")
	}
	appId := request.Data["app_id"].(string)

	var tableSlug = "patient_cards"
	clientId := ""
	if request.Data["guid"] != nil {
		clientId = request.Data["guid"].(string)
	} else if request.Data["object_data"] != nil {
		clientId = request.Data["object_data"].(map[string]interface{})["guid"].(string)
	}

	createReq := Request{Data: map[string]interface{}{"cleints_id": clientId}}
	_, err, response := CreateObject(urlConst, tableSlug, appId, createReq)
	if err != nil {
		respByte, _ := json.Marshal(response)
		return string(respByte)
	}

	response.Data = map[string]interface{}{}
	response.Status = "done"
	respByte, _ := json.Marshal(response)
	return string(respByte)
}

func errorResponse(msg string) string {
	response := Response{Status: "error", Data: map[string]interface{}{"message": msg}}
	respByte, _ := json.Marshal(response)
	return string(respByte)
}
