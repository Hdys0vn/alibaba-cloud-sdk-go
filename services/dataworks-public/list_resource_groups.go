package dataworks_public

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

// ListResourceGroups invokes the dataworks_public.ListResourceGroups API synchronously
func (client *Client) ListResourceGroups(request *ListResourceGroupsRequest) (response *ListResourceGroupsResponse, err error) {
	response = CreateListResourceGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListResourceGroupsWithChan invokes the dataworks_public.ListResourceGroups API asynchronously
func (client *Client) ListResourceGroupsWithChan(request *ListResourceGroupsRequest) (<-chan *ListResourceGroupsResponse, <-chan error) {
	responseChan := make(chan *ListResourceGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListResourceGroups(request)
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

// ListResourceGroupsWithCallback invokes the dataworks_public.ListResourceGroups API asynchronously
func (client *Client) ListResourceGroupsWithCallback(request *ListResourceGroupsRequest, callback func(response *ListResourceGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListResourceGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListResourceGroups(request)
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

// ListResourceGroupsRequest is the request struct for api ListResourceGroups
type ListResourceGroupsRequest struct {
	*requests.RpcRequest
	BizExtKey                      string                    `position:"Query" name:"BizExtKey"`
	Tags                           *[]ListResourceGroupsTags `position:"Query" name:"Tags"  type:"Json"`
	TypeNames                      string                    `position:"Query" name:"TypeNames"`
	ResourceManagerResourceGroupId string                    `position:"Query" name:"ResourceManagerResourceGroupId"`
	ResourceGroupType              requests.Integer          `position:"Query" name:"ResourceGroupType"`
	Keyword                        string                    `position:"Query" name:"Keyword"`
}

// ListResourceGroupsTags is a repeated param struct in ListResourceGroupsRequest
type ListResourceGroupsTags struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ListResourceGroupsResponse is the response struct for api ListResourceGroups
type ListResourceGroupsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string                         `json:"RequestId" xml:"RequestId"`
	Success        bool                           `json:"Success" xml:"Success"`
	Data           []DataItemInListResourceGroups `json:"Data" xml:"Data"`
}

// CreateListResourceGroupsRequest creates a request to invoke ListResourceGroups API
func CreateListResourceGroupsRequest() (request *ListResourceGroupsRequest) {
	request = &ListResourceGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListResourceGroups", "", "")
	request.Method = requests.POST
	return
}

// CreateListResourceGroupsResponse creates a response to parse from ListResourceGroups response
func CreateListResourceGroupsResponse() (response *ListResourceGroupsResponse) {
	response = &ListResourceGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}