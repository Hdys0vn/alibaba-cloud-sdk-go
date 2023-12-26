package alidns

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

// UpdateGtmRecoveryPlan invokes the alidns.UpdateGtmRecoveryPlan API synchronously
func (client *Client) UpdateGtmRecoveryPlan(request *UpdateGtmRecoveryPlanRequest) (response *UpdateGtmRecoveryPlanResponse, err error) {
	response = CreateUpdateGtmRecoveryPlanResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGtmRecoveryPlanWithChan invokes the alidns.UpdateGtmRecoveryPlan API asynchronously
func (client *Client) UpdateGtmRecoveryPlanWithChan(request *UpdateGtmRecoveryPlanRequest) (<-chan *UpdateGtmRecoveryPlanResponse, <-chan error) {
	responseChan := make(chan *UpdateGtmRecoveryPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGtmRecoveryPlan(request)
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

// UpdateGtmRecoveryPlanWithCallback invokes the alidns.UpdateGtmRecoveryPlan API asynchronously
func (client *Client) UpdateGtmRecoveryPlanWithCallback(request *UpdateGtmRecoveryPlanRequest, callback func(response *UpdateGtmRecoveryPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGtmRecoveryPlanResponse
		var err error
		defer close(result)
		response, err = client.UpdateGtmRecoveryPlan(request)
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

// UpdateGtmRecoveryPlanRequest is the request struct for api UpdateGtmRecoveryPlan
type UpdateGtmRecoveryPlanRequest struct {
	*requests.RpcRequest
	FaultAddrPool  string           `position:"Query" name:"FaultAddrPool"`
	Remark         string           `position:"Query" name:"Remark"`
	UserClientIp   string           `position:"Query" name:"UserClientIp"`
	Name           string           `position:"Query" name:"Name"`
	RecoveryPlanId requests.Integer `position:"Query" name:"RecoveryPlanId"`
	Lang           string           `position:"Query" name:"Lang"`
}

// UpdateGtmRecoveryPlanResponse is the response struct for api UpdateGtmRecoveryPlan
type UpdateGtmRecoveryPlanResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateGtmRecoveryPlanRequest creates a request to invoke UpdateGtmRecoveryPlan API
func CreateUpdateGtmRecoveryPlanRequest() (request *UpdateGtmRecoveryPlanRequest) {
	request = &UpdateGtmRecoveryPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "UpdateGtmRecoveryPlan", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGtmRecoveryPlanResponse creates a response to parse from UpdateGtmRecoveryPlan response
func CreateUpdateGtmRecoveryPlanResponse() (response *UpdateGtmRecoveryPlanResponse) {
	response = &UpdateGtmRecoveryPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
