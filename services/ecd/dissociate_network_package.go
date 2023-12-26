package ecd

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

// DissociateNetworkPackage invokes the ecd.DissociateNetworkPackage API synchronously
func (client *Client) DissociateNetworkPackage(request *DissociateNetworkPackageRequest) (response *DissociateNetworkPackageResponse, err error) {
	response = CreateDissociateNetworkPackageResponse()
	err = client.DoAction(request, response)
	return
}

// DissociateNetworkPackageWithChan invokes the ecd.DissociateNetworkPackage API asynchronously
func (client *Client) DissociateNetworkPackageWithChan(request *DissociateNetworkPackageRequest) (<-chan *DissociateNetworkPackageResponse, <-chan error) {
	responseChan := make(chan *DissociateNetworkPackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DissociateNetworkPackage(request)
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

// DissociateNetworkPackageWithCallback invokes the ecd.DissociateNetworkPackage API asynchronously
func (client *Client) DissociateNetworkPackageWithCallback(request *DissociateNetworkPackageRequest, callback func(response *DissociateNetworkPackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DissociateNetworkPackageResponse
		var err error
		defer close(result)
		response, err = client.DissociateNetworkPackage(request)
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

// DissociateNetworkPackageRequest is the request struct for api DissociateNetworkPackage
type DissociateNetworkPackageRequest struct {
	*requests.RpcRequest
	NetworkPackageId string `position:"Query" name:"NetworkPackageId"`
}

// DissociateNetworkPackageResponse is the response struct for api DissociateNetworkPackage
type DissociateNetworkPackageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDissociateNetworkPackageRequest creates a request to invoke DissociateNetworkPackage API
func CreateDissociateNetworkPackageRequest() (request *DissociateNetworkPackageRequest) {
	request = &DissociateNetworkPackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DissociateNetworkPackage", "", "")
	request.Method = requests.POST
	return
}

// CreateDissociateNetworkPackageResponse creates a response to parse from DissociateNetworkPackage response
func CreateDissociateNetworkPackageResponse() (response *DissociateNetworkPackageResponse) {
	response = &DissociateNetworkPackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
