package ga

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

// ReplaceBandwidthPackage invokes the ga.ReplaceBandwidthPackage API synchronously
func (client *Client) ReplaceBandwidthPackage(request *ReplaceBandwidthPackageRequest) (response *ReplaceBandwidthPackageResponse, err error) {
	response = CreateReplaceBandwidthPackageResponse()
	err = client.DoAction(request, response)
	return
}

// ReplaceBandwidthPackageWithChan invokes the ga.ReplaceBandwidthPackage API asynchronously
func (client *Client) ReplaceBandwidthPackageWithChan(request *ReplaceBandwidthPackageRequest) (<-chan *ReplaceBandwidthPackageResponse, <-chan error) {
	responseChan := make(chan *ReplaceBandwidthPackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReplaceBandwidthPackage(request)
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

// ReplaceBandwidthPackageWithCallback invokes the ga.ReplaceBandwidthPackage API asynchronously
func (client *Client) ReplaceBandwidthPackageWithCallback(request *ReplaceBandwidthPackageRequest, callback func(response *ReplaceBandwidthPackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReplaceBandwidthPackageResponse
		var err error
		defer close(result)
		response, err = client.ReplaceBandwidthPackage(request)
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

// ReplaceBandwidthPackageRequest is the request struct for api ReplaceBandwidthPackage
type ReplaceBandwidthPackageRequest struct {
	*requests.RpcRequest
	BandwidthPackageId       string `position:"Query" name:"BandwidthPackageId"`
	TargetBandwidthPackageId string `position:"Query" name:"TargetBandwidthPackageId"`
}

// ReplaceBandwidthPackageResponse is the response struct for api ReplaceBandwidthPackage
type ReplaceBandwidthPackageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReplaceBandwidthPackageRequest creates a request to invoke ReplaceBandwidthPackage API
func CreateReplaceBandwidthPackageRequest() (request *ReplaceBandwidthPackageRequest) {
	request = &ReplaceBandwidthPackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "ReplaceBandwidthPackage", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReplaceBandwidthPackageResponse creates a response to parse from ReplaceBandwidthPackage response
func CreateReplaceBandwidthPackageResponse() (response *ReplaceBandwidthPackageResponse) {
	response = &ReplaceBandwidthPackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
