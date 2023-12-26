package ecs

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

// ImportKeyPair invokes the ecs.ImportKeyPair API synchronously
func (client *Client) ImportKeyPair(request *ImportKeyPairRequest) (response *ImportKeyPairResponse, err error) {
	response = CreateImportKeyPairResponse()
	err = client.DoAction(request, response)
	return
}

// ImportKeyPairWithChan invokes the ecs.ImportKeyPair API asynchronously
func (client *Client) ImportKeyPairWithChan(request *ImportKeyPairRequest) (<-chan *ImportKeyPairResponse, <-chan error) {
	responseChan := make(chan *ImportKeyPairResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportKeyPair(request)
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

// ImportKeyPairWithCallback invokes the ecs.ImportKeyPair API asynchronously
func (client *Client) ImportKeyPairWithCallback(request *ImportKeyPairRequest, callback func(response *ImportKeyPairResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportKeyPairResponse
		var err error
		defer close(result)
		response, err = client.ImportKeyPair(request)
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

// ImportKeyPairRequest is the request struct for api ImportKeyPair
type ImportKeyPairRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer    `position:"Query" name:"ResourceOwnerId"`
	KeyPairName          string              `position:"Query" name:"KeyPairName"`
	ResourceGroupId      string              `position:"Query" name:"ResourceGroupId"`
	Tag                  *[]ImportKeyPairTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount string              `position:"Query" name:"ResourceOwnerAccount"`
	PublicKeyBody        string              `position:"Query" name:"PublicKeyBody"`
	OwnerId              requests.Integer    `position:"Query" name:"OwnerId"`
}

// ImportKeyPairTag is a repeated param struct in ImportKeyPairRequest
type ImportKeyPairTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ImportKeyPairResponse is the response struct for api ImportKeyPair
type ImportKeyPairResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	KeyPairName        string `json:"KeyPairName" xml:"KeyPairName"`
	KeyPairFingerPrint string `json:"KeyPairFingerPrint" xml:"KeyPairFingerPrint"`
}

// CreateImportKeyPairRequest creates a request to invoke ImportKeyPair API
func CreateImportKeyPairRequest() (request *ImportKeyPairRequest) {
	request = &ImportKeyPairRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ImportKeyPair", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateImportKeyPairResponse creates a response to parse from ImportKeyPair response
func CreateImportKeyPairResponse() (response *ImportKeyPairResponse) {
	response = &ImportKeyPairResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
