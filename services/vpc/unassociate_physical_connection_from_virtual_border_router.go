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

// UnassociatePhysicalConnectionFromVirtualBorderRouter invokes the vpc.UnassociatePhysicalConnectionFromVirtualBorderRouter API synchronously
func (client *Client) UnassociatePhysicalConnectionFromVirtualBorderRouter(request *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) (response *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse, err error) {
	response = CreateUnassociatePhysicalConnectionFromVirtualBorderRouterResponse()
	err = client.DoAction(request, response)
	return
}

// UnassociatePhysicalConnectionFromVirtualBorderRouterWithChan invokes the vpc.UnassociatePhysicalConnectionFromVirtualBorderRouter API asynchronously
func (client *Client) UnassociatePhysicalConnectionFromVirtualBorderRouterWithChan(request *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) (<-chan *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse, <-chan error) {
	responseChan := make(chan *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnassociatePhysicalConnectionFromVirtualBorderRouter(request)
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

// UnassociatePhysicalConnectionFromVirtualBorderRouterWithCallback invokes the vpc.UnassociatePhysicalConnectionFromVirtualBorderRouter API asynchronously
func (client *Client) UnassociatePhysicalConnectionFromVirtualBorderRouterWithCallback(request *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest, callback func(response *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse
		var err error
		defer close(result)
		response, err = client.UnassociatePhysicalConnectionFromVirtualBorderRouter(request)
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

// UnassociatePhysicalConnectionFromVirtualBorderRouterRequest is the request struct for api UnassociatePhysicalConnectionFromVirtualBorderRouter
type UnassociatePhysicalConnectionFromVirtualBorderRouterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	VbrId                string           `position:"Query" name:"VbrId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PhysicalConnectionId string           `position:"Query" name:"PhysicalConnectionId"`
}

// UnassociatePhysicalConnectionFromVirtualBorderRouterResponse is the response struct for api UnassociatePhysicalConnectionFromVirtualBorderRouter
type UnassociatePhysicalConnectionFromVirtualBorderRouterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnassociatePhysicalConnectionFromVirtualBorderRouterRequest creates a request to invoke UnassociatePhysicalConnectionFromVirtualBorderRouter API
func CreateUnassociatePhysicalConnectionFromVirtualBorderRouterRequest() (request *UnassociatePhysicalConnectionFromVirtualBorderRouterRequest) {
	request = &UnassociatePhysicalConnectionFromVirtualBorderRouterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "UnassociatePhysicalConnectionFromVirtualBorderRouter", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnassociatePhysicalConnectionFromVirtualBorderRouterResponse creates a response to parse from UnassociatePhysicalConnectionFromVirtualBorderRouter response
func CreateUnassociatePhysicalConnectionFromVirtualBorderRouterResponse() (response *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse) {
	response = &UnassociatePhysicalConnectionFromVirtualBorderRouterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
