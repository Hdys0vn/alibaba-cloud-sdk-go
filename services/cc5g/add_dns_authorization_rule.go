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

// AddDNSAuthorizationRule invokes the cc5g.AddDNSAuthorizationRule API synchronously
func (client *Client) AddDNSAuthorizationRule(request *AddDNSAuthorizationRuleRequest) (response *AddDNSAuthorizationRuleResponse, err error) {
	response = CreateAddDNSAuthorizationRuleResponse()
	err = client.DoAction(request, response)
	return
}

// AddDNSAuthorizationRuleWithChan invokes the cc5g.AddDNSAuthorizationRule API asynchronously
func (client *Client) AddDNSAuthorizationRuleWithChan(request *AddDNSAuthorizationRuleRequest) (<-chan *AddDNSAuthorizationRuleResponse, <-chan error) {
	responseChan := make(chan *AddDNSAuthorizationRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDNSAuthorizationRule(request)
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

// AddDNSAuthorizationRuleWithCallback invokes the cc5g.AddDNSAuthorizationRule API asynchronously
func (client *Client) AddDNSAuthorizationRuleWithCallback(request *AddDNSAuthorizationRuleRequest, callback func(response *AddDNSAuthorizationRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDNSAuthorizationRuleResponse
		var err error
		defer close(result)
		response, err = client.AddDNSAuthorizationRule(request)
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

// AddDNSAuthorizationRuleRequest is the request struct for api AddDNSAuthorizationRule
type AddDNSAuthorizationRuleRequest struct {
	*requests.RpcRequest
	ClientToken              string           `position:"Query" name:"ClientToken"`
	Description              string           `position:"Query" name:"Description"`
	DryRun                   requests.Boolean `position:"Query" name:"DryRun"`
	SourceDNSIp              string           `position:"Query" name:"SourceDNSIp"`
	DestinationIp            string           `position:"Query" name:"DestinationIp"`
	WirelessCloudConnectorId string           `position:"Query" name:"WirelessCloudConnectorId"`
	Name                     string           `position:"Query" name:"Name"`
}

// AddDNSAuthorizationRuleResponse is the response struct for api AddDNSAuthorizationRule
type AddDNSAuthorizationRuleResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	AuthorizationRuleId string `json:"AuthorizationRuleId" xml:"AuthorizationRuleId"`
}

// CreateAddDNSAuthorizationRuleRequest creates a request to invoke AddDNSAuthorizationRule API
func CreateAddDNSAuthorizationRuleRequest() (request *AddDNSAuthorizationRuleRequest) {
	request = &AddDNSAuthorizationRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "AddDNSAuthorizationRule", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddDNSAuthorizationRuleResponse creates a response to parse from AddDNSAuthorizationRule response
func CreateAddDNSAuthorizationRuleResponse() (response *AddDNSAuthorizationRuleResponse) {
	response = &AddDNSAuthorizationRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}