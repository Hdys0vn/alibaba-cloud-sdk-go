package resourcemanager

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

// CreateResourceAccount invokes the resourcemanager.CreateResourceAccount API synchronously
func (client *Client) CreateResourceAccount(request *CreateResourceAccountRequest) (response *CreateResourceAccountResponse, err error) {
	response = CreateCreateResourceAccountResponse()
	err = client.DoAction(request, response)
	return
}

// CreateResourceAccountWithChan invokes the resourcemanager.CreateResourceAccount API asynchronously
func (client *Client) CreateResourceAccountWithChan(request *CreateResourceAccountRequest) (<-chan *CreateResourceAccountResponse, <-chan error) {
	responseChan := make(chan *CreateResourceAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateResourceAccount(request)
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

// CreateResourceAccountWithCallback invokes the resourcemanager.CreateResourceAccount API asynchronously
func (client *Client) CreateResourceAccountWithCallback(request *CreateResourceAccountRequest, callback func(response *CreateResourceAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateResourceAccountResponse
		var err error
		defer close(result)
		response, err = client.CreateResourceAccount(request)
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

// CreateResourceAccountRequest is the request struct for api CreateResourceAccount
type CreateResourceAccountRequest struct {
	*requests.RpcRequest
	Tag               *[]CreateResourceAccountTag `position:"Query" name:"Tag"  type:"Repeated"`
	AccountNamePrefix string                      `position:"Query" name:"AccountNamePrefix"`
	ResellAccountType string                      `position:"Query" name:"ResellAccountType"`
	ParentFolderId    string                      `position:"Query" name:"ParentFolderId"`
	DisplayName       string                      `position:"Query" name:"DisplayName"`
	PayerAccountId    string                      `position:"Query" name:"PayerAccountId"`
}

// CreateResourceAccountTag is a repeated param struct in CreateResourceAccountRequest
type CreateResourceAccountTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateResourceAccountResponse is the response struct for api CreateResourceAccount
type CreateResourceAccountResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Account   Account `json:"Account" xml:"Account"`
}

// CreateCreateResourceAccountRequest creates a request to invoke CreateResourceAccount API
func CreateCreateResourceAccountRequest() (request *CreateResourceAccountRequest) {
	request = &CreateResourceAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "CreateResourceAccount", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateResourceAccountResponse creates a response to parse from CreateResourceAccount response
func CreateCreateResourceAccountResponse() (response *CreateResourceAccountResponse) {
	response = &CreateResourceAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}