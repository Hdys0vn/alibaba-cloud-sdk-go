package sas

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

// DescribeServiceLinkedRoleStatus invokes the sas.DescribeServiceLinkedRoleStatus API synchronously
func (client *Client) DescribeServiceLinkedRoleStatus(request *DescribeServiceLinkedRoleStatusRequest) (response *DescribeServiceLinkedRoleStatusResponse, err error) {
	response = CreateDescribeServiceLinkedRoleStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeServiceLinkedRoleStatusWithChan invokes the sas.DescribeServiceLinkedRoleStatus API asynchronously
func (client *Client) DescribeServiceLinkedRoleStatusWithChan(request *DescribeServiceLinkedRoleStatusRequest) (<-chan *DescribeServiceLinkedRoleStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeServiceLinkedRoleStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeServiceLinkedRoleStatus(request)
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

// DescribeServiceLinkedRoleStatusWithCallback invokes the sas.DescribeServiceLinkedRoleStatus API asynchronously
func (client *Client) DescribeServiceLinkedRoleStatusWithCallback(request *DescribeServiceLinkedRoleStatusRequest, callback func(response *DescribeServiceLinkedRoleStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeServiceLinkedRoleStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeServiceLinkedRoleStatus(request)
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

// DescribeServiceLinkedRoleStatusRequest is the request struct for api DescribeServiceLinkedRoleStatus
type DescribeServiceLinkedRoleStatusRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeServiceLinkedRoleStatusResponse is the response struct for api DescribeServiceLinkedRoleStatus
type DescribeServiceLinkedRoleStatusResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	RoleStatus RoleStatus `json:"RoleStatus" xml:"RoleStatus"`
}

// CreateDescribeServiceLinkedRoleStatusRequest creates a request to invoke DescribeServiceLinkedRoleStatus API
func CreateDescribeServiceLinkedRoleStatusRequest() (request *DescribeServiceLinkedRoleStatusRequest) {
	request = &DescribeServiceLinkedRoleStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeServiceLinkedRoleStatus", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeServiceLinkedRoleStatusResponse creates a response to parse from DescribeServiceLinkedRoleStatus response
func CreateDescribeServiceLinkedRoleStatusResponse() (response *DescribeServiceLinkedRoleStatusResponse) {
	response = &DescribeServiceLinkedRoleStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
