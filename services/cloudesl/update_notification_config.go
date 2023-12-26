package cloudesl

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

// UpdateNotificationConfig invokes the cloudesl.UpdateNotificationConfig API synchronously
func (client *Client) UpdateNotificationConfig(request *UpdateNotificationConfigRequest) (response *UpdateNotificationConfigResponse, err error) {
	response = CreateUpdateNotificationConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateNotificationConfigWithChan invokes the cloudesl.UpdateNotificationConfig API asynchronously
func (client *Client) UpdateNotificationConfigWithChan(request *UpdateNotificationConfigRequest) (<-chan *UpdateNotificationConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateNotificationConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateNotificationConfig(request)
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

// UpdateNotificationConfigWithCallback invokes the cloudesl.UpdateNotificationConfig API asynchronously
func (client *Client) UpdateNotificationConfigWithCallback(request *UpdateNotificationConfigRequest, callback func(response *UpdateNotificationConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateNotificationConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateNotificationConfig(request)
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

// UpdateNotificationConfigRequest is the request struct for api UpdateNotificationConfig
type UpdateNotificationConfigRequest struct {
	*requests.RpcRequest
	Endpoint string           `position:"Body" name:"Endpoint"`
	Enable   requests.Boolean `position:"Body" name:"Enable"`
	Tag      string           `position:"Body" name:"Tag"`
	GroupId  string           `position:"Body" name:"GroupId"`
	Topic    string           `position:"Body" name:"Topic"`
}

// UpdateNotificationConfigResponse is the response struct for api UpdateNotificationConfig
type UpdateNotificationConfigResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
}

// CreateUpdateNotificationConfigRequest creates a request to invoke UpdateNotificationConfig API
func CreateUpdateNotificationConfigRequest() (request *UpdateNotificationConfigRequest) {
	request = &UpdateNotificationConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "UpdateNotificationConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateNotificationConfigResponse creates a response to parse from UpdateNotificationConfig response
func CreateUpdateNotificationConfigResponse() (response *UpdateNotificationConfigResponse) {
	response = &UpdateNotificationConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
