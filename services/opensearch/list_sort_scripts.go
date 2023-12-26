package opensearch

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

// ListSortScripts invokes the opensearch.ListSortScripts API synchronously
func (client *Client) ListSortScripts(request *ListSortScriptsRequest) (response *ListSortScriptsResponse, err error) {
	response = CreateListSortScriptsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSortScriptsWithChan invokes the opensearch.ListSortScripts API asynchronously
func (client *Client) ListSortScriptsWithChan(request *ListSortScriptsRequest) (<-chan *ListSortScriptsResponse, <-chan error) {
	responseChan := make(chan *ListSortScriptsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSortScripts(request)
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

// ListSortScriptsWithCallback invokes the opensearch.ListSortScripts API asynchronously
func (client *Client) ListSortScriptsWithCallback(request *ListSortScriptsRequest, callback func(response *ListSortScriptsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSortScriptsResponse
		var err error
		defer close(result)
		response, err = client.ListSortScripts(request)
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

// ListSortScriptsRequest is the request struct for api ListSortScripts
type ListSortScriptsRequest struct {
	*requests.RoaRequest
	AppVersionId     string `position:"Path" name:"appVersionId"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// ListSortScriptsResponse is the response struct for api ListSortScripts
type ListSortScriptsResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"requestId" xml:"requestId"`
	Result    []ScriptInfo `json:"result" xml:"result"`
}

// CreateListSortScriptsRequest creates a request to invoke ListSortScripts API
func CreateListSortScriptsRequest() (request *ListSortScriptsRequest) {
	request = &ListSortScriptsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListSortScripts", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appVersionId]/sort-scripts", "", "")
	request.Method = requests.GET
	return
}

// CreateListSortScriptsResponse creates a response to parse from ListSortScripts response
func CreateListSortScriptsResponse() (response *ListSortScriptsResponse) {
	response = &ListSortScriptsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
