package config

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

// ListPreManagedRules invokes the config.ListPreManagedRules API synchronously
func (client *Client) ListPreManagedRules(request *ListPreManagedRulesRequest) (response *ListPreManagedRulesResponse, err error) {
	response = CreateListPreManagedRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListPreManagedRulesWithChan invokes the config.ListPreManagedRules API asynchronously
func (client *Client) ListPreManagedRulesWithChan(request *ListPreManagedRulesRequest) (<-chan *ListPreManagedRulesResponse, <-chan error) {
	responseChan := make(chan *ListPreManagedRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPreManagedRules(request)
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

// ListPreManagedRulesWithCallback invokes the config.ListPreManagedRules API asynchronously
func (client *Client) ListPreManagedRulesWithCallback(request *ListPreManagedRulesRequest, callback func(response *ListPreManagedRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPreManagedRulesResponse
		var err error
		defer close(result)
		response, err = client.ListPreManagedRules(request)
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

// ListPreManagedRulesRequest is the request struct for api ListPreManagedRules
type ListPreManagedRulesRequest struct {
	*requests.RpcRequest
	ResourceTypes      *[]string        `position:"Body" name:"ResourceTypes"  type:"Json"`
	PageNumber         requests.Integer `position:"Body" name:"PageNumber"`
	PageSize           requests.Integer `position:"Body" name:"PageSize"`
	ResourceTypeFormat string           `position:"Body" name:"ResourceTypeFormat"`
}

// ListPreManagedRulesResponse is the response struct for api ListPreManagedRules
type ListPreManagedRulesResponse struct {
	*responses.BaseResponse
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	PageNumber   int64         `json:"PageNumber" xml:"PageNumber"`
	PageSize     int64         `json:"PageSize" xml:"PageSize"`
	ManagedRules []ManagedRule `json:"ManagedRules" xml:"ManagedRules"`
}

// CreateListPreManagedRulesRequest creates a request to invoke ListPreManagedRules API
func CreateListPreManagedRulesRequest() (request *ListPreManagedRulesRequest) {
	request = &ListPreManagedRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "ListPreManagedRules", "", "")
	request.Method = requests.POST
	return
}

// CreateListPreManagedRulesResponse creates a response to parse from ListPreManagedRules response
func CreateListPreManagedRulesResponse() (response *ListPreManagedRulesResponse) {
	response = &ListPreManagedRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
