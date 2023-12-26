package iot

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateSpeech invokes the iot.CreateSpeech API synchronously
func (client *Client) CreateSpeech(request *CreateSpeechRequest) (response *CreateSpeechResponse, err error) {
	response = CreateCreateSpeechResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSpeechWithChan invokes the iot.CreateSpeech API asynchronously
func (client *Client) CreateSpeechWithChan(request *CreateSpeechRequest) (<-chan *CreateSpeechResponse, <-chan error) {
	responseChan := make(chan *CreateSpeechResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSpeech(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateSpeechWithCallback invokes the iot.CreateSpeech API asynchronously
func (client *Client) CreateSpeechWithCallback(request *CreateSpeechRequest, callback func(response *CreateSpeechResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSpeechResponse
		var err error
		defer close(result)
		response, err = client.CreateSpeech(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateSpeechRequest is the request struct for api CreateSpeech
type CreateSpeechRequest struct {
	*requests.RpcRequest
	Voice           string           `position:"Body" name:"Voice"`
	ProjectCode     string           `position:"Body" name:"ProjectCode"`
	AudioFormat     string           `position:"Body" name:"AudioFormat"`
	IotInstanceId   string           `position:"Body" name:"IotInstanceId"`
	Text            string           `position:"Body" name:"Text"`
	SoundCodeConfig string           `position:"Body" name:"SoundCodeConfig"`
	SpeechType      string           `position:"Body" name:"SpeechType"`
	EnableSoundCode requests.Boolean `position:"Body" name:"EnableSoundCode"`
	Volume          requests.Integer `position:"Body" name:"Volume"`
	BizCode         string           `position:"Body" name:"BizCode"`
	ApiProduct      string           `position:"Body" name:"ApiProduct"`
	ApiRevision     string           `position:"Body" name:"ApiRevision"`
	SpeechRate      requests.Integer `position:"Body" name:"SpeechRate"`
}

// CreateSpeechResponse is the response struct for api CreateSpeech
type CreateSpeechResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         string `json:"Data" xml:"Data"`
}

// CreateCreateSpeechRequest creates a request to invoke CreateSpeech API
func CreateCreateSpeechRequest() (request *CreateSpeechRequest) {
	request = &CreateSpeechRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateSpeech", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateSpeechResponse creates a response to parse from CreateSpeech response
func CreateCreateSpeechResponse() (response *CreateSpeechResponse) {
	response = &CreateSpeechResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
