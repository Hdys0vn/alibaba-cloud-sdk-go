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

// ListBusiRegions invokes the ga.ListBusiRegions API synchronously
func (client *Client) ListBusiRegions(request *ListBusiRegionsRequest) (response *ListBusiRegionsResponse, err error) {
	response = CreateListBusiRegionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListBusiRegionsWithChan invokes the ga.ListBusiRegions API asynchronously
func (client *Client) ListBusiRegionsWithChan(request *ListBusiRegionsRequest) (<-chan *ListBusiRegionsResponse, <-chan error) {
	responseChan := make(chan *ListBusiRegionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBusiRegions(request)
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

// ListBusiRegionsWithCallback invokes the ga.ListBusiRegions API asynchronously
func (client *Client) ListBusiRegionsWithCallback(request *ListBusiRegionsRequest, callback func(response *ListBusiRegionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBusiRegionsResponse
		var err error
		defer close(result)
		response, err = client.ListBusiRegions(request)
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

// ListBusiRegionsRequest is the request struct for api ListBusiRegions
type ListBusiRegionsRequest struct {
	*requests.RpcRequest
}

// ListBusiRegionsResponse is the response struct for api ListBusiRegions
type ListBusiRegionsResponse struct {
	*responses.BaseResponse
	RequestId string        `json:"RequestId" xml:"RequestId"`
	Regions   []RegionsItem `json:"Regions" xml:"Regions"`
}

// CreateListBusiRegionsRequest creates a request to invoke ListBusiRegions API
func CreateListBusiRegionsRequest() (request *ListBusiRegionsRequest) {
	request = &ListBusiRegionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "ListBusiRegions", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListBusiRegionsResponse creates a response to parse from ListBusiRegions response
func CreateListBusiRegionsResponse() (response *ListBusiRegionsResponse) {
	response = &ListBusiRegionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
