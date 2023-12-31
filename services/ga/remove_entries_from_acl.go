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

// RemoveEntriesFromAcl invokes the ga.RemoveEntriesFromAcl API synchronously
func (client *Client) RemoveEntriesFromAcl(request *RemoveEntriesFromAclRequest) (response *RemoveEntriesFromAclResponse, err error) {
	response = CreateRemoveEntriesFromAclResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveEntriesFromAclWithChan invokes the ga.RemoveEntriesFromAcl API asynchronously
func (client *Client) RemoveEntriesFromAclWithChan(request *RemoveEntriesFromAclRequest) (<-chan *RemoveEntriesFromAclResponse, <-chan error) {
	responseChan := make(chan *RemoveEntriesFromAclResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveEntriesFromAcl(request)
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

// RemoveEntriesFromAclWithCallback invokes the ga.RemoveEntriesFromAcl API asynchronously
func (client *Client) RemoveEntriesFromAclWithCallback(request *RemoveEntriesFromAclRequest, callback func(response *RemoveEntriesFromAclResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveEntriesFromAclResponse
		var err error
		defer close(result)
		response, err = client.RemoveEntriesFromAcl(request)
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

// RemoveEntriesFromAclRequest is the request struct for api RemoveEntriesFromAcl
type RemoveEntriesFromAclRequest struct {
	*requests.RpcRequest
	AclId       string                            `position:"Query" name:"AclId"`
	DryRun      requests.Boolean                  `position:"Query" name:"DryRun"`
	ClientToken string                            `position:"Query" name:"ClientToken"`
	AclEntries  *[]RemoveEntriesFromAclAclEntries `position:"Query" name:"AclEntries"  type:"Repeated"`
}

// RemoveEntriesFromAclAclEntries is a repeated param struct in RemoveEntriesFromAclRequest
type RemoveEntriesFromAclAclEntries struct {
	Entry string `name:"Entry"`
}

// RemoveEntriesFromAclResponse is the response struct for api RemoveEntriesFromAcl
type RemoveEntriesFromAclResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	AclId     string `json:"AclId" xml:"AclId"`
}

// CreateRemoveEntriesFromAclRequest creates a request to invoke RemoveEntriesFromAcl API
func CreateRemoveEntriesFromAclRequest() (request *RemoveEntriesFromAclRequest) {
	request = &RemoveEntriesFromAclRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "RemoveEntriesFromAcl", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveEntriesFromAclResponse creates a response to parse from RemoveEntriesFromAcl response
func CreateRemoveEntriesFromAclResponse() (response *RemoveEntriesFromAclResponse) {
	response = &RemoveEntriesFromAclResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
