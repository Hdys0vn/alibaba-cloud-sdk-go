package slb

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

// SetLoadBalancerHTTPListenerAttribute invokes the slb.SetLoadBalancerHTTPListenerAttribute API synchronously
func (client *Client) SetLoadBalancerHTTPListenerAttribute(request *SetLoadBalancerHTTPListenerAttributeRequest) (response *SetLoadBalancerHTTPListenerAttributeResponse, err error) {
	response = CreateSetLoadBalancerHTTPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// SetLoadBalancerHTTPListenerAttributeWithChan invokes the slb.SetLoadBalancerHTTPListenerAttribute API asynchronously
func (client *Client) SetLoadBalancerHTTPListenerAttributeWithChan(request *SetLoadBalancerHTTPListenerAttributeRequest) (<-chan *SetLoadBalancerHTTPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerHTTPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerHTTPListenerAttribute(request)
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

// SetLoadBalancerHTTPListenerAttributeWithCallback invokes the slb.SetLoadBalancerHTTPListenerAttribute API asynchronously
func (client *Client) SetLoadBalancerHTTPListenerAttributeWithCallback(request *SetLoadBalancerHTTPListenerAttributeRequest, callback func(response *SetLoadBalancerHTTPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerHTTPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerHTTPListenerAttribute(request)
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

// SetLoadBalancerHTTPListenerAttributeRequest is the request struct for api SetLoadBalancerHTTPListenerAttribute
type SetLoadBalancerHTTPListenerAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HealthCheckTimeout         requests.Integer `position:"Query" name:"HealthCheckTimeout"`
	XForwardedFor              string           `position:"Query" name:"XForwardedFor"`
	HealthCheckURI             string           `position:"Query" name:"HealthCheckURI"`
	XForwardedForSLBPORT       string           `position:"Query" name:"XForwardedFor_SLBPORT"`
	AclStatus                  string           `position:"Query" name:"AclStatus"`
	AclType                    string           `position:"Query" name:"AclType"`
	HealthCheck                string           `position:"Query" name:"HealthCheck"`
	VpcIds                     string           `position:"Query" name:"VpcIds"`
	VServerGroupId             string           `position:"Query" name:"VServerGroupId"`
	AclId                      string           `position:"Query" name:"AclId"`
	ForwardCode                requests.Integer `position:"Query" name:"ForwardCode"`
	Cookie                     string           `position:"Query" name:"Cookie"`
	HealthCheckMethod          string           `position:"Query" name:"HealthCheckMethod"`
	HealthCheckDomain          string           `position:"Query" name:"HealthCheckDomain"`
	RequestTimeout             requests.Integer `position:"Query" name:"RequestTimeout"`
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	Tags                       string           `position:"Query" name:"Tags"`
	LoadBalancerId             string           `position:"Query" name:"LoadBalancerId"`
	XForwardedForSLBIP         string           `position:"Query" name:"XForwardedFor_SLBIP"`
	HealthCheckInterval        requests.Integer `position:"Query" name:"HealthCheckInterval"`
	XForwardedForSLBID         string           `position:"Query" name:"XForwardedFor_SLBID"`
	AccessKeyId                string           `position:"Query" name:"access_key_id"`
	XForwardedForClientSrcPort string           `position:"Query" name:"XForwardedFor_ClientSrcPort"`
	Description                string           `position:"Query" name:"Description"`
	UnhealthyThreshold         requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold           requests.Integer `position:"Query" name:"HealthyThreshold"`
	Scheduler                  string           `position:"Query" name:"Scheduler"`
	MaxConnection              requests.Integer `position:"Query" name:"MaxConnection"`
	CookieTimeout              requests.Integer `position:"Query" name:"CookieTimeout"`
	StickySessionType          string           `position:"Query" name:"StickySessionType"`
	ListenerPort               requests.Integer `position:"Query" name:"ListenerPort"`
	HealthCheckType            string           `position:"Query" name:"HealthCheckType"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                  requests.Integer `position:"Query" name:"Bandwidth"`
	StickySession              string           `position:"Query" name:"StickySession"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
	Gzip                       string           `position:"Query" name:"Gzip"`
	IdleTimeout                requests.Integer `position:"Query" name:"IdleTimeout"`
	XForwardedForProto         string           `position:"Query" name:"XForwardedFor_proto"`
	HealthCheckConnectPort     requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode        string           `position:"Query" name:"HealthCheckHttpCode"`
	VServerGroup               string           `position:"Query" name:"VServerGroup"`
}

// SetLoadBalancerHTTPListenerAttributeResponse is the response struct for api SetLoadBalancerHTTPListenerAttribute
type SetLoadBalancerHTTPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetLoadBalancerHTTPListenerAttributeRequest creates a request to invoke SetLoadBalancerHTTPListenerAttribute API
func CreateSetLoadBalancerHTTPListenerAttributeRequest() (request *SetLoadBalancerHTTPListenerAttributeRequest) {
	request = &SetLoadBalancerHTTPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerHTTPListenerAttribute", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetLoadBalancerHTTPListenerAttributeResponse creates a response to parse from SetLoadBalancerHTTPListenerAttribute response
func CreateSetLoadBalancerHTTPListenerAttributeResponse() (response *SetLoadBalancerHTTPListenerAttributeResponse) {
	response = &SetLoadBalancerHTTPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
