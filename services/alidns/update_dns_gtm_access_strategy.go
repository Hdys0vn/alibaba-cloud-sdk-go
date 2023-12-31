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

// UpdateDnsGtmAccessStrategy invokes the alidns.UpdateDnsGtmAccessStrategy API synchronously
func (client *Client) UpdateDnsGtmAccessStrategy(request *UpdateDnsGtmAccessStrategyRequest) (response *UpdateDnsGtmAccessStrategyResponse, err error) {
	response = CreateUpdateDnsGtmAccessStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDnsGtmAccessStrategyWithChan invokes the alidns.UpdateDnsGtmAccessStrategy API asynchronously
func (client *Client) UpdateDnsGtmAccessStrategyWithChan(request *UpdateDnsGtmAccessStrategyRequest) (<-chan *UpdateDnsGtmAccessStrategyResponse, <-chan error) {
	responseChan := make(chan *UpdateDnsGtmAccessStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDnsGtmAccessStrategy(request)
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

// UpdateDnsGtmAccessStrategyWithCallback invokes the alidns.UpdateDnsGtmAccessStrategy API asynchronously
func (client *Client) UpdateDnsGtmAccessStrategyWithCallback(request *UpdateDnsGtmAccessStrategyRequest, callback func(response *UpdateDnsGtmAccessStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDnsGtmAccessStrategyResponse
		var err error
		defer close(result)
		response, err = client.UpdateDnsGtmAccessStrategy(request)
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

// UpdateDnsGtmAccessStrategyRequest is the request struct for api UpdateDnsGtmAccessStrategy
type UpdateDnsGtmAccessStrategyRequest struct {
	*requests.RpcRequest
	DefaultLbaStrategy          string                                        `position:"Query" name:"DefaultLbaStrategy"`
	FailoverAddrPoolType        string                                        `position:"Query" name:"FailoverAddrPoolType"`
	DefaultAddrPoolType         string                                        `position:"Query" name:"DefaultAddrPoolType"`
	FailoverMaxReturnAddrNum    requests.Integer                              `position:"Query" name:"FailoverMaxReturnAddrNum"`
	FailoverLbaStrategy         string                                        `position:"Query" name:"FailoverLbaStrategy"`
	DefaultAddrPool             *[]UpdateDnsGtmAccessStrategyDefaultAddrPool  `position:"Query" name:"DefaultAddrPool"  type:"Repeated"`
	FailoverMinAvailableAddrNum requests.Integer                              `position:"Query" name:"FailoverMinAvailableAddrNum"`
	DefaultMaxReturnAddrNum     requests.Integer                              `position:"Query" name:"DefaultMaxReturnAddrNum"`
	DefaultMinAvailableAddrNum  requests.Integer                              `position:"Query" name:"DefaultMinAvailableAddrNum"`
	Lang                        string                                        `position:"Query" name:"Lang"`
	Lines                       string                                        `position:"Query" name:"Lines"`
	StrategyName                string                                        `position:"Query" name:"StrategyName"`
	DefaultLatencyOptimization  string                                        `position:"Query" name:"DefaultLatencyOptimization"`
	FailoverLatencyOptimization string                                        `position:"Query" name:"FailoverLatencyOptimization"`
	UserClientIp                string                                        `position:"Query" name:"UserClientIp"`
	StrategyId                  string                                        `position:"Query" name:"StrategyId"`
	FailoverAddrPool            *[]UpdateDnsGtmAccessStrategyFailoverAddrPool `position:"Query" name:"FailoverAddrPool"  type:"Repeated"`
	AccessMode                  string                                        `position:"Query" name:"AccessMode"`
}

// UpdateDnsGtmAccessStrategyDefaultAddrPool is a repeated param struct in UpdateDnsGtmAccessStrategyRequest
type UpdateDnsGtmAccessStrategyDefaultAddrPool struct {
	Id        string `name:"Id"`
	LbaWeight string `name:"LbaWeight"`
}

// UpdateDnsGtmAccessStrategyFailoverAddrPool is a repeated param struct in UpdateDnsGtmAccessStrategyRequest
type UpdateDnsGtmAccessStrategyFailoverAddrPool struct {
	Id        string `name:"Id"`
	LbaWeight string `name:"LbaWeight"`
}

// UpdateDnsGtmAccessStrategyResponse is the response struct for api UpdateDnsGtmAccessStrategy
type UpdateDnsGtmAccessStrategyResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	StrategyId string `json:"StrategyId" xml:"StrategyId"`
}

// CreateUpdateDnsGtmAccessStrategyRequest creates a request to invoke UpdateDnsGtmAccessStrategy API
func CreateUpdateDnsGtmAccessStrategyRequest() (request *UpdateDnsGtmAccessStrategyRequest) {
	request = &UpdateDnsGtmAccessStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "UpdateDnsGtmAccessStrategy", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateDnsGtmAccessStrategyResponse creates a response to parse from UpdateDnsGtmAccessStrategy response
func CreateUpdateDnsGtmAccessStrategyResponse() (response *UpdateDnsGtmAccessStrategyResponse) {
	response = &UpdateDnsGtmAccessStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
