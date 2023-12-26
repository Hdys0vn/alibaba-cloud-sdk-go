package rds

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

// CreateDBProxyEndpointAddress invokes the rds.CreateDBProxyEndpointAddress API synchronously
func (client *Client) CreateDBProxyEndpointAddress(request *CreateDBProxyEndpointAddressRequest) (response *CreateDBProxyEndpointAddressResponse, err error) {
	response = CreateCreateDBProxyEndpointAddressResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDBProxyEndpointAddressWithChan invokes the rds.CreateDBProxyEndpointAddress API asynchronously
func (client *Client) CreateDBProxyEndpointAddressWithChan(request *CreateDBProxyEndpointAddressRequest) (<-chan *CreateDBProxyEndpointAddressResponse, <-chan error) {
	responseChan := make(chan *CreateDBProxyEndpointAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBProxyEndpointAddress(request)
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

// CreateDBProxyEndpointAddressWithCallback invokes the rds.CreateDBProxyEndpointAddress API asynchronously
func (client *Client) CreateDBProxyEndpointAddressWithCallback(request *CreateDBProxyEndpointAddressRequest, callback func(response *CreateDBProxyEndpointAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBProxyEndpointAddressResponse
		var err error
		defer close(result)
		response, err = client.CreateDBProxyEndpointAddress(request)
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

// CreateDBProxyEndpointAddressRequest is the request struct for api CreateDBProxyEndpointAddress
type CreateDBProxyEndpointAddressRequest struct {
	*requests.RpcRequest
	ConnectionStringPrefix      string `position:"Query" name:"ConnectionStringPrefix"`
	DBProxyConnectStringNetType string `position:"Query" name:"DBProxyConnectStringNetType"`
	ResourceGroupId             string `position:"Query" name:"ResourceGroupId"`
	DBInstanceId                string `position:"Query" name:"DBInstanceId"`
	DBProxyNewConnectStringPort string `position:"Query" name:"DBProxyNewConnectStringPort"`
	DBProxyEngineType           string `position:"Query" name:"DBProxyEngineType"`
	VSwitchId                   string `position:"Query" name:"VSwitchId"`
	DBProxyEndpointId           string `position:"Query" name:"DBProxyEndpointId"`
	VPCId                       string `position:"Query" name:"VPCId"`
}

// CreateDBProxyEndpointAddressResponse is the response struct for api CreateDBProxyEndpointAddress
type CreateDBProxyEndpointAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDBProxyEndpointAddressRequest creates a request to invoke CreateDBProxyEndpointAddress API
func CreateCreateDBProxyEndpointAddressRequest() (request *CreateDBProxyEndpointAddressRequest) {
	request = &CreateDBProxyEndpointAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateDBProxyEndpointAddress", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDBProxyEndpointAddressResponse creates a response to parse from CreateDBProxyEndpointAddress response
func CreateCreateDBProxyEndpointAddressResponse() (response *CreateDBProxyEndpointAddressResponse) {
	response = &CreateDBProxyEndpointAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
