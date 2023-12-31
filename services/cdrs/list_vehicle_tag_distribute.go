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

// ListVehicleTagDistribute invokes the cdrs.ListVehicleTagDistribute API synchronously
func (client *Client) ListVehicleTagDistribute(request *ListVehicleTagDistributeRequest) (response *ListVehicleTagDistributeResponse, err error) {
	response = CreateListVehicleTagDistributeResponse()
	err = client.DoAction(request, response)
	return
}

// ListVehicleTagDistributeWithChan invokes the cdrs.ListVehicleTagDistribute API asynchronously
func (client *Client) ListVehicleTagDistributeWithChan(request *ListVehicleTagDistributeRequest) (<-chan *ListVehicleTagDistributeResponse, <-chan error) {
	responseChan := make(chan *ListVehicleTagDistributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVehicleTagDistribute(request)
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

// ListVehicleTagDistributeWithCallback invokes the cdrs.ListVehicleTagDistribute API asynchronously
func (client *Client) ListVehicleTagDistributeWithCallback(request *ListVehicleTagDistributeRequest, callback func(response *ListVehicleTagDistributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVehicleTagDistributeResponse
		var err error
		defer close(result)
		response, err = client.ListVehicleTagDistribute(request)
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

// ListVehicleTagDistributeRequest is the request struct for api ListVehicleTagDistribute
type ListVehicleTagDistributeRequest struct {
	*requests.RpcRequest
	CorpId    string `position:"Body" name:"CorpId"`
	EndTime   string `position:"Body" name:"EndTime"`
	StartTime string `position:"Body" name:"StartTime"`
	TagCode   string `position:"Body" name:"TagCode"`
}

// ListVehicleTagDistributeResponse is the response struct for api ListVehicleTagDistribute
type ListVehicleTagDistributeResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListVehicleTagDistributeRequest creates a request to invoke ListVehicleTagDistribute API
func CreateListVehicleTagDistributeRequest() (request *ListVehicleTagDistributeRequest) {
	request = &ListVehicleTagDistributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListVehicleTagDistribute", "", "")
	request.Method = requests.POST
	return
}

// CreateListVehicleTagDistributeResponse creates a response to parse from ListVehicleTagDistribute response
func CreateListVehicleTagDistributeResponse() (response *ListVehicleTagDistributeResponse) {
	response = &ListVehicleTagDistributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
