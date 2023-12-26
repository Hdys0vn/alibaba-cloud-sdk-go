package cusanalytic_sc_online

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

// ListVisitors invokes the cusanalytic_sc_online.ListVisitors API synchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/listvisitors.html
func (client *Client) ListVisitors(request *ListVisitorsRequest) (response *ListVisitorsResponse, err error) {
	response = CreateListVisitorsResponse()
	err = client.DoAction(request, response)
	return
}

// ListVisitorsWithChan invokes the cusanalytic_sc_online.ListVisitors API asynchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/listvisitors.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVisitorsWithChan(request *ListVisitorsRequest) (<-chan *ListVisitorsResponse, <-chan error) {
	responseChan := make(chan *ListVisitorsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVisitors(request)
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

// ListVisitorsWithCallback invokes the cusanalytic_sc_online.ListVisitors API asynchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/listvisitors.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVisitorsWithCallback(request *ListVisitorsRequest, callback func(response *ListVisitorsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVisitorsResponse
		var err error
		defer close(result)
		response, err = client.ListVisitors(request)
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

// ListVisitorsRequest is the request struct for api ListVisitors
type ListVisitorsRequest struct {
	*requests.RpcRequest
	Gender      string           `position:"Body" name:"Gender"`
	UkId        requests.Integer `position:"Body" name:"UkId"`
	LocationIds string           `position:"Body" name:"LocationIds"`
	StartTime   string           `position:"Body" name:"StartTime"`
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	EnterCount  requests.Integer `position:"Body" name:"EnterCount"`
	PageIndex   requests.Integer `position:"Body" name:"PageIndex"`
	AgeStart    requests.Integer `position:"Body" name:"AgeStart"`
	AgeEnd      requests.Integer `position:"Body" name:"AgeEnd"`
	PkId        string           `position:"Body" name:"PkId"`
	EndTime     string           `position:"Body" name:"EndTime"`
	StoreIds    string           `position:"Body" name:"StoreIds"`
}

// ListVisitorsResponse is the response struct for api ListVisitors
type ListVisitorsResponse struct {
	*responses.BaseResponse
	Total        int64        `json:"Total" xml:"Total"`
	PageIndex    int          `json:"PageIndex" xml:"PageIndex"`
	PageSize     int          `json:"PageSize" xml:"PageSize"`
	VisitorItems VisitorItems `json:"VisitorItems" xml:"VisitorItems"`
}

// CreateListVisitorsRequest creates a request to invoke ListVisitors API
func CreateListVisitorsRequest() (request *ListVisitorsRequest) {
	request = &ListVisitorsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cusanalytic_sc_online", "2019-05-24", "ListVisitors", "", "")
	return
}

// CreateListVisitorsResponse creates a response to parse from ListVisitors response
func CreateListVisitorsResponse() (response *ListVisitorsResponse) {
	response = &ListVisitorsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
