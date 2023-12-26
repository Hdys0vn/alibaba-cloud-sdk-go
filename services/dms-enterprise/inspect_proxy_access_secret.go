package dms_enterprise

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

// InspectProxyAccessSecret invokes the dms_enterprise.InspectProxyAccessSecret API synchronously
func (client *Client) InspectProxyAccessSecret(request *InspectProxyAccessSecretRequest) (response *InspectProxyAccessSecretResponse, err error) {
	response = CreateInspectProxyAccessSecretResponse()
	err = client.DoAction(request, response)
	return
}

// InspectProxyAccessSecretWithChan invokes the dms_enterprise.InspectProxyAccessSecret API asynchronously
func (client *Client) InspectProxyAccessSecretWithChan(request *InspectProxyAccessSecretRequest) (<-chan *InspectProxyAccessSecretResponse, <-chan error) {
	responseChan := make(chan *InspectProxyAccessSecretResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InspectProxyAccessSecret(request)
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

// InspectProxyAccessSecretWithCallback invokes the dms_enterprise.InspectProxyAccessSecret API asynchronously
func (client *Client) InspectProxyAccessSecretWithCallback(request *InspectProxyAccessSecretRequest, callback func(response *InspectProxyAccessSecretResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InspectProxyAccessSecretResponse
		var err error
		defer close(result)
		response, err = client.InspectProxyAccessSecret(request)
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

// InspectProxyAccessSecretRequest is the request struct for api InspectProxyAccessSecret
type InspectProxyAccessSecretRequest struct {
	*requests.RpcRequest
	ProxyAccessId requests.Integer `position:"Query" name:"ProxyAccessId"`
	Tid           requests.Integer `position:"Query" name:"Tid"`
}

// InspectProxyAccessSecretResponse is the response struct for api InspectProxyAccessSecret
type InspectProxyAccessSecretResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	AccessSecret string `json:"AccessSecret" xml:"AccessSecret"`
}

// CreateInspectProxyAccessSecretRequest creates a request to invoke InspectProxyAccessSecret API
func CreateInspectProxyAccessSecretRequest() (request *InspectProxyAccessSecretRequest) {
	request = &InspectProxyAccessSecretRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "InspectProxyAccessSecret", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInspectProxyAccessSecretResponse creates a response to parse from InspectProxyAccessSecret response
func CreateInspectProxyAccessSecretResponse() (response *InspectProxyAccessSecretResponse) {
	response = &InspectProxyAccessSecretResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
