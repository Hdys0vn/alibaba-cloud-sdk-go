package cloudwf

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

// ShopCreatemarketing invokes the cloudwf.ShopCreatemarketing API synchronously
// api document: https://help.aliyun.com/api/cloudwf/shopcreatemarketing.html
func (client *Client) ShopCreatemarketing(request *ShopCreatemarketingRequest) (response *ShopCreatemarketingResponse, err error) {
	response = CreateShopCreatemarketingResponse()
	err = client.DoAction(request, response)
	return
}

// ShopCreatemarketingWithChan invokes the cloudwf.ShopCreatemarketing API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/shopcreatemarketing.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ShopCreatemarketingWithChan(request *ShopCreatemarketingRequest) (<-chan *ShopCreatemarketingResponse, <-chan error) {
	responseChan := make(chan *ShopCreatemarketingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ShopCreatemarketing(request)
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

// ShopCreatemarketingWithCallback invokes the cloudwf.ShopCreatemarketing API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/shopcreatemarketing.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ShopCreatemarketingWithCallback(request *ShopCreatemarketingRequest, callback func(response *ShopCreatemarketingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ShopCreatemarketingResponse
		var err error
		defer close(result)
		response, err = client.ShopCreatemarketing(request)
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

// ShopCreatemarketingRequest is the request struct for api ShopCreatemarketing
type ShopCreatemarketingRequest struct {
	*requests.RpcRequest
	Etime string           `position:"Query" name:"Etime"`
	Name  string           `position:"Query" name:"Name"`
	Stime string           `position:"Query" name:"Stime"`
	Sid   requests.Integer `position:"Query" name:"Sid"`
}

// ShopCreatemarketingResponse is the response struct for api ShopCreatemarketing
type ShopCreatemarketingResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateShopCreatemarketingRequest creates a request to invoke ShopCreatemarketing API
func CreateShopCreatemarketingRequest() (request *ShopCreatemarketingRequest) {
	request = &ShopCreatemarketingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ShopCreatemarketing", "cloudwf", "openAPI")
	return
}

// CreateShopCreatemarketingResponse creates a response to parse from ShopCreatemarketing response
func CreateShopCreatemarketingResponse() (response *ShopCreatemarketingResponse) {
	response = &ShopCreatemarketingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
