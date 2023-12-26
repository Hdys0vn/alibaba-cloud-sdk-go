package quickbi_public

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

// UpdateWorkspaceUserRole invokes the quickbi_public.UpdateWorkspaceUserRole API synchronously
func (client *Client) UpdateWorkspaceUserRole(request *UpdateWorkspaceUserRoleRequest) (response *UpdateWorkspaceUserRoleResponse, err error) {
	response = CreateUpdateWorkspaceUserRoleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateWorkspaceUserRoleWithChan invokes the quickbi_public.UpdateWorkspaceUserRole API asynchronously
func (client *Client) UpdateWorkspaceUserRoleWithChan(request *UpdateWorkspaceUserRoleRequest) (<-chan *UpdateWorkspaceUserRoleResponse, <-chan error) {
	responseChan := make(chan *UpdateWorkspaceUserRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateWorkspaceUserRole(request)
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

// UpdateWorkspaceUserRoleWithCallback invokes the quickbi_public.UpdateWorkspaceUserRole API asynchronously
func (client *Client) UpdateWorkspaceUserRoleWithCallback(request *UpdateWorkspaceUserRoleRequest, callback func(response *UpdateWorkspaceUserRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateWorkspaceUserRoleResponse
		var err error
		defer close(result)
		response, err = client.UpdateWorkspaceUserRole(request)
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

// UpdateWorkspaceUserRoleRequest is the request struct for api UpdateWorkspaceUserRole
type UpdateWorkspaceUserRoleRequest struct {
	*requests.RpcRequest
	RoleId      requests.Integer `position:"Query" name:"RoleId"`
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	SignType    string           `position:"Query" name:"SignType"`
	UserId      string           `position:"Query" name:"UserId"`
	WorkspaceId string           `position:"Query" name:"WorkspaceId"`
}

// UpdateWorkspaceUserRoleResponse is the response struct for api UpdateWorkspaceUserRole
type UpdateWorkspaceUserRoleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdateWorkspaceUserRoleRequest creates a request to invoke UpdateWorkspaceUserRole API
func CreateUpdateWorkspaceUserRoleRequest() (request *UpdateWorkspaceUserRoleRequest) {
	request = &UpdateWorkspaceUserRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "UpdateWorkspaceUserRole", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateWorkspaceUserRoleResponse creates a response to parse from UpdateWorkspaceUserRole response
func CreateUpdateWorkspaceUserRoleResponse() (response *UpdateWorkspaceUserRoleResponse) {
	response = &UpdateWorkspaceUserRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
