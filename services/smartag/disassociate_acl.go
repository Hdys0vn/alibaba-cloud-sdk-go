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

// DisassociateACL invokes the smartag.DisassociateACL API synchronously
func (client *Client) DisassociateACL(request *DisassociateACLRequest) (response *DisassociateACLResponse, err error) {
	response = CreateDisassociateACLResponse()
	err = client.DoAction(request, response)
	return
}

// DisassociateACLWithChan invokes the smartag.DisassociateACL API asynchronously
func (client *Client) DisassociateACLWithChan(request *DisassociateACLRequest) (<-chan *DisassociateACLResponse, <-chan error) {
	responseChan := make(chan *DisassociateACLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisassociateACL(request)
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

// DisassociateACLWithCallback invokes the smartag.DisassociateACL API asynchronously
func (client *Client) DisassociateACLWithCallback(request *DisassociateACLRequest, callback func(response *DisassociateACLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisassociateACLResponse
		var err error
		defer close(result)
		response, err = client.DisassociateACL(request)
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

// DisassociateACLRequest is the request struct for api DisassociateACL
type DisassociateACLRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
}

// DisassociateACLResponse is the response struct for api DisassociateACL
type DisassociateACLResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisassociateACLRequest creates a request to invoke DisassociateACL API
func CreateDisassociateACLRequest() (request *DisassociateACLRequest) {
	request = &DisassociateACLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DisassociateACL", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisassociateACLResponse creates a response to parse from DisassociateACL response
func CreateDisassociateACLResponse() (response *DisassociateACLResponse) {
	response = &DisassociateACLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
