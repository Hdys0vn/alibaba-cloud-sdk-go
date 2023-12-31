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

// RemoveBandwidthPackageIps invokes the ecs.RemoveBandwidthPackageIps API synchronously
func (client *Client) RemoveBandwidthPackageIps(request *RemoveBandwidthPackageIpsRequest) (response *RemoveBandwidthPackageIpsResponse, err error) {
	response = CreateRemoveBandwidthPackageIpsResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveBandwidthPackageIpsWithChan invokes the ecs.RemoveBandwidthPackageIps API asynchronously
func (client *Client) RemoveBandwidthPackageIpsWithChan(request *RemoveBandwidthPackageIpsRequest) (<-chan *RemoveBandwidthPackageIpsResponse, <-chan error) {
	responseChan := make(chan *RemoveBandwidthPackageIpsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveBandwidthPackageIps(request)
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

// RemoveBandwidthPackageIpsWithCallback invokes the ecs.RemoveBandwidthPackageIps API asynchronously
func (client *Client) RemoveBandwidthPackageIpsWithCallback(request *RemoveBandwidthPackageIpsRequest, callback func(response *RemoveBandwidthPackageIpsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveBandwidthPackageIpsResponse
		var err error
		defer close(result)
		response, err = client.RemoveBandwidthPackageIps(request)
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

// RemoveBandwidthPackageIpsRequest is the request struct for api RemoveBandwidthPackageIps
type RemoveBandwidthPackageIpsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	RemovedIpAddresses   *[]string        `position:"Query" name:"RemovedIpAddresses"  type:"Repeated"`
	BandwidthPackageId   string           `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// RemoveBandwidthPackageIpsResponse is the response struct for api RemoveBandwidthPackageIps
type RemoveBandwidthPackageIpsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveBandwidthPackageIpsRequest creates a request to invoke RemoveBandwidthPackageIps API
func CreateRemoveBandwidthPackageIpsRequest() (request *RemoveBandwidthPackageIpsRequest) {
	request = &RemoveBandwidthPackageIpsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "RemoveBandwidthPackageIps", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveBandwidthPackageIpsResponse creates a response to parse from RemoveBandwidthPackageIps response
func CreateRemoveBandwidthPackageIpsResponse() (response *RemoveBandwidthPackageIpsResponse) {
	response = &RemoveBandwidthPackageIpsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
