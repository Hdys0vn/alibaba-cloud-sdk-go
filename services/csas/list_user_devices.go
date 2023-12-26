package csas

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

// ListUserDevices invokes the csas.ListUserDevices API synchronously
func (client *Client) ListUserDevices(request *ListUserDevicesRequest) (response *ListUserDevicesResponse, err error) {
	response = CreateListUserDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// ListUserDevicesWithChan invokes the csas.ListUserDevices API asynchronously
func (client *Client) ListUserDevicesWithChan(request *ListUserDevicesRequest) (<-chan *ListUserDevicesResponse, <-chan error) {
	responseChan := make(chan *ListUserDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUserDevices(request)
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

// ListUserDevicesWithCallback invokes the csas.ListUserDevices API asynchronously
func (client *Client) ListUserDevicesWithCallback(request *ListUserDevicesRequest, callback func(response *ListUserDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUserDevicesResponse
		var err error
		defer close(result)
		response, err = client.ListUserDevices(request)
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

// ListUserDevicesRequest is the request struct for api ListUserDevices
type ListUserDevicesRequest struct {
	*requests.RpcRequest
	Mac            string           `position:"Query" name:"Mac"`
	DeviceTypes    *[]string        `position:"Query" name:"DeviceTypes"  type:"Repeated"`
	Hostname       string           `position:"Query" name:"Hostname"`
	AppStatuses    *[]string        `position:"Query" name:"AppStatuses"  type:"Repeated"`
	DlpStatuses    *[]string        `position:"Query" name:"DlpStatuses"  type:"Repeated"`
	SourceIp       string           `position:"Query" name:"SourceIp"`
	SaseUserId     string           `position:"Query" name:"SaseUserId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	NacStatuses    *[]string        `position:"Query" name:"NacStatuses"  type:"Repeated"`
	Department     string           `position:"Query" name:"Department"`
	IaStatuses     *[]string        `position:"Query" name:"IaStatuses"  type:"Repeated"`
	DeviceBelong   string           `position:"Query" name:"DeviceBelong"`
	CurrentPage    requests.Integer `position:"Query" name:"CurrentPage"`
	SharingStatus  requests.Boolean `position:"Query" name:"SharingStatus"`
	DeviceTags     *[]string        `position:"Query" name:"DeviceTags"  type:"Repeated"`
	DeviceStatuses *[]string        `position:"Query" name:"DeviceStatuses"  type:"Repeated"`
	PaStatuses     *[]string        `position:"Query" name:"PaStatuses"  type:"Repeated"`
	SortBy         string           `position:"Query" name:"SortBy"`
	Username       string           `position:"Query" name:"Username"`
}

// ListUserDevicesResponse is the response struct for api ListUserDevices
type ListUserDevicesResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	TotalNum  int64      `json:"TotalNum" xml:"TotalNum"`
	Devices   []DataList `json:"Devices" xml:"Devices"`
}

// CreateListUserDevicesRequest creates a request to invoke ListUserDevices API
func CreateListUserDevicesRequest() (request *ListUserDevicesRequest) {
	request = &ListUserDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "ListUserDevices", "", "")
	request.Method = requests.GET
	return
}

// CreateListUserDevicesResponse creates a response to parse from ListUserDevices response
func CreateListUserDevicesResponse() (response *ListUserDevicesResponse) {
	response = &ListUserDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}