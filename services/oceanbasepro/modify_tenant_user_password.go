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

// ModifyTenantUserPassword invokes the oceanbasepro.ModifyTenantUserPassword API synchronously
func (client *Client) ModifyTenantUserPassword(request *ModifyTenantUserPasswordRequest) (response *ModifyTenantUserPasswordResponse, err error) {
	response = CreateModifyTenantUserPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyTenantUserPasswordWithChan invokes the oceanbasepro.ModifyTenantUserPassword API asynchronously
func (client *Client) ModifyTenantUserPasswordWithChan(request *ModifyTenantUserPasswordRequest) (<-chan *ModifyTenantUserPasswordResponse, <-chan error) {
	responseChan := make(chan *ModifyTenantUserPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyTenantUserPassword(request)
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

// ModifyTenantUserPasswordWithCallback invokes the oceanbasepro.ModifyTenantUserPassword API asynchronously
func (client *Client) ModifyTenantUserPasswordWithCallback(request *ModifyTenantUserPasswordRequest, callback func(response *ModifyTenantUserPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyTenantUserPasswordResponse
		var err error
		defer close(result)
		response, err = client.ModifyTenantUserPassword(request)
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

// ModifyTenantUserPasswordRequest is the request struct for api ModifyTenantUserPassword
type ModifyTenantUserPasswordRequest struct {
	*requests.RpcRequest
	UserPassword   string `position:"Body" name:"UserPassword"`
	InstanceId     string `position:"Body" name:"InstanceId"`
	TenantId       string `position:"Body" name:"TenantId"`
	EncryptionType string `position:"Body" name:"EncryptionType"`
	UserName       string `position:"Body" name:"UserName"`
}

// ModifyTenantUserPasswordResponse is the response struct for api ModifyTenantUserPassword
type ModifyTenantUserPasswordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyTenantUserPasswordRequest creates a request to invoke ModifyTenantUserPassword API
func CreateModifyTenantUserPasswordRequest() (request *ModifyTenantUserPasswordRequest) {
	request = &ModifyTenantUserPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "ModifyTenantUserPassword", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyTenantUserPasswordResponse creates a response to parse from ModifyTenantUserPassword response
func CreateModifyTenantUserPasswordResponse() (response *ModifyTenantUserPasswordResponse) {
	response = &ModifyTenantUserPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}