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

// UpdateStateConfiguration invokes the oos.UpdateStateConfiguration API synchronously
func (client *Client) UpdateStateConfiguration(request *UpdateStateConfigurationRequest) (response *UpdateStateConfigurationResponse, err error) {
	response = CreateUpdateStateConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateStateConfigurationWithChan invokes the oos.UpdateStateConfiguration API asynchronously
func (client *Client) UpdateStateConfigurationWithChan(request *UpdateStateConfigurationRequest) (<-chan *UpdateStateConfigurationResponse, <-chan error) {
	responseChan := make(chan *UpdateStateConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateStateConfiguration(request)
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

// UpdateStateConfigurationWithCallback invokes the oos.UpdateStateConfiguration API asynchronously
func (client *Client) UpdateStateConfigurationWithCallback(request *UpdateStateConfigurationRequest, callback func(response *UpdateStateConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateStateConfigurationResponse
		var err error
		defer close(result)
		response, err = client.UpdateStateConfiguration(request)
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

// UpdateStateConfigurationRequest is the request struct for api UpdateStateConfiguration
type UpdateStateConfigurationRequest struct {
	*requests.RpcRequest
	ScheduleType         string `position:"Query" name:"ScheduleType"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Description          string `position:"Query" name:"Description"`
	Targets              string `position:"Query" name:"Targets"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ScheduleExpression   string `position:"Query" name:"ScheduleExpression"`
	ConfigureMode        string `position:"Query" name:"ConfigureMode"`
	Tags                 string `position:"Query" name:"Tags"`
	Parameters           string `position:"Query" name:"Parameters"`
	StateConfigurationId string `position:"Query" name:"StateConfigurationId"`
}

// UpdateStateConfigurationResponse is the response struct for api UpdateStateConfiguration
type UpdateStateConfigurationResponse struct {
	*responses.BaseResponse
	RequestId          string                   `json:"RequestId" xml:"RequestId"`
	StateConfiguration []StateConfigurationItem `json:"StateConfiguration" xml:"StateConfiguration"`
}

// CreateUpdateStateConfigurationRequest creates a request to invoke UpdateStateConfiguration API
func CreateUpdateStateConfigurationRequest() (request *UpdateStateConfigurationRequest) {
	request = &UpdateStateConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "UpdateStateConfiguration", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateStateConfigurationResponse creates a response to parse from UpdateStateConfiguration response
func CreateUpdateStateConfigurationResponse() (response *UpdateStateConfigurationResponse) {
	response = &UpdateStateConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
