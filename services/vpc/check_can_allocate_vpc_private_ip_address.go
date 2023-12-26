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

// CheckCanAllocateVpcPrivateIpAddress invokes the vpc.CheckCanAllocateVpcPrivateIpAddress API synchronously
func (client *Client) CheckCanAllocateVpcPrivateIpAddress(request *CheckCanAllocateVpcPrivateIpAddressRequest) (response *CheckCanAllocateVpcPrivateIpAddressResponse, err error) {
	response = CreateCheckCanAllocateVpcPrivateIpAddressResponse()
	err = client.DoAction(request, response)
	return
}

// CheckCanAllocateVpcPrivateIpAddressWithChan invokes the vpc.CheckCanAllocateVpcPrivateIpAddress API asynchronously
func (client *Client) CheckCanAllocateVpcPrivateIpAddressWithChan(request *CheckCanAllocateVpcPrivateIpAddressRequest) (<-chan *CheckCanAllocateVpcPrivateIpAddressResponse, <-chan error) {
	responseChan := make(chan *CheckCanAllocateVpcPrivateIpAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckCanAllocateVpcPrivateIpAddress(request)
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

// CheckCanAllocateVpcPrivateIpAddressWithCallback invokes the vpc.CheckCanAllocateVpcPrivateIpAddress API asynchronously
func (client *Client) CheckCanAllocateVpcPrivateIpAddressWithCallback(request *CheckCanAllocateVpcPrivateIpAddressRequest, callback func(response *CheckCanAllocateVpcPrivateIpAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckCanAllocateVpcPrivateIpAddressResponse
		var err error
		defer close(result)
		response, err = client.CheckCanAllocateVpcPrivateIpAddress(request)
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

// CheckCanAllocateVpcPrivateIpAddressRequest is the request struct for api CheckCanAllocateVpcPrivateIpAddress
type CheckCanAllocateVpcPrivateIpAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	IpVersion            string           `position:"Query" name:"IpVersion"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress     string           `position:"Query" name:"PrivateIpAddress"`
}

// CheckCanAllocateVpcPrivateIpAddressResponse is the response struct for api CheckCanAllocateVpcPrivateIpAddress
type CheckCanAllocateVpcPrivateIpAddressResponse struct {
	*responses.BaseResponse
	CanAllocate bool   `json:"CanAllocate" xml:"CanAllocate"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
}

// CreateCheckCanAllocateVpcPrivateIpAddressRequest creates a request to invoke CheckCanAllocateVpcPrivateIpAddress API
func CreateCheckCanAllocateVpcPrivateIpAddressRequest() (request *CheckCanAllocateVpcPrivateIpAddressRequest) {
	request = &CheckCanAllocateVpcPrivateIpAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CheckCanAllocateVpcPrivateIpAddress", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckCanAllocateVpcPrivateIpAddressResponse creates a response to parse from CheckCanAllocateVpcPrivateIpAddress response
func CreateCheckCanAllocateVpcPrivateIpAddressResponse() (response *CheckCanAllocateVpcPrivateIpAddressResponse) {
	response = &CheckCanAllocateVpcPrivateIpAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
