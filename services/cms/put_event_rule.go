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

// PutEventRule invokes the cms.PutEventRule API synchronously
func (client *Client) PutEventRule(request *PutEventRuleRequest) (response *PutEventRuleResponse, err error) {
	response = CreatePutEventRuleResponse()
	err = client.DoAction(request, response)
	return
}

// PutEventRuleWithChan invokes the cms.PutEventRule API asynchronously
func (client *Client) PutEventRuleWithChan(request *PutEventRuleRequest) (<-chan *PutEventRuleResponse, <-chan error) {
	responseChan := make(chan *PutEventRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutEventRule(request)
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

// PutEventRuleWithCallback invokes the cms.PutEventRule API asynchronously
func (client *Client) PutEventRuleWithCallback(request *PutEventRuleRequest, callback func(response *PutEventRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutEventRuleResponse
		var err error
		defer close(result)
		response, err = client.PutEventRule(request)
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

// PutEventRuleRequest is the request struct for api PutEventRule
type PutEventRuleRequest struct {
	*requests.RpcRequest
	Description  string                      `position:"Query" name:"Description"`
	RuleName     string                      `position:"Query" name:"RuleName"`
	SilenceTime  requests.Integer            `position:"Query" name:"SilenceTime"`
	State        string                      `position:"Query" name:"State"`
	GroupId      string                      `position:"Query" name:"GroupId"`
	EventPattern *[]PutEventRuleEventPattern `position:"Query" name:"EventPattern"  type:"Repeated"`
	EventType    string                      `position:"Query" name:"EventType"`
}

// PutEventRuleEventPattern is a repeated param struct in PutEventRuleRequest
type PutEventRuleEventPattern struct {
	LevelList     *[]string                             `name:"LevelList" type:"Repeated"`
	KeywordFilter PutEventRuleEventPatternKeywordFilter `name:"KeywordFilter" type:"Struct"`
	Product       string                                `name:"Product"`
	StatusList    *[]string                             `name:"StatusList" type:"Repeated"`
	NameList      *[]string                             `name:"NameList" type:"Repeated"`
	CustomFilters string                                `name:"CustomFilters"`
	EventTypeList *[]string                             `name:"EventTypeList" type:"Repeated"`
	SQLFilter     string                                `name:"SQLFilter"`
}

// PutEventRuleEventPatternKeywordFilter is a repeated param struct in PutEventRuleRequest
type PutEventRuleEventPatternKeywordFilter struct {
	Keywords *[]string `name:"Keywords" type:"Repeated"`
	Relation string    `name:"Relation"`
}

// PutEventRuleResponse is the response struct for api PutEventRule
type PutEventRuleResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreatePutEventRuleRequest creates a request to invoke PutEventRule API
func CreatePutEventRuleRequest() (request *PutEventRuleRequest) {
	request = &PutEventRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "PutEventRule", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePutEventRuleResponse creates a response to parse from PutEventRule response
func CreatePutEventRuleResponse() (response *PutEventRuleResponse) {
	response = &PutEventRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
