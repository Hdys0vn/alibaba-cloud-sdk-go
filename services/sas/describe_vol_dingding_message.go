package sas

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

// DescribeVolDingdingMessage invokes the sas.DescribeVolDingdingMessage API synchronously
func (client *Client) DescribeVolDingdingMessage(request *DescribeVolDingdingMessageRequest) (response *DescribeVolDingdingMessageResponse, err error) {
	response = CreateDescribeVolDingdingMessageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVolDingdingMessageWithChan invokes the sas.DescribeVolDingdingMessage API asynchronously
func (client *Client) DescribeVolDingdingMessageWithChan(request *DescribeVolDingdingMessageRequest) (<-chan *DescribeVolDingdingMessageResponse, <-chan error) {
	responseChan := make(chan *DescribeVolDingdingMessageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVolDingdingMessage(request)
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

// DescribeVolDingdingMessageWithCallback invokes the sas.DescribeVolDingdingMessage API asynchronously
func (client *Client) DescribeVolDingdingMessageWithCallback(request *DescribeVolDingdingMessageRequest, callback func(response *DescribeVolDingdingMessageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVolDingdingMessageResponse
		var err error
		defer close(result)
		response, err = client.DescribeVolDingdingMessage(request)
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

// DescribeVolDingdingMessageRequest is the request struct for api DescribeVolDingdingMessage
type DescribeVolDingdingMessageRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
}

// DescribeVolDingdingMessageResponse is the response struct for api DescribeVolDingdingMessage
type DescribeVolDingdingMessageResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DingdingUrl string `json:"DingdingUrl" xml:"DingdingUrl"`
}

// CreateDescribeVolDingdingMessageRequest creates a request to invoke DescribeVolDingdingMessage API
func CreateDescribeVolDingdingMessageRequest() (request *DescribeVolDingdingMessageRequest) {
	request = &DescribeVolDingdingMessageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeVolDingdingMessage", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVolDingdingMessageResponse creates a response to parse from DescribeVolDingdingMessage response
func CreateDescribeVolDingdingMessageResponse() (response *DescribeVolDingdingMessageResponse) {
	response = &DescribeVolDingdingMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}