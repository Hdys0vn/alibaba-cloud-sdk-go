package nas

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

// CreateAccessRule invokes the nas.CreateAccessRule API synchronously
func (client *Client) CreateAccessRule(request *CreateAccessRuleRequest) (response *CreateAccessRuleResponse, err error) {
	response = CreateCreateAccessRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAccessRuleWithChan invokes the nas.CreateAccessRule API asynchronously
func (client *Client) CreateAccessRuleWithChan(request *CreateAccessRuleRequest) (<-chan *CreateAccessRuleResponse, <-chan error) {
	responseChan := make(chan *CreateAccessRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAccessRule(request)
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

// CreateAccessRuleWithCallback invokes the nas.CreateAccessRule API asynchronously
func (client *Client) CreateAccessRuleWithCallback(request *CreateAccessRuleRequest, callback func(response *CreateAccessRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAccessRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateAccessRule(request)
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

// CreateAccessRuleRequest is the request struct for api CreateAccessRule
type CreateAccessRuleRequest struct {
	*requests.RpcRequest
	RWAccessType     string           `position:"Query" name:"RWAccessType"`
	UserAccessType   string           `position:"Query" name:"UserAccessType"`
	FileSystemType   string           `position:"Query" name:"FileSystemType"`
	Ipv6SourceCidrIp string           `position:"Query" name:"Ipv6SourceCidrIp"`
	SourceCidrIp     string           `position:"Query" name:"SourceCidrIp"`
	Priority         requests.Integer `position:"Query" name:"Priority"`
	AccessGroupName  string           `position:"Query" name:"AccessGroupName"`
}

// CreateAccessRuleResponse is the response struct for api CreateAccessRule
type CreateAccessRuleResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	AccessRuleId string `json:"AccessRuleId" xml:"AccessRuleId"`
}

// CreateCreateAccessRuleRequest creates a request to invoke CreateAccessRule API
func CreateCreateAccessRuleRequest() (request *CreateAccessRuleRequest) {
	request = &CreateAccessRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "CreateAccessRule", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAccessRuleResponse creates a response to parse from CreateAccessRule response
func CreateCreateAccessRuleResponse() (response *CreateAccessRuleResponse) {
	response = &CreateAccessRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
