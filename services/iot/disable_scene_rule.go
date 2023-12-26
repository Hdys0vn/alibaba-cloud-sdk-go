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

// DisableSceneRule invokes the iot.DisableSceneRule API synchronously
func (client *Client) DisableSceneRule(request *DisableSceneRuleRequest) (response *DisableSceneRuleResponse, err error) {
	response = CreateDisableSceneRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DisableSceneRuleWithChan invokes the iot.DisableSceneRule API asynchronously
func (client *Client) DisableSceneRuleWithChan(request *DisableSceneRuleRequest) (<-chan *DisableSceneRuleResponse, <-chan error) {
	responseChan := make(chan *DisableSceneRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableSceneRule(request)
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

// DisableSceneRuleWithCallback invokes the iot.DisableSceneRule API asynchronously
func (client *Client) DisableSceneRuleWithCallback(request *DisableSceneRuleRequest, callback func(response *DisableSceneRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableSceneRuleResponse
		var err error
		defer close(result)
		response, err = client.DisableSceneRule(request)
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

// DisableSceneRuleRequest is the request struct for api DisableSceneRule
type DisableSceneRuleRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	RuleId        string `position:"Query" name:"RuleId"`
}

// DisableSceneRuleResponse is the response struct for api DisableSceneRule
type DisableSceneRuleResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string `json:"Code" xml:"Code"`
}

// CreateDisableSceneRuleRequest creates a request to invoke DisableSceneRule API
func CreateDisableSceneRuleRequest() (request *DisableSceneRuleRequest) {
	request = &DisableSceneRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DisableSceneRule", "", "")
	request.Method = requests.POST
	return
}

// CreateDisableSceneRuleResponse creates a response to parse from DisableSceneRule response
func CreateDisableSceneRuleResponse() (response *DisableSceneRuleResponse) {
	response = &DisableSceneRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
