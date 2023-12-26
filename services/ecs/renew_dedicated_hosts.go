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

// RenewDedicatedHosts invokes the ecs.RenewDedicatedHosts API synchronously
func (client *Client) RenewDedicatedHosts(request *RenewDedicatedHostsRequest) (response *RenewDedicatedHostsResponse, err error) {
	response = CreateRenewDedicatedHostsResponse()
	err = client.DoAction(request, response)
	return
}

// RenewDedicatedHostsWithChan invokes the ecs.RenewDedicatedHosts API asynchronously
func (client *Client) RenewDedicatedHostsWithChan(request *RenewDedicatedHostsRequest) (<-chan *RenewDedicatedHostsResponse, <-chan error) {
	responseChan := make(chan *RenewDedicatedHostsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenewDedicatedHosts(request)
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

// RenewDedicatedHostsWithCallback invokes the ecs.RenewDedicatedHosts API asynchronously
func (client *Client) RenewDedicatedHostsWithCallback(request *RenewDedicatedHostsRequest, callback func(response *RenewDedicatedHostsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenewDedicatedHostsResponse
		var err error
		defer close(result)
		response, err = client.RenewDedicatedHosts(request)
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

// RenewDedicatedHostsRequest is the request struct for api RenewDedicatedHosts
type RenewDedicatedHostsRequest struct {
	*requests.RpcRequest
	DedicatedHostIds     string           `position:"Query" name:"DedicatedHostIds"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Period               requests.Integer `position:"Query" name:"Period"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PeriodUnit           string           `position:"Query" name:"PeriodUnit"`
}

// RenewDedicatedHostsResponse is the response struct for api RenewDedicatedHosts
type RenewDedicatedHostsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRenewDedicatedHostsRequest creates a request to invoke RenewDedicatedHosts API
func CreateRenewDedicatedHostsRequest() (request *RenewDedicatedHostsRequest) {
	request = &RenewDedicatedHostsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "RenewDedicatedHosts", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRenewDedicatedHostsResponse creates a response to parse from RenewDedicatedHosts response
func CreateRenewDedicatedHostsResponse() (response *RenewDedicatedHostsResponse) {
	response = &RenewDedicatedHostsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
