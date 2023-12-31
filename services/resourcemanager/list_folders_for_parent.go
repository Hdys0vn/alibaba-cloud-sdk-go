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

// ListFoldersForParent invokes the resourcemanager.ListFoldersForParent API synchronously
func (client *Client) ListFoldersForParent(request *ListFoldersForParentRequest) (response *ListFoldersForParentResponse, err error) {
	response = CreateListFoldersForParentResponse()
	err = client.DoAction(request, response)
	return
}

// ListFoldersForParentWithChan invokes the resourcemanager.ListFoldersForParent API asynchronously
func (client *Client) ListFoldersForParentWithChan(request *ListFoldersForParentRequest) (<-chan *ListFoldersForParentResponse, <-chan error) {
	responseChan := make(chan *ListFoldersForParentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFoldersForParent(request)
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

// ListFoldersForParentWithCallback invokes the resourcemanager.ListFoldersForParent API asynchronously
func (client *Client) ListFoldersForParentWithCallback(request *ListFoldersForParentRequest, callback func(response *ListFoldersForParentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFoldersForParentResponse
		var err error
		defer close(result)
		response, err = client.ListFoldersForParent(request)
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

// ListFoldersForParentRequest is the request struct for api ListFoldersForParent
type ListFoldersForParentRequest struct {
	*requests.RpcRequest
	QueryKeyword   string           `position:"Query" name:"QueryKeyword"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	ParentFolderId string           `position:"Query" name:"ParentFolderId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
}

// ListFoldersForParentResponse is the response struct for api ListFoldersForParent
type ListFoldersForParentResponse struct {
	*responses.BaseResponse
	TotalCount int                           `json:"TotalCount" xml:"TotalCount"`
	RequestId  string                        `json:"RequestId" xml:"RequestId"`
	PageSize   int                           `json:"PageSize" xml:"PageSize"`
	PageNumber int                           `json:"PageNumber" xml:"PageNumber"`
	Folders    FoldersInListFoldersForParent `json:"Folders" xml:"Folders"`
}

// CreateListFoldersForParentRequest creates a request to invoke ListFoldersForParent API
func CreateListFoldersForParentRequest() (request *ListFoldersForParentRequest) {
	request = &ListFoldersForParentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "ListFoldersForParent", "", "")
	request.Method = requests.POST
	return
}

// CreateListFoldersForParentResponse creates a response to parse from ListFoldersForParent response
func CreateListFoldersForParentResponse() (response *ListFoldersForParentResponse) {
	response = &ListFoldersForParentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
