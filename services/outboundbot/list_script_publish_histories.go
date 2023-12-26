package outboundbot

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

// ListScriptPublishHistories invokes the outboundbot.ListScriptPublishHistories API synchronously
func (client *Client) ListScriptPublishHistories(request *ListScriptPublishHistoriesRequest) (response *ListScriptPublishHistoriesResponse, err error) {
	response = CreateListScriptPublishHistoriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListScriptPublishHistoriesWithChan invokes the outboundbot.ListScriptPublishHistories API asynchronously
func (client *Client) ListScriptPublishHistoriesWithChan(request *ListScriptPublishHistoriesRequest) (<-chan *ListScriptPublishHistoriesResponse, <-chan error) {
	responseChan := make(chan *ListScriptPublishHistoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListScriptPublishHistories(request)
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

// ListScriptPublishHistoriesWithCallback invokes the outboundbot.ListScriptPublishHistories API asynchronously
func (client *Client) ListScriptPublishHistoriesWithCallback(request *ListScriptPublishHistoriesRequest, callback func(response *ListScriptPublishHistoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListScriptPublishHistoriesResponse
		var err error
		defer close(result)
		response, err = client.ListScriptPublishHistories(request)
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

// ListScriptPublishHistoriesRequest is the request struct for api ListScriptPublishHistories
type ListScriptPublishHistoriesRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	ScriptId   string           `position:"Query" name:"ScriptId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListScriptPublishHistoriesResponse is the response struct for api ListScriptPublishHistories
type ListScriptPublishHistoriesResponse struct {
	*responses.BaseResponse
	HttpStatusCode         int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code                   string                 `json:"Code" xml:"Code"`
	Message                string                 `json:"Message" xml:"Message"`
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	Success                bool                   `json:"Success" xml:"Success"`
	ScriptPublishHistories ScriptPublishHistories `json:"ScriptPublishHistories" xml:"ScriptPublishHistories"`
}

// CreateListScriptPublishHistoriesRequest creates a request to invoke ListScriptPublishHistories API
func CreateListScriptPublishHistoriesRequest() (request *ListScriptPublishHistoriesRequest) {
	request = &ListScriptPublishHistoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ListScriptPublishHistories", "", "")
	request.Method = requests.POST
	return
}

// CreateListScriptPublishHistoriesResponse creates a response to parse from ListScriptPublishHistories response
func CreateListScriptPublishHistoriesResponse() (response *ListScriptPublishHistoriesResponse) {
	response = &ListScriptPublishHistoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
