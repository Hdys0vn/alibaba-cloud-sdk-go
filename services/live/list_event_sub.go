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

// ListEventSub invokes the live.ListEventSub API synchronously
func (client *Client) ListEventSub(request *ListEventSubRequest) (response *ListEventSubResponse, err error) {
	response = CreateListEventSubResponse()
	err = client.DoAction(request, response)
	return
}

// ListEventSubWithChan invokes the live.ListEventSub API asynchronously
func (client *Client) ListEventSubWithChan(request *ListEventSubRequest) (<-chan *ListEventSubResponse, <-chan error) {
	responseChan := make(chan *ListEventSubResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEventSub(request)
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

// ListEventSubWithCallback invokes the live.ListEventSub API asynchronously
func (client *Client) ListEventSubWithCallback(request *ListEventSubRequest, callback func(response *ListEventSubResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEventSubResponse
		var err error
		defer close(result)
		response, err = client.ListEventSub(request)
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

// ListEventSubRequest is the request struct for api ListEventSub
type ListEventSubRequest struct {
	*requests.RpcRequest
	AppId string `position:"Query" name:"AppId"`
}

// ListEventSubResponse is the response struct for api ListEventSub
type ListEventSubResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Subscribers []Data `json:"Subscribers" xml:"Subscribers"`
}

// CreateListEventSubRequest creates a request to invoke ListEventSub API
func CreateListEventSubRequest() (request *ListEventSubRequest) {
	request = &ListEventSubRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ListEventSub", "live", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListEventSubResponse creates a response to parse from ListEventSub response
func CreateListEventSubResponse() (response *ListEventSubResponse) {
	response = &ListEventSubResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}