package cloudapi

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

// SetWildcardDomainPatterns invokes the cloudapi.SetWildcardDomainPatterns API synchronously
func (client *Client) SetWildcardDomainPatterns(request *SetWildcardDomainPatternsRequest) (response *SetWildcardDomainPatternsResponse, err error) {
	response = CreateSetWildcardDomainPatternsResponse()
	err = client.DoAction(request, response)
	return
}

// SetWildcardDomainPatternsWithChan invokes the cloudapi.SetWildcardDomainPatterns API asynchronously
func (client *Client) SetWildcardDomainPatternsWithChan(request *SetWildcardDomainPatternsRequest) (<-chan *SetWildcardDomainPatternsResponse, <-chan error) {
	responseChan := make(chan *SetWildcardDomainPatternsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetWildcardDomainPatterns(request)
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

// SetWildcardDomainPatternsWithCallback invokes the cloudapi.SetWildcardDomainPatterns API asynchronously
func (client *Client) SetWildcardDomainPatternsWithCallback(request *SetWildcardDomainPatternsRequest, callback func(response *SetWildcardDomainPatternsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetWildcardDomainPatternsResponse
		var err error
		defer close(result)
		response, err = client.SetWildcardDomainPatterns(request)
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

// SetWildcardDomainPatternsRequest is the request struct for api SetWildcardDomainPatterns
type SetWildcardDomainPatternsRequest struct {
	*requests.RpcRequest
	GroupId                string `position:"Query" name:"GroupId"`
	DomainName             string `position:"Query" name:"DomainName"`
	SecurityToken          string `position:"Query" name:"SecurityToken"`
	WildcardDomainPatterns string `position:"Query" name:"WildcardDomainPatterns"`
}

// SetWildcardDomainPatternsResponse is the response struct for api SetWildcardDomainPatterns
type SetWildcardDomainPatternsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetWildcardDomainPatternsRequest creates a request to invoke SetWildcardDomainPatterns API
func CreateSetWildcardDomainPatternsRequest() (request *SetWildcardDomainPatternsRequest) {
	request = &SetWildcardDomainPatternsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "SetWildcardDomainPatterns", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetWildcardDomainPatternsResponse creates a response to parse from SetWildcardDomainPatterns response
func CreateSetWildcardDomainPatternsResponse() (response *SetWildcardDomainPatternsResponse) {
	response = &SetWildcardDomainPatternsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}