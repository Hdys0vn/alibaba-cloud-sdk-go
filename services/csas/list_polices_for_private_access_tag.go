package csas

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

// ListPolicesForPrivateAccessTag invokes the csas.ListPolicesForPrivateAccessTag API synchronously
func (client *Client) ListPolicesForPrivateAccessTag(request *ListPolicesForPrivateAccessTagRequest) (response *ListPolicesForPrivateAccessTagResponse, err error) {
	response = CreateListPolicesForPrivateAccessTagResponse()
	err = client.DoAction(request, response)
	return
}

// ListPolicesForPrivateAccessTagWithChan invokes the csas.ListPolicesForPrivateAccessTag API asynchronously
func (client *Client) ListPolicesForPrivateAccessTagWithChan(request *ListPolicesForPrivateAccessTagRequest) (<-chan *ListPolicesForPrivateAccessTagResponse, <-chan error) {
	responseChan := make(chan *ListPolicesForPrivateAccessTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPolicesForPrivateAccessTag(request)
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

// ListPolicesForPrivateAccessTagWithCallback invokes the csas.ListPolicesForPrivateAccessTag API asynchronously
func (client *Client) ListPolicesForPrivateAccessTagWithCallback(request *ListPolicesForPrivateAccessTagRequest, callback func(response *ListPolicesForPrivateAccessTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPolicesForPrivateAccessTagResponse
		var err error
		defer close(result)
		response, err = client.ListPolicesForPrivateAccessTag(request)
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

// ListPolicesForPrivateAccessTagRequest is the request struct for api ListPolicesForPrivateAccessTag
type ListPolicesForPrivateAccessTagRequest struct {
	*requests.RpcRequest
	TagIds   *[]string `position:"Query" name:"TagIds"  type:"Repeated"`
	SourceIp string    `position:"Query" name:"SourceIp"`
}

// ListPolicesForPrivateAccessTagResponse is the response struct for api ListPolicesForPrivateAccessTag
type ListPolicesForPrivateAccessTagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Tags      []Tag  `json:"Tags" xml:"Tags"`
}

// CreateListPolicesForPrivateAccessTagRequest creates a request to invoke ListPolicesForPrivateAccessTag API
func CreateListPolicesForPrivateAccessTagRequest() (request *ListPolicesForPrivateAccessTagRequest) {
	request = &ListPolicesForPrivateAccessTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "ListPolicesForPrivateAccessTag", "", "")
	request.Method = requests.GET
	return
}

// CreateListPolicesForPrivateAccessTagResponse creates a response to parse from ListPolicesForPrivateAccessTag response
func CreateListPolicesForPrivateAccessTagResponse() (response *ListPolicesForPrivateAccessTagResponse) {
	response = &ListPolicesForPrivateAccessTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
