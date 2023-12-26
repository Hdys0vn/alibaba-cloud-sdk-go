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

// DeleteDBInstanceEndpointAddress invokes the rds.DeleteDBInstanceEndpointAddress API synchronously
func (client *Client) DeleteDBInstanceEndpointAddress(request *DeleteDBInstanceEndpointAddressRequest) (response *DeleteDBInstanceEndpointAddressResponse, err error) {
	response = CreateDeleteDBInstanceEndpointAddressResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDBInstanceEndpointAddressWithChan invokes the rds.DeleteDBInstanceEndpointAddress API asynchronously
func (client *Client) DeleteDBInstanceEndpointAddressWithChan(request *DeleteDBInstanceEndpointAddressRequest) (<-chan *DeleteDBInstanceEndpointAddressResponse, <-chan error) {
	responseChan := make(chan *DeleteDBInstanceEndpointAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDBInstanceEndpointAddress(request)
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

// DeleteDBInstanceEndpointAddressWithCallback invokes the rds.DeleteDBInstanceEndpointAddress API asynchronously
func (client *Client) DeleteDBInstanceEndpointAddressWithCallback(request *DeleteDBInstanceEndpointAddressRequest, callback func(response *DeleteDBInstanceEndpointAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDBInstanceEndpointAddressResponse
		var err error
		defer close(result)
		response, err = client.DeleteDBInstanceEndpointAddress(request)
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

// DeleteDBInstanceEndpointAddressRequest is the request struct for api DeleteDBInstanceEndpointAddress
type DeleteDBInstanceEndpointAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ConnectionString     string           `position:"Body" name:"ConnectionString"`
	DBInstanceEndpointId string           `position:"Body" name:"DBInstanceEndpointId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DeleteDBInstanceEndpointAddressResponse is the response struct for api DeleteDBInstanceEndpointAddress
type DeleteDBInstanceEndpointAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDeleteDBInstanceEndpointAddressRequest creates a request to invoke DeleteDBInstanceEndpointAddress API
func CreateDeleteDBInstanceEndpointAddressRequest() (request *DeleteDBInstanceEndpointAddressRequest) {
	request = &DeleteDBInstanceEndpointAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DeleteDBInstanceEndpointAddress", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteDBInstanceEndpointAddressResponse creates a response to parse from DeleteDBInstanceEndpointAddress response
func CreateDeleteDBInstanceEndpointAddressResponse() (response *DeleteDBInstanceEndpointAddressResponse) {
	response = &DeleteDBInstanceEndpointAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
