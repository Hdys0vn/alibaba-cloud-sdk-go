package elasticsearch

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

// ListAlternativeSnapshotRepos invokes the elasticsearch.ListAlternativeSnapshotRepos API synchronously
func (client *Client) ListAlternativeSnapshotRepos(request *ListAlternativeSnapshotReposRequest) (response *ListAlternativeSnapshotReposResponse, err error) {
	response = CreateListAlternativeSnapshotReposResponse()
	err = client.DoAction(request, response)
	return
}

// ListAlternativeSnapshotReposWithChan invokes the elasticsearch.ListAlternativeSnapshotRepos API asynchronously
func (client *Client) ListAlternativeSnapshotReposWithChan(request *ListAlternativeSnapshotReposRequest) (<-chan *ListAlternativeSnapshotReposResponse, <-chan error) {
	responseChan := make(chan *ListAlternativeSnapshotReposResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAlternativeSnapshotRepos(request)
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

// ListAlternativeSnapshotReposWithCallback invokes the elasticsearch.ListAlternativeSnapshotRepos API asynchronously
func (client *Client) ListAlternativeSnapshotReposWithCallback(request *ListAlternativeSnapshotReposRequest, callback func(response *ListAlternativeSnapshotReposResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAlternativeSnapshotReposResponse
		var err error
		defer close(result)
		response, err = client.ListAlternativeSnapshotRepos(request)
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

// ListAlternativeSnapshotReposRequest is the request struct for api ListAlternativeSnapshotRepos
type ListAlternativeSnapshotReposRequest struct {
	*requests.RoaRequest
	InstanceId      string           `position:"Path" name:"InstanceId"`
	AlreadySetItems requests.Boolean `position:"Query" name:"alreadySetItems"`
}

// ListAlternativeSnapshotReposResponse is the response struct for api ListAlternativeSnapshotRepos
type ListAlternativeSnapshotReposResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    []Repo `json:"Result" xml:"Result"`
}

// CreateListAlternativeSnapshotReposRequest creates a request to invoke ListAlternativeSnapshotRepos API
func CreateListAlternativeSnapshotReposRequest() (request *ListAlternativeSnapshotReposRequest) {
	request = &ListAlternativeSnapshotReposRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListAlternativeSnapshotRepos", "/openapi/instances/[InstanceId]/alternative-snapshot-repos", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListAlternativeSnapshotReposResponse creates a response to parse from ListAlternativeSnapshotRepos response
func CreateListAlternativeSnapshotReposResponse() (response *ListAlternativeSnapshotReposResponse) {
	response = &ListAlternativeSnapshotReposResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
