package oam

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

// RemoveRoleFromRole invokes the oam.RemoveRoleFromRole API synchronously
// api document: https://help.aliyun.com/api/oam/removerolefromrole.html
func (client *Client) RemoveRoleFromRole(request *RemoveRoleFromRoleRequest) (response *RemoveRoleFromRoleResponse, err error) {
	response = CreateRemoveRoleFromRoleResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveRoleFromRoleWithChan invokes the oam.RemoveRoleFromRole API asynchronously
// api document: https://help.aliyun.com/api/oam/removerolefromrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveRoleFromRoleWithChan(request *RemoveRoleFromRoleRequest) (<-chan *RemoveRoleFromRoleResponse, <-chan error) {
	responseChan := make(chan *RemoveRoleFromRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveRoleFromRole(request)
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

// RemoveRoleFromRoleWithCallback invokes the oam.RemoveRoleFromRole API asynchronously
// api document: https://help.aliyun.com/api/oam/removerolefromrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveRoleFromRoleWithCallback(request *RemoveRoleFromRoleRequest, callback func(response *RemoveRoleFromRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveRoleFromRoleResponse
		var err error
		defer close(result)
		response, err = client.RemoveRoleFromRole(request)
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

// RemoveRoleFromRoleRequest is the request struct for api RemoveRoleFromRole
type RemoveRoleFromRoleRequest struct {
	*requests.RpcRequest
	RoleName     string `position:"Query" name:"RoleName"`
	BaseRoleName string `position:"Query" name:"BaseRoleName"`
}

// RemoveRoleFromRoleResponse is the response struct for api RemoveRoleFromRole
type RemoveRoleFromRoleResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
}

// CreateRemoveRoleFromRoleRequest creates a request to invoke RemoveRoleFromRole API
func CreateRemoveRoleFromRoleRequest() (request *RemoveRoleFromRoleRequest) {
	request = &RemoveRoleFromRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "RemoveRoleFromRole", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveRoleFromRoleResponse creates a response to parse from RemoveRoleFromRole response
func CreateRemoveRoleFromRoleResponse() (response *RemoveRoleFromRoleResponse) {
	response = &RemoveRoleFromRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
