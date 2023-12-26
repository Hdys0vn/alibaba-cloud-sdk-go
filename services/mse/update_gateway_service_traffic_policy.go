package mse

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

// UpdateGatewayServiceTrafficPolicy invokes the mse.UpdateGatewayServiceTrafficPolicy API synchronously
func (client *Client) UpdateGatewayServiceTrafficPolicy(request *UpdateGatewayServiceTrafficPolicyRequest) (response *UpdateGatewayServiceTrafficPolicyResponse, err error) {
	response = CreateUpdateGatewayServiceTrafficPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGatewayServiceTrafficPolicyWithChan invokes the mse.UpdateGatewayServiceTrafficPolicy API asynchronously
func (client *Client) UpdateGatewayServiceTrafficPolicyWithChan(request *UpdateGatewayServiceTrafficPolicyRequest) (<-chan *UpdateGatewayServiceTrafficPolicyResponse, <-chan error) {
	responseChan := make(chan *UpdateGatewayServiceTrafficPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGatewayServiceTrafficPolicy(request)
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

// UpdateGatewayServiceTrafficPolicyWithCallback invokes the mse.UpdateGatewayServiceTrafficPolicy API asynchronously
func (client *Client) UpdateGatewayServiceTrafficPolicyWithCallback(request *UpdateGatewayServiceTrafficPolicyRequest, callback func(response *UpdateGatewayServiceTrafficPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGatewayServiceTrafficPolicyResponse
		var err error
		defer close(result)
		response, err = client.UpdateGatewayServiceTrafficPolicy(request)
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

// UpdateGatewayServiceTrafficPolicyRequest is the request struct for api UpdateGatewayServiceTrafficPolicy
type UpdateGatewayServiceTrafficPolicyRequest struct {
	*requests.RpcRequest
	MseSessionId         string                                                `position:"Query" name:"MseSessionId"`
	GatewayUniqueId      string                                                `position:"Query" name:"GatewayUniqueId"`
	GatewayId            requests.Integer                                      `position:"Query" name:"GatewayId"`
	GatewayTrafficPolicy UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicy `position:"Query" name:"GatewayTrafficPolicy"  type:"Struct"`
	AcceptLanguage       string                                                `position:"Query" name:"AcceptLanguage"`
	ServiceId            requests.Integer                                      `position:"Query" name:"ServiceId"`
}

// UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicy is a repeated param struct in UpdateGatewayServiceTrafficPolicyRequest
type UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicy struct {
	TlsSetting           UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyTlsSetting           `name:"TlsSetting" type:"Struct"`
	LoadBalancerSettings UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyLoadBalancerSettings `name:"LoadBalancerSettings" type:"Struct"`
}

// UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyTlsSetting is a repeated param struct in UpdateGatewayServiceTrafficPolicyRequest
type UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyTlsSetting struct {
	TlsMode       string `name:"TlsMode"`
	CaCertContent string `name:"CaCertContent"`
	CertId        string `name:"CertId"`
	Sni           string `name:"Sni"`
}

// UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyLoadBalancerSettings is a repeated param struct in UpdateGatewayServiceTrafficPolicyRequest
type UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyLoadBalancerSettings struct {
	WarmupDuration         string                                                                                          `name:"WarmupDuration"`
	LoadbalancerType       string                                                                                          `name:"LoadbalancerType"`
	ConsistentHashLBConfig UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig `name:"ConsistentHashLBConfig" type:"Struct"`
}

// UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig is a repeated param struct in UpdateGatewayServiceTrafficPolicyRequest
type UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfig struct {
	HttpCookie           UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie `name:"HttpCookie" type:"Struct"`
	ParameterName        string                                                                                                    `name:"ParameterName"`
	ConsistentHashLBType string                                                                                                    `name:"ConsistentHashLBType"`
}

// UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie is a repeated param struct in UpdateGatewayServiceTrafficPolicyRequest
type UpdateGatewayServiceTrafficPolicyGatewayTrafficPolicyLoadBalancerSettingsConsistentHashLBConfigHttpCookie struct {
	Path string `name:"Path"`
	Name string `name:"Name"`
	TTL  string `name:"TTL"`
}

// UpdateGatewayServiceTrafficPolicyResponse is the response struct for api UpdateGatewayServiceTrafficPolicy
type UpdateGatewayServiceTrafficPolicyResponse struct {
	*responses.BaseResponse
	RequestId      string                                  `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int                                     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                                  `json:"Message" xml:"Message"`
	Code           int                                     `json:"Code" xml:"Code"`
	Success        bool                                    `json:"Success" xml:"Success"`
	Data           DataInUpdateGatewayServiceTrafficPolicy `json:"Data" xml:"Data"`
}

// CreateUpdateGatewayServiceTrafficPolicyRequest creates a request to invoke UpdateGatewayServiceTrafficPolicy API
func CreateUpdateGatewayServiceTrafficPolicyRequest() (request *UpdateGatewayServiceTrafficPolicyRequest) {
	request = &UpdateGatewayServiceTrafficPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateGatewayServiceTrafficPolicy", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateGatewayServiceTrafficPolicyResponse creates a response to parse from UpdateGatewayServiceTrafficPolicy response
func CreateUpdateGatewayServiceTrafficPolicyResponse() (response *UpdateGatewayServiceTrafficPolicyResponse) {
	response = &UpdateGatewayServiceTrafficPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
