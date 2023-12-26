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

// CreateLoadBalancer invokes the nlb.CreateLoadBalancer API synchronously
func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
	response = CreateCreateLoadBalancerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLoadBalancerWithChan invokes the nlb.CreateLoadBalancer API asynchronously
func (client *Client) CreateLoadBalancerWithChan(request *CreateLoadBalancerRequest) (<-chan *CreateLoadBalancerResponse, <-chan error) {
	responseChan := make(chan *CreateLoadBalancerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLoadBalancer(request)
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

// CreateLoadBalancerWithCallback invokes the nlb.CreateLoadBalancer API asynchronously
func (client *Client) CreateLoadBalancerWithCallback(request *CreateLoadBalancerRequest, callback func(response *CreateLoadBalancerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLoadBalancerResponse
		var err error
		defer close(result)
		response, err = client.CreateLoadBalancer(request)
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

// CreateLoadBalancerRequest is the request struct for api CreateLoadBalancer
type CreateLoadBalancerRequest struct {
	*requests.RpcRequest
	ClientToken                  string                                         `position:"Body" name:"ClientToken"`
	ModificationProtectionConfig CreateLoadBalancerModificationProtectionConfig `position:"Body" name:"ModificationProtectionConfig"  type:"Struct"`
	LoadBalancerBillingConfig    CreateLoadBalancerLoadBalancerBillingConfig    `position:"Body" name:"LoadBalancerBillingConfig"  type:"Struct"`
	BizFlag                      string                                         `position:"Body" name:"BizFlag"`
	DeletionProtectionConfig     CreateLoadBalancerDeletionProtectionConfig     `position:"Body" name:"DeletionProtectionConfig"  type:"Struct"`
	AddressIpVersion             string                                         `position:"Body" name:"AddressIpVersion"`
	ResourceGroupId              string                                         `position:"Body" name:"ResourceGroupId"`
	LoadBalancerName             string                                         `position:"Body" name:"LoadBalancerName"`
	AddressType                  string                                         `position:"Body" name:"AddressType"`
	Tag                          *[]CreateLoadBalancerTag                       `position:"Body" name:"Tag"  type:"Repeated"`
	BandwidthPackageId           string                                         `position:"Body" name:"BandwidthPackageId"`
	DryRun                       requests.Boolean                               `position:"Body" name:"DryRun"`
	ZoneMappings                 *[]CreateLoadBalancerZoneMappings              `position:"Body" name:"ZoneMappings"  type:"Repeated"`
	SecurityGroupIds             *[]string                                      `position:"Body" name:"SecurityGroupIds"  type:"Repeated"`
	LoadBalancerType             string                                         `position:"Body" name:"LoadBalancerType"`
	VpcId                        string                                         `position:"Body" name:"VpcId"`
}

// CreateLoadBalancerZoneMappings is a repeated param struct in CreateLoadBalancerRequest
type CreateLoadBalancerZoneMappings struct {
	VSwitchId          string `name:"VSwitchId"`
	ZoneId             string `name:"ZoneId"`
	PrivateIPv4Address string `name:"PrivateIPv4Address"`
	AllocationId       string `name:"AllocationId"`
}

// CreateLoadBalancerModificationProtectionConfig is a repeated param struct in CreateLoadBalancerRequest
type CreateLoadBalancerModificationProtectionConfig struct {
	Status string `name:"Status"`
	Reason string `name:"Reason"`
}

// CreateLoadBalancerLoadBalancerBillingConfig is a repeated param struct in CreateLoadBalancerRequest
type CreateLoadBalancerLoadBalancerBillingConfig struct {
	PayType string `name:"PayType"`
}

// CreateLoadBalancerDeletionProtectionConfig is a repeated param struct in CreateLoadBalancerRequest
type CreateLoadBalancerDeletionProtectionConfig struct {
	Enabled string `name:"Enabled"`
	Reason  string `name:"Reason"`
}

// CreateLoadBalancerTag is a repeated param struct in CreateLoadBalancerRequest
type CreateLoadBalancerTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateLoadBalancerResponse is the response struct for api CreateLoadBalancer
type CreateLoadBalancerResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	LoadbalancerId string `json:"LoadbalancerId" xml:"LoadbalancerId"`
	JobId          string `json:"JobId" xml:"JobId"`
	OrderId        int64  `json:"OrderId" xml:"OrderId"`
}

// CreateCreateLoadBalancerRequest creates a request to invoke CreateLoadBalancer API
func CreateCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
	request = &CreateLoadBalancerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "CreateLoadBalancer", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLoadBalancerResponse creates a response to parse from CreateLoadBalancer response
func CreateCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
	response = &CreateLoadBalancerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
