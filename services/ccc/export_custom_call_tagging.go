package ccc

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

// ExportCustomCallTagging invokes the ccc.ExportCustomCallTagging API synchronously
func (client *Client) ExportCustomCallTagging(request *ExportCustomCallTaggingRequest) (response *ExportCustomCallTaggingResponse, err error) {
	response = CreateExportCustomCallTaggingResponse()
	err = client.DoAction(request, response)
	return
}

// ExportCustomCallTaggingWithChan invokes the ccc.ExportCustomCallTagging API asynchronously
func (client *Client) ExportCustomCallTaggingWithChan(request *ExportCustomCallTaggingRequest) (<-chan *ExportCustomCallTaggingResponse, <-chan error) {
	responseChan := make(chan *ExportCustomCallTaggingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportCustomCallTagging(request)
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

// ExportCustomCallTaggingWithCallback invokes the ccc.ExportCustomCallTagging API asynchronously
func (client *Client) ExportCustomCallTaggingWithCallback(request *ExportCustomCallTaggingRequest, callback func(response *ExportCustomCallTaggingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportCustomCallTaggingResponse
		var err error
		defer close(result)
		response, err = client.ExportCustomCallTagging(request)
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

// ExportCustomCallTaggingRequest is the request struct for api ExportCustomCallTagging
type ExportCustomCallTaggingRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ExportCustomCallTaggingResponse is the response struct for api ExportCustomCallTagging
type ExportCustomCallTaggingResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	Data           string `json:"Data" xml:"Data"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateExportCustomCallTaggingRequest creates a request to invoke ExportCustomCallTagging API
func CreateExportCustomCallTaggingRequest() (request *ExportCustomCallTaggingRequest) {
	request = &ExportCustomCallTaggingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ExportCustomCallTagging", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportCustomCallTaggingResponse creates a response to parse from ExportCustomCallTagging response
func CreateExportCustomCallTaggingResponse() (response *ExportCustomCallTaggingResponse) {
	response = &ExportCustomCallTaggingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
