package edas

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

// AbortChangeOrder invokes the edas.AbortChangeOrder API synchronously
func (client *Client) AbortChangeOrder(request *AbortChangeOrderRequest) (response *AbortChangeOrderResponse, err error) {
	response = CreateAbortChangeOrderResponse()
	err = client.DoAction(request, response)
	return
}

// AbortChangeOrderWithChan invokes the edas.AbortChangeOrder API asynchronously
func (client *Client) AbortChangeOrderWithChan(request *AbortChangeOrderRequest) (<-chan *AbortChangeOrderResponse, <-chan error) {
	responseChan := make(chan *AbortChangeOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AbortChangeOrder(request)
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

// AbortChangeOrderWithCallback invokes the edas.AbortChangeOrder API asynchronously
func (client *Client) AbortChangeOrderWithCallback(request *AbortChangeOrderRequest, callback func(response *AbortChangeOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AbortChangeOrderResponse
		var err error
		defer close(result)
		response, err = client.AbortChangeOrder(request)
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

// AbortChangeOrderRequest is the request struct for api AbortChangeOrder
type AbortChangeOrderRequest struct {
	*requests.RoaRequest
	ChangeOrderId string `position:"Query" name:"ChangeOrderId"`
}

// AbortChangeOrderResponse is the response struct for api AbortChangeOrder
type AbortChangeOrderResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAbortChangeOrderRequest creates a request to invoke AbortChangeOrder API
func CreateAbortChangeOrderRequest() (request *AbortChangeOrderRequest) {
	request = &AbortChangeOrderRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "AbortChangeOrder", "/pop/v5/changeorder/change_order_abort", "Edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateAbortChangeOrderResponse creates a response to parse from AbortChangeOrder response
func CreateAbortChangeOrderResponse() (response *AbortChangeOrderResponse) {
	response = &AbortChangeOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
