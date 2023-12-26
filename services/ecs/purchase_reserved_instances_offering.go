package ecs

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

// PurchaseReservedInstancesOffering invokes the ecs.PurchaseReservedInstancesOffering API synchronously
func (client *Client) PurchaseReservedInstancesOffering(request *PurchaseReservedInstancesOfferingRequest) (response *PurchaseReservedInstancesOfferingResponse, err error) {
	response = CreatePurchaseReservedInstancesOfferingResponse()
	err = client.DoAction(request, response)
	return
}

// PurchaseReservedInstancesOfferingWithChan invokes the ecs.PurchaseReservedInstancesOffering API asynchronously
func (client *Client) PurchaseReservedInstancesOfferingWithChan(request *PurchaseReservedInstancesOfferingRequest) (<-chan *PurchaseReservedInstancesOfferingResponse, <-chan error) {
	responseChan := make(chan *PurchaseReservedInstancesOfferingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PurchaseReservedInstancesOffering(request)
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

// PurchaseReservedInstancesOfferingWithCallback invokes the ecs.PurchaseReservedInstancesOffering API asynchronously
func (client *Client) PurchaseReservedInstancesOfferingWithCallback(request *PurchaseReservedInstancesOfferingRequest, callback func(response *PurchaseReservedInstancesOfferingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PurchaseReservedInstancesOfferingResponse
		var err error
		defer close(result)
		response, err = client.PurchaseReservedInstancesOffering(request)
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

// PurchaseReservedInstancesOfferingRequest is the request struct for api PurchaseReservedInstancesOffering
type PurchaseReservedInstancesOfferingRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                        `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string                                  `position:"Query" name:"ClientToken"`
	Description          string                                  `position:"Query" name:"Description"`
	Platform             string                                  `position:"Query" name:"Platform"`
	ResourceGroupId      string                                  `position:"Query" name:"ResourceGroupId"`
	Scope                string                                  `position:"Query" name:"Scope"`
	InstanceType         string                                  `position:"Query" name:"InstanceType"`
	Tag                  *[]PurchaseReservedInstancesOfferingTag `position:"Query" name:"Tag"  type:"Repeated"`
	AutoRenewPeriod      requests.Integer                        `position:"Query" name:"AutoRenewPeriod"`
	Period               requests.Integer                        `position:"Query" name:"Period"`
	ResourceOwnerAccount string                                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                  `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                        `position:"Query" name:"OwnerId"`
	PeriodUnit           string                                  `position:"Query" name:"PeriodUnit"`
	OfferingType         string                                  `position:"Query" name:"OfferingType"`
	AutoRenew            requests.Boolean                        `position:"Query" name:"AutoRenew"`
	ZoneId               string                                  `position:"Query" name:"ZoneId"`
	ReservedInstanceName string                                  `position:"Query" name:"ReservedInstanceName"`
	InstanceAmount       requests.Integer                        `position:"Query" name:"InstanceAmount"`
}

// PurchaseReservedInstancesOfferingTag is a repeated param struct in PurchaseReservedInstancesOfferingRequest
type PurchaseReservedInstancesOfferingTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// PurchaseReservedInstancesOfferingResponse is the response struct for api PurchaseReservedInstancesOffering
type PurchaseReservedInstancesOfferingResponse struct {
	*responses.BaseResponse
	RequestId              string                                                    `json:"RequestId" xml:"RequestId"`
	ReservedInstanceIdSets ReservedInstanceIdSetsInPurchaseReservedInstancesOffering `json:"ReservedInstanceIdSets" xml:"ReservedInstanceIdSets"`
}

// CreatePurchaseReservedInstancesOfferingRequest creates a request to invoke PurchaseReservedInstancesOffering API
func CreatePurchaseReservedInstancesOfferingRequest() (request *PurchaseReservedInstancesOfferingRequest) {
	request = &PurchaseReservedInstancesOfferingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "PurchaseReservedInstancesOffering", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePurchaseReservedInstancesOfferingResponse creates a response to parse from PurchaseReservedInstancesOffering response
func CreatePurchaseReservedInstancesOfferingResponse() (response *PurchaseReservedInstancesOfferingResponse) {
	response = &PurchaseReservedInstancesOfferingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
