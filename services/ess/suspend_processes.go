package ess

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

// SuspendProcesses invokes the ess.SuspendProcesses API synchronously
func (client *Client) SuspendProcesses(request *SuspendProcessesRequest) (response *SuspendProcessesResponse, err error) {
	response = CreateSuspendProcessesResponse()
	err = client.DoAction(request, response)
	return
}

// SuspendProcessesWithChan invokes the ess.SuspendProcesses API asynchronously
func (client *Client) SuspendProcessesWithChan(request *SuspendProcessesRequest) (<-chan *SuspendProcessesResponse, <-chan error) {
	responseChan := make(chan *SuspendProcessesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SuspendProcesses(request)
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

// SuspendProcessesWithCallback invokes the ess.SuspendProcesses API asynchronously
func (client *Client) SuspendProcessesWithCallback(request *SuspendProcessesRequest, callback func(response *SuspendProcessesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SuspendProcessesResponse
		var err error
		defer close(result)
		response, err = client.SuspendProcesses(request)
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

// SuspendProcessesRequest is the request struct for api SuspendProcesses
type SuspendProcessesRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	Process              *[]string        `position:"Query" name:"Process"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// SuspendProcessesResponse is the response struct for api SuspendProcesses
type SuspendProcessesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSuspendProcessesRequest creates a request to invoke SuspendProcesses API
func CreateSuspendProcessesRequest() (request *SuspendProcessesRequest) {
	request = &SuspendProcessesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "SuspendProcesses", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSuspendProcessesResponse creates a response to parse from SuspendProcesses response
func CreateSuspendProcessesResponse() (response *SuspendProcessesResponse) {
	response = &SuspendProcessesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
