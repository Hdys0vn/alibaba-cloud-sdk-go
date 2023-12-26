package iot

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

// UpdateSceneRule invokes the iot.UpdateSceneRule API synchronously
func (client *Client) UpdateSceneRule(request *UpdateSceneRuleRequest) (response *UpdateSceneRuleResponse, err error) {
	response = CreateUpdateSceneRuleResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSceneRuleWithChan invokes the iot.UpdateSceneRule API asynchronously
func (client *Client) UpdateSceneRuleWithChan(request *UpdateSceneRuleRequest) (<-chan *UpdateSceneRuleResponse, <-chan error) {
	responseChan := make(chan *UpdateSceneRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSceneRule(request)
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

// UpdateSceneRuleWithCallback invokes the iot.UpdateSceneRule API asynchronously
func (client *Client) UpdateSceneRuleWithCallback(request *UpdateSceneRuleRequest, callback func(response *UpdateSceneRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSceneRuleResponse
		var err error
		defer close(result)
		response, err = client.UpdateSceneRule(request)
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

// UpdateSceneRuleRequest is the request struct for api UpdateSceneRule
type UpdateSceneRuleRequest struct {
	*requests.RpcRequest
	RuleName        string `position:"Query" name:"RuleName"`
	IotInstanceId   string `position:"Query" name:"IotInstanceId"`
	RuleDescription string `position:"Query" name:"RuleDescription"`
	RuleContent     string `position:"Query" name:"RuleContent"`
	ApiProduct      string `position:"Body" name:"ApiProduct"`
	ApiRevision     string `position:"Body" name:"ApiRevision"`
	RuleId          string `position:"Query" name:"RuleId"`
}

// UpdateSceneRuleResponse is the response struct for api UpdateSceneRule
type UpdateSceneRuleResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateSceneRuleRequest creates a request to invoke UpdateSceneRule API
func CreateUpdateSceneRuleRequest() (request *UpdateSceneRuleRequest) {
	request = &UpdateSceneRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateSceneRule", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateSceneRuleResponse creates a response to parse from UpdateSceneRule response
func CreateUpdateSceneRuleResponse() (response *UpdateSceneRuleResponse) {
	response = &UpdateSceneRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
