package cloudfw

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

// DeleteControlPolicy invokes the cloudfw.DeleteControlPolicy API synchronously
func (client *Client) DeleteControlPolicy(request *DeleteControlPolicyRequest) (response *DeleteControlPolicyResponse, err error) {
	response = CreateDeleteControlPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteControlPolicyWithChan invokes the cloudfw.DeleteControlPolicy API asynchronously
func (client *Client) DeleteControlPolicyWithChan(request *DeleteControlPolicyRequest) (<-chan *DeleteControlPolicyResponse, <-chan error) {
	responseChan := make(chan *DeleteControlPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteControlPolicy(request)
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

// DeleteControlPolicyWithCallback invokes the cloudfw.DeleteControlPolicy API asynchronously
func (client *Client) DeleteControlPolicyWithCallback(request *DeleteControlPolicyRequest, callback func(response *DeleteControlPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteControlPolicyResponse
		var err error
		defer close(result)
		response, err = client.DeleteControlPolicy(request)
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

// DeleteControlPolicyRequest is the request struct for api DeleteControlPolicy
type DeleteControlPolicyRequest struct {
	*requests.RpcRequest
	AclUuid   string `position:"Query" name:"AclUuid"`
	SourceIp  string `position:"Query" name:"SourceIp"`
	Lang      string `position:"Query" name:"Lang"`
	Direction string `position:"Query" name:"Direction"`
}

// DeleteControlPolicyResponse is the response struct for api DeleteControlPolicy
type DeleteControlPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteControlPolicyRequest creates a request to invoke DeleteControlPolicy API
func CreateDeleteControlPolicyRequest() (request *DeleteControlPolicyRequest) {
	request = &DeleteControlPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DeleteControlPolicy", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteControlPolicyResponse creates a response to parse from DeleteControlPolicy response
func CreateDeleteControlPolicyResponse() (response *DeleteControlPolicyResponse) {
	response = &DeleteControlPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
