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

// GetRuleAction invokes the iot.GetRuleAction API synchronously
func (client *Client) GetRuleAction(request *GetRuleActionRequest) (response *GetRuleActionResponse, err error) {
	response = CreateGetRuleActionResponse()
	err = client.DoAction(request, response)
	return
}

// GetRuleActionWithChan invokes the iot.GetRuleAction API asynchronously
func (client *Client) GetRuleActionWithChan(request *GetRuleActionRequest) (<-chan *GetRuleActionResponse, <-chan error) {
	responseChan := make(chan *GetRuleActionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRuleAction(request)
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

// GetRuleActionWithCallback invokes the iot.GetRuleAction API asynchronously
func (client *Client) GetRuleActionWithCallback(request *GetRuleActionRequest, callback func(response *GetRuleActionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRuleActionResponse
		var err error
		defer close(result)
		response, err = client.GetRuleAction(request)
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

// GetRuleActionRequest is the request struct for api GetRuleAction
type GetRuleActionRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ActionId      requests.Integer `position:"Query" name:"ActionId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// GetRuleActionResponse is the response struct for api GetRuleAction
type GetRuleActionResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Code           string         `json:"Code" xml:"Code"`
	ErrorMessage   string         `json:"ErrorMessage" xml:"ErrorMessage"`
	RuleActionInfo RuleActionInfo `json:"RuleActionInfo" xml:"RuleActionInfo"`
}

// CreateGetRuleActionRequest creates a request to invoke GetRuleAction API
func CreateGetRuleActionRequest() (request *GetRuleActionRequest) {
	request = &GetRuleActionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetRuleAction", "", "")
	request.Method = requests.POST
	return
}

// CreateGetRuleActionResponse creates a response to parse from GetRuleAction response
func CreateGetRuleActionResponse() (response *GetRuleActionResponse) {
	response = &GetRuleActionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
