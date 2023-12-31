package dms_enterprise

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

// ListUserTenants invokes the dms_enterprise.ListUserTenants API synchronously
func (client *Client) ListUserTenants(request *ListUserTenantsRequest) (response *ListUserTenantsResponse, err error) {
	response = CreateListUserTenantsResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserTenantsWithChan invokes the dms_enterprise.ListUserTenants API asynchronously
func (client *Client) ListUserTenantsWithChan(request *ListUserTenantsRequest) (<-chan *ListUserTenantsResponse, <-chan error) {
	responseChan := make(chan *ListUserTenantsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserTenants(request)
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

// ListUserTenantsWithCallback invokes the dms_enterprise.ListUserTenants API asynchronously
func (client *Client) ListUserTenantsWithCallback(request *ListUserTenantsRequest, callback func(response *ListUserTenantsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserTenantsResponse
		var err error
		defer close(result)
		response, err = client.ListUserTenants(request)
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

// ListUserTenantsRequest is the request struct for api ListUserTenants
type ListUserTenantsRequest struct {
	*requests.RpcRequest
	Tid requests.Integer `position:"Query" name:"Tid"`
}

// ListUserTenantsResponse is the response struct for api ListUserTenants
type ListUserTenantsResponse struct {
	*responses.BaseResponse
	RequestId    string   `json:"RequestId" xml:"RequestId"`
	ErrorCode    string   `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string   `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool     `json:"Success" xml:"Success"`
	TenantList   []Tenant `json:"TenantList" xml:"TenantList"`
}

// CreateListUserTenantsRequest creates a request to invoke ListUserTenants API
func CreateListUserTenantsRequest() (request *ListUserTenantsRequest) {
	request = &ListUserTenantsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListUserTenants", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListUserTenantsResponse creates a response to parse from ListUserTenants response
func CreateListUserTenantsResponse() (response *ListUserTenantsResponse) {
	response = &ListUserTenantsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
