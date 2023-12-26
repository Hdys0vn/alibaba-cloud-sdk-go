package resourcemanager

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

// ListControlPolicies invokes the resourcemanager.ListControlPolicies API synchronously
func (client *Client) ListControlPolicies(request *ListControlPoliciesRequest) (response *ListControlPoliciesResponse, err error) {
	response = CreateListControlPoliciesResponse()
	err = client.DoAction(request, response)
	return
}

// ListControlPoliciesWithChan invokes the resourcemanager.ListControlPolicies API asynchronously
func (client *Client) ListControlPoliciesWithChan(request *ListControlPoliciesRequest) (<-chan *ListControlPoliciesResponse, <-chan error) {
	responseChan := make(chan *ListControlPoliciesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListControlPolicies(request)
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

// ListControlPoliciesWithCallback invokes the resourcemanager.ListControlPolicies API asynchronously
func (client *Client) ListControlPoliciesWithCallback(request *ListControlPoliciesRequest, callback func(response *ListControlPoliciesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListControlPoliciesResponse
		var err error
		defer close(result)
		response, err = client.ListControlPolicies(request)
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

// ListControlPoliciesRequest is the request struct for api ListControlPolicies
type ListControlPoliciesRequest struct {
	*requests.RpcRequest
	PolicyType string           `position:"Query" name:"PolicyType"`
	Language   string           `position:"Query" name:"Language"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListControlPoliciesResponse is the response struct for api ListControlPolicies
type ListControlPoliciesResponse struct {
	*responses.BaseResponse
	TotalCount      int             `json:"TotalCount" xml:"TotalCount"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	PageSize        int             `json:"PageSize" xml:"PageSize"`
	PageNumber      int             `json:"PageNumber" xml:"PageNumber"`
	ControlPolicies ControlPolicies `json:"ControlPolicies" xml:"ControlPolicies"`
}

// CreateListControlPoliciesRequest creates a request to invoke ListControlPolicies API
func CreateListControlPoliciesRequest() (request *ListControlPoliciesRequest) {
	request = &ListControlPoliciesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "ListControlPolicies", "", "")
	request.Method = requests.POST
	return
}

// CreateListControlPoliciesResponse creates a response to parse from ListControlPolicies response
func CreateListControlPoliciesResponse() (response *ListControlPoliciesResponse) {
	response = &ListControlPoliciesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
