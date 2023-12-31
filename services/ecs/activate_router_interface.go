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

// ActivateRouterInterface invokes the ecs.ActivateRouterInterface API synchronously
func (client *Client) ActivateRouterInterface(request *ActivateRouterInterfaceRequest) (response *ActivateRouterInterfaceResponse, err error) {
	response = CreateActivateRouterInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

// ActivateRouterInterfaceWithChan invokes the ecs.ActivateRouterInterface API asynchronously
func (client *Client) ActivateRouterInterfaceWithChan(request *ActivateRouterInterfaceRequest) (<-chan *ActivateRouterInterfaceResponse, <-chan error) {
	responseChan := make(chan *ActivateRouterInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ActivateRouterInterface(request)
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

// ActivateRouterInterfaceWithCallback invokes the ecs.ActivateRouterInterface API asynchronously
func (client *Client) ActivateRouterInterfaceWithCallback(request *ActivateRouterInterfaceRequest, callback func(response *ActivateRouterInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ActivateRouterInterfaceResponse
		var err error
		defer close(result)
		response, err = client.ActivateRouterInterface(request)
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

// ActivateRouterInterfaceRequest is the request struct for api ActivateRouterInterface
type ActivateRouterInterfaceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string           `position:"Query" name:"RouterInterfaceId"`
}

// ActivateRouterInterfaceResponse is the response struct for api ActivateRouterInterface
type ActivateRouterInterfaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateActivateRouterInterfaceRequest creates a request to invoke ActivateRouterInterface API
func CreateActivateRouterInterfaceRequest() (request *ActivateRouterInterfaceRequest) {
	request = &ActivateRouterInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ActivateRouterInterface", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateActivateRouterInterfaceResponse creates a response to parse from ActivateRouterInterface response
func CreateActivateRouterInterfaceResponse() (response *ActivateRouterInterfaceResponse) {
	response = &ActivateRouterInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
