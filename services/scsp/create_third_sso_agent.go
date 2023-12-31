package scsp

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

// CreateThirdSsoAgent invokes the scsp.CreateThirdSsoAgent API synchronously
func (client *Client) CreateThirdSsoAgent(request *CreateThirdSsoAgentRequest) (response *CreateThirdSsoAgentResponse, err error) {
	response = CreateCreateThirdSsoAgentResponse()
	err = client.DoAction(request, response)
	return
}

// CreateThirdSsoAgentWithChan invokes the scsp.CreateThirdSsoAgent API asynchronously
func (client *Client) CreateThirdSsoAgentWithChan(request *CreateThirdSsoAgentRequest) (<-chan *CreateThirdSsoAgentResponse, <-chan error) {
	responseChan := make(chan *CreateThirdSsoAgentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateThirdSsoAgent(request)
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

// CreateThirdSsoAgentWithCallback invokes the scsp.CreateThirdSsoAgent API asynchronously
func (client *Client) CreateThirdSsoAgentWithCallback(request *CreateThirdSsoAgentRequest, callback func(response *CreateThirdSsoAgentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateThirdSsoAgentResponse
		var err error
		defer close(result)
		response, err = client.CreateThirdSsoAgent(request)
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

// CreateThirdSsoAgentRequest is the request struct for api CreateThirdSsoAgent
type CreateThirdSsoAgentRequest struct {
	*requests.RpcRequest
	ClientToken   string                              `position:"Body"`
	InstanceId    string                              `position:"Body"`
	ClientId      string                              `position:"Body"`
	AccountId     string                              `position:"Body"`
	AccountName   string                              `position:"Body"`
	DisplayName   string                              `position:"Body"`
	SkillGroupIds *[]CreateThirdSsoAgentSkillGroupIds `position:"Body" name:"SkillGroupIds"  type:"Repeated"`
	RoleIds       *[]CreateThirdSsoAgentRoleIds       `position:"Body" name:"RoleIds"  type:"Repeated"`
}

// CreateThirdSsoAgentSkillGroupIds is a repeated param struct in CreateThirdSsoAgentRequest
type CreateThirdSsoAgentSkillGroupIds struct {
	SkillGroupIds string `name:"SkillGroupIds"`
}

// CreateThirdSsoAgentRoleIds is a repeated param struct in CreateThirdSsoAgentRequest
type CreateThirdSsoAgentRoleIds struct {
	RoleIds string `name:"RoleIds"`
}

// CreateThirdSsoAgentResponse is the response struct for api CreateThirdSsoAgent
type CreateThirdSsoAgentResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int64  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           int64  `json:"Data" xml:"Data"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateCreateThirdSsoAgentRequest creates a request to invoke CreateThirdSsoAgent API
func CreateCreateThirdSsoAgentRequest() (request *CreateThirdSsoAgentRequest) {
	request = &CreateThirdSsoAgentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scsp", "2020-07-02", "CreateThirdSsoAgent", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateThirdSsoAgentResponse creates a response to parse from CreateThirdSsoAgent response
func CreateCreateThirdSsoAgentResponse() (response *CreateThirdSsoAgentResponse) {
	response = &CreateThirdSsoAgentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
