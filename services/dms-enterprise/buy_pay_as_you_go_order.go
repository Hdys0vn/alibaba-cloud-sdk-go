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

// BuyPayAsYouGoOrder invokes the dms_enterprise.BuyPayAsYouGoOrder API synchronously
func (client *Client) BuyPayAsYouGoOrder(request *BuyPayAsYouGoOrderRequest) (response *BuyPayAsYouGoOrderResponse, err error) {
	response = CreateBuyPayAsYouGoOrderResponse()
	err = client.DoAction(request, response)
	return
}

// BuyPayAsYouGoOrderWithChan invokes the dms_enterprise.BuyPayAsYouGoOrder API asynchronously
func (client *Client) BuyPayAsYouGoOrderWithChan(request *BuyPayAsYouGoOrderRequest) (<-chan *BuyPayAsYouGoOrderResponse, <-chan error) {
	responseChan := make(chan *BuyPayAsYouGoOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BuyPayAsYouGoOrder(request)
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

// BuyPayAsYouGoOrderWithCallback invokes the dms_enterprise.BuyPayAsYouGoOrder API asynchronously
func (client *Client) BuyPayAsYouGoOrderWithCallback(request *BuyPayAsYouGoOrderRequest, callback func(response *BuyPayAsYouGoOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BuyPayAsYouGoOrderResponse
		var err error
		defer close(result)
		response, err = client.BuyPayAsYouGoOrder(request)
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

// BuyPayAsYouGoOrderRequest is the request struct for api BuyPayAsYouGoOrder
type BuyPayAsYouGoOrderRequest struct {
	*requests.RpcRequest
	Tid           requests.Integer `position:"Query" name:"Tid"`
	CommodityType string           `position:"Query" name:"CommodityType"`
	InsNum        requests.Integer `position:"Query" name:"InsNum"`
	VersionType   string           `position:"Query" name:"VersionType"`
}

// BuyPayAsYouGoOrderResponse is the response struct for api BuyPayAsYouGoOrder
type BuyPayAsYouGoOrderResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	InstanceId   string `json:"InstanceId" xml:"InstanceId"`
}

// CreateBuyPayAsYouGoOrderRequest creates a request to invoke BuyPayAsYouGoOrder API
func CreateBuyPayAsYouGoOrderRequest() (request *BuyPayAsYouGoOrderRequest) {
	request = &BuyPayAsYouGoOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "BuyPayAsYouGoOrder", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBuyPayAsYouGoOrderResponse creates a response to parse from BuyPayAsYouGoOrder response
func CreateBuyPayAsYouGoOrderResponse() (response *BuyPayAsYouGoOrderResponse) {
	response = &BuyPayAsYouGoOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
