package aegis

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

// DescribeSecureSuggestion invokes the aegis.DescribeSecureSuggestion API synchronously
// api document: https://help.aliyun.com/api/aegis/describesecuresuggestion.html
func (client *Client) DescribeSecureSuggestion(request *DescribeSecureSuggestionRequest) (response *DescribeSecureSuggestionResponse, err error) {
	response = CreateDescribeSecureSuggestionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSecureSuggestionWithChan invokes the aegis.DescribeSecureSuggestion API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesecuresuggestion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSecureSuggestionWithChan(request *DescribeSecureSuggestionRequest) (<-chan *DescribeSecureSuggestionResponse, <-chan error) {
	responseChan := make(chan *DescribeSecureSuggestionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecureSuggestion(request)
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

// DescribeSecureSuggestionWithCallback invokes the aegis.DescribeSecureSuggestion API asynchronously
// api document: https://help.aliyun.com/api/aegis/describesecuresuggestion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSecureSuggestionWithCallback(request *DescribeSecureSuggestionRequest, callback func(response *DescribeSecureSuggestionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecureSuggestionResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecureSuggestion(request)
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

// DescribeSecureSuggestionRequest is the request struct for api DescribeSecureSuggestion
type DescribeSecureSuggestionRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
}

// DescribeSecureSuggestionResponse is the response struct for api DescribeSecureSuggestion
type DescribeSecureSuggestionResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	TotalCount  int          `json:"TotalCount" xml:"TotalCount"`
	Suggestions []Suggestion `json:"Suggestions" xml:"Suggestions"`
}

// CreateDescribeSecureSuggestionRequest creates a request to invoke DescribeSecureSuggestion API
func CreateDescribeSecureSuggestionRequest() (request *DescribeSecureSuggestionRequest) {
	request = &DescribeSecureSuggestionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeSecureSuggestion", "vipaegis", "openAPI")
	return
}

// CreateDescribeSecureSuggestionResponse creates a response to parse from DescribeSecureSuggestion response
func CreateDescribeSecureSuggestionResponse() (response *DescribeSecureSuggestionResponse) {
	response = &DescribeSecureSuggestionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
