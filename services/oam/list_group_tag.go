package oam

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

// ListGroupTag invokes the oam.ListGroupTag API synchronously
// api document: https://help.aliyun.com/api/oam/listgrouptag.html
func (client *Client) ListGroupTag(request *ListGroupTagRequest) (response *ListGroupTagResponse, err error) {
	response = CreateListGroupTagResponse()
	err = client.DoAction(request, response)
	return
}

// ListGroupTagWithChan invokes the oam.ListGroupTag API asynchronously
// api document: https://help.aliyun.com/api/oam/listgrouptag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListGroupTagWithChan(request *ListGroupTagRequest) (<-chan *ListGroupTagResponse, <-chan error) {
	responseChan := make(chan *ListGroupTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGroupTag(request)
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

// ListGroupTagWithCallback invokes the oam.ListGroupTag API asynchronously
// api document: https://help.aliyun.com/api/oam/listgrouptag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListGroupTagWithCallback(request *ListGroupTagRequest, callback func(response *ListGroupTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGroupTagResponse
		var err error
		defer close(result)
		response, err = client.ListGroupTag(request)
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

// ListGroupTagRequest is the request struct for api ListGroupTag
type ListGroupTagRequest struct {
	*requests.RpcRequest
	GroupTagId string `position:"Query" name:"GroupTagId"`
	Name       string `position:"Query" name:"Name"`
	UserId     string `position:"Query" name:"UserId"`
}

// ListGroupTagResponse is the response struct for api ListGroupTag
type ListGroupTagResponse struct {
	*responses.BaseResponse
	Code      string             `json:"Code" xml:"Code"`
	Message   string             `json:"Message" xml:"Message"`
	RequestId string             `json:"RequestId" xml:"RequestId"`
	Data      DataInListGroupTag `json:"Data" xml:"Data"`
}

// CreateListGroupTagRequest creates a request to invoke ListGroupTag API
func CreateListGroupTagRequest() (request *ListGroupTagRequest) {
	request = &ListGroupTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "ListGroupTag", "", "")
	request.Method = requests.POST
	return
}

// CreateListGroupTagResponse creates a response to parse from ListGroupTag response
func CreateListGroupTagResponse() (response *ListGroupTagResponse) {
	response = &ListGroupTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
