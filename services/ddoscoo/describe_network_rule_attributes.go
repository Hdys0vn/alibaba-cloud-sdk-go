package ddoscoo

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

// DescribeNetworkRuleAttributes invokes the ddoscoo.DescribeNetworkRuleAttributes API synchronously
func (client *Client) DescribeNetworkRuleAttributes(request *DescribeNetworkRuleAttributesRequest) (response *DescribeNetworkRuleAttributesResponse, err error) {
	response = CreateDescribeNetworkRuleAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkRuleAttributesWithChan invokes the ddoscoo.DescribeNetworkRuleAttributes API asynchronously
func (client *Client) DescribeNetworkRuleAttributesWithChan(request *DescribeNetworkRuleAttributesRequest) (<-chan *DescribeNetworkRuleAttributesResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkRuleAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkRuleAttributes(request)
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

// DescribeNetworkRuleAttributesWithCallback invokes the ddoscoo.DescribeNetworkRuleAttributes API asynchronously
func (client *Client) DescribeNetworkRuleAttributesWithCallback(request *DescribeNetworkRuleAttributesRequest, callback func(response *DescribeNetworkRuleAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkRuleAttributesResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkRuleAttributes(request)
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

// DescribeNetworkRuleAttributesRequest is the request struct for api DescribeNetworkRuleAttributes
type DescribeNetworkRuleAttributesRequest struct {
	*requests.RpcRequest
	NetworkRules string `position:"Query" name:"NetworkRules"`
	SourceIp     string `position:"Query" name:"SourceIp"`
}

// DescribeNetworkRuleAttributesResponse is the response struct for api DescribeNetworkRuleAttributes
type DescribeNetworkRuleAttributesResponse struct {
	*responses.BaseResponse
	RequestId             string                 `json:"RequestId" xml:"RequestId"`
	NetworkRuleAttributes []NetworkRuleAttribute `json:"NetworkRuleAttributes" xml:"NetworkRuleAttributes"`
}

// CreateDescribeNetworkRuleAttributesRequest creates a request to invoke DescribeNetworkRuleAttributes API
func CreateDescribeNetworkRuleAttributesRequest() (request *DescribeNetworkRuleAttributesRequest) {
	request = &DescribeNetworkRuleAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeNetworkRuleAttributes", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeNetworkRuleAttributesResponse creates a response to parse from DescribeNetworkRuleAttributes response
func CreateDescribeNetworkRuleAttributesResponse() (response *DescribeNetworkRuleAttributesResponse) {
	response = &DescribeNetworkRuleAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
