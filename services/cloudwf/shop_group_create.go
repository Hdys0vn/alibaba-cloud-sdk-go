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

// ShopGroupCreate invokes the cloudwf.ShopGroupCreate API synchronously
// api document: https://help.aliyun.com/api/cloudwf/shopgroupcreate.html
func (client *Client) ShopGroupCreate(request *ShopGroupCreateRequest) (response *ShopGroupCreateResponse, err error) {
	response = CreateShopGroupCreateResponse()
	err = client.DoAction(request, response)
	return
}

// ShopGroupCreateWithChan invokes the cloudwf.ShopGroupCreate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/shopgroupcreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ShopGroupCreateWithChan(request *ShopGroupCreateRequest) (<-chan *ShopGroupCreateResponse, <-chan error) {
	responseChan := make(chan *ShopGroupCreateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ShopGroupCreate(request)
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

// ShopGroupCreateWithCallback invokes the cloudwf.ShopGroupCreate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/shopgroupcreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ShopGroupCreateWithCallback(request *ShopGroupCreateRequest, callback func(response *ShopGroupCreateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ShopGroupCreateResponse
		var err error
		defer close(result)
		response, err = client.ShopGroupCreate(request)
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

// ShopGroupCreateRequest is the request struct for api ShopGroupCreate
type ShopGroupCreateRequest struct {
	*requests.RpcRequest
	ShopIds     string           `position:"Query" name:"ShopIds"`
	Name        string           `position:"Query" name:"Name"`
	Description string           `position:"Query" name:"Description"`
	Bid         requests.Integer `position:"Query" name:"Bid"`
}

// ShopGroupCreateResponse is the response struct for api ShopGroupCreate
type ShopGroupCreateResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateShopGroupCreateRequest creates a request to invoke ShopGroupCreate API
func CreateShopGroupCreateRequest() (request *ShopGroupCreateRequest) {
	request = &ShopGroupCreateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupCreate", "cloudwf", "openAPI")
	return
}

// CreateShopGroupCreateResponse creates a response to parse from ShopGroupCreate response
func CreateShopGroupCreateResponse() (response *ShopGroupCreateResponse) {
	response = &ShopGroupCreateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
