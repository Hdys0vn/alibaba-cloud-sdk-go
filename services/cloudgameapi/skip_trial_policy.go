package cloudgameapi

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

// SkipTrialPolicy invokes the cloudgameapi.SkipTrialPolicy API synchronously
func (client *Client) SkipTrialPolicy(request *SkipTrialPolicyRequest) (response *SkipTrialPolicyResponse, err error) {
	response = CreateSkipTrialPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// SkipTrialPolicyWithChan invokes the cloudgameapi.SkipTrialPolicy API asynchronously
func (client *Client) SkipTrialPolicyWithChan(request *SkipTrialPolicyRequest) (<-chan *SkipTrialPolicyResponse, <-chan error) {
	responseChan := make(chan *SkipTrialPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SkipTrialPolicy(request)
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

// SkipTrialPolicyWithCallback invokes the cloudgameapi.SkipTrialPolicy API asynchronously
func (client *Client) SkipTrialPolicyWithCallback(request *SkipTrialPolicyRequest, callback func(response *SkipTrialPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SkipTrialPolicyResponse
		var err error
		defer close(result)
		response, err = client.SkipTrialPolicy(request)
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

// SkipTrialPolicyRequest is the request struct for api SkipTrialPolicy
type SkipTrialPolicyRequest struct {
	*requests.RpcRequest
	GameSessionId string `position:"Query" name:"GameSessionId"`
}

// SkipTrialPolicyResponse is the response struct for api SkipTrialPolicy
type SkipTrialPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSkipTrialPolicyRequest creates a request to invoke SkipTrialPolicy API
func CreateSkipTrialPolicyRequest() (request *SkipTrialPolicyRequest) {
	request = &SkipTrialPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudGameAPI", "2020-07-28", "SkipTrialPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateSkipTrialPolicyResponse creates a response to parse from SkipTrialPolicy response
func CreateSkipTrialPolicyResponse() (response *SkipTrialPolicyResponse) {
	response = &SkipTrialPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
