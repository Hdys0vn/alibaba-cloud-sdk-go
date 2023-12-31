package dms_enterprise

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

// ListDAGVersions invokes the dms_enterprise.ListDAGVersions API synchronously
func (client *Client) ListDAGVersions(request *ListDAGVersionsRequest) (response *ListDAGVersionsResponse, err error) {
	response = CreateListDAGVersionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDAGVersionsWithChan invokes the dms_enterprise.ListDAGVersions API asynchronously
func (client *Client) ListDAGVersionsWithChan(request *ListDAGVersionsRequest) (<-chan *ListDAGVersionsResponse, <-chan error) {
	responseChan := make(chan *ListDAGVersionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDAGVersions(request)
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

// ListDAGVersionsWithCallback invokes the dms_enterprise.ListDAGVersions API asynchronously
func (client *Client) ListDAGVersionsWithCallback(request *ListDAGVersionsRequest, callback func(response *ListDAGVersionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDAGVersionsResponse
		var err error
		defer close(result)
		response, err = client.ListDAGVersions(request)
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

// ListDAGVersionsRequest is the request struct for api ListDAGVersions
type ListDAGVersionsRequest struct {
	*requests.RpcRequest
	DagId     requests.Integer `position:"Query" name:"DagId"`
	Tid       requests.Integer `position:"Query" name:"Tid"`
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	PageIndex requests.Integer `position:"Query" name:"PageIndex"`
}

// ListDAGVersionsResponse is the response struct for api ListDAGVersions
type ListDAGVersionsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	ErrorCode      string         `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string         `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool           `json:"Success" xml:"Success"`
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	DagVersionList DagVersionList `json:"DagVersionList" xml:"DagVersionList"`
}

// CreateListDAGVersionsRequest creates a request to invoke ListDAGVersions API
func CreateListDAGVersionsRequest() (request *ListDAGVersionsRequest) {
	request = &ListDAGVersionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListDAGVersions", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDAGVersionsResponse creates a response to parse from ListDAGVersions response
func CreateListDAGVersionsResponse() (response *ListDAGVersionsResponse) {
	response = &ListDAGVersionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
