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

// ListOTAJobByFirmware invokes the iot.ListOTAJobByFirmware API synchronously
func (client *Client) ListOTAJobByFirmware(request *ListOTAJobByFirmwareRequest) (response *ListOTAJobByFirmwareResponse, err error) {
	response = CreateListOTAJobByFirmwareResponse()
	err = client.DoAction(request, response)
	return
}

// ListOTAJobByFirmwareWithChan invokes the iot.ListOTAJobByFirmware API asynchronously
func (client *Client) ListOTAJobByFirmwareWithChan(request *ListOTAJobByFirmwareRequest) (<-chan *ListOTAJobByFirmwareResponse, <-chan error) {
	responseChan := make(chan *ListOTAJobByFirmwareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOTAJobByFirmware(request)
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

// ListOTAJobByFirmwareWithCallback invokes the iot.ListOTAJobByFirmware API asynchronously
func (client *Client) ListOTAJobByFirmwareWithCallback(request *ListOTAJobByFirmwareRequest, callback func(response *ListOTAJobByFirmwareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOTAJobByFirmwareResponse
		var err error
		defer close(result)
		response, err = client.ListOTAJobByFirmware(request)
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

// ListOTAJobByFirmwareRequest is the request struct for api ListOTAJobByFirmware
type ListOTAJobByFirmwareRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	FirmwareId    string           `position:"Query" name:"FirmwareId"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// ListOTAJobByFirmwareResponse is the response struct for api ListOTAJobByFirmware
type ListOTAJobByFirmwareResponse struct {
	*responses.BaseResponse
	RequestId    string                     `json:"RequestId" xml:"RequestId"`
	Success      bool                       `json:"Success" xml:"Success"`
	Code         string                     `json:"Code" xml:"Code"`
	ErrorMessage string                     `json:"ErrorMessage" xml:"ErrorMessage"`
	Total        int                        `json:"Total" xml:"Total"`
	PageSize     int                        `json:"PageSize" xml:"PageSize"`
	PageCount    int                        `json:"PageCount" xml:"PageCount"`
	CurrentPage  int                        `json:"CurrentPage" xml:"CurrentPage"`
	Data         DataInListOTAJobByFirmware `json:"Data" xml:"Data"`
}

// CreateListOTAJobByFirmwareRequest creates a request to invoke ListOTAJobByFirmware API
func CreateListOTAJobByFirmwareRequest() (request *ListOTAJobByFirmwareRequest) {
	request = &ListOTAJobByFirmwareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "ListOTAJobByFirmware", "", "")
	request.Method = requests.POST
	return
}

// CreateListOTAJobByFirmwareResponse creates a response to parse from ListOTAJobByFirmware response
func CreateListOTAJobByFirmwareResponse() (response *ListOTAJobByFirmwareResponse) {
	response = &ListOTAJobByFirmwareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
