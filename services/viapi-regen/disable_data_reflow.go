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

// DisableDataReflow invokes the viapi_regen.DisableDataReflow API synchronously
func (client *Client) DisableDataReflow(request *DisableDataReflowRequest) (response *DisableDataReflowResponse, err error) {
	response = CreateDisableDataReflowResponse()
	err = client.DoAction(request, response)
	return
}

// DisableDataReflowWithChan invokes the viapi_regen.DisableDataReflow API asynchronously
func (client *Client) DisableDataReflowWithChan(request *DisableDataReflowRequest) (<-chan *DisableDataReflowResponse, <-chan error) {
	responseChan := make(chan *DisableDataReflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableDataReflow(request)
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

// DisableDataReflowWithCallback invokes the viapi_regen.DisableDataReflow API asynchronously
func (client *Client) DisableDataReflowWithCallback(request *DisableDataReflowRequest, callback func(response *DisableDataReflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableDataReflowResponse
		var err error
		defer close(result)
		response, err = client.DisableDataReflow(request)
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

// DisableDataReflowRequest is the request struct for api DisableDataReflow
type DisableDataReflowRequest struct {
	*requests.RpcRequest
	ServiceId requests.Integer `position:"Body" name:"ServiceId"`
}

// DisableDataReflowResponse is the response struct for api DisableDataReflow
type DisableDataReflowResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDisableDataReflowRequest creates a request to invoke DisableDataReflow API
func CreateDisableDataReflowRequest() (request *DisableDataReflowRequest) {
	request = &DisableDataReflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "DisableDataReflow", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableDataReflowResponse creates a response to parse from DisableDataReflow response
func CreateDisableDataReflowResponse() (response *DisableDataReflowResponse) {
	response = &DisableDataReflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
