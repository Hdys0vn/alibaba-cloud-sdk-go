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

// AssignPrivilegesToRole invokes the cloudcallcenter.AssignPrivilegesToRole API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/assignprivilegestorole.html
func (client *Client) AssignPrivilegesToRole(request *AssignPrivilegesToRoleRequest) (response *AssignPrivilegesToRoleResponse, err error) {
	response = CreateAssignPrivilegesToRoleResponse()
	err = client.DoAction(request, response)
	return
}

// AssignPrivilegesToRoleWithChan invokes the cloudcallcenter.AssignPrivilegesToRole API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/assignprivilegestorole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssignPrivilegesToRoleWithChan(request *AssignPrivilegesToRoleRequest) (<-chan *AssignPrivilegesToRoleResponse, <-chan error) {
	responseChan := make(chan *AssignPrivilegesToRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssignPrivilegesToRole(request)
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

// AssignPrivilegesToRoleWithCallback invokes the cloudcallcenter.AssignPrivilegesToRole API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/assignprivilegestorole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssignPrivilegesToRoleWithCallback(request *AssignPrivilegesToRoleRequest, callback func(response *AssignPrivilegesToRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssignPrivilegesToRoleResponse
		var err error
		defer close(result)
		response, err = client.AssignPrivilegesToRole(request)
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

// AssignPrivilegesToRoleRequest is the request struct for api AssignPrivilegesToRole
type AssignPrivilegesToRoleRequest struct {
	*requests.RpcRequest
	InstanceId  string    `position:"Query" name:"InstanceId"`
	PrivilegeId *[]string `position:"Query" name:"PrivilegeId"  type:"Repeated"`
	RoleId      string    `position:"Query" name:"RoleId"`
}

// AssignPrivilegesToRoleResponse is the response struct for api AssignPrivilegesToRole
type AssignPrivilegesToRoleResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Role           Role   `json:"Role" xml:"Role"`
}

// CreateAssignPrivilegesToRoleRequest creates a request to invoke AssignPrivilegesToRole API
func CreateAssignPrivilegesToRoleRequest() (request *AssignPrivilegesToRoleRequest) {
	request = &AssignPrivilegesToRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "AssignPrivilegesToRole", "", "")
	request.Method = requests.POST
	return
}

// CreateAssignPrivilegesToRoleResponse creates a response to parse from AssignPrivilegesToRole response
func CreateAssignPrivilegesToRoleResponse() (response *AssignPrivilegesToRoleResponse) {
	response = &AssignPrivilegesToRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
