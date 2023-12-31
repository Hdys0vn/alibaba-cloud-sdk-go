package dg

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

// ModifyGateway invokes the dg.ModifyGateway API synchronously
func (client *Client) ModifyGateway(request *ModifyGatewayRequest) (response *ModifyGatewayResponse, err error) {
	response = CreateModifyGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyGatewayWithChan invokes the dg.ModifyGateway API asynchronously
func (client *Client) ModifyGatewayWithChan(request *ModifyGatewayRequest) (<-chan *ModifyGatewayResponse, <-chan error) {
	responseChan := make(chan *ModifyGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyGateway(request)
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

// ModifyGatewayWithCallback invokes the dg.ModifyGateway API asynchronously
func (client *Client) ModifyGatewayWithCallback(request *ModifyGatewayRequest, callback func(response *ModifyGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyGatewayResponse
		var err error
		defer close(result)
		response, err = client.ModifyGateway(request)
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

// ModifyGatewayRequest is the request struct for api ModifyGateway
type ModifyGatewayRequest struct {
	*requests.RpcRequest
	GatewayDesc string `position:"Body" name:"GatewayDesc"`
	GatewayName string `position:"Body" name:"GatewayName"`
	GatewayId   string `position:"Body" name:"GatewayId"`
}

// ModifyGatewayResponse is the response struct for api ModifyGateway
type ModifyGatewayResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateModifyGatewayRequest creates a request to invoke ModifyGateway API
func CreateModifyGatewayRequest() (request *ModifyGatewayRequest) {
	request = &ModifyGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dg", "2019-03-27", "ModifyGateway", "dg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyGatewayResponse creates a response to parse from ModifyGateway response
func CreateModifyGatewayResponse() (response *ModifyGatewayResponse) {
	response = &ModifyGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
