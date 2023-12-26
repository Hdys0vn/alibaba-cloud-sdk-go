package linkvisual

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

// QueryEventRecordPlanDeviceByDevice invokes the linkvisual.QueryEventRecordPlanDeviceByDevice API synchronously
func (client *Client) QueryEventRecordPlanDeviceByDevice(request *QueryEventRecordPlanDeviceByDeviceRequest) (response *QueryEventRecordPlanDeviceByDeviceResponse, err error) {
	response = CreateQueryEventRecordPlanDeviceByDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEventRecordPlanDeviceByDeviceWithChan invokes the linkvisual.QueryEventRecordPlanDeviceByDevice API asynchronously
func (client *Client) QueryEventRecordPlanDeviceByDeviceWithChan(request *QueryEventRecordPlanDeviceByDeviceRequest) (<-chan *QueryEventRecordPlanDeviceByDeviceResponse, <-chan error) {
	responseChan := make(chan *QueryEventRecordPlanDeviceByDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEventRecordPlanDeviceByDevice(request)
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

// QueryEventRecordPlanDeviceByDeviceWithCallback invokes the linkvisual.QueryEventRecordPlanDeviceByDevice API asynchronously
func (client *Client) QueryEventRecordPlanDeviceByDeviceWithCallback(request *QueryEventRecordPlanDeviceByDeviceRequest, callback func(response *QueryEventRecordPlanDeviceByDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEventRecordPlanDeviceByDeviceResponse
		var err error
		defer close(result)
		response, err = client.QueryEventRecordPlanDeviceByDevice(request)
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

// QueryEventRecordPlanDeviceByDeviceRequest is the request struct for api QueryEventRecordPlanDeviceByDevice
type QueryEventRecordPlanDeviceByDeviceRequest struct {
	*requests.RpcRequest
	IotId         string           `position:"Query" name:"IotId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	StreamType    requests.Integer `position:"Query" name:"StreamType"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
}

// QueryEventRecordPlanDeviceByDeviceResponse is the response struct for api QueryEventRecordPlanDeviceByDevice
type QueryEventRecordPlanDeviceByDeviceResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryEventRecordPlanDeviceByDeviceRequest creates a request to invoke QueryEventRecordPlanDeviceByDevice API
func CreateQueryEventRecordPlanDeviceByDeviceRequest() (request *QueryEventRecordPlanDeviceByDeviceRequest) {
	request = &QueryEventRecordPlanDeviceByDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryEventRecordPlanDeviceByDevice", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryEventRecordPlanDeviceByDeviceResponse creates a response to parse from QueryEventRecordPlanDeviceByDevice response
func CreateQueryEventRecordPlanDeviceByDeviceResponse() (response *QueryEventRecordPlanDeviceByDeviceResponse) {
	response = &QueryEventRecordPlanDeviceByDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
