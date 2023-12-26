package idrsservice

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

// GetTaskGroup invokes the idrsservice.GetTaskGroup API synchronously
func (client *Client) GetTaskGroup(request *GetTaskGroupRequest) (response *GetTaskGroupResponse, err error) {
	response = CreateGetTaskGroupResponse()
	err = client.DoAction(request, response)
	return
}

// GetTaskGroupWithChan invokes the idrsservice.GetTaskGroup API asynchronously
func (client *Client) GetTaskGroupWithChan(request *GetTaskGroupRequest) (<-chan *GetTaskGroupResponse, <-chan error) {
	responseChan := make(chan *GetTaskGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTaskGroup(request)
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

// GetTaskGroupWithCallback invokes the idrsservice.GetTaskGroup API asynchronously
func (client *Client) GetTaskGroupWithCallback(request *GetTaskGroupRequest, callback func(response *GetTaskGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTaskGroupResponse
		var err error
		defer close(result)
		response, err = client.GetTaskGroup(request)
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

// GetTaskGroupRequest is the request struct for api GetTaskGroup
type GetTaskGroupRequest struct {
	*requests.RpcRequest
	Id string `position:"Query" name:"Id"`
}

// GetTaskGroupResponse is the response struct for api GetTaskGroup
type GetTaskGroupResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetTaskGroupRequest creates a request to invoke GetTaskGroup API
func CreateGetTaskGroupRequest() (request *GetTaskGroupRequest) {
	request = &GetTaskGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "GetTaskGroup", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetTaskGroupResponse creates a response to parse from GetTaskGroup response
func CreateGetTaskGroupResponse() (response *GetTaskGroupResponse) {
	response = &GetTaskGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
