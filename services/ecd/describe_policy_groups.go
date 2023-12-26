package ecd

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

// DescribePolicyGroups invokes the ecd.DescribePolicyGroups API synchronously
func (client *Client) DescribePolicyGroups(request *DescribePolicyGroupsRequest) (response *DescribePolicyGroupsResponse, err error) {
	response = CreateDescribePolicyGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePolicyGroupsWithChan invokes the ecd.DescribePolicyGroups API asynchronously
func (client *Client) DescribePolicyGroupsWithChan(request *DescribePolicyGroupsRequest) (<-chan *DescribePolicyGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribePolicyGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePolicyGroups(request)
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

// DescribePolicyGroupsWithCallback invokes the ecd.DescribePolicyGroups API asynchronously
func (client *Client) DescribePolicyGroupsWithCallback(request *DescribePolicyGroupsRequest, callback func(response *DescribePolicyGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePolicyGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribePolicyGroups(request)
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

// DescribePolicyGroupsRequest is the request struct for api DescribePolicyGroups
type DescribePolicyGroupsRequest struct {
	*requests.RpcRequest
	NextToken     string           `position:"Query" name:"NextToken"`
	MaxResults    requests.Integer `position:"Query" name:"MaxResults"`
	PolicyGroupId *[]string        `position:"Query" name:"PolicyGroupId"  type:"Repeated"`
}

// DescribePolicyGroupsResponse is the response struct for api DescribePolicyGroups
type DescribePolicyGroupsResponse struct {
	*responses.BaseResponse
	NextToken            string                `json:"NextToken" xml:"NextToken"`
	RequestId            string                `json:"RequestId" xml:"RequestId"`
	DescribePolicyGroups []DescribePolicyGroup `json:"DescribePolicyGroups" xml:"DescribePolicyGroups"`
}

// CreateDescribePolicyGroupsRequest creates a request to invoke DescribePolicyGroups API
func CreateDescribePolicyGroupsRequest() (request *DescribePolicyGroupsRequest) {
	request = &DescribePolicyGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribePolicyGroups", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribePolicyGroupsResponse creates a response to parse from DescribePolicyGroups response
func CreateDescribePolicyGroupsResponse() (response *DescribePolicyGroupsResponse) {
	response = &DescribePolicyGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
