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

// DebugBeginVnDialogue invokes the cloudcallcenter.DebugBeginVnDialogue API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugbeginvndialogue.html
func (client *Client) DebugBeginVnDialogue(request *DebugBeginVnDialogueRequest) (response *DebugBeginVnDialogueResponse, err error) {
	response = CreateDebugBeginVnDialogueResponse()
	err = client.DoAction(request, response)
	return
}

// DebugBeginVnDialogueWithChan invokes the cloudcallcenter.DebugBeginVnDialogue API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugbeginvndialogue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugBeginVnDialogueWithChan(request *DebugBeginVnDialogueRequest) (<-chan *DebugBeginVnDialogueResponse, <-chan error) {
	responseChan := make(chan *DebugBeginVnDialogueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DebugBeginVnDialogue(request)
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

// DebugBeginVnDialogueWithCallback invokes the cloudcallcenter.DebugBeginVnDialogue API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/debugbeginvndialogue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugBeginVnDialogueWithCallback(request *DebugBeginVnDialogueRequest, callback func(response *DebugBeginVnDialogueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DebugBeginVnDialogueResponse
		var err error
		defer close(result)
		response, err = client.DebugBeginVnDialogue(request)
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

// DebugBeginVnDialogueRequest is the request struct for api DebugBeginVnDialogue
type DebugBeginVnDialogueRequest struct {
	*requests.RpcRequest
	CallingNumber  string `position:"Query" name:"CallingNumber"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	CalledNumber   string `position:"Query" name:"CalledNumber"`
	ConversationId string `position:"Query" name:"ConversationId"`
	InitialContext string `position:"Query" name:"InitialContext"`
}

// DebugBeginVnDialogueResponse is the response struct for api DebugBeginVnDialogue
type DebugBeginVnDialogueResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TextResponse  string `json:"TextResponse" xml:"TextResponse"`
	Interruptible bool   `json:"Interruptible" xml:"Interruptible"`
	Action        string `json:"Action" xml:"Action"`
	ActionParams  string `json:"ActionParams" xml:"ActionParams"`
}

// CreateDebugBeginVnDialogueRequest creates a request to invoke DebugBeginVnDialogue API
func CreateDebugBeginVnDialogueRequest() (request *DebugBeginVnDialogueRequest) {
	request = &DebugBeginVnDialogueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DebugBeginVnDialogue", "", "")
	request.Method = requests.GET
	return
}

// CreateDebugBeginVnDialogueResponse creates a response to parse from DebugBeginVnDialogue response
func CreateDebugBeginVnDialogueResponse() (response *DebugBeginVnDialogueResponse) {
	response = &DebugBeginVnDialogueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
