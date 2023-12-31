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

// GetInstanceCheckpoint invokes the foas.GetInstanceCheckpoint API synchronously
func (client *Client) GetInstanceCheckpoint(request *GetInstanceCheckpointRequest) (response *GetInstanceCheckpointResponse, err error) {
	response = CreateGetInstanceCheckpointResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceCheckpointWithChan invokes the foas.GetInstanceCheckpoint API asynchronously
func (client *Client) GetInstanceCheckpointWithChan(request *GetInstanceCheckpointRequest) (<-chan *GetInstanceCheckpointResponse, <-chan error) {
	responseChan := make(chan *GetInstanceCheckpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceCheckpoint(request)
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

// GetInstanceCheckpointWithCallback invokes the foas.GetInstanceCheckpoint API asynchronously
func (client *Client) GetInstanceCheckpointWithCallback(request *GetInstanceCheckpointRequest, callback func(response *GetInstanceCheckpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceCheckpointResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceCheckpoint(request)
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

// GetInstanceCheckpointRequest is the request struct for api GetInstanceCheckpoint
type GetInstanceCheckpointRequest struct {
	*requests.RoaRequest
	ProjectName string           `position:"Path" name:"projectName"`
	InstanceId  requests.Integer `position:"Path" name:"instanceId"`
	JobName     string           `position:"Path" name:"jobName"`
}

// GetInstanceCheckpointResponse is the response struct for api GetInstanceCheckpoint
type GetInstanceCheckpointResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Checkpoints string `json:"Checkpoints" xml:"Checkpoints"`
}

// CreateGetInstanceCheckpointRequest creates a request to invoke GetInstanceCheckpoint API
func CreateGetInstanceCheckpointRequest() (request *GetInstanceCheckpointRequest) {
	request = &GetInstanceCheckpointRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "GetInstanceCheckpoint", "/api/v2/projects/[projectName]/jobs/[jobName]/instances/[instanceId]/checkpoints", "foas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetInstanceCheckpointResponse creates a response to parse from GetInstanceCheckpoint response
func CreateGetInstanceCheckpointResponse() (response *GetInstanceCheckpointResponse) {
	response = &GetInstanceCheckpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
