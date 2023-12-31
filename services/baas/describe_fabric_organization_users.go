package baas

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

// DescribeFabricOrganizationUsers invokes the baas.DescribeFabricOrganizationUsers API synchronously
// api document: https://help.aliyun.com/api/baas/describefabricorganizationusers.html
func (client *Client) DescribeFabricOrganizationUsers(request *DescribeFabricOrganizationUsersRequest) (response *DescribeFabricOrganizationUsersResponse, err error) {
	response = CreateDescribeFabricOrganizationUsersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFabricOrganizationUsersWithChan invokes the baas.DescribeFabricOrganizationUsers API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricorganizationusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricOrganizationUsersWithChan(request *DescribeFabricOrganizationUsersRequest) (<-chan *DescribeFabricOrganizationUsersResponse, <-chan error) {
	responseChan := make(chan *DescribeFabricOrganizationUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFabricOrganizationUsers(request)
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

// DescribeFabricOrganizationUsersWithCallback invokes the baas.DescribeFabricOrganizationUsers API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricorganizationusers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricOrganizationUsersWithCallback(request *DescribeFabricOrganizationUsersRequest, callback func(response *DescribeFabricOrganizationUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFabricOrganizationUsersResponse
		var err error
		defer close(result)
		response, err = client.DescribeFabricOrganizationUsers(request)
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

// DescribeFabricOrganizationUsersRequest is the request struct for api DescribeFabricOrganizationUsers
type DescribeFabricOrganizationUsersRequest struct {
	*requests.RpcRequest
	OrganizationId string `position:"Query" name:"OrganizationId"`
	Location       string `position:"Body" name:"Location"`
}

// DescribeFabricOrganizationUsersResponse is the response struct for api DescribeFabricOrganizationUsers
type DescribeFabricOrganizationUsersResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Success   bool         `json:"Success" xml:"Success"`
	ErrorCode int          `json:"ErrorCode" xml:"ErrorCode"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateDescribeFabricOrganizationUsersRequest creates a request to invoke DescribeFabricOrganizationUsers API
func CreateDescribeFabricOrganizationUsersRequest() (request *DescribeFabricOrganizationUsersRequest) {
	request = &DescribeFabricOrganizationUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeFabricOrganizationUsers", "baas", "openAPI")
	return
}

// CreateDescribeFabricOrganizationUsersResponse creates a response to parse from DescribeFabricOrganizationUsers response
func CreateDescribeFabricOrganizationUsersResponse() (response *DescribeFabricOrganizationUsersResponse) {
	response = &DescribeFabricOrganizationUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
