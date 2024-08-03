package ollama_client

import (
	"encoding/json"
	"net/http"
	"log"
	"bytes"
)

type Response struct {
	Model string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response string `json:"response"`
	Done bool `json:"done"`
	DoneReason string `json:"dont_reason"`
	Context []int `json:"context"`
	TotalDuration int `json:"total_duration"`
	LoadDuration int `json:"load_duration"`
	PromptEvalCount int `json:"promp_eval_count"`
	PromptEvalDuration int `json:"promp_eval_duration"`
	EvalCount int `json:"eval_count"`
	EvalDuration int `json:"eval_duration"`
}

type Request struct {
	Model string `json:"model"`
	Stream bool `json:"stream"`
	Prompt string `json:"prompt"`
}

func Generate(query string) string {
	url := "http://ollama:11434/api/generate"
	
	body := &Request{
		Model: "phi3",
		Stream: false,
		Prompt: query,
	}
	
	bodyJson, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("could not marshal JSON: %s", err)
	}

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	response := &Response{}
	derr := json.NewDecoder(res.Body).Decode(response)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}
	
	return response.Response
}