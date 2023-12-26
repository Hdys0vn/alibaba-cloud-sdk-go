package sgw

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

// DescribeStorageBundle invokes the sgw.DescribeStorageBundle API synchronously
func (client *Client) DescribeStorageBundle(request *DescribeStorageBundleRequest) (response *DescribeStorageBundleResponse, err error) {
	response = CreateDescribeStorageBundleResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStorageBundleWithChan invokes the sgw.DescribeStorageBundle API asynchronously
func (client *Client) DescribeStorageBundleWithChan(request *DescribeStorageBundleRequest) (<-chan *DescribeStorageBundleResponse, <-chan error) {
	responseChan := make(chan *DescribeStorageBundleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStorageBundle(request)
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

// DescribeStorageBundleWithCallback invokes the sgw.DescribeStorageBundle API asynchronously
func (client *Client) DescribeStorageBundleWithCallback(request *DescribeStorageBundleRequest, callback func(response *DescribeStorageBundleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStorageBundleResponse
		var err error
		defer close(result)
		response, err = client.DescribeStorageBundle(request)
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

// DescribeStorageBundleRequest is the request struct for api DescribeStorageBundle
type DescribeStorageBundleRequest struct {
	*requests.RpcRequest
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	StorageBundleId string `position:"Query" name:"StorageBundleId"`
}

// DescribeStorageBundleResponse is the response struct for api DescribeStorageBundle
type DescribeStorageBundleResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	Success               bool   `json:"Success" xml:"Success"`
	Code                  string `json:"Code" xml:"Code"`
	Message               string `json:"Message" xml:"Message"`
	StorageBundleId       string `json:"StorageBundleId" xml:"StorageBundleId"`
	Name                  string `json:"Name" xml:"Name"`
	Description           string `json:"Description" xml:"Description"`
	BackendBucketRegionId string `json:"BackendBucketRegionId" xml:"BackendBucketRegionId"`
	Location              string `json:"Location" xml:"Location"`
	CreatedTime           int64  `json:"CreatedTime" xml:"CreatedTime"`
	ResourceGroupId       string `json:"ResourceGroupId" xml:"ResourceGroupId"`
}

// CreateDescribeStorageBundleRequest creates a request to invoke DescribeStorageBundle API
func CreateDescribeStorageBundleRequest() (request *DescribeStorageBundleRequest) {
	request = &DescribeStorageBundleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeStorageBundle", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeStorageBundleResponse creates a response to parse from DescribeStorageBundle response
func CreateDescribeStorageBundleResponse() (response *DescribeStorageBundleResponse) {
	response = &DescribeStorageBundleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
