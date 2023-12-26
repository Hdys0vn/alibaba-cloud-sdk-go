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

// QueryDeviceOriginalEventData invokes the iot.QueryDeviceOriginalEventData API synchronously
func (client *Client) QueryDeviceOriginalEventData(request *QueryDeviceOriginalEventDataRequest) (response *QueryDeviceOriginalEventDataResponse, err error) {
	response = CreateQueryDeviceOriginalEventDataResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceOriginalEventDataWithChan invokes the iot.QueryDeviceOriginalEventData API asynchronously
func (client *Client) QueryDeviceOriginalEventDataWithChan(request *QueryDeviceOriginalEventDataRequest) (<-chan *QueryDeviceOriginalEventDataResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceOriginalEventDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceOriginalEventData(request)
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

// QueryDeviceOriginalEventDataWithCallback invokes the iot.QueryDeviceOriginalEventData API asynchronously
func (client *Client) QueryDeviceOriginalEventDataWithCallback(request *QueryDeviceOriginalEventDataRequest, callback func(response *QueryDeviceOriginalEventDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceOriginalEventDataResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceOriginalEventData(request)
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

// QueryDeviceOriginalEventDataRequest is the request struct for api QueryDeviceOriginalEventData
type QueryDeviceOriginalEventDataRequest struct {
	*requests.RpcRequest
	NextPageToken string           `position:"Query" name:"NextPageToken"`
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	IotId         string           `position:"Query" name:"IotId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	Identifier    string           `position:"Query" name:"Identifier"`
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	Asc           requests.Integer `position:"Query" name:"Asc"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
}

// QueryDeviceOriginalEventDataResponse is the response struct for api QueryDeviceOriginalEventData
type QueryDeviceOriginalEventDataResponse struct {
	*responses.BaseResponse
	RequestId    string                             `json:"RequestId" xml:"RequestId"`
	Success      bool                               `json:"Success" xml:"Success"`
	Code         string                             `json:"Code" xml:"Code"`
	ErrorMessage string                             `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryDeviceOriginalEventData `json:"Data" xml:"Data"`
}

// CreateQueryDeviceOriginalEventDataRequest creates a request to invoke QueryDeviceOriginalEventData API
func CreateQueryDeviceOriginalEventDataRequest() (request *QueryDeviceOriginalEventDataRequest) {
	request = &QueryDeviceOriginalEventDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceOriginalEventData", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceOriginalEventDataResponse creates a response to parse from QueryDeviceOriginalEventData response
func CreateQueryDeviceOriginalEventDataResponse() (response *QueryDeviceOriginalEventDataResponse) {
	response = &QueryDeviceOriginalEventDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
