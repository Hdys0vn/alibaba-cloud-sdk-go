package alb

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

// EnableLoadBalancerAccessLog invokes the alb.EnableLoadBalancerAccessLog API synchronously
func (client *Client) EnableLoadBalancerAccessLog(request *EnableLoadBalancerAccessLogRequest) (response *EnableLoadBalancerAccessLogResponse, err error) {
	response = CreateEnableLoadBalancerAccessLogResponse()
	err = client.DoAction(request, response)
	return
}

// EnableLoadBalancerAccessLogWithChan invokes the alb.EnableLoadBalancerAccessLog API asynchronously
func (client *Client) EnableLoadBalancerAccessLogWithChan(request *EnableLoadBalancerAccessLogRequest) (<-chan *EnableLoadBalancerAccessLogResponse, <-chan error) {
	responseChan := make(chan *EnableLoadBalancerAccessLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableLoadBalancerAccessLog(request)
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

// EnableLoadBalancerAccessLogWithCallback invokes the alb.EnableLoadBalancerAccessLog API asynchronously
func (client *Client) EnableLoadBalancerAccessLogWithCallback(request *EnableLoadBalancerAccessLogRequest, callback func(response *EnableLoadBalancerAccessLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableLoadBalancerAccessLogResponse
		var err error
		defer close(result)
		response, err = client.EnableLoadBalancerAccessLog(request)
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

// EnableLoadBalancerAccessLogRequest is the request struct for api EnableLoadBalancerAccessLog
type EnableLoadBalancerAccessLogRequest struct {
	*requests.RpcRequest
	LogProject     string           `position:"Query" name:"LogProject"`
	ClientToken    string           `position:"Query" name:"ClientToken"`
	DryRun         requests.Boolean `position:"Query" name:"DryRun"`
	LoadBalancerId string           `position:"Query" name:"LoadBalancerId"`
	LogStore       string           `position:"Query" name:"LogStore"`
}

// EnableLoadBalancerAccessLogResponse is the response struct for api EnableLoadBalancerAccessLog
type EnableLoadBalancerAccessLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableLoadBalancerAccessLogRequest creates a request to invoke EnableLoadBalancerAccessLog API
func CreateEnableLoadBalancerAccessLogRequest() (request *EnableLoadBalancerAccessLogRequest) {
	request = &EnableLoadBalancerAccessLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "EnableLoadBalancerAccessLog", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableLoadBalancerAccessLogResponse creates a response to parse from EnableLoadBalancerAccessLog response
func CreateEnableLoadBalancerAccessLogResponse() (response *EnableLoadBalancerAccessLogResponse) {
	response = &EnableLoadBalancerAccessLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}