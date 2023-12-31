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

// ListCityMapImageDetails invokes the cdrs.ListCityMapImageDetails API synchronously
func (client *Client) ListCityMapImageDetails(request *ListCityMapImageDetailsRequest) (response *ListCityMapImageDetailsResponse, err error) {
	response = CreateListCityMapImageDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCityMapImageDetailsWithChan invokes the cdrs.ListCityMapImageDetails API asynchronously
func (client *Client) ListCityMapImageDetailsWithChan(request *ListCityMapImageDetailsRequest) (<-chan *ListCityMapImageDetailsResponse, <-chan error) {
	responseChan := make(chan *ListCityMapImageDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCityMapImageDetails(request)
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

// ListCityMapImageDetailsWithCallback invokes the cdrs.ListCityMapImageDetails API asynchronously
func (client *Client) ListCityMapImageDetailsWithCallback(request *ListCityMapImageDetailsRequest, callback func(response *ListCityMapImageDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCityMapImageDetailsResponse
		var err error
		defer close(result)
		response, err = client.ListCityMapImageDetails(request)
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

// ListCityMapImageDetailsRequest is the request struct for api ListCityMapImageDetails
type ListCityMapImageDetailsRequest struct {
	*requests.RpcRequest
	TimeInterval string           `position:"Body" name:"TimeInterval"`
	RecordNumber requests.Integer `position:"Body" name:"RecordNumber"`
	DataSourceId string           `position:"Body" name:"DataSourceId"`
}

// ListCityMapImageDetailsResponse is the response struct for api ListCityMapImageDetails
type ListCityMapImageDetailsResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListCityMapImageDetailsRequest creates a request to invoke ListCityMapImageDetails API
func CreateListCityMapImageDetailsRequest() (request *ListCityMapImageDetailsRequest) {
	request = &ListCityMapImageDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListCityMapImageDetails", "", "")
	request.Method = requests.POST
	return
}

// CreateListCityMapImageDetailsResponse creates a response to parse from ListCityMapImageDetails response
func CreateListCityMapImageDetailsResponse() (response *ListCityMapImageDetailsResponse) {
	response = &ListCityMapImageDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
