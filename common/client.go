package common

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Agent struct {
  ClientConfig *ClientConfig
  HttpClient *http.Client
  Logger *logrus.Logger
}

func (agent *Agent) Call(
  ctx context.Context, 
  method,
  url string, 
  body any,
  response any,
) error {
  fullUrl := BASE_URL + url
  
  agent.Logger.Infoln(fullUrl)

  var bodyBuffer bytes.Buffer
  err := json.NewEncoder(&bodyBuffer).Encode(body)

  req, err := http.NewRequest(method, fullUrl, &bodyBuffer)

  if err != nil {

  }

  req.WithContext(ctx)
  req.SetBasicAuth(agent.ClientConfig.ApiKey, "")
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Accept", "application/json")

  res, err := agent.HttpClient.Do(req)

  if err != nil {

  }

  defer res.Body.Close()

  agent.Logger.Infoln(res.Body)

  // decode body
  if res.StatusCode >= 200 && res.StatusCode <= 209 {
    bodyByte, err := ioutil.ReadAll(res.Body)
    if err != nil {
      agent.Logger.Infoln(err)
    }
    json.Unmarshal(bodyByte, &response)
    // json.NewDecoder(res.Body).Decode(&response)
  } else if res.StatusCode >= 400 && res.StatusCode <= 409 {

  } else if res.StatusCode >= 500 && res.StatusCode <= 509 {

  }

  return nil
}

func NewAgent(clientConfig *ClientConfig) *Agent {
  return &Agent{
    ClientConfig: clientConfig,
    HttpClient: &http.Client{},
    Logger: logrus.New(),
  }
}
