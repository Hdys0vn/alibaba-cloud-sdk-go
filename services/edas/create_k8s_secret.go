package edas

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

// CreateK8sSecret invokes the edas.CreateK8sSecret API synchronously
func (client *Client) CreateK8sSecret(request *CreateK8sSecretRequest) (response *CreateK8sSecretResponse, err error) {
	response = CreateCreateK8sSecretResponse()
	err = client.DoAction(request, response)
	return
}

// CreateK8sSecretWithChan invokes the edas.CreateK8sSecret API asynchronously
func (client *Client) CreateK8sSecretWithChan(request *CreateK8sSecretRequest) (<-chan *CreateK8sSecretResponse, <-chan error) {
	responseChan := make(chan *CreateK8sSecretResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateK8sSecret(request)
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

// CreateK8sSecretWithCallback invokes the edas.CreateK8sSecret API asynchronously
func (client *Client) CreateK8sSecretWithCallback(request *CreateK8sSecretRequest, callback func(response *CreateK8sSecretResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateK8sSecretResponse
		var err error
		defer close(result)
		response, err = client.CreateK8sSecret(request)
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

// CreateK8sSecretRequest is the request struct for api CreateK8sSecret
type CreateK8sSecretRequest struct {
	*requests.RoaRequest
	Base64Encoded requests.Boolean `position:"Body" name:"Base64Encoded"`
	Data          string           `position:"Body" name:"Data"`
	Name          string           `position:"Body" name:"Name"`
	Namespace     string           `position:"Body" name:"Namespace"`
	ClusterId     string           `position:"Body" name:"ClusterId"`
	CertId        string           `position:"Body" name:"CertId"`
	Type          string           `position:"Body" name:"Type"`
	CertRegionId  string           `position:"Body" name:"CertRegionId"`
}

// CreateK8sSecretResponse is the response struct for api CreateK8sSecret
type CreateK8sSecretResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateK8sSecretRequest creates a request to invoke CreateK8sSecret API
func CreateCreateK8sSecretRequest() (request *CreateK8sSecretRequest) {
	request = &CreateK8sSecretRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "CreateK8sSecret", "/pop/v5/k8s/acs/k8s_secret", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateK8sSecretResponse creates a response to parse from CreateK8sSecret response
func CreateCreateK8sSecretResponse() (response *CreateK8sSecretResponse) {
	response = &CreateK8sSecretResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
