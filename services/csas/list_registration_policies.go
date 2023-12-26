package csas

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

// ListRegistrationPolicies invokes the csas.ListRegistrationPolicies API synchronously
func (client *Client) ListRegistrationPolicies(request *ListRegistrationPoliciesRequest) (response *ListRegistrationPoliciesResponse, err error) {
	response = CreateListRegistrationPoliciesResponse()
	err = client.DoAction(request, response)
	return
}

// ListRegistrationPoliciesWithChan invokes the csas.ListRegistrationPolicies API asynchronously
func (client *Client) ListRegistrationPoliciesWithChan(request *ListRegistrationPoliciesRequest) (<-chan *ListRegistrationPoliciesResponse, <-chan error) {
	responseChan := make(chan *ListRegistrationPoliciesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRegistrationPolicies(request)
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

// ListRegistrationPoliciesWithCallback invokes the csas.ListRegistrationPolicies API asynchronously
func (client *Client) ListRegistrationPoliciesWithCallback(request *ListRegistrationPoliciesRequest, callback func(response *ListRegistrationPoliciesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRegistrationPoliciesResponse
		var err error
		defer close(result)
		response, err = client.ListRegistrationPolicies(request)
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

// ListRegistrationPoliciesRequest is the request struct for api ListRegistrationPolicies
type ListRegistrationPoliciesRequest struct {
	*requests.RpcRequest
	PolicyIds         *[]string        `position:"Query" name:"PolicyIds"  type:"Repeated"`
	MatchMode         string           `position:"Query" name:"MatchMode"`
	SourceIp          string           `position:"Query" name:"SourceIp"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	UserGroupId       string           `position:"Query" name:"UserGroupId"`
	CurrentPage       requests.Integer `position:"Query" name:"CurrentPage"`
	PersonalLimitType string           `position:"Query" name:"PersonalLimitType"`
	Name              string           `position:"Query" name:"Name"`
	CompanyLimitType  string           `position:"Query" name:"CompanyLimitType"`
	Status            string           `position:"Query" name:"Status"`
}

// ListRegistrationPoliciesResponse is the response struct for api ListRegistrationPolicies
type ListRegistrationPoliciesResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	TotalNum  string     `json:"TotalNum" xml:"TotalNum"`
	Policies  []DataList `json:"Policies" xml:"Policies"`
}

// CreateListRegistrationPoliciesRequest creates a request to invoke ListRegistrationPolicies API
func CreateListRegistrationPoliciesRequest() (request *ListRegistrationPoliciesRequest) {
	request = &ListRegistrationPoliciesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "ListRegistrationPolicies", "", "")
	request.Method = requests.GET
	return
}

// CreateListRegistrationPoliciesResponse creates a response to parse from ListRegistrationPolicies response
func CreateListRegistrationPoliciesResponse() (response *ListRegistrationPoliciesResponse) {
	response = &ListRegistrationPoliciesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}