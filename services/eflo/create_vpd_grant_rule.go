package eflo

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

// CreateVpdGrantRule invokes the eflo.CreateVpdGrantRule API synchronously
func (client *Client) CreateVpdGrantRule(request *CreateVpdGrantRuleRequest) (response *CreateVpdGrantRuleResponse, err error) {
	response = CreateCreateVpdGrantRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVpdGrantRuleWithChan invokes the eflo.CreateVpdGrantRule API asynchronously
func (client *Client) CreateVpdGrantRuleWithChan(request *CreateVpdGrantRuleRequest) (<-chan *CreateVpdGrantRuleResponse, <-chan error) {
	responseChan := make(chan *CreateVpdGrantRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVpdGrantRule(request)
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

// CreateVpdGrantRuleWithCallback invokes the eflo.CreateVpdGrantRule API asynchronously
func (client *Client) CreateVpdGrantRuleWithCallback(request *CreateVpdGrantRuleRequest, callback func(response *CreateVpdGrantRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVpdGrantRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateVpdGrantRule(request)
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

// CreateVpdGrantRuleRequest is the request struct for api CreateVpdGrantRule
type CreateVpdGrantRuleRequest struct {
	*requests.RpcRequest
	ErId          string `position:"Body" name:"ErId"`
	GrantTenantId string `position:"Body" name:"GrantTenantId"`
	InstanceId    string `position:"Body" name:"InstanceId"`
}

// CreateVpdGrantRuleResponse is the response struct for api CreateVpdGrantRule
type CreateVpdGrantRuleResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateCreateVpdGrantRuleRequest creates a request to invoke CreateVpdGrantRule API
func CreateCreateVpdGrantRuleRequest() (request *CreateVpdGrantRuleRequest) {
	request = &CreateVpdGrantRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "CreateVpdGrantRule", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVpdGrantRuleResponse creates a response to parse from CreateVpdGrantRule response
func CreateCreateVpdGrantRuleResponse() (response *CreateVpdGrantRuleResponse) {
	response = &CreateVpdGrantRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
