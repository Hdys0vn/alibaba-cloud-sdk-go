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

// RejectUserSolution invokes the companyreg.RejectUserSolution API synchronously
func (client *Client) RejectUserSolution(request *RejectUserSolutionRequest) (response *RejectUserSolutionResponse, err error) {
	response = CreateRejectUserSolutionResponse()
	err = client.DoAction(request, response)
	return
}

// RejectUserSolutionWithChan invokes the companyreg.RejectUserSolution API asynchronously
func (client *Client) RejectUserSolutionWithChan(request *RejectUserSolutionRequest) (<-chan *RejectUserSolutionResponse, <-chan error) {
	responseChan := make(chan *RejectUserSolutionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RejectUserSolution(request)
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

// RejectUserSolutionWithCallback invokes the companyreg.RejectUserSolution API asynchronously
func (client *Client) RejectUserSolutionWithCallback(request *RejectUserSolutionRequest, callback func(response *RejectUserSolutionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RejectUserSolutionResponse
		var err error
		defer close(result)
		response, err = client.RejectUserSolution(request)
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

// RejectUserSolutionRequest is the request struct for api RejectUserSolution
type RejectUserSolutionRequest struct {
	*requests.RpcRequest
	BizType       string `position:"Query" name:"BizType"`
	Note          string `position:"Query" name:"Note"`
	SolutionBizId string `position:"Query" name:"SolutionBizId"`
}

// RejectUserSolutionResponse is the response struct for api RejectUserSolution
type RejectUserSolutionResponse struct {
	*responses.BaseResponse
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateRejectUserSolutionRequest creates a request to invoke RejectUserSolution API
func CreateRejectUserSolutionRequest() (request *RejectUserSolutionRequest) {
	request = &RejectUserSolutionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "RejectUserSolution", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRejectUserSolutionResponse creates a response to parse from RejectUserSolution response
func CreateRejectUserSolutionResponse() (response *RejectUserSolutionResponse) {
	response = &RejectUserSolutionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}