package dms_enterprise

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

// SearchDataTrackResult invokes the dms_enterprise.SearchDataTrackResult API synchronously
func (client *Client) SearchDataTrackResult(request *SearchDataTrackResultRequest) (response *SearchDataTrackResultResponse, err error) {
	response = CreateSearchDataTrackResultResponse()
	err = client.DoAction(request, response)
	return
}

// SearchDataTrackResultWithChan invokes the dms_enterprise.SearchDataTrackResult API asynchronously
func (client *Client) SearchDataTrackResultWithChan(request *SearchDataTrackResultRequest) (<-chan *SearchDataTrackResultResponse, <-chan error) {
	responseChan := make(chan *SearchDataTrackResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchDataTrackResult(request)
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

// SearchDataTrackResultWithCallback invokes the dms_enterprise.SearchDataTrackResult API asynchronously
func (client *Client) SearchDataTrackResultWithCallback(request *SearchDataTrackResultRequest, callback func(response *SearchDataTrackResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchDataTrackResultResponse
		var err error
		defer close(result)
		response, err = client.SearchDataTrackResult(request)
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

// SearchDataTrackResultRequest is the request struct for api SearchDataTrackResult
type SearchDataTrackResultRequest struct {
	*requests.RpcRequest
	FilterStartTime string                            `position:"Query" name:"FilterStartTime"`
	FilterTypeList  *[]string                         `position:"Query" name:"FilterTypeList"  type:"Json"`
	Tid             requests.Integer                  `position:"Query" name:"Tid"`
	OrderId         requests.Integer                  `position:"Query" name:"OrderId"`
	FilterTableList *[]string                         `position:"Query" name:"FilterTableList"  type:"Json"`
	FilterEndTime   string                            `position:"Query" name:"FilterEndTime"`
	ColumnFilter    SearchDataTrackResultColumnFilter `position:"Query" name:"ColumnFilter"  type:"Struct"`
}

// SearchDataTrackResultColumnFilter is a repeated param struct in SearchDataTrackResultRequest
type SearchDataTrackResultColumnFilter struct {
	BetweenStart string    `name:"BetweenStart"`
	BetweenEnd   string    `name:"BetweenEnd"`
	ColumnName   string    `name:"ColumnName"`
	Value        string    `name:"Value"`
	Operator     string    `name:"Operator"`
	InList       *[]string `name:"InList" type:"Repeated"`
}

// SearchDataTrackResultResponse is the response struct for api SearchDataTrackResult
type SearchDataTrackResultResponse struct {
	*responses.BaseResponse
	RequestId    string      `json:"RequestId" xml:"RequestId"`
	Success      bool        `json:"Success" xml:"Success"`
	ErrorMessage string      `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string      `json:"ErrorCode" xml:"ErrorCode"`
	TrackResult  TrackResult `json:"TrackResult" xml:"TrackResult"`
}

// CreateSearchDataTrackResultRequest creates a request to invoke SearchDataTrackResult API
func CreateSearchDataTrackResultRequest() (request *SearchDataTrackResultRequest) {
	request = &SearchDataTrackResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "SearchDataTrackResult", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSearchDataTrackResultResponse creates a response to parse from SearchDataTrackResult response
func CreateSearchDataTrackResultResponse() (response *SearchDataTrackResultResponse) {
	response = &SearchDataTrackResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
