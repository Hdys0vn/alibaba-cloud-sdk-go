package dcdn

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

// DescribeDcdnWafPolicies invokes the dcdn.DescribeDcdnWafPolicies API synchronously
func (client *Client) DescribeDcdnWafPolicies(request *DescribeDcdnWafPoliciesRequest) (response *DescribeDcdnWafPoliciesResponse, err error) {
	response = CreateDescribeDcdnWafPoliciesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnWafPoliciesWithChan invokes the dcdn.DescribeDcdnWafPolicies API asynchronously
func (client *Client) DescribeDcdnWafPoliciesWithChan(request *DescribeDcdnWafPoliciesRequest) (<-chan *DescribeDcdnWafPoliciesResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnWafPoliciesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnWafPolicies(request)
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

// DescribeDcdnWafPoliciesWithCallback invokes the dcdn.DescribeDcdnWafPolicies API asynchronously
func (client *Client) DescribeDcdnWafPoliciesWithCallback(request *DescribeDcdnWafPoliciesRequest, callback func(response *DescribeDcdnWafPoliciesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnWafPoliciesResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnWafPolicies(request)
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

// DescribeDcdnWafPoliciesRequest is the request struct for api DescribeDcdnWafPolicies
type DescribeDcdnWafPoliciesRequest struct {
	*requests.RpcRequest
	QueryArgs  string           `position:"Query" name:"QueryArgs"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeDcdnWafPoliciesResponse is the response struct for api DescribeDcdnWafPolicies
type DescribeDcdnWafPoliciesResponse struct {
	*responses.BaseResponse
	PageSize   int          `json:"PageSize" xml:"PageSize"`
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	PageNumber int          `json:"PageNumber" xml:"PageNumber"`
	TotalCount int          `json:"TotalCount" xml:"TotalCount"`
	Policies   []PolicyItem `json:"Policies" xml:"Policies"`
}

// CreateDescribeDcdnWafPoliciesRequest creates a request to invoke DescribeDcdnWafPolicies API
func CreateDescribeDcdnWafPoliciesRequest() (request *DescribeDcdnWafPoliciesRequest) {
	request = &DescribeDcdnWafPoliciesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnWafPolicies", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnWafPoliciesResponse creates a response to parse from DescribeDcdnWafPolicies response
func CreateDescribeDcdnWafPoliciesResponse() (response *DescribeDcdnWafPoliciesResponse) {
	response = &DescribeDcdnWafPoliciesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
