package ess

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

// DetachLoadBalancers invokes the ess.DetachLoadBalancers API synchronously
func (client *Client) DetachLoadBalancers(request *DetachLoadBalancersRequest) (response *DetachLoadBalancersResponse, err error) {
	response = CreateDetachLoadBalancersResponse()
	err = client.DoAction(request, response)
	return
}

// DetachLoadBalancersWithChan invokes the ess.DetachLoadBalancers API asynchronously
func (client *Client) DetachLoadBalancersWithChan(request *DetachLoadBalancersRequest) (<-chan *DetachLoadBalancersResponse, <-chan error) {
	responseChan := make(chan *DetachLoadBalancersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachLoadBalancers(request)
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

// DetachLoadBalancersWithCallback invokes the ess.DetachLoadBalancers API asynchronously
func (client *Client) DetachLoadBalancersWithCallback(request *DetachLoadBalancersRequest, callback func(response *DetachLoadBalancersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachLoadBalancersResponse
		var err error
		defer close(result)
		response, err = client.DetachLoadBalancers(request)
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

// DetachLoadBalancersRequest is the request struct for api DetachLoadBalancers
type DetachLoadBalancersRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	LoadBalancer         *[]string        `position:"Query" name:"LoadBalancer"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Async                requests.Boolean `position:"Query" name:"Async"`
	ForceDetach          requests.Boolean `position:"Query" name:"ForceDetach"`
}

// DetachLoadBalancersResponse is the response struct for api DetachLoadBalancers
type DetachLoadBalancersResponse struct {
	*responses.BaseResponse
	ScalingActivityId string `json:"ScalingActivityId" xml:"ScalingActivityId"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachLoadBalancersRequest creates a request to invoke DetachLoadBalancers API
func CreateDetachLoadBalancersRequest() (request *DetachLoadBalancersRequest) {
	request = &DetachLoadBalancersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DetachLoadBalancers", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachLoadBalancersResponse creates a response to parse from DetachLoadBalancers response
func CreateDetachLoadBalancersResponse() (response *DetachLoadBalancersResponse) {
	response = &DetachLoadBalancersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
