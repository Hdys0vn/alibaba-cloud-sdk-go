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

// DescribeSecret invokes the sae.DescribeSecret API synchronously
func (client *Client) DescribeSecret(request *DescribeSecretRequest) (response *DescribeSecretResponse, err error) {
	response = CreateDescribeSecretResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSecretWithChan invokes the sae.DescribeSecret API asynchronously
func (client *Client) DescribeSecretWithChan(request *DescribeSecretRequest) (<-chan *DescribeSecretResponse, <-chan error) {
	responseChan := make(chan *DescribeSecretResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecret(request)
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

// DescribeSecretWithCallback invokes the sae.DescribeSecret API asynchronously
func (client *Client) DescribeSecretWithCallback(request *DescribeSecretRequest, callback func(response *DescribeSecretResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecretResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecret(request)
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

// DescribeSecretRequest is the request struct for api DescribeSecret
type DescribeSecretRequest struct {
	*requests.RoaRequest
	NamespaceId string           `position:"Query" name:"NamespaceId"`
	SecretId    requests.Integer `position:"Query" name:"SecretId"`
}

// DescribeSecretResponse is the response struct for api DescribeSecret
type DescribeSecretResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeSecretRequest creates a request to invoke DescribeSecret API
func CreateDescribeSecretRequest() (request *DescribeSecretRequest) {
	request = &DescribeSecretRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeSecret", "/pop/v1/sam/secret/secret", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeSecretResponse creates a response to parse from DescribeSecret response
func CreateDescribeSecretResponse() (response *DescribeSecretResponse) {
	response = &DescribeSecretResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
