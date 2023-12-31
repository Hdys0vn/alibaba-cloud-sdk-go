package cs

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

// ModifyClusterConfiguration invokes the cs.ModifyClusterConfiguration API synchronously
func (client *Client) ModifyClusterConfiguration(request *ModifyClusterConfigurationRequest) (response *ModifyClusterConfigurationResponse, err error) {
	response = CreateModifyClusterConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyClusterConfigurationWithChan invokes the cs.ModifyClusterConfiguration API asynchronously
func (client *Client) ModifyClusterConfigurationWithChan(request *ModifyClusterConfigurationRequest) (<-chan *ModifyClusterConfigurationResponse, <-chan error) {
	responseChan := make(chan *ModifyClusterConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyClusterConfiguration(request)
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

// ModifyClusterConfigurationWithCallback invokes the cs.ModifyClusterConfiguration API asynchronously
func (client *Client) ModifyClusterConfigurationWithCallback(request *ModifyClusterConfigurationRequest, callback func(response *ModifyClusterConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyClusterConfigurationResponse
		var err error
		defer close(result)
		response, err = client.ModifyClusterConfiguration(request)
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

// ModifyClusterConfigurationRequest is the request struct for api ModifyClusterConfiguration
type ModifyClusterConfigurationRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

// ModifyClusterConfigurationResponse is the response struct for api ModifyClusterConfiguration
type ModifyClusterConfigurationResponse struct {
	*responses.BaseResponse
}

// CreateModifyClusterConfigurationRequest creates a request to invoke ModifyClusterConfiguration API
func CreateModifyClusterConfigurationRequest() (request *ModifyClusterConfigurationRequest) {
	request = &ModifyClusterConfigurationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "ModifyClusterConfiguration", "/clusters/[ClusterId]/configuration", "", "")
	request.Method = requests.PUT
	return
}

// CreateModifyClusterConfigurationResponse creates a response to parse from ModifyClusterConfiguration response
func CreateModifyClusterConfigurationResponse() (response *ModifyClusterConfigurationResponse) {
	response = &ModifyClusterConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
