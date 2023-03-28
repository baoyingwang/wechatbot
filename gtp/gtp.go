package gtp

import (
	"bytes"
	"encoding/json"
	"github.com/869413421/wechatbot/config"
	"io/ioutil"
	"log"
	"net/http"
)

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string                   `json:"id"`
	Object  string                   `json:"object"`
	Created int                      `json:"created"`
	Model   string                   `json:"model"`
	Choices []Choice                 `json:"choices"`
	Usage   map[string]interface{}   `json:"usage"`
}

type ChoiceItem struct {
}

// ChatGPTRequestBody 响应体
type ChatGPTRequestBody struct {
	Model            string  `json:"model"`
	Messesges      []Message `json:"messages"`
	//MaxTokens        int     `json:"max_tokens"`
	//Temperature      float32 `json:"temperature"`
	//TopP             int     `json:"top_p"`
	//FrequencyPenalty int     `json:"frequency_penalty"`
	//PresencePenalty  int     `json:"presence_penalty"`
}

type Choice struct {
	Index          int     `json:"index"`
	FinishReason   string  `json:"finish_reason"`
	Messesge       Message `json:"message"`
}

type Message struct {
	Role            string  `json:"role"`
	Content         string  `json:"content"`
}

// Completions gtp文本模型回复
//curl https://api.openai.com/v1/completions
//-H "Content-Type: application/json"
//-H "Authorization: Bearer your chatGPT key"
//-d '{"model": "text-davinci-003", "prompt": "give me good song", "temperature": 0, "max_tokens": 7}'
func Completions(msg string) (string, error) {

	message := Message{
		Role:            "user",
		Content:         msg,
	}
	messages := []Message{ message }
	requestBody := ChatGPTRequestBody{
		Model:            config.LoadConfig().Model,
		Messesges:        messages,
		//MaxTokens:        2048,
		//Temperature:      0.7,
		//TopP:             1,
		//FrequencyPenalty: 0,
		//PresencePenalty:  0,
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	url := config.LoadConfig().Endpoint+"openai/deployments/"+config.LoadConfig().Model+"/chat/completions?api-version="+config.LoadConfig().ApiVersion
	log.Printf("request gtp json string : %v", string(requestData))
	log.Printf("request url : %v", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	apiKey := config.LoadConfig().ApiKey
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", apiKey)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}
	var reply string
	if len(gptResponseBody.Choices) > 0 {
		for _, choice := range gptResponseBody.Choices {
			reply = choice.Messesge.Content
			break
		}
	}
	log.Printf("gpt response text: %s \n", reply)
	return reply, nil
}
