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

// ListConversationsWithUnrecognizedIntents invokes the cloudcallcenter.ListConversationsWithUnrecognizedIntents API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listconversationswithunrecognizedintents.html
func (client *Client) ListConversationsWithUnrecognizedIntents(request *ListConversationsWithUnrecognizedIntentsRequest) (response *ListConversationsWithUnrecognizedIntentsResponse, err error) {
	response = CreateListConversationsWithUnrecognizedIntentsResponse()
	err = client.DoAction(request, response)
	return
}

// ListConversationsWithUnrecognizedIntentsWithChan invokes the cloudcallcenter.ListConversationsWithUnrecognizedIntents API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listconversationswithunrecognizedintents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListConversationsWithUnrecognizedIntentsWithChan(request *ListConversationsWithUnrecognizedIntentsRequest) (<-chan *ListConversationsWithUnrecognizedIntentsResponse, <-chan error) {
	responseChan := make(chan *ListConversationsWithUnrecognizedIntentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListConversationsWithUnrecognizedIntents(request)
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

// ListConversationsWithUnrecognizedIntentsWithCallback invokes the cloudcallcenter.ListConversationsWithUnrecognizedIntents API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listconversationswithunrecognizedintents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListConversationsWithUnrecognizedIntentsWithCallback(request *ListConversationsWithUnrecognizedIntentsRequest, callback func(response *ListConversationsWithUnrecognizedIntentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListConversationsWithUnrecognizedIntentsResponse
		var err error
		defer close(result)
		response, err = client.ListConversationsWithUnrecognizedIntents(request)
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

// ListConversationsWithUnrecognizedIntentsRequest is the request struct for api ListConversationsWithUnrecognizedIntents
type ListConversationsWithUnrecognizedIntentsRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	JobGroupId string `position:"Query" name:"JobGroupId"`
	NodeId     string `position:"Query" name:"NodeId"`
}

// ListConversationsWithUnrecognizedIntentsResponse is the response struct for api ListConversationsWithUnrecognizedIntents
type ListConversationsWithUnrecognizedIntentsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Tasks          []Task `json:"Tasks" xml:"Tasks"`
}

// CreateListConversationsWithUnrecognizedIntentsRequest creates a request to invoke ListConversationsWithUnrecognizedIntents API
func CreateListConversationsWithUnrecognizedIntentsRequest() (request *ListConversationsWithUnrecognizedIntentsRequest) {
	request = &ListConversationsWithUnrecognizedIntentsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListConversationsWithUnrecognizedIntents", "", "")
	request.Method = requests.POST
	return
}

// CreateListConversationsWithUnrecognizedIntentsResponse creates a response to parse from ListConversationsWithUnrecognizedIntents response
func CreateListConversationsWithUnrecognizedIntentsResponse() (response *ListConversationsWithUnrecognizedIntentsResponse) {
	response = &ListConversationsWithUnrecognizedIntentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
