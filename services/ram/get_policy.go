package ram

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

// GetPolicy invokes the ram.GetPolicy API synchronously
func (client *Client) GetPolicy(request *GetPolicyRequest) (response *GetPolicyResponse, err error) {
	response = CreateGetPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// GetPolicyWithChan invokes the ram.GetPolicy API asynchronously
func (client *Client) GetPolicyWithChan(request *GetPolicyRequest) (<-chan *GetPolicyResponse, <-chan error) {
	responseChan := make(chan *GetPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPolicy(request)
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

// GetPolicyWithCallback invokes the ram.GetPolicy API asynchronously
func (client *Client) GetPolicyWithCallback(request *GetPolicyRequest, callback func(response *GetPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPolicyResponse
		var err error
		defer close(result)
		response, err = client.GetPolicy(request)
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

// GetPolicyRequest is the request struct for api GetPolicy
type GetPolicyRequest struct {
	*requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

// GetPolicyResponse is the response struct for api GetPolicy
type GetPolicyResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	Policy               Policy               `json:"Policy" xml:"Policy"`
	DefaultPolicyVersion DefaultPolicyVersion `json:"DefaultPolicyVersion" xml:"DefaultPolicyVersion"`
}

// CreateGetPolicyRequest creates a request to invoke GetPolicy API
func CreateGetPolicyRequest() (request *GetPolicyRequest) {
	request = &GetPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "GetPolicy", "Ram", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetPolicyResponse creates a response to parse from GetPolicy response
func CreateGetPolicyResponse() (response *GetPolicyResponse) {
	response = &GetPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
