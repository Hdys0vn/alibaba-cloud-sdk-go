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

// ListProxies invokes the dms_enterprise.ListProxies API synchronously
func (client *Client) ListProxies(request *ListProxiesRequest) (response *ListProxiesResponse, err error) {
	response = CreateListProxiesResponse()
	err = client.DoAction(request, response)
	return
}

// ListProxiesWithChan invokes the dms_enterprise.ListProxies API asynchronously
func (client *Client) ListProxiesWithChan(request *ListProxiesRequest) (<-chan *ListProxiesResponse, <-chan error) {
	responseChan := make(chan *ListProxiesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProxies(request)
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

// ListProxiesWithCallback invokes the dms_enterprise.ListProxies API asynchronously
func (client *Client) ListProxiesWithCallback(request *ListProxiesRequest, callback func(response *ListProxiesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProxiesResponse
		var err error
		defer close(result)
		response, err = client.ListProxies(request)
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

// ListProxiesRequest is the request struct for api ListProxies
type ListProxiesRequest struct {
	*requests.RpcRequest
	Tid requests.Integer `position:"Query" name:"Tid"`
}

// ListProxiesResponse is the response struct for api ListProxies
type ListProxiesResponse struct {
	*responses.BaseResponse
	RequestId    string          `json:"RequestId" xml:"RequestId"`
	Success      bool            `json:"Success" xml:"Success"`
	ErrorMessage string          `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string          `json:"ErrorCode" xml:"ErrorCode"`
	ProxyList    []ProxyListItem `json:"ProxyList" xml:"ProxyList"`
}

// CreateListProxiesRequest creates a request to invoke ListProxies API
func CreateListProxiesRequest() (request *ListProxiesRequest) {
	request = &ListProxiesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListProxies", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListProxiesResponse creates a response to parse from ListProxies response
func CreateListProxiesResponse() (response *ListProxiesResponse) {
	response = &ListProxiesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
