package companyreg

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

// RejectSolution invokes the companyreg.RejectSolution API synchronously
func (client *Client) RejectSolution(request *RejectSolutionRequest) (response *RejectSolutionResponse, err error) {
	response = CreateRejectSolutionResponse()
	err = client.DoAction(request, response)
	return
}

// RejectSolutionWithChan invokes the companyreg.RejectSolution API asynchronously
func (client *Client) RejectSolutionWithChan(request *RejectSolutionRequest) (<-chan *RejectSolutionResponse, <-chan error) {
	responseChan := make(chan *RejectSolutionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RejectSolution(request)
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

// RejectSolutionWithCallback invokes the companyreg.RejectSolution API asynchronously
func (client *Client) RejectSolutionWithCallback(request *RejectSolutionRequest, callback func(response *RejectSolutionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RejectSolutionResponse
		var err error
		defer close(result)
		response, err = client.RejectSolution(request)
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

// RejectSolutionRequest is the request struct for api RejectSolution
type RejectSolutionRequest struct {
	*requests.RpcRequest
	Note          string `position:"Query" name:"Note"`
	SolutionBizId string `position:"Query" name:"SolutionBizId"`
}

// RejectSolutionResponse is the response struct for api RejectSolution
type RejectSolutionResponse struct {
	*responses.BaseResponse
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateRejectSolutionRequest creates a request to invoke RejectSolution API
func CreateRejectSolutionRequest() (request *RejectSolutionRequest) {
	request = &RejectSolutionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "RejectSolution", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRejectSolutionResponse creates a response to parse from RejectSolution response
func CreateRejectSolutionResponse() (response *RejectSolutionResponse) {
	response = &RejectSolutionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
