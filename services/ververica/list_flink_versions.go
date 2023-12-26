package ververica

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

// ListFlinkVersions invokes the ververica.ListFlinkVersions API synchronously
func (client *Client) ListFlinkVersions(request *ListFlinkVersionsRequest) (response *ListFlinkVersionsResponse, err error) {
	response = CreateListFlinkVersionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListFlinkVersionsWithChan invokes the ververica.ListFlinkVersions API asynchronously
func (client *Client) ListFlinkVersionsWithChan(request *ListFlinkVersionsRequest) (<-chan *ListFlinkVersionsResponse, <-chan error) {
	responseChan := make(chan *ListFlinkVersionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFlinkVersions(request)
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

// ListFlinkVersionsWithCallback invokes the ververica.ListFlinkVersions API asynchronously
func (client *Client) ListFlinkVersionsWithCallback(request *ListFlinkVersionsRequest, callback func(response *ListFlinkVersionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFlinkVersionsResponse
		var err error
		defer close(result)
		response, err = client.ListFlinkVersions(request)
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

// ListFlinkVersionsRequest is the request struct for api ListFlinkVersions
type ListFlinkVersionsRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
}

// ListFlinkVersionsResponse is the response struct for api ListFlinkVersions
type ListFlinkVersionsResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateListFlinkVersionsRequest creates a request to invoke ListFlinkVersions API
func CreateListFlinkVersionsRequest() (request *ListFlinkVersionsRequest) {
	request = &ListFlinkVersionsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "ListFlinkVersions", "/pop/workspaces/[workspace]/api/v1/flink-version-meta.json", "", "")
	request.Method = requests.GET
	return
}

// CreateListFlinkVersionsResponse creates a response to parse from ListFlinkVersions response
func CreateListFlinkVersionsResponse() (response *ListFlinkVersionsResponse) {
	response = &ListFlinkVersionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}