package ecs

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

// StartTerminalSession invokes the ecs.StartTerminalSession API synchronously
func (client *Client) StartTerminalSession(request *StartTerminalSessionRequest) (response *StartTerminalSessionResponse, err error) {
	response = CreateStartTerminalSessionResponse()
	err = client.DoAction(request, response)
	return
}

// StartTerminalSessionWithChan invokes the ecs.StartTerminalSession API asynchronously
func (client *Client) StartTerminalSessionWithChan(request *StartTerminalSessionRequest) (<-chan *StartTerminalSessionResponse, <-chan error) {
	responseChan := make(chan *StartTerminalSessionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartTerminalSession(request)
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

// StartTerminalSessionWithCallback invokes the ecs.StartTerminalSession API asynchronously
func (client *Client) StartTerminalSessionWithCallback(request *StartTerminalSessionRequest, callback func(response *StartTerminalSessionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartTerminalSessionResponse
		var err error
		defer close(result)
		response, err = client.StartTerminalSession(request)
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

// StartTerminalSessionRequest is the request struct for api StartTerminalSession
type StartTerminalSessionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CommandLine          string           `position:"Query" name:"CommandLine"`
	TargetServer         string           `position:"Query" name:"TargetServer"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
	PortNumber           requests.Integer `position:"Query" name:"PortNumber"`
}

// StartTerminalSessionResponse is the response struct for api StartTerminalSession
type StartTerminalSessionResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	SessionId     string `json:"SessionId" xml:"SessionId"`
	SecurityToken string `json:"SecurityToken" xml:"SecurityToken"`
	WebSocketUrl  string `json:"WebSocketUrl" xml:"WebSocketUrl"`
}

// CreateStartTerminalSessionRequest creates a request to invoke StartTerminalSession API
func CreateStartTerminalSessionRequest() (request *StartTerminalSessionRequest) {
	request = &StartTerminalSessionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "StartTerminalSession", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartTerminalSessionResponse creates a response to parse from StartTerminalSession response
func CreateStartTerminalSessionResponse() (response *StartTerminalSessionResponse) {
	response = &StartTerminalSessionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
