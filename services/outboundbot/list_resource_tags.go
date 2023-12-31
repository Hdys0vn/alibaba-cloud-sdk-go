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

// ListResourceTags invokes the outboundbot.ListResourceTags API synchronously
func (client *Client) ListResourceTags(request *ListResourceTagsRequest) (response *ListResourceTagsResponse, err error) {
	response = CreateListResourceTagsResponse()
	err = client.DoAction(request, response)
	return
}

// ListResourceTagsWithChan invokes the outboundbot.ListResourceTags API asynchronously
func (client *Client) ListResourceTagsWithChan(request *ListResourceTagsRequest) (<-chan *ListResourceTagsResponse, <-chan error) {
	responseChan := make(chan *ListResourceTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListResourceTags(request)
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

// ListResourceTagsWithCallback invokes the outboundbot.ListResourceTags API asynchronously
func (client *Client) ListResourceTagsWithCallback(request *ListResourceTagsRequest, callback func(response *ListResourceTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListResourceTagsResponse
		var err error
		defer close(result)
		response, err = client.ListResourceTags(request)
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

// ListResourceTagsRequest is the request struct for api ListResourceTags
type ListResourceTagsRequest struct {
	*requests.RpcRequest
	ResourceType string           `position:"Query" name:"ResourceType"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
}

// ListResourceTagsResponse is the response struct for api ListResourceTags
type ListResourceTagsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string       `json:"Code" xml:"Code"`
	Message        string       `json:"Message" xml:"Message"`
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Success        bool         `json:"Success" xml:"Success"`
	ResourceTags   ResourceTags `json:"ResourceTags" xml:"ResourceTags"`
}

// CreateListResourceTagsRequest creates a request to invoke ListResourceTags API
func CreateListResourceTagsRequest() (request *ListResourceTagsRequest) {
	request = &ListResourceTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ListResourceTags", "", "")
	request.Method = requests.POST
	return
}

// CreateListResourceTagsResponse creates a response to parse from ListResourceTags response
func CreateListResourceTagsResponse() (response *ListResourceTagsResponse) {
	response = &ListResourceTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
