package aegis

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

// DescribeScreenAlarmEventList invokes the aegis.DescribeScreenAlarmEventList API synchronously
// api document: https://help.aliyun.com/api/aegis/describescreenalarmeventlist.html
func (client *Client) DescribeScreenAlarmEventList(request *DescribeScreenAlarmEventListRequest) (response *DescribeScreenAlarmEventListResponse, err error) {
	response = CreateDescribeScreenAlarmEventListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScreenAlarmEventListWithChan invokes the aegis.DescribeScreenAlarmEventList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describescreenalarmeventlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScreenAlarmEventListWithChan(request *DescribeScreenAlarmEventListRequest) (<-chan *DescribeScreenAlarmEventListResponse, <-chan error) {
	responseChan := make(chan *DescribeScreenAlarmEventListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScreenAlarmEventList(request)
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

// DescribeScreenAlarmEventListWithCallback invokes the aegis.DescribeScreenAlarmEventList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describescreenalarmeventlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScreenAlarmEventListWithCallback(request *DescribeScreenAlarmEventListRequest, callback func(response *DescribeScreenAlarmEventListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScreenAlarmEventListResponse
		var err error
		defer close(result)
		response, err = client.DescribeScreenAlarmEventList(request)
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

// DescribeScreenAlarmEventListRequest is the request struct for api DescribeScreenAlarmEventList
type DescribeScreenAlarmEventListRequest struct {
	*requests.RpcRequest
	AlarmEventName string           `position:"Query" name:"AlarmEventName"`
	SourceIp       string           `position:"Query" name:"SourceIp"`
	PageSize       string           `position:"Query" name:"PageSize"`
	AlarmEventType string           `position:"Query" name:"AlarmEventType"`
	Dealed         string           `position:"Query" name:"Dealed"`
	From           string           `position:"Query" name:"From"`
	Remark         string           `position:"Query" name:"Remark"`
	CurrentPage    requests.Integer `position:"Query" name:"CurrentPage"`
	Lang           string           `position:"Query" name:"Lang"`
	Levels         string           `position:"Query" name:"Levels"`
}

// DescribeScreenAlarmEventListResponse is the response struct for api DescribeScreenAlarmEventList
type DescribeScreenAlarmEventListResponse struct {
	*responses.BaseResponse
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	PageInfo   PageInfo         `json:"PageInfo" xml:"PageInfo"`
	SuspEvents []SuspEventsItem `json:"SuspEvents" xml:"SuspEvents"`
}

// CreateDescribeScreenAlarmEventListRequest creates a request to invoke DescribeScreenAlarmEventList API
func CreateDescribeScreenAlarmEventListRequest() (request *DescribeScreenAlarmEventListRequest) {
	request = &DescribeScreenAlarmEventListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeScreenAlarmEventList", "vipaegis", "openAPI")
	return
}

// CreateDescribeScreenAlarmEventListResponse creates a response to parse from DescribeScreenAlarmEventList response
func CreateDescribeScreenAlarmEventListResponse() (response *DescribeScreenAlarmEventListResponse) {
	response = &DescribeScreenAlarmEventListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
