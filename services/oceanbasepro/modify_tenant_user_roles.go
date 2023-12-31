package oceanbasepro

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

// ModifyTenantUserRoles invokes the oceanbasepro.ModifyTenantUserRoles API synchronously
func (client *Client) ModifyTenantUserRoles(request *ModifyTenantUserRolesRequest) (response *ModifyTenantUserRolesResponse, err error) {
	response = CreateModifyTenantUserRolesResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyTenantUserRolesWithChan invokes the oceanbasepro.ModifyTenantUserRoles API asynchronously
func (client *Client) ModifyTenantUserRolesWithChan(request *ModifyTenantUserRolesRequest) (<-chan *ModifyTenantUserRolesResponse, <-chan error) {
	responseChan := make(chan *ModifyTenantUserRolesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyTenantUserRoles(request)
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

// ModifyTenantUserRolesWithCallback invokes the oceanbasepro.ModifyTenantUserRoles API asynchronously
func (client *Client) ModifyTenantUserRolesWithCallback(request *ModifyTenantUserRolesRequest, callback func(response *ModifyTenantUserRolesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyTenantUserRolesResponse
		var err error
		defer close(result)
		response, err = client.ModifyTenantUserRoles(request)
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

// ModifyTenantUserRolesRequest is the request struct for api ModifyTenantUserRoles
type ModifyTenantUserRolesRequest struct {
	*requests.RpcRequest
	UserRole   string `position:"Body" name:"UserRole"`
	InstanceId string `position:"Body" name:"InstanceId"`
	ModifyType string `position:"Body" name:"ModifyType"`
	TenantId   string `position:"Body" name:"TenantId"`
	UserName   string `position:"Body" name:"UserName"`
}

// ModifyTenantUserRolesResponse is the response struct for api ModifyTenantUserRoles
type ModifyTenantUserRolesResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	TenantUser TenantUser `json:"TenantUser" xml:"TenantUser"`
}

// CreateModifyTenantUserRolesRequest creates a request to invoke ModifyTenantUserRoles API
func CreateModifyTenantUserRolesRequest() (request *ModifyTenantUserRolesRequest) {
	request = &ModifyTenantUserRolesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "ModifyTenantUserRoles", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyTenantUserRolesResponse creates a response to parse from ModifyTenantUserRoles response
func CreateModifyTenantUserRolesResponse() (response *ModifyTenantUserRolesResponse) {
	response = &ModifyTenantUserRolesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
