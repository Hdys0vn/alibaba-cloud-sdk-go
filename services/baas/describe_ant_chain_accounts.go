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

// DescribeAntChainAccounts invokes the baas.DescribeAntChainAccounts API synchronously
// api document: https://help.aliyun.com/api/baas/describeantchainaccounts.html
func (client *Client) DescribeAntChainAccounts(request *DescribeAntChainAccountsRequest) (response *DescribeAntChainAccountsResponse, err error) {
	response = CreateDescribeAntChainAccountsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAntChainAccountsWithChan invokes the baas.DescribeAntChainAccounts API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchainaccounts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainAccountsWithChan(request *DescribeAntChainAccountsRequest) (<-chan *DescribeAntChainAccountsResponse, <-chan error) {
	responseChan := make(chan *DescribeAntChainAccountsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAntChainAccounts(request)
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

// DescribeAntChainAccountsWithCallback invokes the baas.DescribeAntChainAccounts API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchainaccounts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainAccountsWithCallback(request *DescribeAntChainAccountsRequest, callback func(response *DescribeAntChainAccountsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAntChainAccountsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAntChainAccounts(request)
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

// DescribeAntChainAccountsRequest is the request struct for api DescribeAntChainAccounts
type DescribeAntChainAccountsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
	AntChainId string           `position:"Body" name:"AntChainId"`
}

// DescribeAntChainAccountsResponse is the response struct for api DescribeAntChainAccounts
type DescribeAntChainAccountsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeAntChainAccountsRequest creates a request to invoke DescribeAntChainAccounts API
func CreateDescribeAntChainAccountsRequest() (request *DescribeAntChainAccountsRequest) {
	request = &DescribeAntChainAccountsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeAntChainAccounts", "baas", "openAPI")
	return
}

// CreateDescribeAntChainAccountsResponse creates a response to parse from DescribeAntChainAccounts response
func CreateDescribeAntChainAccountsResponse() (response *DescribeAntChainAccountsResponse) {
	response = &DescribeAntChainAccountsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
