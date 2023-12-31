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

// ListPersonDetails invokes the cdrs.ListPersonDetails API synchronously
func (client *Client) ListPersonDetails(request *ListPersonDetailsRequest) (response *ListPersonDetailsResponse, err error) {
	response = CreateListPersonDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// ListPersonDetailsWithChan invokes the cdrs.ListPersonDetails API asynchronously
func (client *Client) ListPersonDetailsWithChan(request *ListPersonDetailsRequest) (<-chan *ListPersonDetailsResponse, <-chan error) {
	responseChan := make(chan *ListPersonDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPersonDetails(request)
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

// ListPersonDetailsWithCallback invokes the cdrs.ListPersonDetails API asynchronously
func (client *Client) ListPersonDetailsWithCallback(request *ListPersonDetailsRequest, callback func(response *ListPersonDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPersonDetailsResponse
		var err error
		defer close(result)
		response, err = client.ListPersonDetails(request)
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

// ListPersonDetailsRequest is the request struct for api ListPersonDetails
type ListPersonDetailsRequest struct {
	*requests.RpcRequest
	Schema     string           `position:"Body" name:"Schema"`
	CorpId     string           `position:"Body" name:"CorpId"`
	EndTime    string           `position:"Body" name:"EndTime"`
	StartTime  string           `position:"Body" name:"StartTime"`
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
	PersonId   string           `position:"Body" name:"PersonId"`
}

// ListPersonDetailsResponse is the response struct for api ListPersonDetails
type ListPersonDetailsResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListPersonDetailsRequest creates a request to invoke ListPersonDetails API
func CreateListPersonDetailsRequest() (request *ListPersonDetailsRequest) {
	request = &ListPersonDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListPersonDetails", "", "")
	request.Method = requests.POST
	return
}

// CreateListPersonDetailsResponse creates a response to parse from ListPersonDetails response
func CreateListPersonDetailsResponse() (response *ListPersonDetailsResponse) {
	response = &ListPersonDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
