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

// SubmitSolution invokes the companyreg.SubmitSolution API synchronously
func (client *Client) SubmitSolution(request *SubmitSolutionRequest) (response *SubmitSolutionResponse, err error) {
	response = CreateSubmitSolutionResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitSolutionWithChan invokes the companyreg.SubmitSolution API asynchronously
func (client *Client) SubmitSolutionWithChan(request *SubmitSolutionRequest) (<-chan *SubmitSolutionResponse, <-chan error) {
	responseChan := make(chan *SubmitSolutionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitSolution(request)
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

// SubmitSolutionWithCallback invokes the companyreg.SubmitSolution API asynchronously
func (client *Client) SubmitSolutionWithCallback(request *SubmitSolutionRequest, callback func(response *SubmitSolutionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitSolutionResponse
		var err error
		defer close(result)
		response, err = client.SubmitSolution(request)
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

// SubmitSolutionRequest is the request struct for api SubmitSolution
type SubmitSolutionRequest struct {
	*requests.RpcRequest
	BizType        string `position:"Query" name:"BizType"`
	Solution       string `position:"Query" name:"Solution"`
	IntentionBizId string `position:"Query" name:"IntentionBizId"`
	UserId         string `position:"Query" name:"UserId"`
}

// SubmitSolutionResponse is the response struct for api SubmitSolution
type SubmitSolutionResponse struct {
	*responses.BaseResponse
	SolutionBizId string `json:"SolutionBizId" xml:"SolutionBizId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Success       bool   `json:"Success" xml:"Success"`
	ErrorCode     string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg      string `json:"ErrorMsg" xml:"ErrorMsg"`
	ConfirmUrl    string `json:"ConfirmUrl" xml:"ConfirmUrl"`
}

// CreateSubmitSolutionRequest creates a request to invoke SubmitSolution API
func CreateSubmitSolutionRequest() (request *SubmitSolutionRequest) {
	request = &SubmitSolutionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "SubmitSolution", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitSolutionResponse creates a response to parse from SubmitSolution response
func CreateSubmitSolutionResponse() (response *SubmitSolutionResponse) {
	response = &SubmitSolutionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}