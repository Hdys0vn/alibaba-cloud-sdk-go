package ens

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

// GetDeviceInternetPort invokes the ens.GetDeviceInternetPort API synchronously
func (client *Client) GetDeviceInternetPort(request *GetDeviceInternetPortRequest) (response *GetDeviceInternetPortResponse, err error) {
	response = CreateGetDeviceInternetPortResponse()
	err = client.DoAction(request, response)
	return
}

// GetDeviceInternetPortWithChan invokes the ens.GetDeviceInternetPort API asynchronously
func (client *Client) GetDeviceInternetPortWithChan(request *GetDeviceInternetPortRequest) (<-chan *GetDeviceInternetPortResponse, <-chan error) {
	responseChan := make(chan *GetDeviceInternetPortResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDeviceInternetPort(request)
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

// GetDeviceInternetPortWithCallback invokes the ens.GetDeviceInternetPort API asynchronously
func (client *Client) GetDeviceInternetPortWithCallback(request *GetDeviceInternetPortRequest, callback func(response *GetDeviceInternetPortResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDeviceInternetPortResponse
		var err error
		defer close(result)
		response, err = client.GetDeviceInternetPort(request)
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

// GetDeviceInternetPortRequest is the request struct for api GetDeviceInternetPort
type GetDeviceInternetPortRequest struct {
	*requests.RpcRequest
	NatType    string `position:"Query" name:"NatType"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RuleId     string `position:"Query" name:"RuleId"`
}

// GetDeviceInternetPortResponse is the response struct for api GetDeviceInternetPort
type GetDeviceInternetPortResponse struct {
	*responses.BaseResponse
	RequestId   string            `json:"RequestId" xml:"RequestId"`
	InstanceId  string            `json:"InstanceId" xml:"InstanceId"`
	NetworkInfo []NetworkInfoItem `json:"NetworkInfo" xml:"NetworkInfo"`
}

// CreateGetDeviceInternetPortRequest creates a request to invoke GetDeviceInternetPort API
func CreateGetDeviceInternetPortRequest() (request *GetDeviceInternetPortRequest) {
	request = &GetDeviceInternetPortRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "GetDeviceInternetPort", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetDeviceInternetPortResponse creates a response to parse from GetDeviceInternetPort response
func CreateGetDeviceInternetPortResponse() (response *GetDeviceInternetPortResponse) {
	response = &GetDeviceInternetPortResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}