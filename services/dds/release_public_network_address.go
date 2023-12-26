package dds

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

// ReleasePublicNetworkAddress invokes the dds.ReleasePublicNetworkAddress API synchronously
func (client *Client) ReleasePublicNetworkAddress(request *ReleasePublicNetworkAddressRequest) (response *ReleasePublicNetworkAddressResponse, err error) {
	response = CreateReleasePublicNetworkAddressResponse()
	err = client.DoAction(request, response)
	return
}

// ReleasePublicNetworkAddressWithChan invokes the dds.ReleasePublicNetworkAddress API asynchronously
func (client *Client) ReleasePublicNetworkAddressWithChan(request *ReleasePublicNetworkAddressRequest) (<-chan *ReleasePublicNetworkAddressResponse, <-chan error) {
	responseChan := make(chan *ReleasePublicNetworkAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleasePublicNetworkAddress(request)
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

// ReleasePublicNetworkAddressWithCallback invokes the dds.ReleasePublicNetworkAddress API asynchronously
func (client *Client) ReleasePublicNetworkAddressWithCallback(request *ReleasePublicNetworkAddressRequest, callback func(response *ReleasePublicNetworkAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleasePublicNetworkAddressResponse
		var err error
		defer close(result)
		response, err = client.ReleasePublicNetworkAddress(request)
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

// ReleasePublicNetworkAddressRequest is the request struct for api ReleasePublicNetworkAddress
type ReleasePublicNetworkAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ReleasePublicNetworkAddressResponse is the response struct for api ReleasePublicNetworkAddress
type ReleasePublicNetworkAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleasePublicNetworkAddressRequest creates a request to invoke ReleasePublicNetworkAddress API
func CreateReleasePublicNetworkAddressRequest() (request *ReleasePublicNetworkAddressRequest) {
	request = &ReleasePublicNetworkAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ReleasePublicNetworkAddress", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReleasePublicNetworkAddressResponse creates a response to parse from ReleasePublicNetworkAddress response
func CreateReleasePublicNetworkAddressResponse() (response *ReleasePublicNetworkAddressResponse) {
	response = &ReleasePublicNetworkAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
