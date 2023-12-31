package domain

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

// AcknowledgeTaskResult invokes the domain.AcknowledgeTaskResult API synchronously
func (client *Client) AcknowledgeTaskResult(request *AcknowledgeTaskResultRequest) (response *AcknowledgeTaskResultResponse, err error) {
	response = CreateAcknowledgeTaskResultResponse()
	err = client.DoAction(request, response)
	return
}

// AcknowledgeTaskResultWithChan invokes the domain.AcknowledgeTaskResult API asynchronously
func (client *Client) AcknowledgeTaskResultWithChan(request *AcknowledgeTaskResultRequest) (<-chan *AcknowledgeTaskResultResponse, <-chan error) {
	responseChan := make(chan *AcknowledgeTaskResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AcknowledgeTaskResult(request)
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

// AcknowledgeTaskResultWithCallback invokes the domain.AcknowledgeTaskResult API asynchronously
func (client *Client) AcknowledgeTaskResultWithCallback(request *AcknowledgeTaskResultRequest, callback func(response *AcknowledgeTaskResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AcknowledgeTaskResultResponse
		var err error
		defer close(result)
		response, err = client.AcknowledgeTaskResult(request)
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

// AcknowledgeTaskResultRequest is the request struct for api AcknowledgeTaskResult
type AcknowledgeTaskResultRequest struct {
	*requests.RpcRequest
	TaskDetailNo *[]string `position:"Query" name:"TaskDetailNo"  type:"Repeated"`
	UserClientIp string    `position:"Query" name:"UserClientIp"`
	Lang         string    `position:"Query" name:"Lang"`
}

// AcknowledgeTaskResultResponse is the response struct for api AcknowledgeTaskResult
type AcknowledgeTaskResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    int    `json:"Result" xml:"Result"`
}

// CreateAcknowledgeTaskResultRequest creates a request to invoke AcknowledgeTaskResult API
func CreateAcknowledgeTaskResultRequest() (request *AcknowledgeTaskResultRequest) {
	request = &AcknowledgeTaskResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "AcknowledgeTaskResult", "", "")
	request.Method = requests.POST
	return
}

// CreateAcknowledgeTaskResultResponse creates a response to parse from AcknowledgeTaskResult response
func CreateAcknowledgeTaskResultResponse() (response *AcknowledgeTaskResultResponse) {
	response = &AcknowledgeTaskResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
