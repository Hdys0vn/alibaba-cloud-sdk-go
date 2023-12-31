package cloudcallcenter

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

// ListAgentSummaryReportsByInterval invokes the cloudcallcenter.ListAgentSummaryReportsByInterval API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listagentsummaryreportsbyinterval.html
func (client *Client) ListAgentSummaryReportsByInterval(request *ListAgentSummaryReportsByIntervalRequest) (response *ListAgentSummaryReportsByIntervalResponse, err error) {
	response = CreateListAgentSummaryReportsByIntervalResponse()
	err = client.DoAction(request, response)
	return
}

// ListAgentSummaryReportsByIntervalWithChan invokes the cloudcallcenter.ListAgentSummaryReportsByInterval API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listagentsummaryreportsbyinterval.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAgentSummaryReportsByIntervalWithChan(request *ListAgentSummaryReportsByIntervalRequest) (<-chan *ListAgentSummaryReportsByIntervalResponse, <-chan error) {
	responseChan := make(chan *ListAgentSummaryReportsByIntervalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAgentSummaryReportsByInterval(request)
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

// ListAgentSummaryReportsByIntervalWithCallback invokes the cloudcallcenter.ListAgentSummaryReportsByInterval API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listagentsummaryreportsbyinterval.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAgentSummaryReportsByIntervalWithCallback(request *ListAgentSummaryReportsByIntervalRequest, callback func(response *ListAgentSummaryReportsByIntervalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAgentSummaryReportsByIntervalResponse
		var err error
		defer close(result)
		response, err = client.ListAgentSummaryReportsByInterval(request)
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

// ListAgentSummaryReportsByIntervalRequest is the request struct for api ListAgentSummaryReportsByInterval
type ListAgentSummaryReportsByIntervalRequest struct {
	*requests.RpcRequest
	AgentIds     string           `position:"Query" name:"AgentIds"`
	IntervalType string           `position:"Query" name:"IntervalType"`
	EndTime      string           `position:"Query" name:"EndTime"`
	StartTime    string           `position:"Query" name:"StartTime"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SkillGroupId string           `position:"Query" name:"SkillGroupId"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
}

// ListAgentSummaryReportsByIntervalResponse is the response struct for api ListAgentSummaryReportsByInterval
type ListAgentSummaryReportsByIntervalResponse struct {
	*responses.BaseResponse
	RequestId      string                                  `json:"RequestId" xml:"RequestId"`
	Success        bool                                    `json:"Success" xml:"Success"`
	Code           string                                  `json:"Code" xml:"Code"`
	Message        string                                  `json:"Message" xml:"Message"`
	HttpStatusCode int                                     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInListAgentSummaryReportsByInterval `json:"Data" xml:"Data"`
}

// CreateListAgentSummaryReportsByIntervalRequest creates a request to invoke ListAgentSummaryReportsByInterval API
func CreateListAgentSummaryReportsByIntervalRequest() (request *ListAgentSummaryReportsByIntervalRequest) {
	request = &ListAgentSummaryReportsByIntervalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListAgentSummaryReportsByInterval", "", "")
	request.Method = requests.POST
	return
}

// CreateListAgentSummaryReportsByIntervalResponse creates a response to parse from ListAgentSummaryReportsByInterval response
func CreateListAgentSummaryReportsByIntervalResponse() (response *ListAgentSummaryReportsByIntervalResponse) {
	response = &ListAgentSummaryReportsByIntervalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
