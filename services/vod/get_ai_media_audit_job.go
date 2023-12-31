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

// GetAIMediaAuditJob invokes the vod.GetAIMediaAuditJob API synchronously
func (client *Client) GetAIMediaAuditJob(request *GetAIMediaAuditJobRequest) (response *GetAIMediaAuditJobResponse, err error) {
	response = CreateGetAIMediaAuditJobResponse()
	err = client.DoAction(request, response)
	return
}

// GetAIMediaAuditJobWithChan invokes the vod.GetAIMediaAuditJob API asynchronously
func (client *Client) GetAIMediaAuditJobWithChan(request *GetAIMediaAuditJobRequest) (<-chan *GetAIMediaAuditJobResponse, <-chan error) {
	responseChan := make(chan *GetAIMediaAuditJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAIMediaAuditJob(request)
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

// GetAIMediaAuditJobWithCallback invokes the vod.GetAIMediaAuditJob API asynchronously
func (client *Client) GetAIMediaAuditJobWithCallback(request *GetAIMediaAuditJobRequest, callback func(response *GetAIMediaAuditJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAIMediaAuditJobResponse
		var err error
		defer close(result)
		response, err = client.GetAIMediaAuditJob(request)
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

// GetAIMediaAuditJobRequest is the request struct for api GetAIMediaAuditJob
type GetAIMediaAuditJobRequest struct {
	*requests.RpcRequest
	JobId string `position:"Query" name:"JobId"`
}

// GetAIMediaAuditJobResponse is the response struct for api GetAIMediaAuditJob
type GetAIMediaAuditJobResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	MediaAuditJob MediaAuditJob `json:"MediaAuditJob" xml:"MediaAuditJob"`
}

// CreateGetAIMediaAuditJobRequest creates a request to invoke GetAIMediaAuditJob API
func CreateGetAIMediaAuditJobRequest() (request *GetAIMediaAuditJobRequest) {
	request = &GetAIMediaAuditJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetAIMediaAuditJob", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAIMediaAuditJobResponse creates a response to parse from GetAIMediaAuditJob response
func CreateGetAIMediaAuditJobResponse() (response *GetAIMediaAuditJobResponse) {
	response = &GetAIMediaAuditJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
