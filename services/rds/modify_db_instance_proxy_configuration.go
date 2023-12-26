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

// ModifyDBInstanceProxyConfiguration invokes the rds.ModifyDBInstanceProxyConfiguration API synchronously
func (client *Client) ModifyDBInstanceProxyConfiguration(request *ModifyDBInstanceProxyConfigurationRequest) (response *ModifyDBInstanceProxyConfigurationResponse, err error) {
	response = CreateModifyDBInstanceProxyConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceProxyConfigurationWithChan invokes the rds.ModifyDBInstanceProxyConfiguration API asynchronously
func (client *Client) ModifyDBInstanceProxyConfigurationWithChan(request *ModifyDBInstanceProxyConfigurationRequest) (<-chan *ModifyDBInstanceProxyConfigurationResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceProxyConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceProxyConfiguration(request)
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

// ModifyDBInstanceProxyConfigurationWithCallback invokes the rds.ModifyDBInstanceProxyConfiguration API asynchronously
func (client *Client) ModifyDBInstanceProxyConfigurationWithCallback(request *ModifyDBInstanceProxyConfigurationRequest, callback func(response *ModifyDBInstanceProxyConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceProxyConfigurationResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceProxyConfiguration(request)
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

// ModifyDBInstanceProxyConfigurationRequest is the request struct for api ModifyDBInstanceProxyConfiguration
type ModifyDBInstanceProxyConfigurationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	ProxyConfigurationValue string           `position:"Query" name:"ProxyConfigurationValue"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	ProxyConfigurationKey   string           `position:"Query" name:"ProxyConfigurationKey"`
	DBInstanceId            string           `position:"Query" name:"DBInstanceId"`
}

// ModifyDBInstanceProxyConfigurationResponse is the response struct for api ModifyDBInstanceProxyConfiguration
type ModifyDBInstanceProxyConfigurationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceProxyConfigurationRequest creates a request to invoke ModifyDBInstanceProxyConfiguration API
func CreateModifyDBInstanceProxyConfigurationRequest() (request *ModifyDBInstanceProxyConfigurationRequest) {
	request = &ModifyDBInstanceProxyConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceProxyConfiguration", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstanceProxyConfigurationResponse creates a response to parse from ModifyDBInstanceProxyConfiguration response
func CreateModifyDBInstanceProxyConfigurationResponse() (response *ModifyDBInstanceProxyConfigurationResponse) {
	response = &ModifyDBInstanceProxyConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
