package ens

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

// UpdateEnsSaleControl invokes the ens.UpdateEnsSaleControl API synchronously
func (client *Client) UpdateEnsSaleControl(request *UpdateEnsSaleControlRequest) (response *UpdateEnsSaleControlResponse, err error) {
	response = CreateUpdateEnsSaleControlResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEnsSaleControlWithChan invokes the ens.UpdateEnsSaleControl API asynchronously
func (client *Client) UpdateEnsSaleControlWithChan(request *UpdateEnsSaleControlRequest) (<-chan *UpdateEnsSaleControlResponse, <-chan error) {
	responseChan := make(chan *UpdateEnsSaleControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEnsSaleControl(request)
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

// UpdateEnsSaleControlWithCallback invokes the ens.UpdateEnsSaleControl API asynchronously
func (client *Client) UpdateEnsSaleControlWithCallback(request *UpdateEnsSaleControlRequest, callback func(response *UpdateEnsSaleControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEnsSaleControlResponse
		var err error
		defer close(result)
		response, err = client.UpdateEnsSaleControl(request)
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

// UpdateEnsSaleControlRequest is the request struct for api UpdateEnsSaleControl
type UpdateEnsSaleControlRequest struct {
	*requests.RpcRequest
	SaleControls  *[]UpdateEnsSaleControlSaleControls `position:"Query" name:"SaleControls"  type:"Json"`
	CommodityCode string                              `position:"Query" name:"CommodityCode"`
	AliUidAccount string                              `position:"Query" name:"AliUidAccount"`
}

// UpdateEnsSaleControlSaleControls is a repeated param struct in UpdateEnsSaleControlRequest
type UpdateEnsSaleControlSaleControls struct {
	ModuleCode        string                                                   `name:"ModuleCode"`
	ModuleValue       UpdateEnsSaleControlSaleControlsModuleValue              `name:"ModuleValue" type:"Struct"`
	ConditionControls *[]UpdateEnsSaleControlSaleControlsConditionControlsItem `name:"ConditionControls" type:"Repeated"`
	Description       string                                                   `name:"Description"`
	Operator          string                                                   `name:"Operator"`
	OrderType         string                                                   `name:"OrderType"`
}

// UpdateEnsSaleControlSaleControlsModuleValue is a repeated param struct in UpdateEnsSaleControlRequest
type UpdateEnsSaleControlSaleControlsModuleValue struct {
	ModuleValue    *[]string `name:"ModuleValue" type:"Repeated"`
	ModuleMaxValue string    `name:"ModuleMaxValue"`
	ModuleMinValue string    `name:"ModuleMinValue"`
}

// UpdateEnsSaleControlSaleControlsConditionControlsItem is a repeated param struct in UpdateEnsSaleControlRequest
type UpdateEnsSaleControlSaleControlsConditionControlsItem struct {
	ConditionControlModuleValue string `name:"ConditionControlModuleValue"`
	ConditionControlModuleCode  string `name:"ConditionControlModuleCode"`
}

// UpdateEnsSaleControlResponse is the response struct for api UpdateEnsSaleControl
type UpdateEnsSaleControlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateEnsSaleControlRequest creates a request to invoke UpdateEnsSaleControl API
func CreateUpdateEnsSaleControlRequest() (request *UpdateEnsSaleControlRequest) {
	request = &UpdateEnsSaleControlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "UpdateEnsSaleControl", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateEnsSaleControlResponse creates a response to parse from UpdateEnsSaleControl response
func CreateUpdateEnsSaleControlResponse() (response *UpdateEnsSaleControlResponse) {
	response = &UpdateEnsSaleControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}