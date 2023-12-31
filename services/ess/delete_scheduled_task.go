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

// DeleteScheduledTask invokes the ess.DeleteScheduledTask API synchronously
func (client *Client) DeleteScheduledTask(request *DeleteScheduledTaskRequest) (response *DeleteScheduledTaskResponse, err error) {
	response = CreateDeleteScheduledTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteScheduledTaskWithChan invokes the ess.DeleteScheduledTask API asynchronously
func (client *Client) DeleteScheduledTaskWithChan(request *DeleteScheduledTaskRequest) (<-chan *DeleteScheduledTaskResponse, <-chan error) {
	responseChan := make(chan *DeleteScheduledTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteScheduledTask(request)
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

// DeleteScheduledTaskWithCallback invokes the ess.DeleteScheduledTask API asynchronously
func (client *Client) DeleteScheduledTaskWithCallback(request *DeleteScheduledTaskRequest, callback func(response *DeleteScheduledTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteScheduledTaskResponse
		var err error
		defer close(result)
		response, err = client.DeleteScheduledTask(request)
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

// DeleteScheduledTaskRequest is the request struct for api DeleteScheduledTask
type DeleteScheduledTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ScheduledTaskId      string           `position:"Query" name:"ScheduledTaskId"`
}

// DeleteScheduledTaskResponse is the response struct for api DeleteScheduledTask
type DeleteScheduledTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteScheduledTaskRequest creates a request to invoke DeleteScheduledTask API
func CreateDeleteScheduledTaskRequest() (request *DeleteScheduledTaskRequest) {
	request = &DeleteScheduledTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DeleteScheduledTask", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteScheduledTaskResponse creates a response to parse from DeleteScheduledTask response
func CreateDeleteScheduledTaskResponse() (response *DeleteScheduledTaskResponse) {
	response = &DeleteScheduledTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
