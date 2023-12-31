package vs

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

// BatchStartDevices invokes the vs.BatchStartDevices API synchronously
func (client *Client) BatchStartDevices(request *BatchStartDevicesRequest) (response *BatchStartDevicesResponse, err error) {
	response = CreateBatchStartDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// BatchStartDevicesWithChan invokes the vs.BatchStartDevices API asynchronously
func (client *Client) BatchStartDevicesWithChan(request *BatchStartDevicesRequest) (<-chan *BatchStartDevicesResponse, <-chan error) {
	responseChan := make(chan *BatchStartDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchStartDevices(request)
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

// BatchStartDevicesWithCallback invokes the vs.BatchStartDevices API asynchronously
func (client *Client) BatchStartDevicesWithCallback(request *BatchStartDevicesRequest, callback func(response *BatchStartDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchStartDevicesResponse
		var err error
		defer close(result)
		response, err = client.BatchStartDevices(request)
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

// BatchStartDevicesRequest is the request struct for api BatchStartDevices
type BatchStartDevicesRequest struct {
	*requests.RpcRequest
	Id      string           `position:"Query" name:"Id"`
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// BatchStartDevicesResponse is the response struct for api BatchStartDevices
type BatchStartDevicesResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Results   []Result `json:"Results" xml:"Results"`
}

// CreateBatchStartDevicesRequest creates a request to invoke BatchStartDevices API
func CreateBatchStartDevicesRequest() (request *BatchStartDevicesRequest) {
	request = &BatchStartDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "BatchStartDevices", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchStartDevicesResponse creates a response to parse from BatchStartDevices response
func CreateBatchStartDevicesResponse() (response *BatchStartDevicesResponse) {
	response = &BatchStartDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
