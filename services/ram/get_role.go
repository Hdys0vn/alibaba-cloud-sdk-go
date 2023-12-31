package ram

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

// GetRole invokes the ram.GetRole API synchronously
func (client *Client) GetRole(request *GetRoleRequest) (response *GetRoleResponse, err error) {
	response = CreateGetRoleResponse()
	err = client.DoAction(request, response)
	return
}

// GetRoleWithChan invokes the ram.GetRole API asynchronously
func (client *Client) GetRoleWithChan(request *GetRoleRequest) (<-chan *GetRoleResponse, <-chan error) {
	responseChan := make(chan *GetRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRole(request)
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

// GetRoleWithCallback invokes the ram.GetRole API asynchronously
func (client *Client) GetRoleWithCallback(request *GetRoleRequest, callback func(response *GetRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRoleResponse
		var err error
		defer close(result)
		response, err = client.GetRole(request)
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

// GetRoleRequest is the request struct for api GetRole
type GetRoleRequest struct {
	*requests.RpcRequest
	RoleName string `position:"Query" name:"RoleName"`
}

// GetRoleResponse is the response struct for api GetRole
type GetRoleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Role      Role   `json:"Role" xml:"Role"`
}

// CreateGetRoleRequest creates a request to invoke GetRole API
func CreateGetRoleRequest() (request *GetRoleRequest) {
	request = &GetRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "GetRole", "Ram", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetRoleResponse creates a response to parse from GetRole response
func CreateGetRoleResponse() (response *GetRoleResponse) {
	response = &GetRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
