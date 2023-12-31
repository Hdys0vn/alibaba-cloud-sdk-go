package viapi_regen

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

// GetWorkspace invokes the viapi_regen.GetWorkspace API synchronously
func (client *Client) GetWorkspace(request *GetWorkspaceRequest) (response *GetWorkspaceResponse, err error) {
	response = CreateGetWorkspaceResponse()
	err = client.DoAction(request, response)
	return
}

// GetWorkspaceWithChan invokes the viapi_regen.GetWorkspace API asynchronously
func (client *Client) GetWorkspaceWithChan(request *GetWorkspaceRequest) (<-chan *GetWorkspaceResponse, <-chan error) {
	responseChan := make(chan *GetWorkspaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWorkspace(request)
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

// GetWorkspaceWithCallback invokes the viapi_regen.GetWorkspace API asynchronously
func (client *Client) GetWorkspaceWithCallback(request *GetWorkspaceRequest, callback func(response *GetWorkspaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWorkspaceResponse
		var err error
		defer close(result)
		response, err = client.GetWorkspace(request)
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

// GetWorkspaceRequest is the request struct for api GetWorkspace
type GetWorkspaceRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Body" name:"Id"`
}

// GetWorkspaceResponse is the response struct for api GetWorkspace
type GetWorkspaceResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetWorkspaceRequest creates a request to invoke GetWorkspace API
func CreateGetWorkspaceRequest() (request *GetWorkspaceRequest) {
	request = &GetWorkspaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "GetWorkspace", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetWorkspaceResponse creates a response to parse from GetWorkspace response
func CreateGetWorkspaceResponse() (response *GetWorkspaceResponse) {
	response = &GetWorkspaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
