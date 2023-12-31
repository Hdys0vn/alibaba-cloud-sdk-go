package tag

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

// DetachPolicy invokes the tag.DetachPolicy API synchronously
func (client *Client) DetachPolicy(request *DetachPolicyRequest) (response *DetachPolicyResponse, err error) {
	response = CreateDetachPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DetachPolicyWithChan invokes the tag.DetachPolicy API asynchronously
func (client *Client) DetachPolicyWithChan(request *DetachPolicyRequest) (<-chan *DetachPolicyResponse, <-chan error) {
	responseChan := make(chan *DetachPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachPolicy(request)
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

// DetachPolicyWithCallback invokes the tag.DetachPolicy API asynchronously
func (client *Client) DetachPolicyWithCallback(request *DetachPolicyRequest, callback func(response *DetachPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachPolicyResponse
		var err error
		defer close(result)
		response, err = client.DetachPolicy(request)
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

// DetachPolicyRequest is the request struct for api DetachPolicy
type DetachPolicyRequest struct {
	*requests.RpcRequest
	TargetId             string           `position:"Query" name:"TargetId"`
	TargetType           string           `position:"Query" name:"TargetType"`
	PolicyId             string           `position:"Query" name:"PolicyId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DetachPolicyResponse is the response struct for api DetachPolicy
type DetachPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachPolicyRequest creates a request to invoke DetachPolicy API
func CreateDetachPolicyRequest() (request *DetachPolicyRequest) {
	request = &DetachPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Tag", "2018-08-28", "DetachPolicy", "tag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachPolicyResponse creates a response to parse from DetachPolicy response
func CreateDetachPolicyResponse() (response *DetachPolicyResponse) {
	response = &DetachPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
