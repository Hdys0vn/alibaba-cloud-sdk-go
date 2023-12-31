package cloudcallcenter

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

// ListVnConversations invokes the cloudcallcenter.ListVnConversations API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listvnconversations.html
func (client *Client) ListVnConversations(request *ListVnConversationsRequest) (response *ListVnConversationsResponse, err error) {
	response = CreateListVnConversationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListVnConversationsWithChan invokes the cloudcallcenter.ListVnConversations API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listvnconversations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVnConversationsWithChan(request *ListVnConversationsRequest) (<-chan *ListVnConversationsResponse, <-chan error) {
	responseChan := make(chan *ListVnConversationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVnConversations(request)
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

// ListVnConversationsWithCallback invokes the cloudcallcenter.ListVnConversations API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listvnconversations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVnConversationsWithCallback(request *ListVnConversationsRequest, callback func(response *ListVnConversationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVnConversationsResponse
		var err error
		defer close(result)
		response, err = client.ListVnConversations(request)
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

// ListVnConversationsRequest is the request struct for api ListVnConversations
type ListVnConversationsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListVnConversationsResponse is the response struct for api ListVnConversations
type ListVnConversationsResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	TotalCount    int64          `json:"TotalCount" xml:"TotalCount"`
	PageNumber    int            `json:"PageNumber" xml:"PageNumber"`
	PageSize      int            `json:"PageSize" xml:"PageSize"`
	Conversations []Conversation `json:"Conversations" xml:"Conversations"`
}

// CreateListVnConversationsRequest creates a request to invoke ListVnConversations API
func CreateListVnConversationsRequest() (request *ListVnConversationsRequest) {
	request = &ListVnConversationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListVnConversations", "", "")
	request.Method = requests.GET
	return
}

// CreateListVnConversationsResponse creates a response to parse from ListVnConversations response
func CreateListVnConversationsResponse() (response *ListVnConversationsResponse) {
	response = &ListVnConversationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
