package dataworks_public

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

// RevokeTablePermission invokes the dataworks_public.RevokeTablePermission API synchronously
func (client *Client) RevokeTablePermission(request *RevokeTablePermissionRequest) (response *RevokeTablePermissionResponse, err error) {
	response = CreateRevokeTablePermissionResponse()
	err = client.DoAction(request, response)
	return
}

// RevokeTablePermissionWithChan invokes the dataworks_public.RevokeTablePermission API asynchronously
func (client *Client) RevokeTablePermissionWithChan(request *RevokeTablePermissionRequest) (<-chan *RevokeTablePermissionResponse, <-chan error) {
	responseChan := make(chan *RevokeTablePermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokeTablePermission(request)
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

// RevokeTablePermissionWithCallback invokes the dataworks_public.RevokeTablePermission API asynchronously
func (client *Client) RevokeTablePermissionWithCallback(request *RevokeTablePermissionRequest, callback func(response *RevokeTablePermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokeTablePermissionResponse
		var err error
		defer close(result)
		response, err = client.RevokeTablePermission(request)
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

// RevokeTablePermissionRequest is the request struct for api RevokeTablePermission
type RevokeTablePermissionRequest struct {
	*requests.RpcRequest
	RevokeUserName        string           `position:"Query" name:"RevokeUserName"`
	MaxComputeProjectName string           `position:"Query" name:"MaxComputeProjectName"`
	RevokeUserId          string           `position:"Query" name:"RevokeUserId"`
	TableName             string           `position:"Query" name:"TableName"`
	Actions               string           `position:"Query" name:"Actions"`
	WorkspaceId           requests.Integer `position:"Query" name:"WorkspaceId"`
}

// RevokeTablePermissionResponse is the response struct for api RevokeTablePermission
type RevokeTablePermissionResponse struct {
	*responses.BaseResponse
	RevokeSuccess bool   `json:"RevokeSuccess" xml:"RevokeSuccess"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateRevokeTablePermissionRequest creates a request to invoke RevokeTablePermission API
func CreateRevokeTablePermissionRequest() (request *RevokeTablePermissionRequest) {
	request = &RevokeTablePermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "RevokeTablePermission", "", "")
	request.Method = requests.POST
	return
}

// CreateRevokeTablePermissionResponse creates a response to parse from RevokeTablePermission response
func CreateRevokeTablePermissionResponse() (response *RevokeTablePermissionResponse) {
	response = &RevokeTablePermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}