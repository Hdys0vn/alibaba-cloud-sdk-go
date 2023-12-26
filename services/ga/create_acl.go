package ga

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

// CreateAcl invokes the ga.CreateAcl API synchronously
func (client *Client) CreateAcl(request *CreateAclRequest) (response *CreateAclResponse, err error) {
	response = CreateCreateAclResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAclWithChan invokes the ga.CreateAcl API asynchronously
func (client *Client) CreateAclWithChan(request *CreateAclRequest) (<-chan *CreateAclResponse, <-chan error) {
	responseChan := make(chan *CreateAclResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAcl(request)
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

// CreateAclWithCallback invokes the ga.CreateAcl API asynchronously
func (client *Client) CreateAclWithCallback(request *CreateAclRequest, callback func(response *CreateAclResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAclResponse
		var err error
		defer close(result)
		response, err = client.CreateAcl(request)
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

// CreateAclRequest is the request struct for api CreateAcl
type CreateAclRequest struct {
	*requests.RpcRequest
	DryRun           requests.Boolean       `position:"Query" name:"DryRun"`
	AclName          string                 `position:"Query" name:"AclName"`
	ClientToken      string                 `position:"Query" name:"ClientToken"`
	AclEntries       *[]CreateAclAclEntries `position:"Query" name:"AclEntries"  type:"Repeated"`
	AddressIPVersion string                 `position:"Query" name:"AddressIPVersion"`
}

// CreateAclAclEntries is a repeated param struct in CreateAclRequest
type CreateAclAclEntries struct {
	Entry            string `name:"Entry"`
	EntryDescription string `name:"EntryDescription"`
}

// CreateAclResponse is the response struct for api CreateAcl
type CreateAclResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	AclId     string `json:"AclId" xml:"AclId"`
}

// CreateCreateAclRequest creates a request to invoke CreateAcl API
func CreateCreateAclRequest() (request *CreateAclRequest) {
	request = &CreateAclRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "CreateAcl", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAclResponse creates a response to parse from CreateAcl response
func CreateCreateAclResponse() (response *CreateAclResponse) {
	response = &CreateAclResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}