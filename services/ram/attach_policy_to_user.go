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

// AttachPolicyToUser invokes the ram.AttachPolicyToUser API synchronously
func (client *Client) AttachPolicyToUser(request *AttachPolicyToUserRequest) (response *AttachPolicyToUserResponse, err error) {
	response = CreateAttachPolicyToUserResponse()
	err = client.DoAction(request, response)
	return
}

// AttachPolicyToUserWithChan invokes the ram.AttachPolicyToUser API asynchronously
func (client *Client) AttachPolicyToUserWithChan(request *AttachPolicyToUserRequest) (<-chan *AttachPolicyToUserResponse, <-chan error) {
	responseChan := make(chan *AttachPolicyToUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachPolicyToUser(request)
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

// AttachPolicyToUserWithCallback invokes the ram.AttachPolicyToUser API asynchronously
func (client *Client) AttachPolicyToUserWithCallback(request *AttachPolicyToUserRequest, callback func(response *AttachPolicyToUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachPolicyToUserResponse
		var err error
		defer close(result)
		response, err = client.AttachPolicyToUser(request)
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

// AttachPolicyToUserRequest is the request struct for api AttachPolicyToUser
type AttachPolicyToUserRequest struct {
	*requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	UserName   string `position:"Query" name:"UserName"`
}

// AttachPolicyToUserResponse is the response struct for api AttachPolicyToUser
type AttachPolicyToUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachPolicyToUserRequest creates a request to invoke AttachPolicyToUser API
func CreateAttachPolicyToUserRequest() (request *AttachPolicyToUserRequest) {
	request = &AttachPolicyToUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "AttachPolicyToUser", "Ram", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachPolicyToUserResponse creates a response to parse from AttachPolicyToUser response
func CreateAttachPolicyToUserResponse() (response *AttachPolicyToUserResponse) {
	response = &AttachPolicyToUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
