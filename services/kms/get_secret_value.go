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

// GetSecretValue invokes the kms.GetSecretValue API synchronously
func (client *Client) GetSecretValue(request *GetSecretValueRequest) (response *GetSecretValueResponse, err error) {
	response = CreateGetSecretValueResponse()
	err = client.DoAction(request, response)
	return
}

// GetSecretValueWithChan invokes the kms.GetSecretValue API asynchronously
func (client *Client) GetSecretValueWithChan(request *GetSecretValueRequest) (<-chan *GetSecretValueResponse, <-chan error) {
	responseChan := make(chan *GetSecretValueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSecretValue(request)
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

// GetSecretValueWithCallback invokes the kms.GetSecretValue API asynchronously
func (client *Client) GetSecretValueWithCallback(request *GetSecretValueRequest, callback func(response *GetSecretValueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSecretValueResponse
		var err error
		defer close(result)
		response, err = client.GetSecretValue(request)
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

// GetSecretValueRequest is the request struct for api GetSecretValue
type GetSecretValueRequest struct {
	*requests.RpcRequest
	VersionId           string           `position:"Query" name:"VersionId"`
	VersionStage        string           `position:"Query" name:"VersionStage"`
	SecretName          string           `position:"Query" name:"SecretName"`
	FetchExtendedConfig requests.Boolean `position:"Query" name:"FetchExtendedConfig"`
}

// GetSecretValueResponse is the response struct for api GetSecretValue
type GetSecretValueResponse struct {
	*responses.BaseResponse
	SecretDataType    string                        `json:"SecretDataType" xml:"SecretDataType"`
	CreateTime        string                        `json:"CreateTime" xml:"CreateTime"`
	VersionId         string                        `json:"VersionId" xml:"VersionId"`
	NextRotationDate  string                        `json:"NextRotationDate" xml:"NextRotationDate"`
	SecretData        string                        `json:"SecretData" xml:"SecretData"`
	RotationInterval  string                        `json:"RotationInterval" xml:"RotationInterval"`
	ExtendedConfig    string                        `json:"ExtendedConfig" xml:"ExtendedConfig"`
	LastRotationDate  string                        `json:"LastRotationDate" xml:"LastRotationDate"`
	RequestId         string                        `json:"RequestId" xml:"RequestId"`
	SecretName        string                        `json:"SecretName" xml:"SecretName"`
	AutomaticRotation string                        `json:"AutomaticRotation" xml:"AutomaticRotation"`
	SecretType        string                        `json:"SecretType" xml:"SecretType"`
	VersionStages     VersionStagesInGetSecretValue `json:"VersionStages" xml:"VersionStages"`
}

// CreateGetSecretValueRequest creates a request to invoke GetSecretValue API
func CreateGetSecretValueRequest() (request *GetSecretValueRequest) {
	request = &GetSecretValueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "GetSecretValue", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSecretValueResponse creates a response to parse from GetSecretValue response
func CreateGetSecretValueResponse() (response *GetSecretValueResponse) {
	response = &GetSecretValueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
