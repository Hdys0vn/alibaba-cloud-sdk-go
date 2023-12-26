package nlb

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

// CancelShiftLoadBalancerZones invokes the nlb.CancelShiftLoadBalancerZones API synchronously
func (client *Client) CancelShiftLoadBalancerZones(request *CancelShiftLoadBalancerZonesRequest) (response *CancelShiftLoadBalancerZonesResponse, err error) {
	response = CreateCancelShiftLoadBalancerZonesResponse()
	err = client.DoAction(request, response)
	return
}

// CancelShiftLoadBalancerZonesWithChan invokes the nlb.CancelShiftLoadBalancerZones API asynchronously
func (client *Client) CancelShiftLoadBalancerZonesWithChan(request *CancelShiftLoadBalancerZonesRequest) (<-chan *CancelShiftLoadBalancerZonesResponse, <-chan error) {
	responseChan := make(chan *CancelShiftLoadBalancerZonesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelShiftLoadBalancerZones(request)
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

// CancelShiftLoadBalancerZonesWithCallback invokes the nlb.CancelShiftLoadBalancerZones API asynchronously
func (client *Client) CancelShiftLoadBalancerZonesWithCallback(request *CancelShiftLoadBalancerZonesRequest, callback func(response *CancelShiftLoadBalancerZonesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelShiftLoadBalancerZonesResponse
		var err error
		defer close(result)
		response, err = client.CancelShiftLoadBalancerZones(request)
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

// CancelShiftLoadBalancerZonesRequest is the request struct for api CancelShiftLoadBalancerZones
type CancelShiftLoadBalancerZonesRequest struct {
	*requests.RpcRequest
	ClientToken    string                                      `position:"Body" name:"ClientToken"`
	DryRun         requests.Boolean                            `position:"Body" name:"DryRun"`
	ZoneMappings   *[]CancelShiftLoadBalancerZonesZoneMappings `position:"Body" name:"ZoneMappings"  type:"Repeated"`
	LoadBalancerId string                                      `position:"Body" name:"LoadBalancerId"`
}

// CancelShiftLoadBalancerZonesZoneMappings is a repeated param struct in CancelShiftLoadBalancerZonesRequest
type CancelShiftLoadBalancerZonesZoneMappings struct {
	VSwitchId string `name:"VSwitchId"`
	ZoneId    string `name:"ZoneId"`
}

// CancelShiftLoadBalancerZonesResponse is the response struct for api CancelShiftLoadBalancerZones
type CancelShiftLoadBalancerZonesResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
}

// CreateCancelShiftLoadBalancerZonesRequest creates a request to invoke CancelShiftLoadBalancerZones API
func CreateCancelShiftLoadBalancerZonesRequest() (request *CancelShiftLoadBalancerZonesRequest) {
	request = &CancelShiftLoadBalancerZonesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "CancelShiftLoadBalancerZones", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelShiftLoadBalancerZonesResponse creates a response to parse from CancelShiftLoadBalancerZones response
func CreateCancelShiftLoadBalancerZonesResponse() (response *CancelShiftLoadBalancerZonesResponse) {
	response = &CancelShiftLoadBalancerZonesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
