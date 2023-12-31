package smartag

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

// DeleteACLRule invokes the smartag.DeleteACLRule API synchronously
func (client *Client) DeleteACLRule(request *DeleteACLRuleRequest) (response *DeleteACLRuleResponse, err error) {
	response = CreateDeleteACLRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteACLRuleWithChan invokes the smartag.DeleteACLRule API asynchronously
func (client *Client) DeleteACLRuleWithChan(request *DeleteACLRuleRequest) (<-chan *DeleteACLRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteACLRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteACLRule(request)
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

// DeleteACLRuleWithCallback invokes the smartag.DeleteACLRule API asynchronously
func (client *Client) DeleteACLRuleWithCallback(request *DeleteACLRuleRequest, callback func(response *DeleteACLRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteACLRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteACLRule(request)
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

// DeleteACLRuleRequest is the request struct for api DeleteACLRule
type DeleteACLRuleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AcrId                string           `position:"Query" name:"AcrId"`
}

// DeleteACLRuleResponse is the response struct for api DeleteACLRule
type DeleteACLRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteACLRuleRequest creates a request to invoke DeleteACLRule API
func CreateDeleteACLRuleRequest() (request *DeleteACLRuleRequest) {
	request = &DeleteACLRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DeleteACLRule", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteACLRuleResponse creates a response to parse from DeleteACLRule response
func CreateDeleteACLRuleResponse() (response *DeleteACLRuleResponse) {
	response = &DeleteACLRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
