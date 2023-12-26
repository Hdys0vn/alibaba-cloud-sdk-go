package mpaas

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

// DeleteMPCDById invokes the mpaas.DeleteMPCDById API synchronously
func (client *Client) DeleteMPCDById(request *DeleteMPCDByIdRequest) (response *DeleteMPCDByIdResponse, err error) {
	response = CreateDeleteMPCDByIdResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMPCDByIdWithChan invokes the mpaas.DeleteMPCDById API asynchronously
func (client *Client) DeleteMPCDByIdWithChan(request *DeleteMPCDByIdRequest) (<-chan *DeleteMPCDByIdResponse, <-chan error) {
	responseChan := make(chan *DeleteMPCDByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMPCDById(request)
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

// DeleteMPCDByIdWithCallback invokes the mpaas.DeleteMPCDById API asynchronously
func (client *Client) DeleteMPCDByIdWithCallback(request *DeleteMPCDByIdRequest, callback func(response *DeleteMPCDByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMPCDByIdResponse
		var err error
		defer close(result)
		response, err = client.DeleteMPCDById(request)
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

// DeleteMPCDByIdRequest is the request struct for api DeleteMPCDById
type DeleteMPCDByIdRequest struct {
	*requests.RpcRequest
	AppId       string           `position:"Body" name:"AppId"`
	TenantId    string           `position:"Body" name:"TenantId"`
	ScanTaskId  requests.Integer `position:"Body" name:"ScanTaskId"`
	WorkspaceId string           `position:"Body" name:"WorkspaceId"`
}

// DeleteMPCDByIdResponse is the response struct for api DeleteMPCDById
type DeleteMPCDByIdResponse struct {
	*responses.BaseResponse
}

// CreateDeleteMPCDByIdRequest creates a request to invoke DeleteMPCDById API
func CreateDeleteMPCDByIdRequest() (request *DeleteMPCDByIdRequest) {
	request = &DeleteMPCDByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "DeleteMPCDById", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteMPCDByIdResponse creates a response to parse from DeleteMPCDById response
func CreateDeleteMPCDByIdResponse() (response *DeleteMPCDByIdResponse) {
	response = &DeleteMPCDByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}