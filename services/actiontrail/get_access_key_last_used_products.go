package actiontrail

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

// GetAccessKeyLastUsedProducts invokes the actiontrail.GetAccessKeyLastUsedProducts API synchronously
func (client *Client) GetAccessKeyLastUsedProducts(request *GetAccessKeyLastUsedProductsRequest) (response *GetAccessKeyLastUsedProductsResponse, err error) {
	response = CreateGetAccessKeyLastUsedProductsResponse()
	err = client.DoAction(request, response)
	return
}

// GetAccessKeyLastUsedProductsWithChan invokes the actiontrail.GetAccessKeyLastUsedProducts API asynchronously
func (client *Client) GetAccessKeyLastUsedProductsWithChan(request *GetAccessKeyLastUsedProductsRequest) (<-chan *GetAccessKeyLastUsedProductsResponse, <-chan error) {
	responseChan := make(chan *GetAccessKeyLastUsedProductsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccessKeyLastUsedProducts(request)
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

// GetAccessKeyLastUsedProductsWithCallback invokes the actiontrail.GetAccessKeyLastUsedProducts API asynchronously
func (client *Client) GetAccessKeyLastUsedProductsWithCallback(request *GetAccessKeyLastUsedProductsRequest, callback func(response *GetAccessKeyLastUsedProductsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAccessKeyLastUsedProductsResponse
		var err error
		defer close(result)
		response, err = client.GetAccessKeyLastUsedProducts(request)
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

// GetAccessKeyLastUsedProductsRequest is the request struct for api GetAccessKeyLastUsedProducts
type GetAccessKeyLastUsedProductsRequest struct {
	*requests.RpcRequest
	AccessKey string `position:"Query" name:"AccessKey"`
}

// GetAccessKeyLastUsedProductsResponse is the response struct for api GetAccessKeyLastUsedProducts
type GetAccessKeyLastUsedProductsResponse struct {
	*responses.BaseResponse
	RequestId string         `json:"RequestId" xml:"RequestId"`
	Products  []ProductsItem `json:"Products" xml:"Products"`
}

// CreateGetAccessKeyLastUsedProductsRequest creates a request to invoke GetAccessKeyLastUsedProducts API
func CreateGetAccessKeyLastUsedProductsRequest() (request *GetAccessKeyLastUsedProductsRequest) {
	request = &GetAccessKeyLastUsedProductsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Actiontrail", "2020-07-06", "GetAccessKeyLastUsedProducts", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAccessKeyLastUsedProductsResponse creates a response to parse from GetAccessKeyLastUsedProducts response
func CreateGetAccessKeyLastUsedProductsResponse() (response *GetAccessKeyLastUsedProductsResponse) {
	response = &GetAccessKeyLastUsedProductsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
