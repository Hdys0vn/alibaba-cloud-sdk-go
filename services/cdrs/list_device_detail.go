package cdrs

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

// ListDeviceDetail invokes the cdrs.ListDeviceDetail API synchronously
func (client *Client) ListDeviceDetail(request *ListDeviceDetailRequest) (response *ListDeviceDetailResponse, err error) {
	response = CreateListDeviceDetailResponse()
	err = client.DoAction(request, response)
	return
}

// ListDeviceDetailWithChan invokes the cdrs.ListDeviceDetail API asynchronously
func (client *Client) ListDeviceDetailWithChan(request *ListDeviceDetailRequest) (<-chan *ListDeviceDetailResponse, <-chan error) {
	responseChan := make(chan *ListDeviceDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDeviceDetail(request)
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

// ListDeviceDetailWithCallback invokes the cdrs.ListDeviceDetail API asynchronously
func (client *Client) ListDeviceDetailWithCallback(request *ListDeviceDetailRequest, callback func(response *ListDeviceDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDeviceDetailResponse
		var err error
		defer close(result)
		response, err = client.ListDeviceDetail(request)
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

// ListDeviceDetailRequest is the request struct for api ListDeviceDetail
type ListDeviceDetailRequest struct {
	*requests.RpcRequest
	CorpId       string           `position:"Body" name:"CorpId"`
	PageNumber   requests.Integer `position:"Body" name:"PageNumber"`
	DataSourceId string           `position:"Body" name:"DataSourceId"`
	PageSize     requests.Integer `position:"Body" name:"PageSize"`
}

// ListDeviceDetailResponse is the response struct for api ListDeviceDetail
type ListDeviceDetailResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListDeviceDetailRequest creates a request to invoke ListDeviceDetail API
func CreateListDeviceDetailRequest() (request *ListDeviceDetailRequest) {
	request = &ListDeviceDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListDeviceDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateListDeviceDetailResponse creates a response to parse from ListDeviceDetail response
func CreateListDeviceDetailResponse() (response *ListDeviceDetailResponse) {
	response = &ListDeviceDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
