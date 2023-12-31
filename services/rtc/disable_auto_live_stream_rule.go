package rtc

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

// DisableAutoLiveStreamRule invokes the rtc.DisableAutoLiveStreamRule API synchronously
func (client *Client) DisableAutoLiveStreamRule(request *DisableAutoLiveStreamRuleRequest) (response *DisableAutoLiveStreamRuleResponse, err error) {
	response = CreateDisableAutoLiveStreamRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DisableAutoLiveStreamRuleWithChan invokes the rtc.DisableAutoLiveStreamRule API asynchronously
func (client *Client) DisableAutoLiveStreamRuleWithChan(request *DisableAutoLiveStreamRuleRequest) (<-chan *DisableAutoLiveStreamRuleResponse, <-chan error) {
	responseChan := make(chan *DisableAutoLiveStreamRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableAutoLiveStreamRule(request)
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

// DisableAutoLiveStreamRuleWithCallback invokes the rtc.DisableAutoLiveStreamRule API asynchronously
func (client *Client) DisableAutoLiveStreamRuleWithCallback(request *DisableAutoLiveStreamRuleRequest, callback func(response *DisableAutoLiveStreamRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableAutoLiveStreamRuleResponse
		var err error
		defer close(result)
		response, err = client.DisableAutoLiveStreamRule(request)
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

// DisableAutoLiveStreamRuleRequest is the request struct for api DisableAutoLiveStreamRule
type DisableAutoLiveStreamRuleRequest struct {
	*requests.RpcRequest
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
	RuleId  requests.Integer `position:"Query" name:"RuleId"`
}

// DisableAutoLiveStreamRuleResponse is the response struct for api DisableAutoLiveStreamRule
type DisableAutoLiveStreamRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableAutoLiveStreamRuleRequest creates a request to invoke DisableAutoLiveStreamRule API
func CreateDisableAutoLiveStreamRuleRequest() (request *DisableAutoLiveStreamRuleRequest) {
	request = &DisableAutoLiveStreamRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DisableAutoLiveStreamRule", "", "")
	request.Method = requests.POST
	return
}

// CreateDisableAutoLiveStreamRuleResponse creates a response to parse from DisableAutoLiveStreamRule response
func CreateDisableAutoLiveStreamRuleResponse() (response *DisableAutoLiveStreamRuleResponse) {
	response = &DisableAutoLiveStreamRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
