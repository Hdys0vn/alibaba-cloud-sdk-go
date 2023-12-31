package oam

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

// GrantRoleToOperator invokes the oam.GrantRoleToOperator API synchronously
// api document: https://help.aliyun.com/api/oam/grantroletooperator.html
func (client *Client) GrantRoleToOperator(request *GrantRoleToOperatorRequest) (response *GrantRoleToOperatorResponse, err error) {
	response = CreateGrantRoleToOperatorResponse()
	err = client.DoAction(request, response)
	return
}

// GrantRoleToOperatorWithChan invokes the oam.GrantRoleToOperator API asynchronously
// api document: https://help.aliyun.com/api/oam/grantroletooperator.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GrantRoleToOperatorWithChan(request *GrantRoleToOperatorRequest) (<-chan *GrantRoleToOperatorResponse, <-chan error) {
	responseChan := make(chan *GrantRoleToOperatorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GrantRoleToOperator(request)
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

// GrantRoleToOperatorWithCallback invokes the oam.GrantRoleToOperator API asynchronously
// api document: https://help.aliyun.com/api/oam/grantroletooperator.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GrantRoleToOperatorWithCallback(request *GrantRoleToOperatorRequest, callback func(response *GrantRoleToOperatorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GrantRoleToOperatorResponse
		var err error
		defer close(result)
		response, err = client.GrantRoleToOperator(request)
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

// GrantRoleToOperatorRequest is the request struct for api GrantRoleToOperator
type GrantRoleToOperatorRequest struct {
	*requests.RpcRequest
	ToOperatorName string `position:"Query" name:"ToOperatorName"`
	RoleName       string `position:"Query" name:"RoleName"`
	UserType       string `position:"Query" name:"UserType"`
	GmtExpired     string `position:"Query" name:"GmtExpired"`
}

// GrantRoleToOperatorResponse is the response struct for api GrantRoleToOperator
type GrantRoleToOperatorResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
}

// CreateGrantRoleToOperatorRequest creates a request to invoke GrantRoleToOperator API
func CreateGrantRoleToOperatorRequest() (request *GrantRoleToOperatorRequest) {
	request = &GrantRoleToOperatorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "GrantRoleToOperator", "", "")
	request.Method = requests.POST
	return
}

// CreateGrantRoleToOperatorResponse creates a response to parse from GrantRoleToOperator response
func CreateGrantRoleToOperatorResponse() (response *GrantRoleToOperatorResponse) {
	response = &GrantRoleToOperatorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
