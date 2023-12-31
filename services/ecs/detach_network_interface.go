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

// DetachNetworkInterface invokes the ecs.DetachNetworkInterface API synchronously
func (client *Client) DetachNetworkInterface(request *DetachNetworkInterfaceRequest) (response *DetachNetworkInterfaceResponse, err error) {
	response = CreateDetachNetworkInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

// DetachNetworkInterfaceWithChan invokes the ecs.DetachNetworkInterface API asynchronously
func (client *Client) DetachNetworkInterfaceWithChan(request *DetachNetworkInterfaceRequest) (<-chan *DetachNetworkInterfaceResponse, <-chan error) {
	responseChan := make(chan *DetachNetworkInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachNetworkInterface(request)
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

// DetachNetworkInterfaceWithCallback invokes the ecs.DetachNetworkInterface API asynchronously
func (client *Client) DetachNetworkInterfaceWithCallback(request *DetachNetworkInterfaceRequest, callback func(response *DetachNetworkInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachNetworkInterfaceResponse
		var err error
		defer close(result)
		response, err = client.DetachNetworkInterface(request)
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

// DetachNetworkInterfaceRequest is the request struct for api DetachNetworkInterface
type DetachNetworkInterfaceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TrunkNetworkInstanceId string           `position:"Query" name:"TrunkNetworkInstanceId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId             string           `position:"Query" name:"InstanceId"`
	NetworkInterfaceId     string           `position:"Query" name:"NetworkInterfaceId"`
}

// DetachNetworkInterfaceResponse is the response struct for api DetachNetworkInterface
type DetachNetworkInterfaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachNetworkInterfaceRequest creates a request to invoke DetachNetworkInterface API
func CreateDetachNetworkInterfaceRequest() (request *DetachNetworkInterfaceRequest) {
	request = &DetachNetworkInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DetachNetworkInterface", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachNetworkInterfaceResponse creates a response to parse from DetachNetworkInterface response
func CreateDetachNetworkInterfaceResponse() (response *DetachNetworkInterfaceResponse) {
	response = &DetachNetworkInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
