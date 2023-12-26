package scsp

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

// DisableRole invokes the scsp.DisableRole API synchronously
func (client *Client) DisableRole(request *DisableRoleRequest) (response *DisableRoleResponse, err error) {
	response = CreateDisableRoleResponse()
	err = client.DoAction(request, response)
	return
}

// DisableRoleWithChan invokes the scsp.DisableRole API asynchronously
func (client *Client) DisableRoleWithChan(request *DisableRoleRequest) (<-chan *DisableRoleResponse, <-chan error) {
	responseChan := make(chan *DisableRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableRole(request)
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

// DisableRoleWithCallback invokes the scsp.DisableRole API asynchronously
func (client *Client) DisableRoleWithCallback(request *DisableRoleRequest, callback func(response *DisableRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableRoleResponse
		var err error
		defer close(result)
		response, err = client.DisableRole(request)
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

// DisableRoleRequest is the request struct for api DisableRole
type DisableRoleRequest struct {
	*requests.RpcRequest
	ClientToken string           `position:"Body"`
	InstanceId  string           `position:"Body"`
	RoleId      requests.Integer `position:"Body"`
}

// DisableRoleResponse is the response struct for api DisableRole
type DisableRoleResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDisableRoleRequest creates a request to invoke DisableRole API
func CreateDisableRoleRequest() (request *DisableRoleRequest) {
	request = &DisableRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "DisableRole", "", "")
	request.Method = requests.POST
	return
}

// CreateDisableRoleResponse creates a response to parse from DisableRole response
func CreateDisableRoleResponse() (response *DisableRoleResponse) {
	response = &DisableRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
