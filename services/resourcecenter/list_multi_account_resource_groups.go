package resourcecenter

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

// ListMultiAccountResourceGroups invokes the resourcecenter.ListMultiAccountResourceGroups API synchronously
func (client *Client) ListMultiAccountResourceGroups(request *ListMultiAccountResourceGroupsRequest) (response *ListMultiAccountResourceGroupsResponse, err error) {
	response = CreateListMultiAccountResourceGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListMultiAccountResourceGroupsWithChan invokes the resourcecenter.ListMultiAccountResourceGroups API asynchronously
func (client *Client) ListMultiAccountResourceGroupsWithChan(request *ListMultiAccountResourceGroupsRequest) (<-chan *ListMultiAccountResourceGroupsResponse, <-chan error) {
	responseChan := make(chan *ListMultiAccountResourceGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMultiAccountResourceGroups(request)
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

// ListMultiAccountResourceGroupsWithCallback invokes the resourcecenter.ListMultiAccountResourceGroups API asynchronously
func (client *Client) ListMultiAccountResourceGroupsWithCallback(request *ListMultiAccountResourceGroupsRequest, callback func(response *ListMultiAccountResourceGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMultiAccountResourceGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListMultiAccountResourceGroups(request)
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

// ListMultiAccountResourceGroupsRequest is the request struct for api ListMultiAccountResourceGroups
type ListMultiAccountResourceGroupsRequest struct {
	*requests.RpcRequest
	AccountId        string           `position:"Query" name:"AccountId"`
	ResourceGroupIds *[]string        `position:"Query" name:"ResourceGroupIds"  type:"Repeated"`
	NextToken        string           `position:"Query" name:"NextToken"`
	MaxResults       requests.Integer `position:"Query" name:"MaxResults"`
}

// ListMultiAccountResourceGroupsResponse is the response struct for api ListMultiAccountResourceGroups
type ListMultiAccountResourceGroupsResponse struct {
	*responses.BaseResponse
	NextToken      string          `json:"NextToken" xml:"NextToken"`
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	Success        bool            `json:"Success" xml:"Success"`
	DynamicCode    string          `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string          `json:"DynamicMessage" xml:"DynamicMessage"`
	ErrorCode      string          `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string          `json:"ErrorMessage" xml:"ErrorMessage"`
	ResourceGroups []ResourceGroup `json:"ResourceGroups" xml:"ResourceGroups"`
}

// CreateListMultiAccountResourceGroupsRequest creates a request to invoke ListMultiAccountResourceGroups API
func CreateListMultiAccountResourceGroupsRequest() (request *ListMultiAccountResourceGroupsRequest) {
	request = &ListMultiAccountResourceGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceCenter", "2022-12-01", "ListMultiAccountResourceGroups", "", "")
	request.Method = requests.POST
	return
}

// CreateListMultiAccountResourceGroupsResponse creates a response to parse from ListMultiAccountResourceGroups response
func CreateListMultiAccountResourceGroupsResponse() (response *ListMultiAccountResourceGroupsResponse) {
	response = &ListMultiAccountResourceGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}