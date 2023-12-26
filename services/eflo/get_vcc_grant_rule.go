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

// GetVccGrantRule invokes the eflo.GetVccGrantRule API synchronously
func (client *Client) GetVccGrantRule(request *GetVccGrantRuleRequest) (response *GetVccGrantRuleResponse, err error) {
	response = CreateGetVccGrantRuleResponse()
	err = client.DoAction(request, response)
	return
}

// GetVccGrantRuleWithChan invokes the eflo.GetVccGrantRule API asynchronously
func (client *Client) GetVccGrantRuleWithChan(request *GetVccGrantRuleRequest) (<-chan *GetVccGrantRuleResponse, <-chan error) {
	responseChan := make(chan *GetVccGrantRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVccGrantRule(request)
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

// GetVccGrantRuleWithCallback invokes the eflo.GetVccGrantRule API asynchronously
func (client *Client) GetVccGrantRuleWithCallback(request *GetVccGrantRuleRequest, callback func(response *GetVccGrantRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVccGrantRuleResponse
		var err error
		defer close(result)
		response, err = client.GetVccGrantRule(request)
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

// GetVccGrantRuleRequest is the request struct for api GetVccGrantRule
type GetVccGrantRuleRequest struct {
	*requests.RpcRequest
	ErId          string `position:"Body" name:"ErId"`
	GrantTenantId string `position:"Body" name:"GrantTenantId"`
	InstanceId    string `position:"Body" name:"InstanceId"`
	GrantRuleId   string `position:"Body" name:"GrantRuleId"`
}

// GetVccGrantRuleResponse is the response struct for api GetVccGrantRule
type GetVccGrantRuleResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateGetVccGrantRuleRequest creates a request to invoke GetVccGrantRule API
func CreateGetVccGrantRuleRequest() (request *GetVccGrantRuleRequest) {
	request = &GetVccGrantRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "GetVccGrantRule", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVccGrantRuleResponse creates a response to parse from GetVccGrantRule response
func CreateGetVccGrantRuleResponse() (response *GetVccGrantRuleResponse) {
	response = &GetVccGrantRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}