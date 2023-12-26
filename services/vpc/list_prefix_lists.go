package vpc

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

// ListPrefixLists invokes the vpc.ListPrefixLists API synchronously
func (client *Client) ListPrefixLists(request *ListPrefixListsRequest) (response *ListPrefixListsResponse, err error) {
	response = CreateListPrefixListsResponse()
	err = client.DoAction(request, response)
	return
}

// ListPrefixListsWithChan invokes the vpc.ListPrefixLists API asynchronously
func (client *Client) ListPrefixListsWithChan(request *ListPrefixListsRequest) (<-chan *ListPrefixListsResponse, <-chan error) {
	responseChan := make(chan *ListPrefixListsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPrefixLists(request)
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

// ListPrefixListsWithCallback invokes the vpc.ListPrefixLists API asynchronously
func (client *Client) ListPrefixListsWithCallback(request *ListPrefixListsRequest, callback func(response *ListPrefixListsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPrefixListsResponse
		var err error
		defer close(result)
		response, err = client.ListPrefixLists(request)
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

// ListPrefixListsRequest is the request struct for api ListPrefixLists
type ListPrefixListsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer       `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId      string                 `position:"Query" name:"ResourceGroupId"`
	NextToken            string                 `position:"Query" name:"NextToken"`
	ResourceOwnerAccount string                 `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                 `position:"Query" name:"OwnerAccount"`
	PrefixListIds        *[]string              `position:"Query" name:"PrefixListIds"  type:"Repeated"`
	OwnerId              requests.Integer       `position:"Query" name:"OwnerId"`
	Tags                 *[]ListPrefixListsTags `position:"Query" name:"Tags"  type:"Repeated"`
	PrefixListName       string                 `position:"Query" name:"PrefixListName"`
	MaxResults           requests.Integer       `position:"Query" name:"MaxResults"`
}

// ListPrefixListsTags is a repeated param struct in ListPrefixListsRequest
type ListPrefixListsTags struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// ListPrefixListsResponse is the response struct for api ListPrefixLists
type ListPrefixListsResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	NextToken   string       `json:"NextToken" xml:"NextToken"`
	TotalCount  int64        `json:"TotalCount" xml:"TotalCount"`
	MaxResults  int64        `json:"MaxResults" xml:"MaxResults"`
	PrefixLists []PrefixList `json:"PrefixLists" xml:"PrefixLists"`
}

// CreateListPrefixListsRequest creates a request to invoke ListPrefixLists API
func CreateListPrefixListsRequest() (request *ListPrefixListsRequest) {
	request = &ListPrefixListsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ListPrefixLists", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPrefixListsResponse creates a response to parse from ListPrefixLists response
func CreateListPrefixListsResponse() (response *ListPrefixListsResponse) {
	response = &ListPrefixListsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
