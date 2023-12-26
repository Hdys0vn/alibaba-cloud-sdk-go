package mse

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

// ListAppBySwimmingLaneGroupTag invokes the mse.ListAppBySwimmingLaneGroupTag API synchronously
func (client *Client) ListAppBySwimmingLaneGroupTag(request *ListAppBySwimmingLaneGroupTagRequest) (response *ListAppBySwimmingLaneGroupTagResponse, err error) {
	response = CreateListAppBySwimmingLaneGroupTagResponse()
	err = client.DoAction(request, response)
	return
}

// ListAppBySwimmingLaneGroupTagWithChan invokes the mse.ListAppBySwimmingLaneGroupTag API asynchronously
func (client *Client) ListAppBySwimmingLaneGroupTagWithChan(request *ListAppBySwimmingLaneGroupTagRequest) (<-chan *ListAppBySwimmingLaneGroupTagResponse, <-chan error) {
	responseChan := make(chan *ListAppBySwimmingLaneGroupTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAppBySwimmingLaneGroupTag(request)
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

// ListAppBySwimmingLaneGroupTagWithCallback invokes the mse.ListAppBySwimmingLaneGroupTag API asynchronously
func (client *Client) ListAppBySwimmingLaneGroupTagWithCallback(request *ListAppBySwimmingLaneGroupTagRequest, callback func(response *ListAppBySwimmingLaneGroupTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAppBySwimmingLaneGroupTagResponse
		var err error
		defer close(result)
		response, err = client.ListAppBySwimmingLaneGroupTag(request)
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

// ListAppBySwimmingLaneGroupTagRequest is the request struct for api ListAppBySwimmingLaneGroupTag
type ListAppBySwimmingLaneGroupTagRequest struct {
	*requests.RpcRequest
	MseSessionId   string           `position:"Query" name:"MseSessionId"`
	Tag            string           `position:"Query" name:"Tag"`
	GroupId        requests.Integer `position:"Query" name:"GroupId"`
	Namespace      string           `position:"Query" name:"Namespace"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
}

// ListAppBySwimmingLaneGroupTagResponse is the response struct for api ListAppBySwimmingLaneGroupTag
type ListAppBySwimmingLaneGroupTagResponse struct {
	*responses.BaseResponse
}

// CreateListAppBySwimmingLaneGroupTagRequest creates a request to invoke ListAppBySwimmingLaneGroupTag API
func CreateListAppBySwimmingLaneGroupTagRequest() (request *ListAppBySwimmingLaneGroupTagRequest) {
	request = &ListAppBySwimmingLaneGroupTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListAppBySwimmingLaneGroupTag", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAppBySwimmingLaneGroupTagResponse creates a response to parse from ListAppBySwimmingLaneGroupTag response
func CreateListAppBySwimmingLaneGroupTagResponse() (response *ListAppBySwimmingLaneGroupTagResponse) {
	response = &ListAppBySwimmingLaneGroupTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}