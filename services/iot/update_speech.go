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

// UpdateSpeech invokes the iot.UpdateSpeech API synchronously
func (client *Client) UpdateSpeech(request *UpdateSpeechRequest) (response *UpdateSpeechResponse, err error) {
	response = CreateUpdateSpeechResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSpeechWithChan invokes the iot.UpdateSpeech API asynchronously
func (client *Client) UpdateSpeechWithChan(request *UpdateSpeechRequest) (<-chan *UpdateSpeechResponse, <-chan error) {
	responseChan := make(chan *UpdateSpeechResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSpeech(request)
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

// UpdateSpeechWithCallback invokes the iot.UpdateSpeech API asynchronously
func (client *Client) UpdateSpeechWithCallback(request *UpdateSpeechRequest, callback func(response *UpdateSpeechResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSpeechResponse
		var err error
		defer close(result)
		response, err = client.UpdateSpeech(request)
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

// UpdateSpeechRequest is the request struct for api UpdateSpeech
type UpdateSpeechRequest struct {
	*requests.RpcRequest
	Voice           string           `position:"Body" name:"Voice"`
	ProjectCode     string           `position:"Body" name:"ProjectCode"`
	IotInstanceId   string           `position:"Body" name:"IotInstanceId"`
	SoundCodeConfig string           `position:"Body" name:"SoundCodeConfig"`
	EnableSoundCode requests.Boolean `position:"Body" name:"EnableSoundCode"`
	Volume          requests.Integer `position:"Body" name:"Volume"`
	ApiProduct      string           `position:"Body" name:"ApiProduct"`
	ApiRevision     string           `position:"Body" name:"ApiRevision"`
	SpeechRate      requests.Integer `position:"Body" name:"SpeechRate"`
	SpeechCode      string           `position:"Body" name:"SpeechCode"`
}

// UpdateSpeechResponse is the response struct for api UpdateSpeech
type UpdateSpeechResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateSpeechRequest creates a request to invoke UpdateSpeech API
func CreateUpdateSpeechRequest() (request *UpdateSpeechRequest) {
	request = &UpdateSpeechRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateSpeech", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateSpeechResponse creates a response to parse from UpdateSpeech response
func CreateUpdateSpeechResponse() (response *UpdateSpeechResponse) {
	response = &UpdateSpeechResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
