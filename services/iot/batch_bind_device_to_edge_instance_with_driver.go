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

// BatchBindDeviceToEdgeInstanceWithDriver invokes the iot.BatchBindDeviceToEdgeInstanceWithDriver API synchronously
func (client *Client) BatchBindDeviceToEdgeInstanceWithDriver(request *BatchBindDeviceToEdgeInstanceWithDriverRequest) (response *BatchBindDeviceToEdgeInstanceWithDriverResponse, err error) {
	response = CreateBatchBindDeviceToEdgeInstanceWithDriverResponse()
	err = client.DoAction(request, response)
	return
}

// BatchBindDeviceToEdgeInstanceWithDriverWithChan invokes the iot.BatchBindDeviceToEdgeInstanceWithDriver API asynchronously
func (client *Client) BatchBindDeviceToEdgeInstanceWithDriverWithChan(request *BatchBindDeviceToEdgeInstanceWithDriverRequest) (<-chan *BatchBindDeviceToEdgeInstanceWithDriverResponse, <-chan error) {
	responseChan := make(chan *BatchBindDeviceToEdgeInstanceWithDriverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchBindDeviceToEdgeInstanceWithDriver(request)
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

// BatchBindDeviceToEdgeInstanceWithDriverWithCallback invokes the iot.BatchBindDeviceToEdgeInstanceWithDriver API asynchronously
func (client *Client) BatchBindDeviceToEdgeInstanceWithDriverWithCallback(request *BatchBindDeviceToEdgeInstanceWithDriverRequest, callback func(response *BatchBindDeviceToEdgeInstanceWithDriverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchBindDeviceToEdgeInstanceWithDriverResponse
		var err error
		defer close(result)
		response, err = client.BatchBindDeviceToEdgeInstanceWithDriver(request)
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

// BatchBindDeviceToEdgeInstanceWithDriverRequest is the request struct for api BatchBindDeviceToEdgeInstanceWithDriver
type BatchBindDeviceToEdgeInstanceWithDriverRequest struct {
	*requests.RpcRequest
	DriverId      string    `position:"Query" name:"DriverId"`
	IotIds        *[]string `position:"Query" name:"IotIds"  type:"Repeated"`
	IotInstanceId string    `position:"Query" name:"IotInstanceId"`
	InstanceId    string    `position:"Query" name:"InstanceId"`
	ApiProduct    string    `position:"Body" name:"ApiProduct"`
	ApiRevision   string    `position:"Body" name:"ApiRevision"`
}

// BatchBindDeviceToEdgeInstanceWithDriverResponse is the response struct for api BatchBindDeviceToEdgeInstanceWithDriver
type BatchBindDeviceToEdgeInstanceWithDriverResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateBatchBindDeviceToEdgeInstanceWithDriverRequest creates a request to invoke BatchBindDeviceToEdgeInstanceWithDriver API
func CreateBatchBindDeviceToEdgeInstanceWithDriverRequest() (request *BatchBindDeviceToEdgeInstanceWithDriverRequest) {
	request = &BatchBindDeviceToEdgeInstanceWithDriverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchBindDeviceToEdgeInstanceWithDriver", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchBindDeviceToEdgeInstanceWithDriverResponse creates a response to parse from BatchBindDeviceToEdgeInstanceWithDriver response
func CreateBatchBindDeviceToEdgeInstanceWithDriverResponse() (response *BatchBindDeviceToEdgeInstanceWithDriverResponse) {
	response = &BatchBindDeviceToEdgeInstanceWithDriverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
