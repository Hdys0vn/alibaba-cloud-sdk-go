package csas

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

// CreatePrivateAccessPolicy invokes the csas.CreatePrivateAccessPolicy API synchronously
func (client *Client) CreatePrivateAccessPolicy(request *CreatePrivateAccessPolicyRequest) (response *CreatePrivateAccessPolicyResponse, err error) {
	response = CreateCreatePrivateAccessPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePrivateAccessPolicyWithChan invokes the csas.CreatePrivateAccessPolicy API asynchronously
func (client *Client) CreatePrivateAccessPolicyWithChan(request *CreatePrivateAccessPolicyRequest) (<-chan *CreatePrivateAccessPolicyResponse, <-chan error) {
	responseChan := make(chan *CreatePrivateAccessPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePrivateAccessPolicy(request)
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

// CreatePrivateAccessPolicyWithCallback invokes the csas.CreatePrivateAccessPolicy API asynchronously
func (client *Client) CreatePrivateAccessPolicyWithCallback(request *CreatePrivateAccessPolicyRequest, callback func(response *CreatePrivateAccessPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePrivateAccessPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreatePrivateAccessPolicy(request)
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

// CreatePrivateAccessPolicyRequest is the request struct for api CreatePrivateAccessPolicy
type CreatePrivateAccessPolicyRequest struct {
	*requests.RpcRequest
	Description          string                                           `position:"Body" name:"Description"`
	SourceIp             string                                           `position:"Query" name:"SourceIp"`
	CustomUserAttributes *[]CreatePrivateAccessPolicyCustomUserAttributes `position:"Body" name:"CustomUserAttributes"  type:"Repeated"`
	TagIds               *[]string                                        `position:"Body" name:"TagIds"  type:"Repeated"`
	UserGroupIds         *[]string                                        `position:"Body" name:"UserGroupIds"  type:"Repeated"`
	PolicyAction         string                                           `position:"Body" name:"PolicyAction"`
	Priority             requests.Integer                                 `position:"Body" name:"Priority"`
	ApplicationIds       *[]string                                        `position:"Body" name:"ApplicationIds"  type:"Repeated"`
	UserGroupMode        string                                           `position:"Body" name:"UserGroupMode"`
	Name                 string                                           `position:"Body" name:"Name"`
	ApplicationType      string                                           `position:"Body" name:"ApplicationType"`
	Status               string                                           `position:"Body" name:"Status"`
}

// CreatePrivateAccessPolicyCustomUserAttributes is a repeated param struct in CreatePrivateAccessPolicyRequest
type CreatePrivateAccessPolicyCustomUserAttributes struct {
	UserGroupType string `name:"UserGroupType"`
	IdpId         string `name:"IdpId"`
	Value         string `name:"Value"`
	Relation      string `name:"Relation"`
}

// CreatePrivateAccessPolicyResponse is the response struct for api CreatePrivateAccessPolicy
type CreatePrivateAccessPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	PolicyId  string `json:"PolicyId" xml:"PolicyId"`
}

// CreateCreatePrivateAccessPolicyRequest creates a request to invoke CreatePrivateAccessPolicy API
func CreateCreatePrivateAccessPolicyRequest() (request *CreatePrivateAccessPolicyRequest) {
	request = &CreatePrivateAccessPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "CreatePrivateAccessPolicy", "", "")
	request.Method = requests.POST
	return
}

// CreateCreatePrivateAccessPolicyResponse creates a response to parse from CreatePrivateAccessPolicy response
func CreateCreatePrivateAccessPolicyResponse() (response *CreatePrivateAccessPolicyResponse) {
	response = &CreatePrivateAccessPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}