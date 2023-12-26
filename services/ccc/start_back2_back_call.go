package ccc

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

// StartBack2BackCall invokes the ccc.StartBack2BackCall API synchronously
func (client *Client) StartBack2BackCall(request *StartBack2BackCallRequest) (response *StartBack2BackCallResponse, err error) {
	response = CreateStartBack2BackCallResponse()
	err = client.DoAction(request, response)
	return
}

// StartBack2BackCallWithChan invokes the ccc.StartBack2BackCall API asynchronously
func (client *Client) StartBack2BackCallWithChan(request *StartBack2BackCallRequest) (<-chan *StartBack2BackCallResponse, <-chan error) {
	responseChan := make(chan *StartBack2BackCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartBack2BackCall(request)
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

// StartBack2BackCallWithCallback invokes the ccc.StartBack2BackCall API asynchronously
func (client *Client) StartBack2BackCallWithCallback(request *StartBack2BackCallRequest, callback func(response *StartBack2BackCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartBack2BackCallResponse
		var err error
		defer close(result)
		response, err = client.StartBack2BackCall(request)
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

// StartBack2BackCallRequest is the request struct for api StartBack2BackCall
type StartBack2BackCallRequest struct {
	*requests.RpcRequest
	Callee           string           `position:"Query" name:"Callee"`
	Broker           string           `position:"Query" name:"Broker"`
	AdditionalBroker string           `position:"Query" name:"AdditionalBroker"`
	Tags             string           `position:"Query" name:"Tags"`
	TimeoutSeconds   requests.Integer `position:"Query" name:"TimeoutSeconds"`
	Caller           string           `position:"Query" name:"Caller"`
	InstanceId       string           `position:"Query" name:"InstanceId"`
}

// StartBack2BackCallResponse is the response struct for api StartBack2BackCall
type StartBack2BackCallResponse struct {
	*responses.BaseResponse
	Code           string   `json:"Code" xml:"Code"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string   `json:"Message" xml:"Message"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
	Data           Data     `json:"Data" xml:"Data"`
}

// CreateStartBack2BackCallRequest creates a request to invoke StartBack2BackCall API
func CreateStartBack2BackCallRequest() (request *StartBack2BackCallRequest) {
	request = &StartBack2BackCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "StartBack2BackCall", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartBack2BackCallResponse creates a response to parse from StartBack2BackCall response
func CreateStartBack2BackCallResponse() (response *StartBack2BackCallResponse) {
	response = &StartBack2BackCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
