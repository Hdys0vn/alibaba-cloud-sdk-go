package privatelink

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

// DisableVpcEndpointConnection invokes the privatelink.DisableVpcEndpointConnection API synchronously
func (client *Client) DisableVpcEndpointConnection(request *DisableVpcEndpointConnectionRequest) (response *DisableVpcEndpointConnectionResponse, err error) {
	response = CreateDisableVpcEndpointConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// DisableVpcEndpointConnectionWithChan invokes the privatelink.DisableVpcEndpointConnection API asynchronously
func (client *Client) DisableVpcEndpointConnectionWithChan(request *DisableVpcEndpointConnectionRequest) (<-chan *DisableVpcEndpointConnectionResponse, <-chan error) {
	responseChan := make(chan *DisableVpcEndpointConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableVpcEndpointConnection(request)
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

// DisableVpcEndpointConnectionWithCallback invokes the privatelink.DisableVpcEndpointConnection API asynchronously
func (client *Client) DisableVpcEndpointConnectionWithCallback(request *DisableVpcEndpointConnectionRequest, callback func(response *DisableVpcEndpointConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableVpcEndpointConnectionResponse
		var err error
		defer close(result)
		response, err = client.DisableVpcEndpointConnection(request)
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

// DisableVpcEndpointConnectionRequest is the request struct for api DisableVpcEndpointConnection
type DisableVpcEndpointConnectionRequest struct {
	*requests.RpcRequest
	ClientToken string           `position:"Query" name:"ClientToken"`
	EndpointId  string           `position:"Query" name:"EndpointId"`
	DryRun      requests.Boolean `position:"Query" name:"DryRun"`
	ServiceId   string           `position:"Query" name:"ServiceId"`
}

// DisableVpcEndpointConnectionResponse is the response struct for api DisableVpcEndpointConnection
type DisableVpcEndpointConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableVpcEndpointConnectionRequest creates a request to invoke DisableVpcEndpointConnection API
func CreateDisableVpcEndpointConnectionRequest() (request *DisableVpcEndpointConnectionRequest) {
	request = &DisableVpcEndpointConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Privatelink", "2020-04-15", "DisableVpcEndpointConnection", "privatelink", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableVpcEndpointConnectionResponse creates a response to parse from DisableVpcEndpointConnection response
func CreateDisableVpcEndpointConnectionResponse() (response *DisableVpcEndpointConnectionResponse) {
	response = &DisableVpcEndpointConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
