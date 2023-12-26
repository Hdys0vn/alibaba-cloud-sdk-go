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

// GetPermission invokes the aiworkspace.GetPermission API synchronously
func (client *Client) GetPermission(request *GetPermissionRequest) (response *GetPermissionResponse, err error) {
	response = CreateGetPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// GetPermissionWithChan invokes the aiworkspace.GetPermission API asynchronously
func (client *Client) GetPermissionWithChan(request *GetPermissionRequest) (<-chan *GetPermissionResponse, <-chan error) {
	responseChan := make(chan *GetPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPermission(request)
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

// GetPermissionWithCallback invokes the aiworkspace.GetPermission API asynchronously
func (client *Client) GetPermissionWithCallback(request *GetPermissionRequest, callback func(response *GetPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPermissionResponse
		var err error
		defer close(result)
		response, err = client.GetPermission(request)
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

// GetPermissionRequest is the request struct for api GetPermission
type GetPermissionRequest struct {
	*requests.RoaRequest
	PermissionCode string `position:"Path" name:"PermissionCode"`
	Creator        string `position:"Query" name:"Creator"`
	Accessibility  string `position:"Query" name:"Accessibility"`
	WorkspaceId    string `position:"Path" name:"WorkspaceId"`
}

// GetPermissionResponse is the response struct for api GetPermission
type GetPermissionResponse struct {
	*responses.BaseResponse
	RequestId       string                `json:"RequestId" xml:"RequestId"`
	PermissionCode  string                `json:"PermissionCode" xml:"PermissionCode"`
	PermissionRules []PermissionRulesItem `json:"PermissionRules" xml:"PermissionRules"`
}

// CreateGetPermissionRequest creates a request to invoke GetPermission API
func CreateGetPermissionRequest() (request *GetPermissionRequest) {
	request = &GetPermissionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("AIWorkSpace", "2021-02-04", "GetPermission", "/api/v1/workspaces/[WorkspaceId]/permissions/[PermissionCode]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetPermissionResponse creates a response to parse from GetPermission response
func CreateGetPermissionResponse() (response *GetPermissionResponse) {
	response = &GetPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
