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

// CreateAntChainAccountWithKeyPairAutoCreation invokes the baas.CreateAntChainAccountWithKeyPairAutoCreation API synchronously
// api document: https://help.aliyun.com/api/baas/createantchainaccountwithkeypairautocreation.html
func (client *Client) CreateAntChainAccountWithKeyPairAutoCreation(request *CreateAntChainAccountWithKeyPairAutoCreationRequest) (response *CreateAntChainAccountWithKeyPairAutoCreationResponse, err error) {
	response = CreateCreateAntChainAccountWithKeyPairAutoCreationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAntChainAccountWithKeyPairAutoCreationWithChan invokes the baas.CreateAntChainAccountWithKeyPairAutoCreation API asynchronously
// api document: https://help.aliyun.com/api/baas/createantchainaccountwithkeypairautocreation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAntChainAccountWithKeyPairAutoCreationWithChan(request *CreateAntChainAccountWithKeyPairAutoCreationRequest) (<-chan *CreateAntChainAccountWithKeyPairAutoCreationResponse, <-chan error) {
	responseChan := make(chan *CreateAntChainAccountWithKeyPairAutoCreationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAntChainAccountWithKeyPairAutoCreation(request)
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

// CreateAntChainAccountWithKeyPairAutoCreationWithCallback invokes the baas.CreateAntChainAccountWithKeyPairAutoCreation API asynchronously
// api document: https://help.aliyun.com/api/baas/createantchainaccountwithkeypairautocreation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAntChainAccountWithKeyPairAutoCreationWithCallback(request *CreateAntChainAccountWithKeyPairAutoCreationRequest, callback func(response *CreateAntChainAccountWithKeyPairAutoCreationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAntChainAccountWithKeyPairAutoCreationResponse
		var err error
		defer close(result)
		response, err = client.CreateAntChainAccountWithKeyPairAutoCreation(request)
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

// CreateAntChainAccountWithKeyPairAutoCreationRequest is the request struct for api CreateAntChainAccountWithKeyPairAutoCreation
type CreateAntChainAccountWithKeyPairAutoCreationRequest struct {
	*requests.RpcRequest
	Password        string `position:"Body" name:"Password"`
	AntChainId      string `position:"Body" name:"AntChainId"`
	RecoverPassword string `position:"Body" name:"RecoverPassword"`
	Account         string `position:"Body" name:"Account"`
}

// CreateAntChainAccountWithKeyPairAutoCreationResponse is the response struct for api CreateAntChainAccountWithKeyPairAutoCreation
type CreateAntChainAccountWithKeyPairAutoCreationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateCreateAntChainAccountWithKeyPairAutoCreationRequest creates a request to invoke CreateAntChainAccountWithKeyPairAutoCreation API
func CreateCreateAntChainAccountWithKeyPairAutoCreationRequest() (request *CreateAntChainAccountWithKeyPairAutoCreationRequest) {
	request = &CreateAntChainAccountWithKeyPairAutoCreationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "CreateAntChainAccountWithKeyPairAutoCreation", "baas", "openAPI")
	return
}

// CreateCreateAntChainAccountWithKeyPairAutoCreationResponse creates a response to parse from CreateAntChainAccountWithKeyPairAutoCreation response
func CreateCreateAntChainAccountWithKeyPairAutoCreationResponse() (response *CreateAntChainAccountWithKeyPairAutoCreationResponse) {
	response = &CreateAntChainAccountWithKeyPairAutoCreationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
