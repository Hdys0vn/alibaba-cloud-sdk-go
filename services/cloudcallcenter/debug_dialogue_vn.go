package cloudcallcenter

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

// DebugDialogueVn invokes the cloudcallcenter.DebugDialogueVn API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugdialoguevn.html
func (client *Client) DebugDialogueVn(request *DebugDialogueVnRequest) (response *DebugDialogueVnResponse, err error) {
	response = CreateDebugDialogueVnResponse()
	err = client.DoAction(request, response)
	return
}

// DebugDialogueVnWithChan invokes the cloudcallcenter.DebugDialogueVn API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugdialoguevn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugDialogueVnWithChan(request *DebugDialogueVnRequest) (<-chan *DebugDialogueVnResponse, <-chan error) {
	responseChan := make(chan *DebugDialogueVnResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DebugDialogueVn(request)
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

// DebugDialogueVnWithCallback invokes the cloudcallcenter.DebugDialogueVn API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugdialoguevn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugDialogueVnWithCallback(request *DebugDialogueVnRequest, callback func(response *DebugDialogueVnResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DebugDialogueVnResponse
		var err error
		defer close(result)
		response, err = client.DebugDialogueVn(request)
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

// DebugDialogueVnRequest is the request struct for api DebugDialogueVn
type DebugDialogueVnRequest struct {
	*requests.RpcRequest
	ConversationId    string `position:"Query" name:"ConversationId"`
	InstanceId        string `position:"Query" name:"InstanceId"`
	AdditionalContext string `position:"Query" name:"AdditionalContext"`
	Utterance         string `position:"Query" name:"Utterance"`
}

// DebugDialogueVnResponse is the response struct for api DebugDialogueVn
type DebugDialogueVnResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TextResponse  string `json:"TextResponse" xml:"TextResponse"`
	Interruptible bool   `json:"Interruptible" xml:"Interruptible"`
	Action        string `json:"Action" xml:"Action"`
	ActionParams  string `json:"ActionParams" xml:"ActionParams"`
}

// CreateDebugDialogueVnRequest creates a request to invoke DebugDialogueVn API
func CreateDebugDialogueVnRequest() (request *DebugDialogueVnRequest) {
	request = &DebugDialogueVnRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DebugDialogueVn", "", "")
	request.Method = requests.GET
	return
}

// CreateDebugDialogueVnResponse creates a response to parse from DebugDialogueVn response
func CreateDebugDialogueVnResponse() (response *DebugDialogueVnResponse) {
	response = &DebugDialogueVnResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
