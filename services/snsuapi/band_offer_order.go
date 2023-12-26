package snsuapi

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

// BandOfferOrder invokes the snsuapi.BandOfferOrder API synchronously
// api document: https://help.aliyun.com/api/snsuapi/bandofferorder.html
func (client *Client) BandOfferOrder(request *BandOfferOrderRequest) (response *BandOfferOrderResponse, err error) {
	response = CreateBandOfferOrderResponse()
	err = client.DoAction(request, response)
	return
}

// BandOfferOrderWithChan invokes the snsuapi.BandOfferOrder API asynchronously
// api document: https://help.aliyun.com/api/snsuapi/bandofferorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BandOfferOrderWithChan(request *BandOfferOrderRequest) (<-chan *BandOfferOrderResponse, <-chan error) {
	responseChan := make(chan *BandOfferOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BandOfferOrder(request)
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

// BandOfferOrderWithCallback invokes the snsuapi.BandOfferOrder API asynchronously
// api document: https://help.aliyun.com/api/snsuapi/bandofferorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BandOfferOrderWithCallback(request *BandOfferOrderRequest, callback func(response *BandOfferOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BandOfferOrderResponse
		var err error
		defer close(result)
		response, err = client.BandOfferOrder(request)
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

// BandOfferOrderRequest is the request struct for api BandOfferOrder
type BandOfferOrderRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BandId               string           `position:"Query" name:"BandId"`
	OfferId              string           `position:"Query" name:"OfferId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// BandOfferOrderResponse is the response struct for api BandOfferOrder
type BandOfferOrderResponse struct {
	*responses.BaseResponse
	RequestId     string       `json:"RequestId" xml:"RequestId"`
	ResultCode    string       `json:"ResultCode" xml:"ResultCode"`
	ResultMessage string       `json:"ResultMessage" xml:"ResultMessage"`
	ResultModule  ResultModule `json:"ResultModule" xml:"ResultModule"`
}

// CreateBandOfferOrderRequest creates a request to invoke BandOfferOrder API
func CreateBandOfferOrderRequest() (request *BandOfferOrderRequest) {
	request = &BandOfferOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Snsuapi", "2018-07-09", "BandOfferOrder", "snsuapi", "openAPI")
	return
}

// CreateBandOfferOrderResponse creates a response to parse from BandOfferOrder response
func CreateBandOfferOrderResponse() (response *BandOfferOrderResponse) {
	response = &BandOfferOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
