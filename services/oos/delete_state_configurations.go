package oos

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

// DeleteStateConfigurations invokes the oos.DeleteStateConfigurations API synchronously
func (client *Client) DeleteStateConfigurations(request *DeleteStateConfigurationsRequest) (response *DeleteStateConfigurationsResponse, err error) {
	response = CreateDeleteStateConfigurationsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteStateConfigurationsWithChan invokes the oos.DeleteStateConfigurations API asynchronously
func (client *Client) DeleteStateConfigurationsWithChan(request *DeleteStateConfigurationsRequest) (<-chan *DeleteStateConfigurationsResponse, <-chan error) {
	responseChan := make(chan *DeleteStateConfigurationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteStateConfigurations(request)
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

// DeleteStateConfigurationsWithCallback invokes the oos.DeleteStateConfigurations API asynchronously
func (client *Client) DeleteStateConfigurationsWithCallback(request *DeleteStateConfigurationsRequest, callback func(response *DeleteStateConfigurationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteStateConfigurationsResponse
		var err error
		defer close(result)
		response, err = client.DeleteStateConfigurations(request)
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

// DeleteStateConfigurationsRequest is the request struct for api DeleteStateConfigurations
type DeleteStateConfigurationsRequest struct {
	*requests.RpcRequest
	StateConfigurationIds string `position:"Query" name:"StateConfigurationIds"`
	ClientToken           string `position:"Query" name:"ClientToken"`
}

// DeleteStateConfigurationsResponse is the response struct for api DeleteStateConfigurations
type DeleteStateConfigurationsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteStateConfigurationsRequest creates a request to invoke DeleteStateConfigurations API
func CreateDeleteStateConfigurationsRequest() (request *DeleteStateConfigurationsRequest) {
	request = &DeleteStateConfigurationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "DeleteStateConfigurations", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteStateConfigurationsResponse creates a response to parse from DeleteStateConfigurations response
func CreateDeleteStateConfigurationsResponse() (response *DeleteStateConfigurationsResponse) {
	response = &DeleteStateConfigurationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
