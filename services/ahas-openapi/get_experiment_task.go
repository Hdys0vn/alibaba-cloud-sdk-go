package ahas_openapi

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

// GetExperimentTask invokes the ahas_openapi.GetExperimentTask API synchronously
func (client *Client) GetExperimentTask(request *GetExperimentTaskRequest) (response *GetExperimentTaskResponse, err error) {
	response = CreateGetExperimentTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetExperimentTaskWithChan invokes the ahas_openapi.GetExperimentTask API asynchronously
func (client *Client) GetExperimentTaskWithChan(request *GetExperimentTaskRequest) (<-chan *GetExperimentTaskResponse, <-chan error) {
	responseChan := make(chan *GetExperimentTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetExperimentTask(request)
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

// GetExperimentTaskWithCallback invokes the ahas_openapi.GetExperimentTask API asynchronously
func (client *Client) GetExperimentTaskWithCallback(request *GetExperimentTaskRequest, callback func(response *GetExperimentTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetExperimentTaskResponse
		var err error
		defer close(result)
		response, err = client.GetExperimentTask(request)
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

// GetExperimentTaskRequest is the request struct for api GetExperimentTask
type GetExperimentTaskRequest struct {
	*requests.RpcRequest
	AhasRegionId     string `position:"Query" name:"AhasRegionId"`
	NameSpace        string `position:"Query" name:"NameSpace"`
	ExperimentTaskId string `position:"Query" name:"ExperimentTaskId"`
}

// GetExperimentTaskResponse is the response struct for api GetExperimentTask
type GetExperimentTaskResponse struct {
	*responses.BaseResponse
	TaskId         string           `json:"TaskId" xml:"TaskId"`
	RequestId      string           `json:"RequestId" xml:"RequestId"`
	ExperimentName string           `json:"ExperimentName" xml:"ExperimentName"`
	State          string           `json:"State" xml:"State"`
	ExperimentId   string           `json:"ExperimentId" xml:"ExperimentId"`
	HttpStatusCode int              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	StartTime      int64            `json:"StartTime" xml:"StartTime"`
	Success        bool             `json:"Success" xml:"Success"`
	Result         string           `json:"Result" xml:"Result"`
	Namespace      string           `json:"Namespace" xml:"Namespace"`
	Activities     []ActivitiesItem `json:"Activities" xml:"Activities"`
}

// CreateGetExperimentTaskRequest creates a request to invoke GetExperimentTask API
func CreateGetExperimentTaskRequest() (request *GetExperimentTaskRequest) {
	request = &GetExperimentTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "GetExperimentTask", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetExperimentTaskResponse creates a response to parse from GetExperimentTask response
func CreateGetExperimentTaskResponse() (response *GetExperimentTaskResponse) {
	response = &GetExperimentTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
