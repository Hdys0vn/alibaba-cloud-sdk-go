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

// ExportMeasurementData invokes the ens.ExportMeasurementData API synchronously
func (client *Client) ExportMeasurementData(request *ExportMeasurementDataRequest) (response *ExportMeasurementDataResponse, err error) {
	response = CreateExportMeasurementDataResponse()
	err = client.DoAction(request, response)
	return
}

// ExportMeasurementDataWithChan invokes the ens.ExportMeasurementData API asynchronously
func (client *Client) ExportMeasurementDataWithChan(request *ExportMeasurementDataRequest) (<-chan *ExportMeasurementDataResponse, <-chan error) {
	responseChan := make(chan *ExportMeasurementDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportMeasurementData(request)
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

// ExportMeasurementDataWithCallback invokes the ens.ExportMeasurementData API asynchronously
func (client *Client) ExportMeasurementDataWithCallback(request *ExportMeasurementDataRequest, callback func(response *ExportMeasurementDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportMeasurementDataResponse
		var err error
		defer close(result)
		response, err = client.ExportMeasurementData(request)
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

// ExportMeasurementDataRequest is the request struct for api ExportMeasurementData
type ExportMeasurementDataRequest struct {
	*requests.RpcRequest
	StartDate string `position:"Query" name:"StartDate"`
	EndDate   string `position:"Query" name:"EndDate"`
}

// ExportMeasurementDataResponse is the response struct for api ExportMeasurementData
type ExportMeasurementDataResponse struct {
	*responses.BaseResponse
	FilePath  string `json:"FilePath" xml:"FilePath"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateExportMeasurementDataRequest creates a request to invoke ExportMeasurementData API
func CreateExportMeasurementDataRequest() (request *ExportMeasurementDataRequest) {
	request = &ExportMeasurementDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ExportMeasurementData", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportMeasurementDataResponse creates a response to parse from ExportMeasurementData response
func CreateExportMeasurementDataResponse() (response *ExportMeasurementDataResponse) {
	response = &ExportMeasurementDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
