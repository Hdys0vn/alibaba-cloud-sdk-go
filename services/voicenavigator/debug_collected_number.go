package voicenavigator

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

// DebugCollectedNumber invokes the voicenavigator.DebugCollectedNumber API synchronously
func (client *Client) DebugCollectedNumber(request *DebugCollectedNumberRequest) (response *DebugCollectedNumberResponse, err error) {
	response = CreateDebugCollectedNumberResponse()
	err = client.DoAction(request, response)
	return
}

// DebugCollectedNumberWithChan invokes the voicenavigator.DebugCollectedNumber API asynchronously
func (client *Client) DebugCollectedNumberWithChan(request *DebugCollectedNumberRequest) (<-chan *DebugCollectedNumberResponse, <-chan error) {
	responseChan := make(chan *DebugCollectedNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DebugCollectedNumber(request)
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

// DebugCollectedNumberWithCallback invokes the voicenavigator.DebugCollectedNumber API asynchronously
func (client *Client) DebugCollectedNumberWithCallback(request *DebugCollectedNumberRequest, callback func(response *DebugCollectedNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DebugCollectedNumberResponse
		var err error
		defer close(result)
		response, err = client.DebugCollectedNumber(request)
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

// DebugCollectedNumberRequest is the request struct for api DebugCollectedNumber
type DebugCollectedNumberRequest struct {
	*requests.RpcRequest
	ConversationId string `position:"Query" name:"ConversationId"`
	Number         string `position:"Query" name:"Number"`
	InstanceId     string `position:"Query" name:"InstanceId"`
}

// DebugCollectedNumberResponse is the response struct for api DebugCollectedNumber
type DebugCollectedNumberResponse struct {
	*responses.BaseResponse
	Action        string `json:"Action" xml:"Action"`
	Interruptible bool   `json:"Interruptible" xml:"Interruptible"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ActionParams  string `json:"ActionParams" xml:"ActionParams"`
	TextResponse  string `json:"TextResponse" xml:"TextResponse"`
}

// CreateDebugCollectedNumberRequest creates a request to invoke DebugCollectedNumber API
func CreateDebugCollectedNumberRequest() (request *DebugCollectedNumberRequest) {
	request = &DebugCollectedNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "DebugCollectedNumber", "voicebot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDebugCollectedNumberResponse creates a response to parse from DebugCollectedNumber response
func CreateDebugCollectedNumberResponse() (response *DebugCollectedNumberResponse) {
	response = &DebugCollectedNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
