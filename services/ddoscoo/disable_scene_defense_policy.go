package ddoscoo

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

// DisableSceneDefensePolicy invokes the ddoscoo.DisableSceneDefensePolicy API synchronously
func (client *Client) DisableSceneDefensePolicy(request *DisableSceneDefensePolicyRequest) (response *DisableSceneDefensePolicyResponse, err error) {
	response = CreateDisableSceneDefensePolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DisableSceneDefensePolicyWithChan invokes the ddoscoo.DisableSceneDefensePolicy API asynchronously
func (client *Client) DisableSceneDefensePolicyWithChan(request *DisableSceneDefensePolicyRequest) (<-chan *DisableSceneDefensePolicyResponse, <-chan error) {
	responseChan := make(chan *DisableSceneDefensePolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableSceneDefensePolicy(request)
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

// DisableSceneDefensePolicyWithCallback invokes the ddoscoo.DisableSceneDefensePolicy API asynchronously
func (client *Client) DisableSceneDefensePolicyWithCallback(request *DisableSceneDefensePolicyRequest, callback func(response *DisableSceneDefensePolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableSceneDefensePolicyResponse
		var err error
		defer close(result)
		response, err = client.DisableSceneDefensePolicy(request)
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

// DisableSceneDefensePolicyRequest is the request struct for api DisableSceneDefensePolicy
type DisableSceneDefensePolicyRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	PolicyId string `position:"Query" name:"PolicyId"`
}

// DisableSceneDefensePolicyResponse is the response struct for api DisableSceneDefensePolicy
type DisableSceneDefensePolicyResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableSceneDefensePolicyRequest creates a request to invoke DisableSceneDefensePolicy API
func CreateDisableSceneDefensePolicyRequest() (request *DisableSceneDefensePolicyRequest) {
	request = &DisableSceneDefensePolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DisableSceneDefensePolicy", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableSceneDefensePolicyResponse creates a response to parse from DisableSceneDefensePolicy response
func CreateDisableSceneDefensePolicyResponse() (response *DisableSceneDefensePolicyResponse) {
	response = &DisableSceneDefensePolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}