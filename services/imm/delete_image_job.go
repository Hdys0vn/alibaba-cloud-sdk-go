package imm

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

// DeleteImageJob invokes the imm.DeleteImageJob API synchronously
func (client *Client) DeleteImageJob(request *DeleteImageJobRequest) (response *DeleteImageJobResponse, err error) {
	response = CreateDeleteImageJobResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteImageJobWithChan invokes the imm.DeleteImageJob API asynchronously
func (client *Client) DeleteImageJobWithChan(request *DeleteImageJobRequest) (<-chan *DeleteImageJobResponse, <-chan error) {
	responseChan := make(chan *DeleteImageJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteImageJob(request)
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

// DeleteImageJobWithCallback invokes the imm.DeleteImageJob API asynchronously
func (client *Client) DeleteImageJobWithCallback(request *DeleteImageJobRequest, callback func(response *DeleteImageJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteImageJobResponse
		var err error
		defer close(result)
		response, err = client.DeleteImageJob(request)
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

// DeleteImageJobRequest is the request struct for api DeleteImageJob
type DeleteImageJobRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	JobId   string `position:"Query" name:"JobId"`
	JobType string `position:"Query" name:"JobType"`
}

// DeleteImageJobResponse is the response struct for api DeleteImageJob
type DeleteImageJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteImageJobRequest creates a request to invoke DeleteImageJob API
func CreateDeleteImageJobRequest() (request *DeleteImageJobRequest) {
	request = &DeleteImageJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "DeleteImageJob", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteImageJobResponse creates a response to parse from DeleteImageJob response
func CreateDeleteImageJobResponse() (response *DeleteImageJobResponse) {
	response = &DeleteImageJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
