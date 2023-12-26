package paifeaturestore

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

// ListLabelTables invokes the paifeaturestore.ListLabelTables API synchronously
func (client *Client) ListLabelTables(request *ListLabelTablesRequest) (response *ListLabelTablesResponse, err error) {
	response = CreateListLabelTablesResponse()
	err = client.DoAction(request, response)
	return
}

// ListLabelTablesWithChan invokes the paifeaturestore.ListLabelTables API asynchronously
func (client *Client) ListLabelTablesWithChan(request *ListLabelTablesRequest) (<-chan *ListLabelTablesResponse, <-chan error) {
	responseChan := make(chan *ListLabelTablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLabelTables(request)
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

// ListLabelTablesWithCallback invokes the paifeaturestore.ListLabelTables API asynchronously
func (client *Client) ListLabelTablesWithCallback(request *ListLabelTablesRequest, callback func(response *ListLabelTablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLabelTablesResponse
		var err error
		defer close(result)
		response, err = client.ListLabelTables(request)
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

// ListLabelTablesRequest is the request struct for api ListLabelTables
type ListLabelTablesRequest struct {
	*requests.RoaRequest
	Owner      string           `position:"Query" name:"Owner"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Path" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Name       string           `position:"Query" name:"Name"`
	SortBy     string           `position:"Query" name:"SortBy"`
	ProjectId  string           `position:"Query" name:"ProjectId"`
	Order      string           `position:"Query" name:"Order"`
}

// ListLabelTablesResponse is the response struct for api ListLabelTables
type ListLabelTablesResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	TotalCount  int64        `json:"TotalCount" xml:"TotalCount"`
	LabelTables []LabelTable `json:"LabelTables" xml:"LabelTables"`
}

// CreateListLabelTablesRequest creates a request to invoke ListLabelTables API
func CreateListLabelTablesRequest() (request *ListLabelTablesRequest) {
	request = &ListLabelTablesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "ListLabelTables", "/api/v1/instances/[InstanceId]/labeltables", "", "")
	request.Method = requests.GET
	return
}

// CreateListLabelTablesResponse creates a response to parse from ListLabelTables response
func CreateListLabelTablesResponse() (response *ListLabelTablesResponse) {
	response = &ListLabelTablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}