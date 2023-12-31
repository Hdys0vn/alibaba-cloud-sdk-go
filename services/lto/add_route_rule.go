package lto

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

// AddRouteRule invokes the lto.AddRouteRule API synchronously
func (client *Client) AddRouteRule(request *AddRouteRuleRequest) (response *AddRouteRuleResponse, err error) {
	response = CreateAddRouteRuleResponse()
	err = client.DoAction(request, response)
	return
}

// AddRouteRuleWithChan invokes the lto.AddRouteRule API asynchronously
func (client *Client) AddRouteRuleWithChan(request *AddRouteRuleRequest) (<-chan *AddRouteRuleResponse, <-chan error) {
	responseChan := make(chan *AddRouteRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddRouteRule(request)
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

// AddRouteRuleWithCallback invokes the lto.AddRouteRule API asynchronously
func (client *Client) AddRouteRuleWithCallback(request *AddRouteRuleRequest, callback func(response *AddRouteRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddRouteRuleResponse
		var err error
		defer close(result)
		response, err = client.AddRouteRule(request)
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

// AddRouteRuleRequest is the request struct for api AddRouteRule
type AddRouteRuleRequest struct {
	*requests.RpcRequest
	ContractTemplateId string `position:"Query" name:"ContractTemplateId"`
	ContractName       string `position:"Query" name:"ContractName"`
	PrivacyRuleId      string `position:"Query" name:"PrivacyRuleId"`
	Remark             string `position:"Query" name:"Remark"`
	BizChainId         string `position:"Query" name:"BizChainId"`
	InvokeType         string `position:"Query" name:"InvokeType"`
	DeviceGroupId      string `position:"Query" name:"DeviceGroupId"`
	ChainUpMode        string `position:"Query" name:"ChainUpMode"`
}

// AddRouteRuleResponse is the response struct for api AddRouteRule
type AddRouteRuleResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateAddRouteRuleRequest creates a request to invoke AddRouteRule API
func CreateAddRouteRuleRequest() (request *AddRouteRuleRequest) {
	request = &AddRouteRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "AddRouteRule", "", "")
	request.Method = requests.POST
	return
}

// CreateAddRouteRuleResponse creates a response to parse from AddRouteRule response
func CreateAddRouteRuleResponse() (response *AddRouteRuleResponse) {
	response = &AddRouteRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
