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

// AddVServerGroupBackendServers invokes the slb.AddVServerGroupBackendServers API synchronously
func (client *Client) AddVServerGroupBackendServers(request *AddVServerGroupBackendServersRequest) (response *AddVServerGroupBackendServersResponse, err error) {
	response = CreateAddVServerGroupBackendServersResponse()
	err = client.DoAction(request, response)
	return
}

// AddVServerGroupBackendServersWithChan invokes the slb.AddVServerGroupBackendServers API asynchronously
func (client *Client) AddVServerGroupBackendServersWithChan(request *AddVServerGroupBackendServersRequest) (<-chan *AddVServerGroupBackendServersResponse, <-chan error) {
	responseChan := make(chan *AddVServerGroupBackendServersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddVServerGroupBackendServers(request)
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

// AddVServerGroupBackendServersWithCallback invokes the slb.AddVServerGroupBackendServers API asynchronously
func (client *Client) AddVServerGroupBackendServersWithCallback(request *AddVServerGroupBackendServersRequest, callback func(response *AddVServerGroupBackendServersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddVServerGroupBackendServersResponse
		var err error
		defer close(result)
		response, err = client.AddVServerGroupBackendServers(request)
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

// AddVServerGroupBackendServersRequest is the request struct for api AddVServerGroupBackendServers
type AddVServerGroupBackendServersRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BackendServers       string           `position:"Query" name:"BackendServers"`
	VServerGroupId       string           `position:"Query" name:"VServerGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
}

// AddVServerGroupBackendServersResponse is the response struct for api AddVServerGroupBackendServers
type AddVServerGroupBackendServersResponse struct {
	*responses.BaseResponse
	VServerGroupId string                                        `json:"VServerGroupId" xml:"VServerGroupId"`
	RequestId      string                                        `json:"RequestId" xml:"RequestId"`
	BackendServers BackendServersInAddVServerGroupBackendServers `json:"BackendServers" xml:"BackendServers"`
}

// CreateAddVServerGroupBackendServersRequest creates a request to invoke AddVServerGroupBackendServers API
func CreateAddVServerGroupBackendServersRequest() (request *AddVServerGroupBackendServersRequest) {
	request = &AddVServerGroupBackendServersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "AddVServerGroupBackendServers", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddVServerGroupBackendServersResponse creates a response to parse from AddVServerGroupBackendServers response
func CreateAddVServerGroupBackendServersResponse() (response *AddVServerGroupBackendServersResponse) {
	response = &AddVServerGroupBackendServersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}