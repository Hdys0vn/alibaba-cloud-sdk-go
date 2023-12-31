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

// RemoveRoleCellFromRole invokes the oam.RemoveRoleCellFromRole API synchronously
// api document: https://help.aliyun.com/api/oam/removerolecellfromrole.html
func (client *Client) RemoveRoleCellFromRole(request *RemoveRoleCellFromRoleRequest) (response *RemoveRoleCellFromRoleResponse, err error) {
	response = CreateRemoveRoleCellFromRoleResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveRoleCellFromRoleWithChan invokes the oam.RemoveRoleCellFromRole API asynchronously
// api document: https://help.aliyun.com/api/oam/removerolecellfromrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveRoleCellFromRoleWithChan(request *RemoveRoleCellFromRoleRequest) (<-chan *RemoveRoleCellFromRoleResponse, <-chan error) {
	responseChan := make(chan *RemoveRoleCellFromRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveRoleCellFromRole(request)
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

// RemoveRoleCellFromRoleWithCallback invokes the oam.RemoveRoleCellFromRole API asynchronously
// api document: https://help.aliyun.com/api/oam/removerolecellfromrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveRoleCellFromRoleWithCallback(request *RemoveRoleCellFromRoleRequest, callback func(response *RemoveRoleCellFromRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveRoleCellFromRoleResponse
		var err error
		defer close(result)
		response, err = client.RemoveRoleCellFromRole(request)
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

// RemoveRoleCellFromRoleRequest is the request struct for api RemoveRoleCellFromRole
type RemoveRoleCellFromRoleRequest struct {
	*requests.RpcRequest
	RoleCellId string `position:"Query" name:"RoleCellId"`
}

// RemoveRoleCellFromRoleResponse is the response struct for api RemoveRoleCellFromRole
type RemoveRoleCellFromRoleResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
}

// CreateRemoveRoleCellFromRoleRequest creates a request to invoke RemoveRoleCellFromRole API
func CreateRemoveRoleCellFromRoleRequest() (request *RemoveRoleCellFromRoleRequest) {
	request = &RemoveRoleCellFromRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "RemoveRoleCellFromRole", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveRoleCellFromRoleResponse creates a response to parse from RemoveRoleCellFromRole response
func CreateRemoveRoleCellFromRoleResponse() (response *RemoveRoleCellFromRoleResponse) {
	response = &RemoveRoleCellFromRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
