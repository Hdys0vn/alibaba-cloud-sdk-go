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

// DeleteLoadBalancer invokes the nlb.DeleteLoadBalancer API synchronously
func (client *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
	response = CreateDeleteLoadBalancerResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLoadBalancerWithChan invokes the nlb.DeleteLoadBalancer API asynchronously
func (client *Client) DeleteLoadBalancerWithChan(request *DeleteLoadBalancerRequest) (<-chan *DeleteLoadBalancerResponse, <-chan error) {
	responseChan := make(chan *DeleteLoadBalancerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLoadBalancer(request)
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

// DeleteLoadBalancerWithCallback invokes the nlb.DeleteLoadBalancer API asynchronously
func (client *Client) DeleteLoadBalancerWithCallback(request *DeleteLoadBalancerRequest, callback func(response *DeleteLoadBalancerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLoadBalancerResponse
		var err error
		defer close(result)
		response, err = client.DeleteLoadBalancer(request)
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

// DeleteLoadBalancerRequest is the request struct for api DeleteLoadBalancer
type DeleteLoadBalancerRequest struct {
	*requests.RpcRequest
	ClientToken    string           `position:"Body" name:"ClientToken"`
	DryRun         requests.Boolean `position:"Body" name:"DryRun"`
	LoadBalancerId string           `position:"Body" name:"LoadBalancerId"`
}

// DeleteLoadBalancerResponse is the response struct for api DeleteLoadBalancer
type DeleteLoadBalancerResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	JobId          string `json:"JobId" xml:"JobId"`
}

// CreateDeleteLoadBalancerRequest creates a request to invoke DeleteLoadBalancer API
func CreateDeleteLoadBalancerRequest() (request *DeleteLoadBalancerRequest) {
	request = &DeleteLoadBalancerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "DeleteLoadBalancer", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLoadBalancerResponse creates a response to parse from DeleteLoadBalancer response
func CreateDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
	response = &DeleteLoadBalancerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
