package push

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

// ListSummaryApps invokes the push.ListSummaryApps API synchronously
func (client *Client) ListSummaryApps(request *ListSummaryAppsRequest) (response *ListSummaryAppsResponse, err error) {
	response = CreateListSummaryAppsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSummaryAppsWithChan invokes the push.ListSummaryApps API asynchronously
func (client *Client) ListSummaryAppsWithChan(request *ListSummaryAppsRequest) (<-chan *ListSummaryAppsResponse, <-chan error) {
	responseChan := make(chan *ListSummaryAppsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSummaryApps(request)
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

// ListSummaryAppsWithCallback invokes the push.ListSummaryApps API asynchronously
func (client *Client) ListSummaryAppsWithCallback(request *ListSummaryAppsRequest, callback func(response *ListSummaryAppsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSummaryAppsResponse
		var err error
		defer close(result)
		response, err = client.ListSummaryApps(request)
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

// ListSummaryAppsRequest is the request struct for api ListSummaryApps
type ListSummaryAppsRequest struct {
	*requests.RpcRequest
}

// ListSummaryAppsResponse is the response struct for api ListSummaryApps
type ListSummaryAppsResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	SummaryAppInfos SummaryAppInfos `json:"SummaryAppInfos" xml:"SummaryAppInfos"`
}

// CreateListSummaryAppsRequest creates a request to invoke ListSummaryApps API
func CreateListSummaryAppsRequest() (request *ListSummaryAppsRequest) {
	request = &ListSummaryAppsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "ListSummaryApps", "", "")
	request.Method = requests.POST
	return
}

// CreateListSummaryAppsResponse creates a response to parse from ListSummaryApps response
func CreateListSummaryAppsResponse() (response *ListSummaryAppsResponse) {
	response = &ListSummaryAppsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
