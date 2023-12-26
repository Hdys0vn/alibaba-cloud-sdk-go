package ververica

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

// CreateSecretValue invokes the ververica.CreateSecretValue API synchronously
func (client *Client) CreateSecretValue(request *CreateSecretValueRequest) (response *CreateSecretValueResponse, err error) {
	response = CreateCreateSecretValueResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSecretValueWithChan invokes the ververica.CreateSecretValue API asynchronously
func (client *Client) CreateSecretValueWithChan(request *CreateSecretValueRequest) (<-chan *CreateSecretValueResponse, <-chan error) {
	responseChan := make(chan *CreateSecretValueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSecretValue(request)
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

// CreateSecretValueWithCallback invokes the ververica.CreateSecretValue API asynchronously
func (client *Client) CreateSecretValueWithCallback(request *CreateSecretValueRequest, callback func(response *CreateSecretValueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSecretValueResponse
		var err error
		defer close(result)
		response, err = client.CreateSecretValue(request)
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

// CreateSecretValueRequest is the request struct for api CreateSecretValue
type CreateSecretValueRequest struct {
	*requests.RoaRequest
	Workspace  string `position:"Path" name:"workspace"`
	ParamsJson string `position:"Body" name:"paramsJson"`
	Namespace  string `position:"Path" name:"namespace"`
}

// CreateSecretValueResponse is the response struct for api CreateSecretValue
type CreateSecretValueResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateCreateSecretValueRequest creates a request to invoke CreateSecretValue API
func CreateCreateSecretValueRequest() (request *CreateSecretValueRequest) {
	request = &CreateSecretValueRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "CreateSecretValue", "/pop/workspaces/[workspace]/api/v1/namespaces/[namespace]/secret-values", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateSecretValueResponse creates a response to parse from CreateSecretValue response
func CreateCreateSecretValueResponse() (response *CreateSecretValueResponse) {
	response = &CreateSecretValueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}