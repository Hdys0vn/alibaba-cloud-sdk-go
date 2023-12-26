package cloudgameapi

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

// DeliveryOrder invokes the cloudgameapi.DeliveryOrder API synchronously
func (client *Client) DeliveryOrder(request *DeliveryOrderRequest) (response *DeliveryOrderResponse, err error) {
	response = CreateDeliveryOrderResponse()
	err = client.DoAction(request, response)
	return
}

// DeliveryOrderWithChan invokes the cloudgameapi.DeliveryOrder API asynchronously
func (client *Client) DeliveryOrderWithChan(request *DeliveryOrderRequest) (<-chan *DeliveryOrderResponse, <-chan error) {
	responseChan := make(chan *DeliveryOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeliveryOrder(request)
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

// DeliveryOrderWithCallback invokes the cloudgameapi.DeliveryOrder API asynchronously
func (client *Client) DeliveryOrderWithCallback(request *DeliveryOrderRequest, callback func(response *DeliveryOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeliveryOrderResponse
		var err error
		defer close(result)
		response, err = client.DeliveryOrder(request)
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

// DeliveryOrderRequest is the request struct for api DeliveryOrder
type DeliveryOrderRequest struct {
	*requests.RpcRequest
	AccountDomain  string `position:"Query" name:"AccountDomain"`
	OrderId        string `position:"Query" name:"OrderId"`
	BuyerAccountId string `position:"Query" name:"BuyerAccountId"`
}

// DeliveryOrderResponse is the response struct for api DeliveryOrder
type DeliveryOrderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDeliveryOrderRequest creates a request to invoke DeliveryOrder API
func CreateDeliveryOrderRequest() (request *DeliveryOrderRequest) {
	request = &DeliveryOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudGameAPI", "2020-07-28", "DeliveryOrder", "", "")
	request.Method = requests.POST
	return
}

// CreateDeliveryOrderResponse creates a response to parse from DeliveryOrder response
func CreateDeliveryOrderResponse() (response *DeliveryOrderResponse) {
	response = &DeliveryOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}