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

// ListPrivacyRuleSharedMember invokes the lto.ListPrivacyRuleSharedMember API synchronously
func (client *Client) ListPrivacyRuleSharedMember(request *ListPrivacyRuleSharedMemberRequest) (response *ListPrivacyRuleSharedMemberResponse, err error) {
	response = CreateListPrivacyRuleSharedMemberResponse()
	err = client.DoAction(request, response)
	return
}

// ListPrivacyRuleSharedMemberWithChan invokes the lto.ListPrivacyRuleSharedMember API asynchronously
func (client *Client) ListPrivacyRuleSharedMemberWithChan(request *ListPrivacyRuleSharedMemberRequest) (<-chan *ListPrivacyRuleSharedMemberResponse, <-chan error) {
	responseChan := make(chan *ListPrivacyRuleSharedMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPrivacyRuleSharedMember(request)
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

// ListPrivacyRuleSharedMemberWithCallback invokes the lto.ListPrivacyRuleSharedMember API asynchronously
func (client *Client) ListPrivacyRuleSharedMemberWithCallback(request *ListPrivacyRuleSharedMemberRequest, callback func(response *ListPrivacyRuleSharedMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPrivacyRuleSharedMemberResponse
		var err error
		defer close(result)
		response, err = client.ListPrivacyRuleSharedMember(request)
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

// ListPrivacyRuleSharedMemberRequest is the request struct for api ListPrivacyRuleSharedMember
type ListPrivacyRuleSharedMemberRequest struct {
	*requests.RpcRequest
	PrivacyRuleId string `position:"Query" name:"PrivacyRuleId"`
}

// ListPrivacyRuleSharedMemberResponse is the response struct for api ListPrivacyRuleSharedMember
type ListPrivacyRuleSharedMemberResponse struct {
	*responses.BaseResponse
	Code           string         `json:"Code" xml:"Code"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string         `json:"Message" xml:"Message"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Data           []BizChainInfo `json:"Data" xml:"Data"`
}

// CreateListPrivacyRuleSharedMemberRequest creates a request to invoke ListPrivacyRuleSharedMember API
func CreateListPrivacyRuleSharedMemberRequest() (request *ListPrivacyRuleSharedMemberRequest) {
	request = &ListPrivacyRuleSharedMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("lto", "2021-07-07", "ListPrivacyRuleSharedMember", "", "")
	request.Method = requests.POST
	return
}

// CreateListPrivacyRuleSharedMemberResponse creates a response to parse from ListPrivacyRuleSharedMember response
func CreateListPrivacyRuleSharedMemberResponse() (response *ListPrivacyRuleSharedMemberResponse) {
	response = &ListPrivacyRuleSharedMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
