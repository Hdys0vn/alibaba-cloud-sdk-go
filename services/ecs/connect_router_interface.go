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

// ConnectRouterInterface invokes the ecs.ConnectRouterInterface API synchronously
func (client *Client) ConnectRouterInterface(request *ConnectRouterInterfaceRequest) (response *ConnectRouterInterfaceResponse, err error) {
	response = CreateConnectRouterInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

// ConnectRouterInterfaceWithChan invokes the ecs.ConnectRouterInterface API asynchronously
func (client *Client) ConnectRouterInterfaceWithChan(request *ConnectRouterInterfaceRequest) (<-chan *ConnectRouterInterfaceResponse, <-chan error) {
	responseChan := make(chan *ConnectRouterInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConnectRouterInterface(request)
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

// ConnectRouterInterfaceWithCallback invokes the ecs.ConnectRouterInterface API asynchronously
func (client *Client) ConnectRouterInterfaceWithCallback(request *ConnectRouterInterfaceRequest, callback func(response *ConnectRouterInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConnectRouterInterfaceResponse
		var err error
		defer close(result)
		response, err = client.ConnectRouterInterface(request)
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

// ConnectRouterInterfaceRequest is the request struct for api ConnectRouterInterface
type ConnectRouterInterfaceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string           `position:"Query" name:"RouterInterfaceId"`
}

// ConnectRouterInterfaceResponse is the response struct for api ConnectRouterInterface
type ConnectRouterInterfaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConnectRouterInterfaceRequest creates a request to invoke ConnectRouterInterface API
func CreateConnectRouterInterfaceRequest() (request *ConnectRouterInterfaceRequest) {
	request = &ConnectRouterInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ConnectRouterInterface", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConnectRouterInterfaceResponse creates a response to parse from ConnectRouterInterface response
func CreateConnectRouterInterfaceResponse() (response *ConnectRouterInterfaceResponse) {
	response = &ConnectRouterInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
