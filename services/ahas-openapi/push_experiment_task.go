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

// PushExperimentTask invokes the ahas_openapi.PushExperimentTask API synchronously
func (client *Client) PushExperimentTask(request *PushExperimentTaskRequest) (response *PushExperimentTaskResponse, err error) {
	response = CreatePushExperimentTaskResponse()
	err = client.DoAction(request, response)
	return
}

// PushExperimentTaskWithChan invokes the ahas_openapi.PushExperimentTask API asynchronously
func (client *Client) PushExperimentTaskWithChan(request *PushExperimentTaskRequest) (<-chan *PushExperimentTaskResponse, <-chan error) {
	responseChan := make(chan *PushExperimentTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushExperimentTask(request)
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

// PushExperimentTaskWithCallback invokes the ahas_openapi.PushExperimentTask API asynchronously
func (client *Client) PushExperimentTaskWithCallback(request *PushExperimentTaskRequest, callback func(response *PushExperimentTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushExperimentTaskResponse
		var err error
		defer close(result)
		response, err = client.PushExperimentTask(request)
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

// PushExperimentTaskRequest is the request struct for api PushExperimentTask
type PushExperimentTaskRequest struct {
	*requests.RpcRequest
	AhasRegionId     string `position:"Query" name:"AhasRegionId"`
	NameSpace        string `position:"Query" name:"NameSpace"`
	ExperimentTaskId string `position:"Query" name:"ExperimentTaskId"`
}

// PushExperimentTaskResponse is the response struct for api PushExperimentTask
type PushExperimentTaskResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreatePushExperimentTaskRequest creates a request to invoke PushExperimentTask API
func CreatePushExperimentTaskRequest() (request *PushExperimentTaskRequest) {
	request = &PushExperimentTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "PushExperimentTask", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePushExperimentTaskResponse creates a response to parse from PushExperimentTask response
func CreatePushExperimentTaskResponse() (response *PushExperimentTaskResponse) {
	response = &PushExperimentTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
