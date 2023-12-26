package arms

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

// DeleteIntegrations invokes the arms.DeleteIntegrations API synchronously
func (client *Client) DeleteIntegrations(request *DeleteIntegrationsRequest) (response *DeleteIntegrationsResponse, err error) {
	response = CreateDeleteIntegrationsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteIntegrationsWithChan invokes the arms.DeleteIntegrations API asynchronously
func (client *Client) DeleteIntegrationsWithChan(request *DeleteIntegrationsRequest) (<-chan *DeleteIntegrationsResponse, <-chan error) {
	responseChan := make(chan *DeleteIntegrationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteIntegrations(request)
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

// DeleteIntegrationsWithCallback invokes the arms.DeleteIntegrations API asynchronously
func (client *Client) DeleteIntegrationsWithCallback(request *DeleteIntegrationsRequest, callback func(response *DeleteIntegrationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteIntegrationsResponse
		var err error
		defer close(result)
		response, err = client.DeleteIntegrations(request)
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

// DeleteIntegrationsRequest is the request struct for api DeleteIntegrations
type DeleteIntegrationsRequest struct {
	*requests.RpcRequest
	IntegrationId requests.Integer `position:"Query" name:"IntegrationId"`
}

// DeleteIntegrationsResponse is the response struct for api DeleteIntegrations
type DeleteIntegrationsResponse struct {
	*responses.BaseResponse
	RequestId                   string `json:"RequestId" xml:"RequestId"`
	DeleteIntegrationsIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
}

// CreateDeleteIntegrationsRequest creates a request to invoke DeleteIntegrations API
func CreateDeleteIntegrationsRequest() (request *DeleteIntegrationsRequest) {
	request = &DeleteIntegrationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DeleteIntegrations", "arms", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDeleteIntegrationsResponse creates a response to parse from DeleteIntegrations response
func CreateDeleteIntegrationsResponse() (response *DeleteIntegrationsResponse) {
	response = &DeleteIntegrationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
