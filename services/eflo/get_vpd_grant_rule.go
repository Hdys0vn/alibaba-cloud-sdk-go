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

// GetVpdGrantRule invokes the eflo.GetVpdGrantRule API synchronously
func (client *Client) GetVpdGrantRule(request *GetVpdGrantRuleRequest) (response *GetVpdGrantRuleResponse, err error) {
	response = CreateGetVpdGrantRuleResponse()
	err = client.DoAction(request, response)
	return
}

// GetVpdGrantRuleWithChan invokes the eflo.GetVpdGrantRule API asynchronously
func (client *Client) GetVpdGrantRuleWithChan(request *GetVpdGrantRuleRequest) (<-chan *GetVpdGrantRuleResponse, <-chan error) {
	responseChan := make(chan *GetVpdGrantRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVpdGrantRule(request)
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

// GetVpdGrantRuleWithCallback invokes the eflo.GetVpdGrantRule API asynchronously
func (client *Client) GetVpdGrantRuleWithCallback(request *GetVpdGrantRuleRequest, callback func(response *GetVpdGrantRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVpdGrantRuleResponse
		var err error
		defer close(result)
		response, err = client.GetVpdGrantRule(request)
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

// GetVpdGrantRuleRequest is the request struct for api GetVpdGrantRule
type GetVpdGrantRuleRequest struct {
	*requests.RpcRequest
	ErId          string `position:"Body" name:"ErId"`
	GrantTenantId string `position:"Body" name:"GrantTenantId"`
	InstanceId    string `position:"Body" name:"InstanceId"`
	GrantRuleId   string `position:"Body" name:"GrantRuleId"`
}

// GetVpdGrantRuleResponse is the response struct for api GetVpdGrantRule
type GetVpdGrantRuleResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateGetVpdGrantRuleRequest creates a request to invoke GetVpdGrantRule API
func CreateGetVpdGrantRuleRequest() (request *GetVpdGrantRuleRequest) {
	request = &GetVpdGrantRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "GetVpdGrantRule", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVpdGrantRuleResponse creates a response to parse from GetVpdGrantRule response
func CreateGetVpdGrantRuleResponse() (response *GetVpdGrantRuleResponse) {
	response = &GetVpdGrantRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
