package cloudcallcenter

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

// RemovePrivilegesFromRole invokes the cloudcallcenter.RemovePrivilegesFromRole API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/removeprivilegesfromrole.html
func (client *Client) RemovePrivilegesFromRole(request *RemovePrivilegesFromRoleRequest) (response *RemovePrivilegesFromRoleResponse, err error) {
	response = CreateRemovePrivilegesFromRoleResponse()
	err = client.DoAction(request, response)
	return
}

// RemovePrivilegesFromRoleWithChan invokes the cloudcallcenter.RemovePrivilegesFromRole API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/removeprivilegesfromrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemovePrivilegesFromRoleWithChan(request *RemovePrivilegesFromRoleRequest) (<-chan *RemovePrivilegesFromRoleResponse, <-chan error) {
	responseChan := make(chan *RemovePrivilegesFromRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemovePrivilegesFromRole(request)
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

// RemovePrivilegesFromRoleWithCallback invokes the cloudcallcenter.RemovePrivilegesFromRole API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/removeprivilegesfromrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemovePrivilegesFromRoleWithCallback(request *RemovePrivilegesFromRoleRequest, callback func(response *RemovePrivilegesFromRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemovePrivilegesFromRoleResponse
		var err error
		defer close(result)
		response, err = client.RemovePrivilegesFromRole(request)
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

// RemovePrivilegesFromRoleRequest is the request struct for api RemovePrivilegesFromRole
type RemovePrivilegesFromRoleRequest struct {
	*requests.RpcRequest
	InstanceId  string    `position:"Query" name:"InstanceId"`
	PrivilegeId *[]string `position:"Query" name:"PrivilegeId"  type:"Repeated"`
	RoleId      string    `position:"Query" name:"RoleId"`
}

// RemovePrivilegesFromRoleResponse is the response struct for api RemovePrivilegesFromRole
type RemovePrivilegesFromRoleResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Role           Role   `json:"Role" xml:"Role"`
}

// CreateRemovePrivilegesFromRoleRequest creates a request to invoke RemovePrivilegesFromRole API
func CreateRemovePrivilegesFromRoleRequest() (request *RemovePrivilegesFromRoleRequest) {
	request = &RemovePrivilegesFromRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "RemovePrivilegesFromRole", "", "")
	request.Method = requests.POST
	return
}

// CreateRemovePrivilegesFromRoleResponse creates a response to parse from RemovePrivilegesFromRole response
func CreateRemovePrivilegesFromRoleResponse() (response *RemovePrivilegesFromRoleResponse) {
	response = &RemovePrivilegesFromRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
