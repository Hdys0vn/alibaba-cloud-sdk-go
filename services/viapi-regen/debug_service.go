package viapi_regen

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

// DebugService invokes the viapi_regen.DebugService API synchronously
func (client *Client) DebugService(request *DebugServiceRequest) (response *DebugServiceResponse, err error) {
	response = CreateDebugServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DebugServiceWithChan invokes the viapi_regen.DebugService API asynchronously
func (client *Client) DebugServiceWithChan(request *DebugServiceRequest) (<-chan *DebugServiceResponse, <-chan error) {
	responseChan := make(chan *DebugServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DebugService(request)
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

// DebugServiceWithCallback invokes the viapi_regen.DebugService API asynchronously
func (client *Client) DebugServiceWithCallback(request *DebugServiceRequest, callback func(response *DebugServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DebugServiceResponse
		var err error
		defer close(result)
		response, err = client.DebugService(request)
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

// DebugServiceRequest is the request struct for api DebugService
type DebugServiceRequest struct {
	*requests.RpcRequest
	Param string           `position:"Body" name:"Param"`
	Id    requests.Integer `position:"Body" name:"Id"`
}

// DebugServiceResponse is the response struct for api DebugService
type DebugServiceResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
}

// CreateDebugServiceRequest creates a request to invoke DebugService API
func CreateDebugServiceRequest() (request *DebugServiceRequest) {
	request = &DebugServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "DebugService", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDebugServiceResponse creates a response to parse from DebugService response
func CreateDebugServiceResponse() (response *DebugServiceResponse) {
	response = &DebugServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}