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

// SuspendJobs invokes the outboundbot.SuspendJobs API synchronously
func (client *Client) SuspendJobs(request *SuspendJobsRequest) (response *SuspendJobsResponse, err error) {
	response = CreateSuspendJobsResponse()
	err = client.DoAction(request, response)
	return
}

// SuspendJobsWithChan invokes the outboundbot.SuspendJobs API asynchronously
func (client *Client) SuspendJobsWithChan(request *SuspendJobsRequest) (<-chan *SuspendJobsResponse, <-chan error) {
	responseChan := make(chan *SuspendJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SuspendJobs(request)
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

// SuspendJobsWithCallback invokes the outboundbot.SuspendJobs API asynchronously
func (client *Client) SuspendJobsWithCallback(request *SuspendJobsRequest, callback func(response *SuspendJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SuspendJobsResponse
		var err error
		defer close(result)
		response, err = client.SuspendJobs(request)
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

// SuspendJobsRequest is the request struct for api SuspendJobs
type SuspendJobsRequest struct {
	*requests.RpcRequest
	All            requests.Boolean `position:"Query" name:"All"`
	JobReferenceId *[]string        `position:"Query" name:"JobReferenceId"  type:"Repeated"`
	JobId          *[]string        `position:"Query" name:"JobId"  type:"Repeated"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
	JobGroupId     string           `position:"Query" name:"JobGroupId"`
	ScenarioId     string           `position:"Query" name:"ScenarioId"`
}

// SuspendJobsResponse is the response struct for api SuspendJobs
type SuspendJobsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateSuspendJobsRequest creates a request to invoke SuspendJobs API
func CreateSuspendJobsRequest() (request *SuspendJobsRequest) {
	request = &SuspendJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "SuspendJobs", "", "")
	request.Method = requests.POST
	return
}

// CreateSuspendJobsResponse creates a response to parse from SuspendJobs response
func CreateSuspendJobsResponse() (response *SuspendJobsResponse) {
	response = &SuspendJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
