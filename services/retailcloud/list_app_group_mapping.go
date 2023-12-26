package retailcloud

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

// ListAppGroupMapping invokes the retailcloud.ListAppGroupMapping API synchronously
func (client *Client) ListAppGroupMapping(request *ListAppGroupMappingRequest) (response *ListAppGroupMappingResponse, err error) {
	response = CreateListAppGroupMappingResponse()
	err = client.DoAction(request, response)
	return
}

// ListAppGroupMappingWithChan invokes the retailcloud.ListAppGroupMapping API asynchronously
func (client *Client) ListAppGroupMappingWithChan(request *ListAppGroupMappingRequest) (<-chan *ListAppGroupMappingResponse, <-chan error) {
	responseChan := make(chan *ListAppGroupMappingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAppGroupMapping(request)
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

// ListAppGroupMappingWithCallback invokes the retailcloud.ListAppGroupMapping API asynchronously
func (client *Client) ListAppGroupMappingWithCallback(request *ListAppGroupMappingRequest, callback func(response *ListAppGroupMappingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAppGroupMappingResponse
		var err error
		defer close(result)
		response, err = client.ListAppGroupMapping(request)
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

// ListAppGroupMappingRequest is the request struct for api ListAppGroupMapping
type ListAppGroupMappingRequest struct {
	*requests.RpcRequest
	BizCode    string           `position:"Query" name:"BizCode"`
	Name       string           `position:"Query" name:"Name"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListAppGroupMappingResponse is the response struct for api ListAppGroupMapping
type ListAppGroupMappingResponse struct {
	*responses.BaseResponse
	Code       int                     `json:"Code" xml:"Code"`
	PageNumber int                     `json:"PageNumber" xml:"PageNumber"`
	RequestId  string                  `json:"RequestId" xml:"RequestId"`
	PageSize   int                     `json:"PageSize" xml:"PageSize"`
	TotalCount int64                   `json:"TotalCount" xml:"TotalCount"`
	ErrorMsg   string                  `json:"ErrorMsg" xml:"ErrorMsg"`
	Data       []AppGroupMappingDetail `json:"Data" xml:"Data"`
}

// CreateListAppGroupMappingRequest creates a request to invoke ListAppGroupMapping API
func CreateListAppGroupMappingRequest() (request *ListAppGroupMappingRequest) {
	request = &ListAppGroupMappingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ListAppGroupMapping", "", "")
	request.Method = requests.POST
	return
}

// CreateListAppGroupMappingResponse creates a response to parse from ListAppGroupMapping response
func CreateListAppGroupMappingResponse() (response *ListAppGroupMappingResponse) {
	response = &ListAppGroupMappingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
