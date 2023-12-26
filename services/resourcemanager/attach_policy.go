package resourcemanager

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

// AttachPolicy invokes the resourcemanager.AttachPolicy API synchronously
func (client *Client) AttachPolicy(request *AttachPolicyRequest) (response *AttachPolicyResponse, err error) {
	response = CreateAttachPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// AttachPolicyWithChan invokes the resourcemanager.AttachPolicy API asynchronously
func (client *Client) AttachPolicyWithChan(request *AttachPolicyRequest) (<-chan *AttachPolicyResponse, <-chan error) {
	responseChan := make(chan *AttachPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachPolicy(request)
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

// AttachPolicyWithCallback invokes the resourcemanager.AttachPolicy API asynchronously
func (client *Client) AttachPolicyWithCallback(request *AttachPolicyRequest, callback func(response *AttachPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachPolicyResponse
		var err error
		defer close(result)
		response, err = client.AttachPolicy(request)
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

// AttachPolicyRequest is the request struct for api AttachPolicy
type AttachPolicyRequest struct {
	*requests.RpcRequest
	PolicyType      string `position:"Query" name:"PolicyType"`
	PrincipalType   string `position:"Query" name:"PrincipalType"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	PolicyName      string `position:"Query" name:"PolicyName"`
	PrincipalName   string `position:"Query" name:"PrincipalName"`
}

// AttachPolicyResponse is the response struct for api AttachPolicy
type AttachPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachPolicyRequest creates a request to invoke AttachPolicy API
func CreateAttachPolicyRequest() (request *AttachPolicyRequest) {
	request = &AttachPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "AttachPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateAttachPolicyResponse creates a response to parse from AttachPolicy response
func CreateAttachPolicyResponse() (response *AttachPolicyResponse) {
	response = &AttachPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
