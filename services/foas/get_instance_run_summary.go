package foas

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

// GetInstanceRunSummary invokes the foas.GetInstanceRunSummary API synchronously
func (client *Client) GetInstanceRunSummary(request *GetInstanceRunSummaryRequest) (response *GetInstanceRunSummaryResponse, err error) {
	response = CreateGetInstanceRunSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceRunSummaryWithChan invokes the foas.GetInstanceRunSummary API asynchronously
func (client *Client) GetInstanceRunSummaryWithChan(request *GetInstanceRunSummaryRequest) (<-chan *GetInstanceRunSummaryResponse, <-chan error) {
	responseChan := make(chan *GetInstanceRunSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceRunSummary(request)
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

// GetInstanceRunSummaryWithCallback invokes the foas.GetInstanceRunSummary API asynchronously
func (client *Client) GetInstanceRunSummaryWithCallback(request *GetInstanceRunSummaryRequest, callback func(response *GetInstanceRunSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceRunSummaryResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceRunSummary(request)
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

// GetInstanceRunSummaryRequest is the request struct for api GetInstanceRunSummary
type GetInstanceRunSummaryRequest struct {
	*requests.RoaRequest
	ProjectName string           `position:"Path" name:"projectName"`
	InstanceId  requests.Integer `position:"Path" name:"instanceId"`
	JobName     string           `position:"Path" name:"jobName"`
}

// GetInstanceRunSummaryResponse is the response struct for api GetInstanceRunSummary
type GetInstanceRunSummaryResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	RunSummary RunSummary `json:"RunSummary" xml:"RunSummary"`
}

// CreateGetInstanceRunSummaryRequest creates a request to invoke GetInstanceRunSummary API
func CreateGetInstanceRunSummaryRequest() (request *GetInstanceRunSummaryRequest) {
	request = &GetInstanceRunSummaryRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "GetInstanceRunSummary", "/api/v2/projects/[projectName]/jobs/[jobName]/instances/[instanceId]/runsummary", "foas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetInstanceRunSummaryResponse creates a response to parse from GetInstanceRunSummary response
func CreateGetInstanceRunSummaryResponse() (response *GetInstanceRunSummaryResponse) {
	response = &GetInstanceRunSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
