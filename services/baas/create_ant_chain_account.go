package baas

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

// CreateAntChainAccount invokes the baas.CreateAntChainAccount API synchronously
// api document: https://help.aliyun.com/api/baas/createantchainaccount.html
func (client *Client) CreateAntChainAccount(request *CreateAntChainAccountRequest) (response *CreateAntChainAccountResponse, err error) {
	response = CreateCreateAntChainAccountResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAntChainAccountWithChan invokes the baas.CreateAntChainAccount API asynchronously
// api document: https://help.aliyun.com/api/baas/createantchainaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAntChainAccountWithChan(request *CreateAntChainAccountRequest) (<-chan *CreateAntChainAccountResponse, <-chan error) {
	responseChan := make(chan *CreateAntChainAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAntChainAccount(request)
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

// CreateAntChainAccountWithCallback invokes the baas.CreateAntChainAccount API asynchronously
// api document: https://help.aliyun.com/api/baas/createantchainaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAntChainAccountWithCallback(request *CreateAntChainAccountRequest, callback func(response *CreateAntChainAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAntChainAccountResponse
		var err error
		defer close(result)
		response, err = client.CreateAntChainAccount(request)
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

// CreateAntChainAccountRequest is the request struct for api CreateAntChainAccount
type CreateAntChainAccountRequest struct {
	*requests.RpcRequest
	AccountRecoverPubKey string `position:"Body" name:"AccountRecoverPubKey"`
	AccountPubKey        string `position:"Body" name:"AccountPubKey"`
	AntChainId           string `position:"Body" name:"AntChainId"`
	Account              string `position:"Body" name:"Account"`
}

// CreateAntChainAccountResponse is the response struct for api CreateAntChainAccount
type CreateAntChainAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateCreateAntChainAccountRequest creates a request to invoke CreateAntChainAccount API
func CreateCreateAntChainAccountRequest() (request *CreateAntChainAccountRequest) {
	request = &CreateAntChainAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "CreateAntChainAccount", "baas", "openAPI")
	return
}

// CreateCreateAntChainAccountResponse creates a response to parse from CreateAntChainAccount response
func CreateCreateAntChainAccountResponse() (response *CreateAntChainAccountResponse) {
	response = &CreateAntChainAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
