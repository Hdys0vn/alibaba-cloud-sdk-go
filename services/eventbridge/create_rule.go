package eventbridge

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

// CreateRule invokes the eventbridge.CreateRule API synchronously
func (client *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
	response = CreateCreateRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRuleWithChan invokes the eventbridge.CreateRule API asynchronously
func (client *Client) CreateRuleWithChan(request *CreateRuleRequest) (<-chan *CreateRuleResponse, <-chan error) {
	responseChan := make(chan *CreateRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRule(request)
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

// CreateRuleWithCallback invokes the eventbridge.CreateRule API asynchronously
func (client *Client) CreateRuleWithCallback(request *CreateRuleRequest, callback func(response *CreateRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateRule(request)
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

// CreateRuleRequest is the request struct for api CreateRule
type CreateRuleRequest struct {
	*requests.RpcRequest
	ClientToken   string                    `position:"Query" name:"ClientToken"`
	Description   string                    `position:"Query" name:"Description"`
	RuleName      string                    `position:"Query" name:"RuleName"`
	EventBusName  string                    `position:"Query" name:"EventBusName"`
	Targets       string                    `position:"Query" name:"Targets"`
	EventTargets  *[]CreateRuleEventTargets `position:"Query" name:"EventTargets"  type:"Json"`
	FilterPattern string                    `position:"Query" name:"FilterPattern"`
	Status        string                    `position:"Query" name:"Status"`
}

// CreateRuleEventTargets is a repeated param struct in CreateRuleRequest
type CreateRuleEventTargets struct {
	Endpoint          string                                 `name:"Endpoint"`
	PushRetryStrategy string                                 `name:"PushRetryStrategy"`
	ParamList         *[]CreateRuleEventTargetsParamListItem `name:"ParamList" type:"Repeated"`
	Id                string                                 `name:"Id"`
	DeadLetterQueue   CreateRuleEventTargetsDeadLetterQueue  `name:"DeadLetterQueue" type:"Struct"`
	Type              string                                 `name:"Type"`
	ErrorsTolerance   string                                 `name:"ErrorsTolerance"`
}

// CreateRuleEventTargetsParamListItem is a repeated param struct in CreateRuleRequest
type CreateRuleEventTargetsParamListItem struct {
	Template    string `name:"Template"`
	ResourceKey string `name:"ResourceKey"`
	Form        string `name:"Form"`
	Value       string `name:"Value"`
}

// CreateRuleEventTargetsDeadLetterQueue is a repeated param struct in CreateRuleRequest
type CreateRuleEventTargetsDeadLetterQueue struct {
	Arn string `name:"Arn"`
}

// CreateRuleResponse is the response struct for api CreateRule
type CreateRuleResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateRuleRequest creates a request to invoke CreateRule API
func CreateCreateRuleRequest() (request *CreateRuleRequest) {
	request = &CreateRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eventbridge", "2020-04-01", "CreateRule", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateRuleResponse creates a response to parse from CreateRule response
func CreateCreateRuleResponse() (response *CreateRuleResponse) {
	response = &CreateRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
