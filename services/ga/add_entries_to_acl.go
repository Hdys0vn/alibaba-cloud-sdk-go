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

// AddEntriesToAcl invokes the ga.AddEntriesToAcl API synchronously
func (client *Client) AddEntriesToAcl(request *AddEntriesToAclRequest) (response *AddEntriesToAclResponse, err error) {
	response = CreateAddEntriesToAclResponse()
	err = client.DoAction(request, response)
	return
}

// AddEntriesToAclWithChan invokes the ga.AddEntriesToAcl API asynchronously
func (client *Client) AddEntriesToAclWithChan(request *AddEntriesToAclRequest) (<-chan *AddEntriesToAclResponse, <-chan error) {
	responseChan := make(chan *AddEntriesToAclResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddEntriesToAcl(request)
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

// AddEntriesToAclWithCallback invokes the ga.AddEntriesToAcl API asynchronously
func (client *Client) AddEntriesToAclWithCallback(request *AddEntriesToAclRequest, callback func(response *AddEntriesToAclResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddEntriesToAclResponse
		var err error
		defer close(result)
		response, err = client.AddEntriesToAcl(request)
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

// AddEntriesToAclRequest is the request struct for api AddEntriesToAcl
type AddEntriesToAclRequest struct {
	*requests.RpcRequest
	AclId       string                       `position:"Query" name:"AclId"`
	DryRun      requests.Boolean             `position:"Query" name:"DryRun"`
	ClientToken string                       `position:"Query" name:"ClientToken"`
	AclEntries  *[]AddEntriesToAclAclEntries `position:"Query" name:"AclEntries"  type:"Repeated"`
}

// AddEntriesToAclAclEntries is a repeated param struct in AddEntriesToAclRequest
type AddEntriesToAclAclEntries struct {
	Entry            string `name:"Entry"`
	EntryDescription string `name:"EntryDescription"`
}

// AddEntriesToAclResponse is the response struct for api AddEntriesToAcl
type AddEntriesToAclResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	AclId     string `json:"AclId" xml:"AclId"`
}

// CreateAddEntriesToAclRequest creates a request to invoke AddEntriesToAcl API
func CreateAddEntriesToAclRequest() (request *AddEntriesToAclRequest) {
	request = &AddEntriesToAclRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "AddEntriesToAcl", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddEntriesToAclResponse creates a response to parse from AddEntriesToAcl response
func CreateAddEntriesToAclResponse() (response *AddEntriesToAclResponse) {
	response = &AddEntriesToAclResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
