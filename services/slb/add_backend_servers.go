package slb

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

// AddBackendServers invokes the slb.AddBackendServers API synchronously
func (client *Client) AddBackendServers(request *AddBackendServersRequest) (response *AddBackendServersResponse, err error) {
	response = CreateAddBackendServersResponse()
	err = client.DoAction(request, response)
	return
}

// AddBackendServersWithChan invokes the slb.AddBackendServers API asynchronously
func (client *Client) AddBackendServersWithChan(request *AddBackendServersRequest) (<-chan *AddBackendServersResponse, <-chan error) {
	responseChan := make(chan *AddBackendServersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddBackendServers(request)
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

// AddBackendServersWithCallback invokes the slb.AddBackendServers API asynchronously
func (client *Client) AddBackendServersWithCallback(request *AddBackendServersRequest, callback func(response *AddBackendServersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddBackendServersResponse
		var err error
		defer close(result)
		response, err = client.AddBackendServers(request)
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

// AddBackendServersRequest is the request struct for api AddBackendServers
type AddBackendServersRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BackendServers       string           `position:"Query" name:"BackendServers"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
}

// AddBackendServersResponse is the response struct for api AddBackendServers
type AddBackendServersResponse struct {
	*responses.BaseResponse
	LoadBalancerId string                            `json:"LoadBalancerId" xml:"LoadBalancerId"`
	RequestId      string                            `json:"RequestId" xml:"RequestId"`
	BackendServers BackendServersInAddBackendServers `json:"BackendServers" xml:"BackendServers"`
}

// CreateAddBackendServersRequest creates a request to invoke AddBackendServers API
func CreateAddBackendServersRequest() (request *AddBackendServersRequest) {
	request = &AddBackendServersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "AddBackendServers", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddBackendServersResponse creates a response to parse from AddBackendServers response
func CreateAddBackendServersResponse() (response *AddBackendServersResponse) {
	response = &AddBackendServersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
