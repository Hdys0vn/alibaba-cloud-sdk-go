package sae

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

// DeleteSecret invokes the sae.DeleteSecret API synchronously
func (client *Client) DeleteSecret(request *DeleteSecretRequest) (response *DeleteSecretResponse, err error) {
	response = CreateDeleteSecretResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSecretWithChan invokes the sae.DeleteSecret API asynchronously
func (client *Client) DeleteSecretWithChan(request *DeleteSecretRequest) (<-chan *DeleteSecretResponse, <-chan error) {
	responseChan := make(chan *DeleteSecretResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSecret(request)
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

// DeleteSecretWithCallback invokes the sae.DeleteSecret API asynchronously
func (client *Client) DeleteSecretWithCallback(request *DeleteSecretRequest, callback func(response *DeleteSecretResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSecretResponse
		var err error
		defer close(result)
		response, err = client.DeleteSecret(request)
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

// DeleteSecretRequest is the request struct for api DeleteSecret
type DeleteSecretRequest struct {
	*requests.RoaRequest
	NamespaceId string           `position:"Query" name:"NamespaceId"`
	SecretId    requests.Integer `position:"Query" name:"SecretId"`
}

// DeleteSecretResponse is the response struct for api DeleteSecret
type DeleteSecretResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDeleteSecretRequest creates a request to invoke DeleteSecret API
func CreateDeleteSecretRequest() (request *DeleteSecretRequest) {
	request = &DeleteSecretRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DeleteSecret", "/pop/v1/sam/secret/secret", "serverless", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteSecretResponse creates a response to parse from DeleteSecret response
func CreateDeleteSecretResponse() (response *DeleteSecretResponse) {
	response = &DeleteSecretResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
