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

// ShopGroupUpdate invokes the cloudwf.ShopGroupUpdate API synchronously
// api document: https://help.aliyun.com/api/cloudwf/shopgroupupdate.html
func (client *Client) ShopGroupUpdate(request *ShopGroupUpdateRequest) (response *ShopGroupUpdateResponse, err error) {
	response = CreateShopGroupUpdateResponse()
	err = client.DoAction(request, response)
	return
}

// ShopGroupUpdateWithChan invokes the cloudwf.ShopGroupUpdate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/shopgroupupdate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ShopGroupUpdateWithChan(request *ShopGroupUpdateRequest) (<-chan *ShopGroupUpdateResponse, <-chan error) {
	responseChan := make(chan *ShopGroupUpdateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ShopGroupUpdate(request)
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

// ShopGroupUpdateWithCallback invokes the cloudwf.ShopGroupUpdate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/shopgroupupdate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ShopGroupUpdateWithCallback(request *ShopGroupUpdateRequest, callback func(response *ShopGroupUpdateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ShopGroupUpdateResponse
		var err error
		defer close(result)
		response, err = client.ShopGroupUpdate(request)
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

// ShopGroupUpdateRequest is the request struct for api ShopGroupUpdate
type ShopGroupUpdateRequest struct {
	*requests.RpcRequest
	Gid         requests.Integer `position:"Query" name:"Gid"`
	ShopIds     string           `position:"Query" name:"ShopIds"`
	Name        string           `position:"Query" name:"Name"`
	Description string           `position:"Query" name:"Description"`
}

// ShopGroupUpdateResponse is the response struct for api ShopGroupUpdate
type ShopGroupUpdateResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateShopGroupUpdateRequest creates a request to invoke ShopGroupUpdate API
func CreateShopGroupUpdateRequest() (request *ShopGroupUpdateRequest) {
	request = &ShopGroupUpdateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupUpdate", "cloudwf", "openAPI")
	return
}

// CreateShopGroupUpdateResponse creates a response to parse from ShopGroupUpdate response
func CreateShopGroupUpdateResponse() (response *ShopGroupUpdateResponse) {
	response = &ShopGroupUpdateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
