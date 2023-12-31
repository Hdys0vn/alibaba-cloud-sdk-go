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

// ListVehicleDetails invokes the cdrs.ListVehicleDetails API synchronously
func (client *Client) ListVehicleDetails(request *ListVehicleDetailsRequest) (response *ListVehicleDetailsResponse, err error) {
	response = CreateListVehicleDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// ListVehicleDetailsWithChan invokes the cdrs.ListVehicleDetails API asynchronously
func (client *Client) ListVehicleDetailsWithChan(request *ListVehicleDetailsRequest) (<-chan *ListVehicleDetailsResponse, <-chan error) {
	responseChan := make(chan *ListVehicleDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVehicleDetails(request)
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

// ListVehicleDetailsWithCallback invokes the cdrs.ListVehicleDetails API asynchronously
func (client *Client) ListVehicleDetailsWithCallback(request *ListVehicleDetailsRequest, callback func(response *ListVehicleDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVehicleDetailsResponse
		var err error
		defer close(result)
		response, err = client.ListVehicleDetails(request)
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

// ListVehicleDetailsRequest is the request struct for api ListVehicleDetails
type ListVehicleDetailsRequest struct {
	*requests.RpcRequest
	PlateId    string `position:"Body" name:"PlateId"`
	CorpId     string `position:"Body" name:"CorpId"`
	EndTime    string `position:"Body" name:"EndTime"`
	StartTime  string `position:"Body" name:"StartTime"`
	PageNumber string `position:"Body" name:"PageNumber"`
	PageSize   string `position:"Body" name:"PageSize"`
}

// ListVehicleDetailsResponse is the response struct for api ListVehicleDetails
type ListVehicleDetailsResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListVehicleDetailsRequest creates a request to invoke ListVehicleDetails API
func CreateListVehicleDetailsRequest() (request *ListVehicleDetailsRequest) {
	request = &ListVehicleDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListVehicleDetails", "", "")
	request.Method = requests.POST
	return
}

// CreateListVehicleDetailsResponse creates a response to parse from ListVehicleDetails response
func CreateListVehicleDetailsResponse() (response *ListVehicleDetailsResponse) {
	response = &ListVehicleDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
