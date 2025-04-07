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
