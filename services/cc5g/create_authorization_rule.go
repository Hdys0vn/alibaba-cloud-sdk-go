package cc5g

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

// CreateAuthorizationRule invokes the cc5g.CreateAuthorizationRule API synchronously
func (client *Client) CreateAuthorizationRule(request *CreateAuthorizationRuleRequest) (response *CreateAuthorizationRuleResponse, err error) {
	response = CreateCreateAuthorizationRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAuthorizationRuleWithChan invokes the cc5g.CreateAuthorizationRule API asynchronously
func (client *Client) CreateAuthorizationRuleWithChan(request *CreateAuthorizationRuleRequest) (<-chan *CreateAuthorizationRuleResponse, <-chan error) {
	responseChan := make(chan *CreateAuthorizationRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAuthorizationRule(request)
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

// CreateAuthorizationRuleWithCallback invokes the cc5g.CreateAuthorizationRule API asynchronously
func (client *Client) CreateAuthorizationRuleWithCallback(request *CreateAuthorizationRuleRequest, callback func(response *CreateAuthorizationRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAuthorizationRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateAuthorizationRule(request)
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

// CreateAuthorizationRuleRequest is the request struct for api CreateAuthorizationRule
type CreateAuthorizationRuleRequest struct {
	*requests.RpcRequest
	ClientToken              string           `position:"Query" name:"ClientToken"`
	SourceCidr               string           `position:"Query" name:"SourceCidr"`
	DestinationType          string           `position:"Query" name:"DestinationType"`
	Destination              string           `position:"Query" name:"Destination"`
	Description              string           `position:"Query" name:"Description"`
	Protocol                 string           `position:"Query" name:"Protocol"`
	Policy                   string           `position:"Query" name:"Policy"`
	DryRun                   requests.Boolean `position:"Query" name:"DryRun"`
	DestinationPort          string           `position:"Query" name:"DestinationPort"`
	WirelessCloudConnectorId string           `position:"Query" name:"WirelessCloudConnectorId"`
	Name                     string           `position:"Query" name:"Name"`
}

// CreateAuthorizationRuleResponse is the response struct for api CreateAuthorizationRule
type CreateAuthorizationRuleResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	AuthorizationRuleId string `json:"AuthorizationRuleId" xml:"AuthorizationRuleId"`
}

// CreateCreateAuthorizationRuleRequest creates a request to invoke CreateAuthorizationRule API
func CreateCreateAuthorizationRuleRequest() (request *CreateAuthorizationRuleRequest) {
	request = &CreateAuthorizationRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "CreateAuthorizationRule", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAuthorizationRuleResponse creates a response to parse from CreateAuthorizationRule response
func CreateCreateAuthorizationRuleResponse() (response *CreateAuthorizationRuleResponse) {
	response = &CreateAuthorizationRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
