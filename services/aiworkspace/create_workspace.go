package aiworkspace

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

// CreateWorkspace invokes the aiworkspace.CreateWorkspace API synchronously
func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (response *CreateWorkspaceResponse, err error) {
	response = CreateCreateWorkspaceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateWorkspaceWithChan invokes the aiworkspace.CreateWorkspace API asynchronously
func (client *Client) CreateWorkspaceWithChan(request *CreateWorkspaceRequest) (<-chan *CreateWorkspaceResponse, <-chan error) {
	responseChan := make(chan *CreateWorkspaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateWorkspace(request)
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

// CreateWorkspaceWithCallback invokes the aiworkspace.CreateWorkspace API asynchronously
func (client *Client) CreateWorkspaceWithCallback(request *CreateWorkspaceRequest, callback func(response *CreateWorkspaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateWorkspaceResponse
		var err error
		defer close(result)
		response, err = client.CreateWorkspace(request)
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

// CreateWorkspaceRequest is the request struct for api CreateWorkspace
type CreateWorkspaceRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateWorkspaceResponse is the response struct for api CreateWorkspace
type CreateWorkspaceResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	WorkspaceId string `json:"WorkspaceId" xml:"WorkspaceId"`
}

// CreateCreateWorkspaceRequest creates a request to invoke CreateWorkspace API
func CreateCreateWorkspaceRequest() (request *CreateWorkspaceRequest) {
	request = &CreateWorkspaceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("AIWorkSpace", "2021-02-04", "CreateWorkspace", "/api/v1/workspaces", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateWorkspaceResponse creates a response to parse from CreateWorkspace response
func CreateCreateWorkspaceResponse() (response *CreateWorkspaceResponse) {
	response = &CreateWorkspaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}