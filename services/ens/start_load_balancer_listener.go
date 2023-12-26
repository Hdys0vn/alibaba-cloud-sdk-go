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

// StartLoadBalancerListener invokes the ens.StartLoadBalancerListener API synchronously
func (client *Client) StartLoadBalancerListener(request *StartLoadBalancerListenerRequest) (response *StartLoadBalancerListenerResponse, err error) {
	response = CreateStartLoadBalancerListenerResponse()
	err = client.DoAction(request, response)
	return
}

// StartLoadBalancerListenerWithChan invokes the ens.StartLoadBalancerListener API asynchronously
func (client *Client) StartLoadBalancerListenerWithChan(request *StartLoadBalancerListenerRequest) (<-chan *StartLoadBalancerListenerResponse, <-chan error) {
	responseChan := make(chan *StartLoadBalancerListenerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartLoadBalancerListener(request)
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

// StartLoadBalancerListenerWithCallback invokes the ens.StartLoadBalancerListener API asynchronously
func (client *Client) StartLoadBalancerListenerWithCallback(request *StartLoadBalancerListenerRequest, callback func(response *StartLoadBalancerListenerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartLoadBalancerListenerResponse
		var err error
		defer close(result)
		response, err = client.StartLoadBalancerListener(request)
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

// StartLoadBalancerListenerRequest is the request struct for api StartLoadBalancerListener
type StartLoadBalancerListenerRequest struct {
	*requests.RpcRequest
	ListenerPort     requests.Integer `position:"Query" name:"ListenerPort"`
	ListenerProtocol string           `position:"Query" name:"ListenerProtocol"`
	LoadBalancerId   string           `position:"Query" name:"LoadBalancerId"`
}

// StartLoadBalancerListenerResponse is the response struct for api StartLoadBalancerListener
type StartLoadBalancerListenerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartLoadBalancerListenerRequest creates a request to invoke StartLoadBalancerListener API
func CreateStartLoadBalancerListenerRequest() (request *StartLoadBalancerListenerRequest) {
	request = &StartLoadBalancerListenerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "StartLoadBalancerListener", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartLoadBalancerListenerResponse creates a response to parse from StartLoadBalancerListener response
func CreateStartLoadBalancerListenerResponse() (response *StartLoadBalancerListenerResponse) {
	response = &StartLoadBalancerListenerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
