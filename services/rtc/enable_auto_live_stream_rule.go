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

// EnableAutoLiveStreamRule invokes the rtc.EnableAutoLiveStreamRule API synchronously
func (client *Client) EnableAutoLiveStreamRule(request *EnableAutoLiveStreamRuleRequest) (response *EnableAutoLiveStreamRuleResponse, err error) {
	response = CreateEnableAutoLiveStreamRuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableAutoLiveStreamRuleWithChan invokes the rtc.EnableAutoLiveStreamRule API asynchronously
func (client *Client) EnableAutoLiveStreamRuleWithChan(request *EnableAutoLiveStreamRuleRequest) (<-chan *EnableAutoLiveStreamRuleResponse, <-chan error) {
	responseChan := make(chan *EnableAutoLiveStreamRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableAutoLiveStreamRule(request)
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

// EnableAutoLiveStreamRuleWithCallback invokes the rtc.EnableAutoLiveStreamRule API asynchronously
func (client *Client) EnableAutoLiveStreamRuleWithCallback(request *EnableAutoLiveStreamRuleRequest, callback func(response *EnableAutoLiveStreamRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableAutoLiveStreamRuleResponse
		var err error
		defer close(result)
		response, err = client.EnableAutoLiveStreamRule(request)
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

// EnableAutoLiveStreamRuleRequest is the request struct for api EnableAutoLiveStreamRule
type EnableAutoLiveStreamRuleRequest struct {
	*requests.RpcRequest
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
	RuleId  requests.Integer `position:"Query" name:"RuleId"`
}

// EnableAutoLiveStreamRuleResponse is the response struct for api EnableAutoLiveStreamRule
type EnableAutoLiveStreamRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableAutoLiveStreamRuleRequest creates a request to invoke EnableAutoLiveStreamRule API
func CreateEnableAutoLiveStreamRuleRequest() (request *EnableAutoLiveStreamRuleRequest) {
	request = &EnableAutoLiveStreamRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "EnableAutoLiveStreamRule", "", "")
	request.Method = requests.POST
	return
}

// CreateEnableAutoLiveStreamRuleResponse creates a response to parse from EnableAutoLiveStreamRule response
func CreateEnableAutoLiveStreamRuleResponse() (response *EnableAutoLiveStreamRuleResponse) {
	response = &EnableAutoLiveStreamRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
