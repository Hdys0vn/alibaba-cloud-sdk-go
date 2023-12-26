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

// DeleteK8sSecret invokes the edas.DeleteK8sSecret API synchronously
func (client *Client) DeleteK8sSecret(request *DeleteK8sSecretRequest) (response *DeleteK8sSecretResponse, err error) {
	response = CreateDeleteK8sSecretResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteK8sSecretWithChan invokes the edas.DeleteK8sSecret API asynchronously
func (client *Client) DeleteK8sSecretWithChan(request *DeleteK8sSecretRequest) (<-chan *DeleteK8sSecretResponse, <-chan error) {
	responseChan := make(chan *DeleteK8sSecretResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteK8sSecret(request)
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

// DeleteK8sSecretWithCallback invokes the edas.DeleteK8sSecret API asynchronously
func (client *Client) DeleteK8sSecretWithCallback(request *DeleteK8sSecretRequest, callback func(response *DeleteK8sSecretResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteK8sSecretResponse
		var err error
		defer close(result)
		response, err = client.DeleteK8sSecret(request)
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

// DeleteK8sSecretRequest is the request struct for api DeleteK8sSecret
type DeleteK8sSecretRequest struct {
	*requests.RoaRequest
	Name      string `position:"Query" name:"Name"`
	Namespace string `position:"Query" name:"Namespace"`
	ClusterId string `position:"Query" name:"ClusterId"`
}

// DeleteK8sSecretResponse is the response struct for api DeleteK8sSecret
type DeleteK8sSecretResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteK8sSecretRequest creates a request to invoke DeleteK8sSecret API
func CreateDeleteK8sSecretRequest() (request *DeleteK8sSecretRequest) {
	request = &DeleteK8sSecretRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "DeleteK8sSecret", "/pop/v5/k8s/acs/k8s_secret", "Edas", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteK8sSecretResponse creates a response to parse from DeleteK8sSecret response
func CreateDeleteK8sSecretResponse() (response *DeleteK8sSecretResponse) {
	response = &DeleteK8sSecretResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
