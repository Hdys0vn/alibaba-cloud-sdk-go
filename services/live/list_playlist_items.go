package live

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

// ListPlaylistItems invokes the live.ListPlaylistItems API synchronously
func (client *Client) ListPlaylistItems(request *ListPlaylistItemsRequest) (response *ListPlaylistItemsResponse, err error) {
	response = CreateListPlaylistItemsResponse()
	err = client.DoAction(request, response)
	return
}

// ListPlaylistItemsWithChan invokes the live.ListPlaylistItems API asynchronously
func (client *Client) ListPlaylistItemsWithChan(request *ListPlaylistItemsRequest) (<-chan *ListPlaylistItemsResponse, <-chan error) {
	responseChan := make(chan *ListPlaylistItemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPlaylistItems(request)
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

// ListPlaylistItemsWithCallback invokes the live.ListPlaylistItems API asynchronously
func (client *Client) ListPlaylistItemsWithCallback(request *ListPlaylistItemsRequest, callback func(response *ListPlaylistItemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPlaylistItemsResponse
		var err error
		defer close(result)
		response, err = client.ListPlaylistItems(request)
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

// ListPlaylistItemsRequest is the request struct for api ListPlaylistItems
type ListPlaylistItemsRequest struct {
	*requests.RpcRequest
	ProgramItemIds string           `position:"Query" name:"ProgramItemIds"`
	ProgramId      string           `position:"Query" name:"ProgramId"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
}

// ListPlaylistItemsResponse is the response struct for api ListPlaylistItems
type ListPlaylistItemsResponse struct {
	*responses.BaseResponse
	Total        int           `json:"Total" xml:"Total"`
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	ProgramItems []ProgramItem `json:"ProgramItems" xml:"ProgramItems"`
}

// CreateListPlaylistItemsRequest creates a request to invoke ListPlaylistItems API
func CreateListPlaylistItemsRequest() (request *ListPlaylistItemsRequest) {
	request = &ListPlaylistItemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ListPlaylistItems", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPlaylistItemsResponse creates a response to parse from ListPlaylistItems response
func CreateListPlaylistItemsResponse() (response *ListPlaylistItemsResponse) {
	response = &ListPlaylistItemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
