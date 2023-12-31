package cloudfw

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

// PutEnableAllFwSwitch invokes the cloudfw.PutEnableAllFwSwitch API synchronously
func (client *Client) PutEnableAllFwSwitch(request *PutEnableAllFwSwitchRequest) (response *PutEnableAllFwSwitchResponse, err error) {
	response = CreatePutEnableAllFwSwitchResponse()
	err = client.DoAction(request, response)
	return
}

// PutEnableAllFwSwitchWithChan invokes the cloudfw.PutEnableAllFwSwitch API asynchronously
func (client *Client) PutEnableAllFwSwitchWithChan(request *PutEnableAllFwSwitchRequest) (<-chan *PutEnableAllFwSwitchResponse, <-chan error) {
	responseChan := make(chan *PutEnableAllFwSwitchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutEnableAllFwSwitch(request)
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

// PutEnableAllFwSwitchWithCallback invokes the cloudfw.PutEnableAllFwSwitch API asynchronously
func (client *Client) PutEnableAllFwSwitchWithCallback(request *PutEnableAllFwSwitchRequest, callback func(response *PutEnableAllFwSwitchResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutEnableAllFwSwitchResponse
		var err error
		defer close(result)
		response, err = client.PutEnableAllFwSwitch(request)
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

// PutEnableAllFwSwitchRequest is the request struct for api PutEnableAllFwSwitch
type PutEnableAllFwSwitchRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// PutEnableAllFwSwitchResponse is the response struct for api PutEnableAllFwSwitch
type PutEnableAllFwSwitchResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePutEnableAllFwSwitchRequest creates a request to invoke PutEnableAllFwSwitch API
func CreatePutEnableAllFwSwitchRequest() (request *PutEnableAllFwSwitchRequest) {
	request = &PutEnableAllFwSwitchRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "PutEnableAllFwSwitch", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePutEnableAllFwSwitchResponse creates a response to parse from PutEnableAllFwSwitch response
func CreatePutEnableAllFwSwitchResponse() (response *PutEnableAllFwSwitchResponse) {
	response = &PutEnableAllFwSwitchResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
