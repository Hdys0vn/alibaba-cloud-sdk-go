package iot

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

// WriteDevicesHotStorageData invokes the iot.WriteDevicesHotStorageData API synchronously
func (client *Client) WriteDevicesHotStorageData(request *WriteDevicesHotStorageDataRequest) (response *WriteDevicesHotStorageDataResponse, err error) {
	response = CreateWriteDevicesHotStorageDataResponse()
	err = client.DoAction(request, response)
	return
}

// WriteDevicesHotStorageDataWithChan invokes the iot.WriteDevicesHotStorageData API asynchronously
func (client *Client) WriteDevicesHotStorageDataWithChan(request *WriteDevicesHotStorageDataRequest) (<-chan *WriteDevicesHotStorageDataResponse, <-chan error) {
	responseChan := make(chan *WriteDevicesHotStorageDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.WriteDevicesHotStorageData(request)
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

// WriteDevicesHotStorageDataWithCallback invokes the iot.WriteDevicesHotStorageData API asynchronously
func (client *Client) WriteDevicesHotStorageDataWithCallback(request *WriteDevicesHotStorageDataRequest, callback func(response *WriteDevicesHotStorageDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *WriteDevicesHotStorageDataResponse
		var err error
		defer close(result)
		response, err = client.WriteDevicesHotStorageData(request)
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

// WriteDevicesHotStorageDataRequest is the request struct for api WriteDevicesHotStorageData
type WriteDevicesHotStorageDataRequest struct {
	*requests.RpcRequest
	RealTenantId      string `position:"Query" name:"RealTenantId"`
	RealTripartiteKey string `position:"Query" name:"RealTripartiteKey"`
	IotId             string `position:"Query" name:"IotId"`
	IotInstanceId     string `position:"Query" name:"IotInstanceId"`
	UserTopic         string `position:"Query" name:"UserTopic"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	ApiProduct        string `position:"Body" name:"ApiProduct"`
	ApiRevision       string `position:"Body" name:"ApiRevision"`
	DeviceName        string `position:"Query" name:"DeviceName"`
	Items             string `position:"Query" name:"Items"`
}

// WriteDevicesHotStorageDataResponse is the response struct for api WriteDevicesHotStorageData
type WriteDevicesHotStorageDataResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string `json:"Code" xml:"Code"`
}

// CreateWriteDevicesHotStorageDataRequest creates a request to invoke WriteDevicesHotStorageData API
func CreateWriteDevicesHotStorageDataRequest() (request *WriteDevicesHotStorageDataRequest) {
	request = &WriteDevicesHotStorageDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "WriteDevicesHotStorageData", "", "")
	request.Method = requests.POST
	return
}

// CreateWriteDevicesHotStorageDataResponse creates a response to parse from WriteDevicesHotStorageData response
func CreateWriteDevicesHotStorageDataResponse() (response *WriteDevicesHotStorageDataResponse) {
	response = &WriteDevicesHotStorageDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
