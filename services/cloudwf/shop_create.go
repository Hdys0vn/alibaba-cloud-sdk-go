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

// ShopCreate invokes the cloudwf.ShopCreate API synchronously
// api document: https://help.aliyun.com/api/cloudwf/shopcreate.html
func (client *Client) ShopCreate(request *ShopCreateRequest) (response *ShopCreateResponse, err error) {
	response = CreateShopCreateResponse()
	err = client.DoAction(request, response)
	return
}

// ShopCreateWithChan invokes the cloudwf.ShopCreate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/shopcreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ShopCreateWithChan(request *ShopCreateRequest) (<-chan *ShopCreateResponse, <-chan error) {
	responseChan := make(chan *ShopCreateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ShopCreate(request)
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

// ShopCreateWithCallback invokes the cloudwf.ShopCreate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/shopcreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ShopCreateWithCallback(request *ShopCreateRequest, callback func(response *ShopCreateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ShopCreateResponse
		var err error
		defer close(result)
		response, err = client.ShopCreate(request)
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

// ShopCreateRequest is the request struct for api ShopCreate
type ShopCreateRequest struct {
	*requests.RpcRequest
	ShopCoordinate    string           `position:"Query" name:"ShopCoordinate"`
	ShopProvince      string           `position:"Query" name:"ShopProvince"`
	ShopTopType       requests.Integer `position:"Query" name:"ShopTopType"`
	ShopAddress       string           `position:"Query" name:"ShopAddress"`
	ShopType          requests.Integer `position:"Query" name:"ShopType"`
	WarnEmail         string           `position:"Query" name:"WarnEmail"`
	ShopTel           string           `position:"Query" name:"ShopTel"`
	WarnpHone         string           `position:"Query" name:"WarnpHone"`
	Warn              requests.Integer `position:"Query" name:"Warn"`
	ShopArea          requests.Integer `position:"Query" name:"ShopArea"`
	ShopRemarks       string           `position:"Query" name:"ShopRemarks"`
	ShopCity          string           `position:"Query" name:"ShopCity"`
	ShopSubtype       requests.Integer `position:"Query" name:"ShopSubtype"`
	ShopBrand         string           `position:"Query" name:"ShopBrand"`
	ShopName          string           `position:"Query" name:"ShopName"`
	ShopCloseWarn     requests.Integer `position:"Query" name:"ShopCloseWarn"`
	Bid               requests.Integer `position:"Query" name:"Bid"`
	ShopManager       string           `position:"Query" name:"ShopManager"`
	ShopBusinessHours string           `position:"Query" name:"ShopBusinessHours"`
}

// ShopCreateResponse is the response struct for api ShopCreate
type ShopCreateResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateShopCreateRequest creates a request to invoke ShopCreate API
func CreateShopCreateRequest() (request *ShopCreateRequest) {
	request = &ShopCreateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ShopCreate", "cloudwf", "openAPI")
	return
}

// CreateShopCreateResponse creates a response to parse from ShopCreate response
func CreateShopCreateResponse() (response *ShopCreateResponse) {
	response = &ShopCreateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
