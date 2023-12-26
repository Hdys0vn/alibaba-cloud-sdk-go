package kms

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

// CreateSecret invokes the kms.CreateSecret API synchronously
func (client *Client) CreateSecret(request *CreateSecretRequest) (response *CreateSecretResponse, err error) {
	response = CreateCreateSecretResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSecretWithChan invokes the kms.CreateSecret API asynchronously
func (client *Client) CreateSecretWithChan(request *CreateSecretRequest) (<-chan *CreateSecretResponse, <-chan error) {
	responseChan := make(chan *CreateSecretResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSecret(request)
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

// CreateSecretWithCallback invokes the kms.CreateSecret API asynchronously
func (client *Client) CreateSecretWithCallback(request *CreateSecretRequest, callback func(response *CreateSecretResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSecretResponse
		var err error
		defer close(result)
		response, err = client.CreateSecret(request)
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

// CreateSecretRequest is the request struct for api CreateSecret
type CreateSecretRequest struct {
	*requests.RpcRequest
	SecretType              string           `position:"Query" name:"SecretType"`
	Description             string           `position:"Query" name:"Description"`
	RotationInterval        string           `position:"Query" name:"RotationInterval"`
	EnableAutomaticRotation requests.Boolean `position:"Query" name:"EnableAutomaticRotation"`
	EncryptionKeyId         string           `position:"Query" name:"EncryptionKeyId"`
	Tags                    string           `position:"Query" name:"Tags"`
	ExtendedConfig          string           `position:"Query" name:"ExtendedConfig"`
	VersionId               string           `position:"Query" name:"VersionId"`
	DKMSInstanceId          string           `position:"Query" name:"DKMSInstanceId"`
	SecretData              string           `position:"Query" name:"SecretData"`
	SecretName              string           `position:"Query" name:"SecretName"`
	SecretDataType          string           `position:"Query" name:"SecretDataType"`
}

// CreateSecretResponse is the response struct for api CreateSecret
type CreateSecretResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	AutomaticRotation string `json:"AutomaticRotation" xml:"AutomaticRotation"`
	SecretName        string `json:"SecretName" xml:"SecretName"`
	VersionId         string `json:"VersionId" xml:"VersionId"`
	NextRotationDate  string `json:"NextRotationDate" xml:"NextRotationDate"`
	SecretType        string `json:"SecretType" xml:"SecretType"`
	RotationInterval  string `json:"RotationInterval" xml:"RotationInterval"`
	Arn               string `json:"Arn" xml:"Arn"`
	ExtendedConfig    string `json:"ExtendedConfig" xml:"ExtendedConfig"`
	DKMSInstanceId    string `json:"DKMSInstanceId" xml:"DKMSInstanceId"`
}

// CreateCreateSecretRequest creates a request to invoke CreateSecret API
func CreateCreateSecretRequest() (request *CreateSecretRequest) {
	request = &CreateSecretRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "CreateSecret", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSecretResponse creates a response to parse from CreateSecret response
func CreateCreateSecretResponse() (response *CreateSecretResponse) {
	response = &CreateSecretResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}