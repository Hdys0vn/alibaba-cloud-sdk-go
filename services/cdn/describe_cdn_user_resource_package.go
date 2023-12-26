package cdn

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

// DescribeCdnUserResourcePackage invokes the cdn.DescribeCdnUserResourcePackage API synchronously
func (client *Client) DescribeCdnUserResourcePackage(request *DescribeCdnUserResourcePackageRequest) (response *DescribeCdnUserResourcePackageResponse, err error) {
	response = CreateDescribeCdnUserResourcePackageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnUserResourcePackageWithChan invokes the cdn.DescribeCdnUserResourcePackage API asynchronously
func (client *Client) DescribeCdnUserResourcePackageWithChan(request *DescribeCdnUserResourcePackageRequest) (<-chan *DescribeCdnUserResourcePackageResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnUserResourcePackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnUserResourcePackage(request)
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

// DescribeCdnUserResourcePackageWithCallback invokes the cdn.DescribeCdnUserResourcePackage API asynchronously
func (client *Client) DescribeCdnUserResourcePackageWithCallback(request *DescribeCdnUserResourcePackageRequest, callback func(response *DescribeCdnUserResourcePackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnUserResourcePackageResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnUserResourcePackage(request)
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

// DescribeCdnUserResourcePackageRequest is the request struct for api DescribeCdnUserResourcePackage
type DescribeCdnUserResourcePackageRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	Status        string           `position:"Query" name:"Status"`
}

// DescribeCdnUserResourcePackageResponse is the response struct for api DescribeCdnUserResourcePackage
type DescribeCdnUserResourcePackageResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	ResourcePackageInfos ResourcePackageInfos `json:"ResourcePackageInfos" xml:"ResourcePackageInfos"`
}

// CreateDescribeCdnUserResourcePackageRequest creates a request to invoke DescribeCdnUserResourcePackage API
func CreateDescribeCdnUserResourcePackageRequest() (request *DescribeCdnUserResourcePackageRequest) {
	request = &DescribeCdnUserResourcePackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCdnUserResourcePackage", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCdnUserResourcePackageResponse creates a response to parse from DescribeCdnUserResourcePackage response
func CreateDescribeCdnUserResourcePackageResponse() (response *DescribeCdnUserResourcePackageResponse) {
	response = &DescribeCdnUserResourcePackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
