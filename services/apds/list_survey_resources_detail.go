package apds

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

// ListSurveyResourcesDetail invokes the apds.ListSurveyResourcesDetail API synchronously
func (client *Client) ListSurveyResourcesDetail(request *ListSurveyResourcesDetailRequest) (response *ListSurveyResourcesDetailResponse, err error) {
	response = CreateListSurveyResourcesDetailResponse()
	err = client.DoAction(request, response)
	return
}

// ListSurveyResourcesDetailWithChan invokes the apds.ListSurveyResourcesDetail API asynchronously
func (client *Client) ListSurveyResourcesDetailWithChan(request *ListSurveyResourcesDetailRequest) (<-chan *ListSurveyResourcesDetailResponse, <-chan error) {
	responseChan := make(chan *ListSurveyResourcesDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSurveyResourcesDetail(request)
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

// ListSurveyResourcesDetailWithCallback invokes the apds.ListSurveyResourcesDetail API asynchronously
func (client *Client) ListSurveyResourcesDetailWithCallback(request *ListSurveyResourcesDetailRequest, callback func(response *ListSurveyResourcesDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSurveyResourcesDetailResponse
		var err error
		defer close(result)
		response, err = client.ListSurveyResourcesDetail(request)
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

// ListSurveyResourcesDetailRequest is the request struct for api ListSurveyResourcesDetail
type ListSurveyResourcesDetailRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// ListSurveyResourcesDetailResponse is the response struct for api ListSurveyResourcesDetail
type ListSurveyResourcesDetailResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Error   string `json:"error" xml:"error"`
	Success bool   `json:"Success" xml:"Success"`
	Data    string `json:"Data" xml:"Data"`
}

// CreateListSurveyResourcesDetailRequest creates a request to invoke ListSurveyResourcesDetail API
func CreateListSurveyResourcesDetailRequest() (request *ListSurveyResourcesDetailRequest) {
	request = &ListSurveyResourcesDetailRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("apds", "2022-03-31", "ListSurveyResourcesDetail", "/okss-services/survey-detail/query", "", "")
	request.Method = requests.POST
	return
}

// CreateListSurveyResourcesDetailResponse creates a response to parse from ListSurveyResourcesDetail response
func CreateListSurveyResourcesDetailResponse() (response *ListSurveyResourcesDetailResponse) {
	response = &ListSurveyResourcesDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
