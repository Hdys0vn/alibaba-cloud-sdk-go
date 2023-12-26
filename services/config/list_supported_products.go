package config

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

// ListSupportedProducts invokes the config.ListSupportedProducts API synchronously
func (client *Client) ListSupportedProducts(request *ListSupportedProductsRequest) (response *ListSupportedProductsResponse, err error) {
	response = CreateListSupportedProductsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSupportedProductsWithChan invokes the config.ListSupportedProducts API asynchronously
func (client *Client) ListSupportedProductsWithChan(request *ListSupportedProductsRequest) (<-chan *ListSupportedProductsResponse, <-chan error) {
	responseChan := make(chan *ListSupportedProductsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSupportedProducts(request)
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

// ListSupportedProductsWithCallback invokes the config.ListSupportedProducts API asynchronously
func (client *Client) ListSupportedProductsWithCallback(request *ListSupportedProductsRequest, callback func(response *ListSupportedProductsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSupportedProductsResponse
		var err error
		defer close(result)
		response, err = client.ListSupportedProducts(request)
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

// ListSupportedProductsRequest is the request struct for api ListSupportedProducts
type ListSupportedProductsRequest struct {
	*requests.RpcRequest
	NextToken  string           `position:"Query" name:"NextToken"`
	MaxResults requests.Integer `position:"Query" name:"MaxResults"`
}

// ListSupportedProductsResponse is the response struct for api ListSupportedProducts
type ListSupportedProductsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	NextToken  string `json:"NextToken" xml:"NextToken"`
	MaxResults string `json:"MaxResults" xml:"MaxResults"`
	Products   []Data `json:"Products" xml:"Products"`
}

// CreateListSupportedProductsRequest creates a request to invoke ListSupportedProducts API
func CreateListSupportedProductsRequest() (request *ListSupportedProductsRequest) {
	request = &ListSupportedProductsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "ListSupportedProducts", "", "")
	request.Method = requests.POST
	return
}

// CreateListSupportedProductsResponse creates a response to parse from ListSupportedProducts response
func CreateListSupportedProductsResponse() (response *ListSupportedProductsResponse) {
	response = &ListSupportedProductsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}