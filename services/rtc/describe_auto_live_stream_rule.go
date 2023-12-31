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

// DescribeAutoLiveStreamRule invokes the rtc.DescribeAutoLiveStreamRule API synchronously
func (client *Client) DescribeAutoLiveStreamRule(request *DescribeAutoLiveStreamRuleRequest) (response *DescribeAutoLiveStreamRuleResponse, err error) {
	response = CreateDescribeAutoLiveStreamRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAutoLiveStreamRuleWithChan invokes the rtc.DescribeAutoLiveStreamRule API asynchronously
func (client *Client) DescribeAutoLiveStreamRuleWithChan(request *DescribeAutoLiveStreamRuleRequest) (<-chan *DescribeAutoLiveStreamRuleResponse, <-chan error) {
	responseChan := make(chan *DescribeAutoLiveStreamRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAutoLiveStreamRule(request)
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

// DescribeAutoLiveStreamRuleWithCallback invokes the rtc.DescribeAutoLiveStreamRule API asynchronously
func (client *Client) DescribeAutoLiveStreamRuleWithCallback(request *DescribeAutoLiveStreamRuleRequest, callback func(response *DescribeAutoLiveStreamRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAutoLiveStreamRuleResponse
		var err error
		defer close(result)
		response, err = client.DescribeAutoLiveStreamRule(request)
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

// DescribeAutoLiveStreamRuleRequest is the request struct for api DescribeAutoLiveStreamRule
type DescribeAutoLiveStreamRuleRequest struct {
	*requests.RpcRequest
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
}

// DescribeAutoLiveStreamRuleResponse is the response struct for api DescribeAutoLiveStreamRule
type DescribeAutoLiveStreamRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Rules     []Rule `json:"Rules" xml:"Rules"`
}

// CreateDescribeAutoLiveStreamRuleRequest creates a request to invoke DescribeAutoLiveStreamRule API
func CreateDescribeAutoLiveStreamRuleRequest() (request *DescribeAutoLiveStreamRuleRequest) {
	request = &DescribeAutoLiveStreamRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeAutoLiveStreamRule", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeAutoLiveStreamRuleResponse creates a response to parse from DescribeAutoLiveStreamRule response
func CreateDescribeAutoLiveStreamRuleResponse() (response *DescribeAutoLiveStreamRuleResponse) {
	response = &DescribeAutoLiveStreamRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
