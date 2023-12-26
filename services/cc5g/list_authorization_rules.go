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

// ListAuthorizationRules invokes the cc5g.ListAuthorizationRules API synchronously
func (client *Client) ListAuthorizationRules(request *ListAuthorizationRulesRequest) (response *ListAuthorizationRulesResponse, err error) {
	response = CreateListAuthorizationRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListAuthorizationRulesWithChan invokes the cc5g.ListAuthorizationRules API asynchronously
func (client *Client) ListAuthorizationRulesWithChan(request *ListAuthorizationRulesRequest) (<-chan *ListAuthorizationRulesResponse, <-chan error) {
	responseChan := make(chan *ListAuthorizationRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAuthorizationRules(request)
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

// ListAuthorizationRulesWithCallback invokes the cc5g.ListAuthorizationRules API asynchronously
func (client *Client) ListAuthorizationRulesWithCallback(request *ListAuthorizationRulesRequest, callback func(response *ListAuthorizationRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAuthorizationRulesResponse
		var err error
		defer close(result)
		response, err = client.ListAuthorizationRules(request)
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

// ListAuthorizationRulesRequest is the request struct for api ListAuthorizationRules
type ListAuthorizationRulesRequest struct {
	*requests.RpcRequest
	DestinationType          string           `position:"Query" name:"DestinationType"`
	Destination              string           `position:"Query" name:"Destination"`
	Type                     string           `position:"Query" name:"Type"`
	Protocol                 string           `position:"Query" name:"Protocol"`
	AuthorizationRuleIds     *[]string        `position:"Query" name:"AuthorizationRuleIds"  type:"Repeated"`
	NextToken                string           `position:"Query" name:"NextToken"`
	Policy                   string           `position:"Query" name:"Policy"`
	Dns                      requests.Boolean `position:"Query" name:"Dns"`
	DestinationPort          string           `position:"Query" name:"DestinationPort"`
	Names                    *[]string        `position:"Query" name:"Names"  type:"Repeated"`
	WirelessCloudConnectorId string           `position:"Query" name:"WirelessCloudConnectorId"`
	MaxResults               requests.Integer `position:"Query" name:"MaxResults"`
	Statuses                 *[]string        `position:"Query" name:"Statuses"  type:"Repeated"`
}

// ListAuthorizationRulesResponse is the response struct for api ListAuthorizationRules
type ListAuthorizationRulesResponse struct {
	*responses.BaseResponse
	RequestId          string              `json:"RequestId" xml:"RequestId"`
	NextToken          string              `json:"NextToken" xml:"NextToken"`
	MaxResults         string              `json:"MaxResults" xml:"MaxResults"`
	TotalCount         string              `json:"TotalCount" xml:"TotalCount"`
	AuthorizationRules []AuthorizationRule `json:"AuthorizationRules" xml:"AuthorizationRules"`
}

// CreateListAuthorizationRulesRequest creates a request to invoke ListAuthorizationRules API
func CreateListAuthorizationRulesRequest() (request *ListAuthorizationRulesRequest) {
	request = &ListAuthorizationRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "ListAuthorizationRules", "fivegcc", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListAuthorizationRulesResponse creates a response to parse from ListAuthorizationRules response
func CreateListAuthorizationRulesResponse() (response *ListAuthorizationRulesResponse) {
	response = &ListAuthorizationRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
