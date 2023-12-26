package mse

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

// UpdateServiceSource invokes the mse.UpdateServiceSource API synchronously
func (client *Client) UpdateServiceSource(request *UpdateServiceSourceRequest) (response *UpdateServiceSourceResponse, err error) {
	response = CreateUpdateServiceSourceResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateServiceSourceWithChan invokes the mse.UpdateServiceSource API asynchronously
func (client *Client) UpdateServiceSourceWithChan(request *UpdateServiceSourceRequest) (<-chan *UpdateServiceSourceResponse, <-chan error) {
	responseChan := make(chan *UpdateServiceSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateServiceSource(request)
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

// UpdateServiceSourceWithCallback invokes the mse.UpdateServiceSource API asynchronously
func (client *Client) UpdateServiceSourceWithCallback(request *UpdateServiceSourceRequest, callback func(response *UpdateServiceSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateServiceSourceResponse
		var err error
		defer close(result)
		response, err = client.UpdateServiceSource(request)
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

// UpdateServiceSourceRequest is the request struct for api UpdateServiceSource
type UpdateServiceSourceRequest struct {
	*requests.RpcRequest
	IngressOptionsRequest UpdateServiceSourceIngressOptionsRequest `position:"Query" name:"IngressOptionsRequest"  type:"Struct"`
	MseSessionId          string                                   `position:"Query" name:"MseSessionId"`
	GatewayUniqueId       string                                   `position:"Query" name:"GatewayUniqueId"`
	Source                string                                   `position:"Query" name:"Source"`
	Type                  string                                   `position:"Query" name:"Type"`
	Id                    requests.Integer                         `position:"Query" name:"Id"`
	PathList              *[]string                                `position:"Query" name:"PathList"  type:"Json"`
	GatewayId             requests.Integer                         `position:"Query" name:"GatewayId"`
	Address               string                                   `position:"Query" name:"Address"`
	Name                  string                                   `position:"Query" name:"Name"`
	AcceptLanguage        string                                   `position:"Query" name:"AcceptLanguage"`
}

// UpdateServiceSourceIngressOptionsRequest is a repeated param struct in UpdateServiceSourceRequest
type UpdateServiceSourceIngressOptionsRequest struct {
	EnableStatus   string `name:"EnableStatus"`
	EnableIngress  string `name:"EnableIngress"`
	WatchNamespace string `name:"WatchNamespace"`
	IngressClass   string `name:"IngressClass"`
}

// UpdateServiceSourceResponse is the response struct for api UpdateServiceSource
type UpdateServiceSourceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           int64  `json:"Data" xml:"Data"`
}

// CreateUpdateServiceSourceRequest creates a request to invoke UpdateServiceSource API
func CreateUpdateServiceSourceRequest() (request *UpdateServiceSourceRequest) {
	request = &UpdateServiceSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateServiceSource", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateServiceSourceResponse creates a response to parse from UpdateServiceSource response
func CreateUpdateServiceSourceResponse() (response *UpdateServiceSourceResponse) {
	response = &UpdateServiceSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
