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

// ListWorkspaceUsers invokes the aiworkspace.ListWorkspaceUsers API synchronously
func (client *Client) ListWorkspaceUsers(request *ListWorkspaceUsersRequest) (response *ListWorkspaceUsersResponse, err error) {
	response = CreateListWorkspaceUsersResponse()
	err = client.DoAction(request, response)
	return
}

// ListWorkspaceUsersWithChan invokes the aiworkspace.ListWorkspaceUsers API asynchronously
func (client *Client) ListWorkspaceUsersWithChan(request *ListWorkspaceUsersRequest) (<-chan *ListWorkspaceUsersResponse, <-chan error) {
	responseChan := make(chan *ListWorkspaceUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListWorkspaceUsers(request)
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

// ListWorkspaceUsersWithCallback invokes the aiworkspace.ListWorkspaceUsers API asynchronously
func (client *Client) ListWorkspaceUsersWithCallback(request *ListWorkspaceUsersRequest, callback func(response *ListWorkspaceUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListWorkspaceUsersResponse
		var err error
		defer close(result)
		response, err = client.ListWorkspaceUsers(request)
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

// ListWorkspaceUsersRequest is the request struct for api ListWorkspaceUsers
type ListWorkspaceUsersRequest struct {
	*requests.RoaRequest
	WorkspaceId string `position:"Path" name:"WorkspaceId"`
}

// ListWorkspaceUsersResponse is the response struct for api ListWorkspaceUsers
type ListWorkspaceUsersResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	TotalCount int64       `json:"TotalCount" xml:"TotalCount"`
	Users      []UsersItem `json:"Users" xml:"Users"`
}

// CreateListWorkspaceUsersRequest creates a request to invoke ListWorkspaceUsers API
func CreateListWorkspaceUsersRequest() (request *ListWorkspaceUsersRequest) {
	request = &ListWorkspaceUsersRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("AIWorkSpace", "2021-02-04", "ListWorkspaceUsers", "/api/v1/workspaces/[WorkspaceId]/users", "", "")
	request.Method = requests.GET
	return
}

// CreateListWorkspaceUsersResponse creates a response to parse from ListWorkspaceUsers response
func CreateListWorkspaceUsersResponse() (response *ListWorkspaceUsersResponse) {
	response = &ListWorkspaceUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
