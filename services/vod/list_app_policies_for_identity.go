package vod

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

// ListAppPoliciesForIdentity invokes the vod.ListAppPoliciesForIdentity API synchronously
func (client *Client) ListAppPoliciesForIdentity(request *ListAppPoliciesForIdentityRequest) (response *ListAppPoliciesForIdentityResponse, err error) {
	response = CreateListAppPoliciesForIdentityResponse()
	err = client.DoAction(request, response)
	return
}

// ListAppPoliciesForIdentityWithChan invokes the vod.ListAppPoliciesForIdentity API asynchronously
func (client *Client) ListAppPoliciesForIdentityWithChan(request *ListAppPoliciesForIdentityRequest) (<-chan *ListAppPoliciesForIdentityResponse, <-chan error) {
	responseChan := make(chan *ListAppPoliciesForIdentityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAppPoliciesForIdentity(request)
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

// ListAppPoliciesForIdentityWithCallback invokes the vod.ListAppPoliciesForIdentity API asynchronously
func (client *Client) ListAppPoliciesForIdentityWithCallback(request *ListAppPoliciesForIdentityRequest, callback func(response *ListAppPoliciesForIdentityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAppPoliciesForIdentityResponse
		var err error
		defer close(result)
		response, err = client.ListAppPoliciesForIdentity(request)
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

// ListAppPoliciesForIdentityRequest is the request struct for api ListAppPoliciesForIdentity
type ListAppPoliciesForIdentityRequest struct {
	*requests.RpcRequest
	IdentityType string `position:"Query" name:"IdentityType"`
	IdentityName string `position:"Query" name:"IdentityName"`
	AppId        string `position:"Query" name:"AppId"`
}

// ListAppPoliciesForIdentityResponse is the response struct for api ListAppPoliciesForIdentity
type ListAppPoliciesForIdentityResponse struct {
	*responses.BaseResponse
	RequestId     string      `json:"RequestId" xml:"RequestId"`
	AppPolicyList []AppPolicy `json:"AppPolicyList" xml:"AppPolicyList"`
}

// CreateListAppPoliciesForIdentityRequest creates a request to invoke ListAppPoliciesForIdentity API
func CreateListAppPoliciesForIdentityRequest() (request *ListAppPoliciesForIdentityRequest) {
	request = &ListAppPoliciesForIdentityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "ListAppPoliciesForIdentity", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAppPoliciesForIdentityResponse creates a response to parse from ListAppPoliciesForIdentity response
func CreateListAppPoliciesForIdentityResponse() (response *ListAppPoliciesForIdentityResponse) {
	response = &ListAppPoliciesForIdentityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}