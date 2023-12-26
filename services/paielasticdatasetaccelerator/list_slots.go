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

// ListSlots invokes the paielasticdatasetaccelerator.ListSlots API synchronously
func (client *Client) ListSlots(request *ListSlotsRequest) (response *ListSlotsResponse, err error) {
	response = CreateListSlotsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSlotsWithChan invokes the paielasticdatasetaccelerator.ListSlots API asynchronously
func (client *Client) ListSlotsWithChan(request *ListSlotsRequest) (<-chan *ListSlotsResponse, <-chan error) {
	responseChan := make(chan *ListSlotsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSlots(request)
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

// ListSlotsWithCallback invokes the paielasticdatasetaccelerator.ListSlots API asynchronously
func (client *Client) ListSlotsWithCallback(request *ListSlotsRequest, callback func(response *ListSlotsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSlotsResponse
		var err error
		defer close(result)
		response, err = client.ListSlots(request)
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

// ListSlotsRequest is the request struct for api ListSlots
type ListSlotsRequest struct {
	*requests.RoaRequest
	Phase       string           `position:"Query" name:"Phase"`
	StorageType string           `position:"Query" name:"StorageType"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	EndpointIds string           `position:"Query" name:"EndpointIds"`
	SlotIds     string           `position:"Query" name:"SlotIds"`
	InstanceIds string           `position:"Query" name:"InstanceIds"`
	Name        string           `position:"Query" name:"Name"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	SortBy      string           `position:"Query" name:"SortBy"`
	StorageUri  string           `position:"Query" name:"StorageUri"`
	Order       string           `position:"Query" name:"Order"`
}

// ListSlotsResponse is the response struct for api ListSlots
type ListSlotsResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	TotalCount int         `json:"TotalCount" xml:"TotalCount"`
	Slots      []SlotsItem `json:"Slots" xml:"Slots"`
}

// CreateListSlotsRequest creates a request to invoke ListSlots API
func CreateListSlotsRequest() (request *ListSlotsRequest) {
	request = &ListSlotsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PAIElasticDatasetAccelerator", "2022-08-01", "ListSlots", "/api/v1/slots", "datasetacc", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListSlotsResponse creates a response to parse from ListSlots response
func CreateListSlotsResponse() (response *ListSlotsResponse) {
	response = &ListSlotsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}