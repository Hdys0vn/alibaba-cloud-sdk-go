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

// GetShareSpeechModelAudio invokes the iot.GetShareSpeechModelAudio API synchronously
func (client *Client) GetShareSpeechModelAudio(request *GetShareSpeechModelAudioRequest) (response *GetShareSpeechModelAudioResponse, err error) {
	response = CreateGetShareSpeechModelAudioResponse()
	err = client.DoAction(request, response)
	return
}

// GetShareSpeechModelAudioWithChan invokes the iot.GetShareSpeechModelAudio API asynchronously
func (client *Client) GetShareSpeechModelAudioWithChan(request *GetShareSpeechModelAudioRequest) (<-chan *GetShareSpeechModelAudioResponse, <-chan error) {
	responseChan := make(chan *GetShareSpeechModelAudioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetShareSpeechModelAudio(request)
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

// GetShareSpeechModelAudioWithCallback invokes the iot.GetShareSpeechModelAudio API asynchronously
func (client *Client) GetShareSpeechModelAudioWithCallback(request *GetShareSpeechModelAudioRequest, callback func(response *GetShareSpeechModelAudioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetShareSpeechModelAudioResponse
		var err error
		defer close(result)
		response, err = client.GetShareSpeechModelAudio(request)
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

// GetShareSpeechModelAudioRequest is the request struct for api GetShareSpeechModelAudio
type GetShareSpeechModelAudioRequest struct {
	*requests.RpcRequest
	IotInstanceId       string    `position:"Body" name:"IotInstanceId"`
	ShareTaskId         string    `position:"Body" name:"ShareTaskId"`
	ApiProduct          string    `position:"Body" name:"ApiProduct"`
	ApiRevision         string    `position:"Body" name:"ApiRevision"`
	SpeechModelCodeList *[]string `position:"Body" name:"SpeechModelCodeList"  type:"Repeated"`
}

// GetShareSpeechModelAudioResponse is the response struct for api GetShareSpeechModelAudio
type GetShareSpeechModelAudioResponse struct {
	*responses.BaseResponse
	RequestId    string                         `json:"RequestId" xml:"RequestId"`
	Success      bool                           `json:"Success" xml:"Success"`
	Code         string                         `json:"Code" xml:"Code"`
	ErrorMessage string                         `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInGetShareSpeechModelAudio `json:"Data" xml:"Data"`
}

// CreateGetShareSpeechModelAudioRequest creates a request to invoke GetShareSpeechModelAudio API
func CreateGetShareSpeechModelAudioRequest() (request *GetShareSpeechModelAudioRequest) {
	request = &GetShareSpeechModelAudioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetShareSpeechModelAudio", "", "")
	request.Method = requests.POST
	return
}

// CreateGetShareSpeechModelAudioResponse creates a response to parse from GetShareSpeechModelAudio response
func CreateGetShareSpeechModelAudioResponse() (response *GetShareSpeechModelAudioResponse) {
	response = &GetShareSpeechModelAudioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}