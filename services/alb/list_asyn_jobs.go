package alb

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

// ListAsynJobs invokes the alb.ListAsynJobs API synchronously
func (client *Client) ListAsynJobs(request *ListAsynJobsRequest) (response *ListAsynJobsResponse, err error) {
	response = CreateListAsynJobsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAsynJobsWithChan invokes the alb.ListAsynJobs API asynchronously
func (client *Client) ListAsynJobsWithChan(request *ListAsynJobsRequest) (<-chan *ListAsynJobsResponse, <-chan error) {
	responseChan := make(chan *ListAsynJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAsynJobs(request)
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

// ListAsynJobsWithCallback invokes the alb.ListAsynJobs API asynchronously
func (client *Client) ListAsynJobsWithCallback(request *ListAsynJobsRequest, callback func(response *ListAsynJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAsynJobsResponse
		var err error
		defer close(result)
		response, err = client.ListAsynJobs(request)
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

// ListAsynJobsRequest is the request struct for api ListAsynJobs
type ListAsynJobsRequest struct {
	*requests.RpcRequest
	NextToken    string           `position:"Query" name:"NextToken"`
	EndTime      requests.Integer `position:"Query" name:"EndTime"`
	BeginTime    requests.Integer `position:"Query" name:"BeginTime"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	ApiName      string           `position:"Query" name:"ApiName"`
	JobIds       *[]string        `position:"Query" name:"JobIds"  type:"Repeated"`
	MaxResults   requests.Integer `position:"Query" name:"MaxResults"`
	ResourceIds  *[]string        `position:"Query" name:"ResourceIds"  type:"Repeated"`
}

// ListAsynJobsResponse is the response struct for api ListAsynJobs
type ListAsynJobsResponse struct {
	*responses.BaseResponse
	MaxResults int64  `json:"MaxResults" xml:"MaxResults"`
	NextToken  string `json:"NextToken" xml:"NextToken"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int64  `json:"TotalCount" xml:"TotalCount"`
	Jobs       []Job  `json:"Jobs" xml:"Jobs"`
}

// CreateListAsynJobsRequest creates a request to invoke ListAsynJobs API
func CreateListAsynJobsRequest() (request *ListAsynJobsRequest) {
	request = &ListAsynJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "ListAsynJobs", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAsynJobsResponse creates a response to parse from ListAsynJobs response
func CreateListAsynJobsResponse() (response *ListAsynJobsResponse) {
	response = &ListAsynJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}