package vpc

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

// DeleteRouterInterface invokes the vpc.DeleteRouterInterface API synchronously
func (client *Client) DeleteRouterInterface(request *DeleteRouterInterfaceRequest) (response *DeleteRouterInterfaceResponse, err error) {
	response = CreateDeleteRouterInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRouterInterfaceWithChan invokes the vpc.DeleteRouterInterface API asynchronously
func (client *Client) DeleteRouterInterfaceWithChan(request *DeleteRouterInterfaceRequest) (<-chan *DeleteRouterInterfaceResponse, <-chan error) {
	responseChan := make(chan *DeleteRouterInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRouterInterface(request)
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

// DeleteRouterInterfaceWithCallback invokes the vpc.DeleteRouterInterface API asynchronously
func (client *Client) DeleteRouterInterfaceWithCallback(request *DeleteRouterInterfaceRequest, callback func(response *DeleteRouterInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRouterInterfaceResponse
		var err error
		defer close(result)
		response, err = client.DeleteRouterInterface(request)
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

// DeleteRouterInterfaceRequest is the request struct for api DeleteRouterInterface
type DeleteRouterInterfaceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	RouterInterfaceId    string           `position:"Query" name:"RouterInterfaceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteRouterInterfaceResponse is the response struct for api DeleteRouterInterface
type DeleteRouterInterfaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteRouterInterfaceRequest creates a request to invoke DeleteRouterInterface API
func CreateDeleteRouterInterfaceRequest() (request *DeleteRouterInterfaceRequest) {
	request = &DeleteRouterInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteRouterInterface", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteRouterInterfaceResponse creates a response to parse from DeleteRouterInterface response
func CreateDeleteRouterInterfaceResponse() (response *DeleteRouterInterfaceResponse) {
	response = &DeleteRouterInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
