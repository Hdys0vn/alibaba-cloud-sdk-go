package edas

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

// ListAliyunRegion invokes the edas.ListAliyunRegion API synchronously
func (client *Client) ListAliyunRegion(request *ListAliyunRegionRequest) (response *ListAliyunRegionResponse, err error) {
	response = CreateListAliyunRegionResponse()
	err = client.DoAction(request, response)
	return
}

// ListAliyunRegionWithChan invokes the edas.ListAliyunRegion API asynchronously
func (client *Client) ListAliyunRegionWithChan(request *ListAliyunRegionRequest) (<-chan *ListAliyunRegionResponse, <-chan error) {
	responseChan := make(chan *ListAliyunRegionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAliyunRegion(request)
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

// ListAliyunRegionWithCallback invokes the edas.ListAliyunRegion API asynchronously
func (client *Client) ListAliyunRegionWithCallback(request *ListAliyunRegionRequest, callback func(response *ListAliyunRegionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAliyunRegionResponse
		var err error
		defer close(result)
		response, err = client.ListAliyunRegion(request)
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

// ListAliyunRegionRequest is the request struct for api ListAliyunRegion
type ListAliyunRegionRequest struct {
	*requests.RoaRequest
}

// ListAliyunRegionResponse is the response struct for api ListAliyunRegion
type ListAliyunRegionResponse struct {
	*responses.BaseResponse
	Code             int                                `json:"Code" xml:"Code"`
	Message          string                             `json:"Message" xml:"Message"`
	RequestId        string                             `json:"RequestId" xml:"RequestId"`
	RegionEntityList RegionEntityListInListAliyunRegion `json:"RegionEntityList" xml:"RegionEntityList"`
}

// CreateListAliyunRegionRequest creates a request to invoke ListAliyunRegion API
func CreateListAliyunRegionRequest() (request *ListAliyunRegionRequest) {
	request = &ListAliyunRegionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListAliyunRegion", "/pop/v5/resource/region_list", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAliyunRegionResponse creates a response to parse from ListAliyunRegion response
func CreateListAliyunRegionResponse() (response *ListAliyunRegionResponse) {
	response = &ListAliyunRegionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
