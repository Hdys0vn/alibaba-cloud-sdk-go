package ecd

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

// ModifyDesktopChargeType invokes the ecd.ModifyDesktopChargeType API synchronously
func (client *Client) ModifyDesktopChargeType(request *ModifyDesktopChargeTypeRequest) (response *ModifyDesktopChargeTypeResponse, err error) {
	response = CreateModifyDesktopChargeTypeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDesktopChargeTypeWithChan invokes the ecd.ModifyDesktopChargeType API asynchronously
func (client *Client) ModifyDesktopChargeTypeWithChan(request *ModifyDesktopChargeTypeRequest) (<-chan *ModifyDesktopChargeTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyDesktopChargeTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDesktopChargeType(request)
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

// ModifyDesktopChargeTypeWithCallback invokes the ecd.ModifyDesktopChargeType API asynchronously
func (client *Client) ModifyDesktopChargeTypeWithCallback(request *ModifyDesktopChargeTypeRequest, callback func(response *ModifyDesktopChargeTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDesktopChargeTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDesktopChargeType(request)
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

// ModifyDesktopChargeTypeRequest is the request struct for api ModifyDesktopChargeType
type ModifyDesktopChargeTypeRequest struct {
	*requests.RpcRequest
	Period      requests.Integer `position:"Query" name:"Period"`
	AutoPay     requests.Boolean `position:"Query" name:"AutoPay"`
	PromotionId string           `position:"Query" name:"PromotionId"`
	PeriodUnit  string           `position:"Query" name:"PeriodUnit"`
	ChargeType  string           `position:"Query" name:"ChargeType"`
	DesktopId   *[]string        `position:"Query" name:"DesktopId"  type:"Repeated"`
}

// ModifyDesktopChargeTypeResponse is the response struct for api ModifyDesktopChargeType
type ModifyDesktopChargeTypeResponse struct {
	*responses.BaseResponse
	OrderId   string   `json:"OrderId" xml:"OrderId"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	DesktopId []string `json:"DesktopId" xml:"DesktopId"`
}

// CreateModifyDesktopChargeTypeRequest creates a request to invoke ModifyDesktopChargeType API
func CreateModifyDesktopChargeTypeRequest() (request *ModifyDesktopChargeTypeRequest) {
	request = &ModifyDesktopChargeTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ModifyDesktopChargeType", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDesktopChargeTypeResponse creates a response to parse from ModifyDesktopChargeType response
func CreateModifyDesktopChargeTypeResponse() (response *ModifyDesktopChargeTypeResponse) {
	response = &ModifyDesktopChargeTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}