package codeup

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

// ListMergeRequestComments invokes the codeup.ListMergeRequestComments API synchronously
func (client *Client) ListMergeRequestComments(request *ListMergeRequestCommentsRequest) (response *ListMergeRequestCommentsResponse, err error) {
	response = CreateListMergeRequestCommentsResponse()
	err = client.DoAction(request, response)
	return
}

// ListMergeRequestCommentsWithChan invokes the codeup.ListMergeRequestComments API asynchronously
func (client *Client) ListMergeRequestCommentsWithChan(request *ListMergeRequestCommentsRequest) (<-chan *ListMergeRequestCommentsResponse, <-chan error) {
	responseChan := make(chan *ListMergeRequestCommentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMergeRequestComments(request)
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

// ListMergeRequestCommentsWithCallback invokes the codeup.ListMergeRequestComments API asynchronously
func (client *Client) ListMergeRequestCommentsWithCallback(request *ListMergeRequestCommentsRequest, callback func(response *ListMergeRequestCommentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMergeRequestCommentsResponse
		var err error
		defer close(result)
		response, err = client.ListMergeRequestComments(request)
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

// ListMergeRequestCommentsRequest is the request struct for api ListMergeRequestComments
type ListMergeRequestCommentsRequest struct {
	*requests.RoaRequest
	OrganizationId string           `position:"Query" name:"OrganizationId"`
	MergeRequestId requests.Integer `position:"Path" name:"MergeRequestId"`
	FromCommit     string           `position:"Query" name:"FromCommit"`
	AccessToken    string           `position:"Query" name:"AccessToken"`
	ToCommit       string           `position:"Query" name:"ToCommit"`
	ProjectId      requests.Integer `position:"Path" name:"ProjectId"`
}

// ListMergeRequestCommentsResponse is the response struct for api ListMergeRequestComments
type ListMergeRequestCommentsResponse struct {
	*responses.BaseResponse
	ErrorCode    string       `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string       `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Success      bool         `json:"Success" xml:"Success"`
	Total        int64        `json:"Total" xml:"Total"`
	Result       []ResultItem `json:"Result" xml:"Result"`
}

// CreateListMergeRequestCommentsRequest creates a request to invoke ListMergeRequestComments API
func CreateListMergeRequestCommentsRequest() (request *ListMergeRequestCommentsRequest) {
	request = &ListMergeRequestCommentsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "ListMergeRequestComments", "/api/v4/projects/[ProjectId]/merge_request/[MergeRequestId]/comments", "", "")
	request.Method = requests.GET
	return
}

// CreateListMergeRequestCommentsResponse creates a response to parse from ListMergeRequestComments response
func CreateListMergeRequestCommentsResponse() (response *ListMergeRequestCommentsResponse) {
	response = &ListMergeRequestCommentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
