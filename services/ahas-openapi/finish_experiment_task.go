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

// FinishExperimentTask invokes the ahas_openapi.FinishExperimentTask API synchronously
func (client *Client) FinishExperimentTask(request *FinishExperimentTaskRequest) (response *FinishExperimentTaskResponse, err error) {
	response = CreateFinishExperimentTaskResponse()
	err = client.DoAction(request, response)
	return
}

// FinishExperimentTaskWithChan invokes the ahas_openapi.FinishExperimentTask API asynchronously
func (client *Client) FinishExperimentTaskWithChan(request *FinishExperimentTaskRequest) (<-chan *FinishExperimentTaskResponse, <-chan error) {
	responseChan := make(chan *FinishExperimentTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FinishExperimentTask(request)
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

// FinishExperimentTaskWithCallback invokes the ahas_openapi.FinishExperimentTask API asynchronously
func (client *Client) FinishExperimentTaskWithCallback(request *FinishExperimentTaskRequest, callback func(response *FinishExperimentTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FinishExperimentTaskResponse
		var err error
		defer close(result)
		response, err = client.FinishExperimentTask(request)
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

// FinishExperimentTaskRequest is the request struct for api FinishExperimentTask
type FinishExperimentTaskRequest struct {
	*requests.RpcRequest
	AhasRegionId     string `position:"Query" name:"AhasRegionId"`
	NameSpace        string `position:"Query" name:"NameSpace"`
	ExperimentTaskId string `position:"Query" name:"ExperimentTaskId"`
}

// FinishExperimentTaskResponse is the response struct for api FinishExperimentTask
type FinishExperimentTaskResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateFinishExperimentTaskRequest creates a request to invoke FinishExperimentTask API
func CreateFinishExperimentTaskRequest() (request *FinishExperimentTaskRequest) {
	request = &FinishExperimentTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "FinishExperimentTask", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateFinishExperimentTaskResponse creates a response to parse from FinishExperimentTask response
func CreateFinishExperimentTaskResponse() (response *FinishExperimentTaskResponse) {
	response = &FinishExperimentTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
