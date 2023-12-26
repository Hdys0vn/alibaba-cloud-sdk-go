package sas

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

// ModifyPropertyScheduleConfig invokes the sas.ModifyPropertyScheduleConfig API synchronously
func (client *Client) ModifyPropertyScheduleConfig(request *ModifyPropertyScheduleConfigRequest) (response *ModifyPropertyScheduleConfigResponse, err error) {
	response = CreateModifyPropertyScheduleConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPropertyScheduleConfigWithChan invokes the sas.ModifyPropertyScheduleConfig API asynchronously
func (client *Client) ModifyPropertyScheduleConfigWithChan(request *ModifyPropertyScheduleConfigRequest) (<-chan *ModifyPropertyScheduleConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyPropertyScheduleConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPropertyScheduleConfig(request)
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

// ModifyPropertyScheduleConfigWithCallback invokes the sas.ModifyPropertyScheduleConfig API asynchronously
func (client *Client) ModifyPropertyScheduleConfigWithCallback(request *ModifyPropertyScheduleConfigRequest, callback func(response *ModifyPropertyScheduleConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPropertyScheduleConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyPropertyScheduleConfig(request)
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

// ModifyPropertyScheduleConfigRequest is the request struct for api ModifyPropertyScheduleConfig
type ModifyPropertyScheduleConfigRequest struct {
	*requests.RpcRequest
	Type         string `position:"Query" name:"Type"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	ScheduleTime string `position:"Query" name:"ScheduleTime"`
}

// ModifyPropertyScheduleConfigResponse is the response struct for api ModifyPropertyScheduleConfig
type ModifyPropertyScheduleConfigResponse struct {
	*responses.BaseResponse
	ModifyResult bool   `json:"ModifyResult" xml:"ModifyResult"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyPropertyScheduleConfigRequest creates a request to invoke ModifyPropertyScheduleConfig API
func CreateModifyPropertyScheduleConfigRequest() (request *ModifyPropertyScheduleConfigRequest) {
	request = &ModifyPropertyScheduleConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyPropertyScheduleConfig", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyPropertyScheduleConfigResponse creates a response to parse from ModifyPropertyScheduleConfig response
func CreateModifyPropertyScheduleConfigResponse() (response *ModifyPropertyScheduleConfigResponse) {
	response = &ModifyPropertyScheduleConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
