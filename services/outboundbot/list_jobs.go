package outboundbot

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

// ListJobs invokes the outboundbot.ListJobs API synchronously
func (client *Client) ListJobs(request *ListJobsRequest) (response *ListJobsResponse, err error) {
	response = CreateListJobsResponse()
	err = client.DoAction(request, response)
	return
}

// ListJobsWithChan invokes the outboundbot.ListJobs API asynchronously
func (client *Client) ListJobsWithChan(request *ListJobsRequest) (<-chan *ListJobsResponse, <-chan error) {
	responseChan := make(chan *ListJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJobs(request)
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

// ListJobsWithCallback invokes the outboundbot.ListJobs API asynchronously
func (client *Client) ListJobsWithCallback(request *ListJobsRequest, callback func(response *ListJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobsResponse
		var err error
		defer close(result)
		response, err = client.ListJobs(request)
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

// ListJobsRequest is the request struct for api ListJobs
type ListJobsRequest struct {
	*requests.RpcRequest
	JobId      *[]string `position:"Query" name:"JobId"  type:"Repeated"`
	InstanceId string    `position:"Query" name:"InstanceId"`
}

// ListJobsResponse is the response struct for api ListJobs
type ListJobsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Jobs           []Job  `json:"Jobs" xml:"Jobs"`
}

// CreateListJobsRequest creates a request to invoke ListJobs API
func CreateListJobsRequest() (request *ListJobsRequest) {
	request = &ListJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ListJobs", "", "")
	request.Method = requests.POST
	return
}

// CreateListJobsResponse creates a response to parse from ListJobs response
func CreateListJobsResponse() (response *ListJobsResponse) {
	response = &ListJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}