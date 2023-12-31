package mts

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

// SubmitBeautifyJobs invokes the mts.SubmitBeautifyJobs API synchronously
func (client *Client) SubmitBeautifyJobs(request *SubmitBeautifyJobsRequest) (response *SubmitBeautifyJobsResponse, err error) {
	response = CreateSubmitBeautifyJobsResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitBeautifyJobsWithChan invokes the mts.SubmitBeautifyJobs API asynchronously
func (client *Client) SubmitBeautifyJobsWithChan(request *SubmitBeautifyJobsRequest) (<-chan *SubmitBeautifyJobsResponse, <-chan error) {
	responseChan := make(chan *SubmitBeautifyJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitBeautifyJobs(request)
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

// SubmitBeautifyJobsWithCallback invokes the mts.SubmitBeautifyJobs API asynchronously
func (client *Client) SubmitBeautifyJobsWithCallback(request *SubmitBeautifyJobsRequest, callback func(response *SubmitBeautifyJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitBeautifyJobsResponse
		var err error
		defer close(result)
		response, err = client.SubmitBeautifyJobs(request)
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

// SubmitBeautifyJobsRequest is the request struct for api SubmitBeautifyJobs
type SubmitBeautifyJobsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BeautifyConfig       string           `position:"Query" name:"BeautifyConfig"`
	UserData             string           `position:"Query" name:"UserData"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	Async                requests.Boolean `position:"Query" name:"Async"`
}

// SubmitBeautifyJobsResponse is the response struct for api SubmitBeautifyJobs
type SubmitBeautifyJobsResponse struct {
	*responses.BaseResponse
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	JobList   JobListInSubmitBeautifyJobs `json:"JobList" xml:"JobList"`
}

// CreateSubmitBeautifyJobsRequest creates a request to invoke SubmitBeautifyJobs API
func CreateSubmitBeautifyJobsRequest() (request *SubmitBeautifyJobsRequest) {
	request = &SubmitBeautifyJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitBeautifyJobs", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitBeautifyJobsResponse creates a response to parse from SubmitBeautifyJobs response
func CreateSubmitBeautifyJobsResponse() (response *SubmitBeautifyJobsResponse) {
	response = &SubmitBeautifyJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
