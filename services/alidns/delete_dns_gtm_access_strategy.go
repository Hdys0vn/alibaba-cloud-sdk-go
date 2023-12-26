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

// DeleteDnsGtmAccessStrategy invokes the alidns.DeleteDnsGtmAccessStrategy API synchronously
func (client *Client) DeleteDnsGtmAccessStrategy(request *DeleteDnsGtmAccessStrategyRequest) (response *DeleteDnsGtmAccessStrategyResponse, err error) {
	response = CreateDeleteDnsGtmAccessStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDnsGtmAccessStrategyWithChan invokes the alidns.DeleteDnsGtmAccessStrategy API asynchronously
func (client *Client) DeleteDnsGtmAccessStrategyWithChan(request *DeleteDnsGtmAccessStrategyRequest) (<-chan *DeleteDnsGtmAccessStrategyResponse, <-chan error) {
	responseChan := make(chan *DeleteDnsGtmAccessStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDnsGtmAccessStrategy(request)
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

// DeleteDnsGtmAccessStrategyWithCallback invokes the alidns.DeleteDnsGtmAccessStrategy API asynchronously
func (client *Client) DeleteDnsGtmAccessStrategyWithCallback(request *DeleteDnsGtmAccessStrategyRequest, callback func(response *DeleteDnsGtmAccessStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDnsGtmAccessStrategyResponse
		var err error
		defer close(result)
		response, err = client.DeleteDnsGtmAccessStrategy(request)
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

// DeleteDnsGtmAccessStrategyRequest is the request struct for api DeleteDnsGtmAccessStrategy
type DeleteDnsGtmAccessStrategyRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	StrategyId   string `position:"Query" name:"StrategyId"`
	Lang         string `position:"Query" name:"Lang"`
}

// DeleteDnsGtmAccessStrategyResponse is the response struct for api DeleteDnsGtmAccessStrategy
type DeleteDnsGtmAccessStrategyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDnsGtmAccessStrategyRequest creates a request to invoke DeleteDnsGtmAccessStrategy API
func CreateDeleteDnsGtmAccessStrategyRequest() (request *DeleteDnsGtmAccessStrategyRequest) {
	request = &DeleteDnsGtmAccessStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DeleteDnsGtmAccessStrategy", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDnsGtmAccessStrategyResponse creates a response to parse from DeleteDnsGtmAccessStrategy response
func CreateDeleteDnsGtmAccessStrategyResponse() (response *DeleteDnsGtmAccessStrategyResponse) {
	response = &DeleteDnsGtmAccessStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
