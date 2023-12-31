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

// ModifyDatabaseUserRoles invokes the oceanbasepro.ModifyDatabaseUserRoles API synchronously
func (client *Client) ModifyDatabaseUserRoles(request *ModifyDatabaseUserRolesRequest) (response *ModifyDatabaseUserRolesResponse, err error) {
	response = CreateModifyDatabaseUserRolesResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDatabaseUserRolesWithChan invokes the oceanbasepro.ModifyDatabaseUserRoles API asynchronously
func (client *Client) ModifyDatabaseUserRolesWithChan(request *ModifyDatabaseUserRolesRequest) (<-chan *ModifyDatabaseUserRolesResponse, <-chan error) {
	responseChan := make(chan *ModifyDatabaseUserRolesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDatabaseUserRoles(request)
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

// ModifyDatabaseUserRolesWithCallback invokes the oceanbasepro.ModifyDatabaseUserRoles API asynchronously
func (client *Client) ModifyDatabaseUserRolesWithCallback(request *ModifyDatabaseUserRolesRequest, callback func(response *ModifyDatabaseUserRolesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDatabaseUserRolesResponse
		var err error
		defer close(result)
		response, err = client.ModifyDatabaseUserRoles(request)
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

// ModifyDatabaseUserRolesRequest is the request struct for api ModifyDatabaseUserRoles
type ModifyDatabaseUserRolesRequest struct {
	*requests.RpcRequest
	Users        string `position:"Body" name:"Users"`
	InstanceId   string `position:"Body" name:"InstanceId"`
	DatabaseName string `position:"Body" name:"DatabaseName"`
	TenantId     string `position:"Body" name:"TenantId"`
}

// ModifyDatabaseUserRolesResponse is the response struct for api ModifyDatabaseUserRoles
type ModifyDatabaseUserRolesResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	TenantUser TenantUser `json:"TenantUser" xml:"TenantUser"`
}

// CreateModifyDatabaseUserRolesRequest creates a request to invoke ModifyDatabaseUserRoles API
func CreateModifyDatabaseUserRolesRequest() (request *ModifyDatabaseUserRolesRequest) {
	request = &ModifyDatabaseUserRolesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "ModifyDatabaseUserRoles", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDatabaseUserRolesResponse creates a response to parse from ModifyDatabaseUserRoles response
func CreateModifyDatabaseUserRolesResponse() (response *ModifyDatabaseUserRolesResponse) {
	response = &ModifyDatabaseUserRolesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
