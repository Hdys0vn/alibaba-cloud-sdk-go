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

// ListWorkspaces invokes the aiworkspace.ListWorkspaces API synchronously
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (response *ListWorkspacesResponse, err error) {
	response = CreateListWorkspacesResponse()
	err = client.DoAction(request, response)
	return
}

// ListWorkspacesWithChan invokes the aiworkspace.ListWorkspaces API asynchronously
func (client *Client) ListWorkspacesWithChan(request *ListWorkspacesRequest) (<-chan *ListWorkspacesResponse, <-chan error) {
	responseChan := make(chan *ListWorkspacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListWorkspaces(request)
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

// ListWorkspacesWithCallback invokes the aiworkspace.ListWorkspaces API asynchronously
func (client *Client) ListWorkspacesWithCallback(request *ListWorkspacesRequest, callback func(response *ListWorkspacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListWorkspacesResponse
		var err error
		defer close(result)
		response, err = client.ListWorkspaces(request)
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

// ListWorkspacesRequest is the request struct for api ListWorkspaces
type ListWorkspacesRequest struct {
	*requests.RoaRequest
	WorkspaceIds  string `position:"Query" name:"WorkspaceIds"`
	ModuleList    string `position:"Query" name:"ModuleList"`
	PageSize      string `position:"Query" name:"PageSize"`
	WorkspaceName string `position:"Query" name:"WorkspaceName"`
	SortBy        string `position:"Query" name:"SortBy"`
	Fields        string `position:"Query" name:"Fields"`
	PageNumber    string `position:"Query" name:"PageNumber"`
	Order         string `position:"Query" name:"Order"`
	Status        string `position:"Query" name:"Status"`
	Option        string `position:"Query" name:"Option"`
	Verbose       string `position:"Query" name:"Verbose"`
}

// ListWorkspacesResponse is the response struct for api ListWorkspaces
type ListWorkspacesResponse struct {
	*responses.BaseResponse
	RequestId      string                 `json:"RequestId" xml:"RequestId"`
	TotalCount     int64                  `json:"TotalCount" xml:"TotalCount"`
	ResourceLimits map[string]interface{} `json:"ResourceLimits" xml:"ResourceLimits"`
	Workspaces     []WorkspacesItem       `json:"Workspaces" xml:"Workspaces"`
}

// CreateListWorkspacesRequest creates a request to invoke ListWorkspaces API
func CreateListWorkspacesRequest() (request *ListWorkspacesRequest) {
	request = &ListWorkspacesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("AIWorkSpace", "2021-02-04", "ListWorkspaces", "/api/v1/workspaces", "", "")
	request.Method = requests.GET
	return
}

// CreateListWorkspacesResponse creates a response to parse from ListWorkspaces response
func CreateListWorkspacesResponse() (response *ListWorkspacesResponse) {
	response = &ListWorkspacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
