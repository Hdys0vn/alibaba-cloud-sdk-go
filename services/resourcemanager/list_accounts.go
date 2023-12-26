package resourcemanager

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

// ListAccounts invokes the resourcemanager.ListAccounts API synchronously
func (client *Client) ListAccounts(request *ListAccountsRequest) (response *ListAccountsResponse, err error) {
	response = CreateListAccountsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAccountsWithChan invokes the resourcemanager.ListAccounts API asynchronously
func (client *Client) ListAccountsWithChan(request *ListAccountsRequest) (<-chan *ListAccountsResponse, <-chan error) {
	responseChan := make(chan *ListAccountsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAccounts(request)
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

// ListAccountsWithCallback invokes the resourcemanager.ListAccounts API asynchronously
func (client *Client) ListAccountsWithCallback(request *ListAccountsRequest, callback func(response *ListAccountsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAccountsResponse
		var err error
		defer close(result)
		response, err = client.ListAccounts(request)
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

// ListAccountsRequest is the request struct for api ListAccounts
type ListAccountsRequest struct {
	*requests.RpcRequest
	PageNumber  requests.Integer   `position:"Query" name:"PageNumber"`
	IncludeTags requests.Boolean   `position:"Query" name:"IncludeTags"`
	PageSize    requests.Integer   `position:"Query" name:"PageSize"`
	Tag         *[]ListAccountsTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// ListAccountsTag is a repeated param struct in ListAccountsRequest
type ListAccountsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListAccountsResponse is the response struct for api ListAccounts
type ListAccountsResponse struct {
	*responses.BaseResponse
	TotalCount int                    `json:"TotalCount" xml:"TotalCount"`
	RequestId  string                 `json:"RequestId" xml:"RequestId"`
	PageSize   int                    `json:"PageSize" xml:"PageSize"`
	PageNumber int                    `json:"PageNumber" xml:"PageNumber"`
	Accounts   AccountsInListAccounts `json:"Accounts" xml:"Accounts"`
}

// CreateListAccountsRequest creates a request to invoke ListAccounts API
func CreateListAccountsRequest() (request *ListAccountsRequest) {
	request = &ListAccountsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "ListAccounts", "", "")
	request.Method = requests.POST
	return
}

// CreateListAccountsResponse creates a response to parse from ListAccounts response
func CreateListAccountsResponse() (response *ListAccountsResponse) {
	response = &ListAccountsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
