package cms

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

// PutCustomEventRule invokes the cms.PutCustomEventRule API synchronously
func (client *Client) PutCustomEventRule(request *PutCustomEventRuleRequest) (response *PutCustomEventRuleResponse, err error) {
	response = CreatePutCustomEventRuleResponse()
	err = client.DoAction(request, response)
	return
}

// PutCustomEventRuleWithChan invokes the cms.PutCustomEventRule API asynchronously
func (client *Client) PutCustomEventRuleWithChan(request *PutCustomEventRuleRequest) (<-chan *PutCustomEventRuleResponse, <-chan error) {
	responseChan := make(chan *PutCustomEventRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutCustomEventRule(request)
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

// PutCustomEventRuleWithCallback invokes the cms.PutCustomEventRule API asynchronously
func (client *Client) PutCustomEventRuleWithCallback(request *PutCustomEventRuleRequest, callback func(response *PutCustomEventRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutCustomEventRuleResponse
		var err error
		defer close(result)
		response, err = client.PutCustomEventRule(request)
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

// PutCustomEventRuleRequest is the request struct for api PutCustomEventRule
type PutCustomEventRuleRequest struct {
	*requests.RpcRequest
	Webhook           string `position:"Query" name:"Webhook"`
	RuleName          string `position:"Query" name:"RuleName"`
	Threshold         string `position:"Query" name:"Threshold"`
	EffectiveInterval string `position:"Query" name:"EffectiveInterval"`
	EventName         string `position:"Query" name:"EventName"`
	EmailSubject      string `position:"Query" name:"EmailSubject"`
	MetricName        string `position:"Query" name:"MetricName"`
	Period            string `position:"Query" name:"Period"`
	ContactGroups     string `position:"Query" name:"ContactGroups"`
	Level             string `position:"Query" name:"Level"`
	GroupId           string `position:"Query" name:"GroupId"`
	RuleId            string `position:"Query" name:"RuleId"`
}

// PutCustomEventRuleResponse is the response struct for api PutCustomEventRule
type PutCustomEventRuleResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreatePutCustomEventRuleRequest creates a request to invoke PutCustomEventRule API
func CreatePutCustomEventRuleRequest() (request *PutCustomEventRuleRequest) {
	request = &PutCustomEventRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "PutCustomEventRule", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePutCustomEventRuleResponse creates a response to parse from PutCustomEventRule response
func CreatePutCustomEventRuleResponse() (response *PutCustomEventRuleResponse) {
	response = &PutCustomEventRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
