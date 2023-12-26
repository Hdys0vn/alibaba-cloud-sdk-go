package scsp

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

// SendOutboundCommand invokes the scsp.SendOutboundCommand API synchronously
func (client *Client) SendOutboundCommand(request *SendOutboundCommandRequest) (response *SendOutboundCommandResponse, err error) {
	response = CreateSendOutboundCommandResponse()
	err = client.DoAction(request, response)
	return
}

// SendOutboundCommandWithChan invokes the scsp.SendOutboundCommand API asynchronously
func (client *Client) SendOutboundCommandWithChan(request *SendOutboundCommandRequest) (<-chan *SendOutboundCommandResponse, <-chan error) {
	responseChan := make(chan *SendOutboundCommandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendOutboundCommand(request)
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

// SendOutboundCommandWithCallback invokes the scsp.SendOutboundCommand API asynchronously
func (client *Client) SendOutboundCommandWithCallback(request *SendOutboundCommandRequest, callback func(response *SendOutboundCommandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendOutboundCommandResponse
		var err error
		defer close(result)
		response, err = client.SendOutboundCommand(request)
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

// SendOutboundCommandRequest is the request struct for api SendOutboundCommand
type SendOutboundCommandRequest struct {
	*requests.RpcRequest
	CallingNumber string `position:"Body"`
	InstanceId    string `position:"Body"`
	AccountName   string `position:"Body"`
	CalledNumber  string `position:"Body"`
	CustomerInfo  string `position:"Body"`
}

// SendOutboundCommandResponse is the response struct for api SendOutboundCommand
type SendOutboundCommandResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSendOutboundCommandRequest creates a request to invoke SendOutboundCommand API
func CreateSendOutboundCommandRequest() (request *SendOutboundCommandRequest) {
	request = &SendOutboundCommandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "SendOutboundCommand", "", "")
	request.Method = requests.POST
	return
}

// CreateSendOutboundCommandResponse creates a response to parse from SendOutboundCommand response
func CreateSendOutboundCommandResponse() (response *SendOutboundCommandResponse) {
	response = &SendOutboundCommandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}