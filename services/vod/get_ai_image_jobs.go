package vod

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

// GetAIImageJobs invokes the vod.GetAIImageJobs API synchronously
func (client *Client) GetAIImageJobs(request *GetAIImageJobsRequest) (response *GetAIImageJobsResponse, err error) {
	response = CreateGetAIImageJobsResponse()
	err = client.DoAction(request, response)
	return
}

// GetAIImageJobsWithChan invokes the vod.GetAIImageJobs API asynchronously
func (client *Client) GetAIImageJobsWithChan(request *GetAIImageJobsRequest) (<-chan *GetAIImageJobsResponse, <-chan error) {
	responseChan := make(chan *GetAIImageJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAIImageJobs(request)
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

// GetAIImageJobsWithCallback invokes the vod.GetAIImageJobs API asynchronously
func (client *Client) GetAIImageJobsWithCallback(request *GetAIImageJobsRequest, callback func(response *GetAIImageJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAIImageJobsResponse
		var err error
		defer close(result)
		response, err = client.GetAIImageJobs(request)
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

// GetAIImageJobsRequest is the request struct for api GetAIImageJobs
type GetAIImageJobsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	JobIds               string `position:"Query" name:"JobIds"`
}

// GetAIImageJobsResponse is the response struct for api GetAIImageJobs
type GetAIImageJobsResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	AIImageJobList []AIImageJob `json:"AIImageJobList" xml:"AIImageJobList"`
}

// CreateGetAIImageJobsRequest creates a request to invoke GetAIImageJobs API
func CreateGetAIImageJobsRequest() (request *GetAIImageJobsRequest) {
	request = &GetAIImageJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetAIImageJobs", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAIImageJobsResponse creates a response to parse from GetAIImageJobs response
func CreateGetAIImageJobsResponse() (response *GetAIImageJobsResponse) {
	response = &GetAIImageJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
