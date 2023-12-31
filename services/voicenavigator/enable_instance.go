package voicenavigator

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

// EnableInstance invokes the voicenavigator.EnableInstance API synchronously
func (client *Client) EnableInstance(request *EnableInstanceRequest) (response *EnableInstanceResponse, err error) {
	response = CreateEnableInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// EnableInstanceWithChan invokes the voicenavigator.EnableInstance API asynchronously
func (client *Client) EnableInstanceWithChan(request *EnableInstanceRequest) (<-chan *EnableInstanceResponse, <-chan error) {
	responseChan := make(chan *EnableInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableInstance(request)
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

// EnableInstanceWithCallback invokes the voicenavigator.EnableInstance API asynchronously
func (client *Client) EnableInstanceWithCallback(request *EnableInstanceRequest, callback func(response *EnableInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableInstanceResponse
		var err error
		defer close(result)
		response, err = client.EnableInstance(request)
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

// EnableInstanceRequest is the request struct for api EnableInstance
type EnableInstanceRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// EnableInstanceResponse is the response struct for api EnableInstance
type EnableInstanceResponse struct {
	*responses.BaseResponse
	Status    string `json:"Status" xml:"Status"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableInstanceRequest creates a request to invoke EnableInstance API
func CreateEnableInstanceRequest() (request *EnableInstanceRequest) {
	request = &EnableInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "EnableInstance", "voicebot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableInstanceResponse creates a response to parse from EnableInstance response
func CreateEnableInstanceResponse() (response *EnableInstanceResponse) {
	response = &EnableInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
