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

// ListUsersOfRole invokes the cloudcallcenter.ListUsersOfRole API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listusersofrole.html
func (client *Client) ListUsersOfRole(request *ListUsersOfRoleRequest) (response *ListUsersOfRoleResponse, err error) {
	response = CreateListUsersOfRoleResponse()
	err = client.DoAction(request, response)
	return
}

// ListUsersOfRoleWithChan invokes the cloudcallcenter.ListUsersOfRole API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listusersofrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUsersOfRoleWithChan(request *ListUsersOfRoleRequest) (<-chan *ListUsersOfRoleResponse, <-chan error) {
	responseChan := make(chan *ListUsersOfRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUsersOfRole(request)
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

// ListUsersOfRoleWithCallback invokes the cloudcallcenter.ListUsersOfRole API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listusersofrole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUsersOfRoleWithCallback(request *ListUsersOfRoleRequest, callback func(response *ListUsersOfRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUsersOfRoleResponse
		var err error
		defer close(result)
		response, err = client.ListUsersOfRole(request)
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

// ListUsersOfRoleRequest is the request struct for api ListUsersOfRole
type ListUsersOfRoleRequest struct {
	*requests.RpcRequest
	InstanceId string           `position:"Query" name:"InstanceId"`
	RoleId     string           `position:"Query" name:"RoleId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListUsersOfRoleResponse is the response struct for api ListUsersOfRole
type ListUsersOfRoleResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Users          Users  `json:"Users" xml:"Users"`
}

// CreateListUsersOfRoleRequest creates a request to invoke ListUsersOfRole API
func CreateListUsersOfRoleRequest() (request *ListUsersOfRoleRequest) {
	request = &ListUsersOfRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListUsersOfRole", "", "")
	request.Method = requests.POST
	return
}

// CreateListUsersOfRoleResponse creates a response to parse from ListUsersOfRole response
func CreateListUsersOfRoleResponse() (response *ListUsersOfRoleResponse) {
	response = &ListUsersOfRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
