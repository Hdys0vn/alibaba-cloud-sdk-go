package vod

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

// DeleteStream invokes the vod.DeleteStream API synchronously
func (client *Client) DeleteStream(request *DeleteStreamRequest) (response *DeleteStreamResponse, err error) {
	response = CreateDeleteStreamResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteStreamWithChan invokes the vod.DeleteStream API asynchronously
func (client *Client) DeleteStreamWithChan(request *DeleteStreamRequest) (<-chan *DeleteStreamResponse, <-chan error) {
	responseChan := make(chan *DeleteStreamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteStream(request)
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

// DeleteStreamWithCallback invokes the vod.DeleteStream API asynchronously
func (client *Client) DeleteStreamWithCallback(request *DeleteStreamRequest, callback func(response *DeleteStreamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteStreamResponse
		var err error
		defer close(result)
		response, err = client.DeleteStream(request)
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

// DeleteStreamRequest is the request struct for api DeleteStream
type DeleteStreamRequest struct {
	*requests.RpcRequest
	VideoId string `position:"Query" name:"VideoId"`
	JobIds  string `position:"Query" name:"JobIds"`
}

// DeleteStreamResponse is the response struct for api DeleteStream
type DeleteStreamResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteStreamRequest creates a request to invoke DeleteStream API
func CreateDeleteStreamRequest() (request *DeleteStreamRequest) {
	request = &DeleteStreamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DeleteStream", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteStreamResponse creates a response to parse from DeleteStream response
func CreateDeleteStreamResponse() (response *DeleteStreamResponse) {
	response = &DeleteStreamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
