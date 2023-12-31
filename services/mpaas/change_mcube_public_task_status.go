package mpaas

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

// ChangeMcubePublicTaskStatus invokes the mpaas.ChangeMcubePublicTaskStatus API synchronously
func (client *Client) ChangeMcubePublicTaskStatus(request *ChangeMcubePublicTaskStatusRequest) (response *ChangeMcubePublicTaskStatusResponse, err error) {
	response = CreateChangeMcubePublicTaskStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeMcubePublicTaskStatusWithChan invokes the mpaas.ChangeMcubePublicTaskStatus API asynchronously
func (client *Client) ChangeMcubePublicTaskStatusWithChan(request *ChangeMcubePublicTaskStatusRequest) (<-chan *ChangeMcubePublicTaskStatusResponse, <-chan error) {
	responseChan := make(chan *ChangeMcubePublicTaskStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeMcubePublicTaskStatus(request)
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

// ChangeMcubePublicTaskStatusWithCallback invokes the mpaas.ChangeMcubePublicTaskStatus API asynchronously
func (client *Client) ChangeMcubePublicTaskStatusWithCallback(request *ChangeMcubePublicTaskStatusRequest, callback func(response *ChangeMcubePublicTaskStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeMcubePublicTaskStatusResponse
		var err error
		defer close(result)
		response, err = client.ChangeMcubePublicTaskStatus(request)
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

// ChangeMcubePublicTaskStatusRequest is the request struct for api ChangeMcubePublicTaskStatus
type ChangeMcubePublicTaskStatusRequest struct {
	*requests.RpcRequest
	TaskStatus  string `position:"Body" name:"TaskStatus"`
	TenantId    string `position:"Body" name:"TenantId"`
	TaskId      string `position:"Body" name:"TaskId"`
	AppId       string `position:"Body" name:"AppId"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// ChangeMcubePublicTaskStatusResponse is the response struct for api ChangeMcubePublicTaskStatus
type ChangeMcubePublicTaskStatusResponse struct {
	*responses.BaseResponse
	ResultMessage string                                     `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string                                     `json:"ResultCode" xml:"ResultCode"`
	RequestId     string                                     `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContentInChangeMcubePublicTaskStatus `json:"ResultContent" xml:"ResultContent"`
}

// CreateChangeMcubePublicTaskStatusRequest creates a request to invoke ChangeMcubePublicTaskStatus API
func CreateChangeMcubePublicTaskStatusRequest() (request *ChangeMcubePublicTaskStatusRequest) {
	request = &ChangeMcubePublicTaskStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "ChangeMcubePublicTaskStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateChangeMcubePublicTaskStatusResponse creates a response to parse from ChangeMcubePublicTaskStatus response
func CreateChangeMcubePublicTaskStatusResponse() (response *ChangeMcubePublicTaskStatusResponse) {
	response = &ChangeMcubePublicTaskStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
