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

// FindUserGatewayById invokes the dg.FindUserGatewayById API synchronously
func (client *Client) FindUserGatewayById(request *FindUserGatewayByIdRequest) (response *FindUserGatewayByIdResponse, err error) {
	response = CreateFindUserGatewayByIdResponse()
	err = client.DoAction(request, response)
	return
}

// FindUserGatewayByIdWithChan invokes the dg.FindUserGatewayById API asynchronously
func (client *Client) FindUserGatewayByIdWithChan(request *FindUserGatewayByIdRequest) (<-chan *FindUserGatewayByIdResponse, <-chan error) {
	responseChan := make(chan *FindUserGatewayByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindUserGatewayById(request)
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

// FindUserGatewayByIdWithCallback invokes the dg.FindUserGatewayById API asynchronously
func (client *Client) FindUserGatewayByIdWithCallback(request *FindUserGatewayByIdRequest, callback func(response *FindUserGatewayByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindUserGatewayByIdResponse
		var err error
		defer close(result)
		response, err = client.FindUserGatewayById(request)
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

// FindUserGatewayByIdRequest is the request struct for api FindUserGatewayById
type FindUserGatewayByIdRequest struct {
	*requests.RpcRequest
	GatewayId string `position:"Body" name:"GatewayId"`
}

// FindUserGatewayByIdResponse is the response struct for api FindUserGatewayById
type FindUserGatewayByIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateFindUserGatewayByIdRequest creates a request to invoke FindUserGatewayById API
func CreateFindUserGatewayByIdRequest() (request *FindUserGatewayByIdRequest) {
	request = &FindUserGatewayByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dg", "2019-03-27", "FindUserGatewayById", "dg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateFindUserGatewayByIdResponse creates a response to parse from FindUserGatewayById response
func CreateFindUserGatewayByIdResponse() (response *FindUserGatewayByIdResponse) {
	response = &FindUserGatewayByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
