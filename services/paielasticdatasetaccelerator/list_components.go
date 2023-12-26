package paielasticdatasetaccelerator

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

// ListComponents invokes the paielasticdatasetaccelerator.ListComponents API synchronously
func (client *Client) ListComponents(request *ListComponentsRequest) (response *ListComponentsResponse, err error) {
	response = CreateListComponentsResponse()
	err = client.DoAction(request, response)
	return
}

// ListComponentsWithChan invokes the paielasticdatasetaccelerator.ListComponents API asynchronously
func (client *Client) ListComponentsWithChan(request *ListComponentsRequest) (<-chan *ListComponentsResponse, <-chan error) {
	responseChan := make(chan *ListComponentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListComponents(request)
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

// ListComponentsWithCallback invokes the paielasticdatasetaccelerator.ListComponents API asynchronously
func (client *Client) ListComponentsWithCallback(request *ListComponentsRequest, callback func(response *ListComponentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListComponentsResponse
		var err error
		defer close(result)
		response, err = client.ListComponents(request)
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

// ListComponentsRequest is the request struct for api ListComponents
type ListComponentsRequest struct {
	*requests.RoaRequest
	ComponentIds string           `position:"Query" name:"ComponentIds"`
	Name         string           `position:"Query" name:"Name"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	SortBy       string           `position:"Query" name:"SortBy"`
	Version      string           `position:"Query" name:"Version"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	Order        string           `position:"Query" name:"Order"`
}

// ListComponentsResponse is the response struct for api ListComponents
type ListComponentsResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	TotalCount int              `json:"TotalCount" xml:"TotalCount"`
	Components []ComponentsItem `json:"Components" xml:"Components"`
}

// CreateListComponentsRequest creates a request to invoke ListComponents API
func CreateListComponentsRequest() (request *ListComponentsRequest) {
	request = &ListComponentsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PAIElasticDatasetAccelerator", "2022-08-01", "ListComponents", "/api/v1/components", "datasetacc", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListComponentsResponse creates a response to parse from ListComponents response
func CreateListComponentsResponse() (response *ListComponentsResponse) {
	response = &ListComponentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
