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

// ListOverviewAccumulateReport invokes the cloudcallcenter.ListOverviewAccumulateReport API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listoverviewaccumulatereport.html
func (client *Client) ListOverviewAccumulateReport(request *ListOverviewAccumulateReportRequest) (response *ListOverviewAccumulateReportResponse, err error) {
	response = CreateListOverviewAccumulateReportResponse()
	err = client.DoAction(request, response)
	return
}

// ListOverviewAccumulateReportWithChan invokes the cloudcallcenter.ListOverviewAccumulateReport API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listoverviewaccumulatereport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListOverviewAccumulateReportWithChan(request *ListOverviewAccumulateReportRequest) (<-chan *ListOverviewAccumulateReportResponse, <-chan error) {
	responseChan := make(chan *ListOverviewAccumulateReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOverviewAccumulateReport(request)
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

// ListOverviewAccumulateReportWithCallback invokes the cloudcallcenter.ListOverviewAccumulateReport API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listoverviewaccumulatereport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListOverviewAccumulateReportWithCallback(request *ListOverviewAccumulateReportRequest, callback func(response *ListOverviewAccumulateReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOverviewAccumulateReportResponse
		var err error
		defer close(result)
		response, err = client.ListOverviewAccumulateReport(request)
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

// ListOverviewAccumulateReportRequest is the request struct for api ListOverviewAccumulateReport
type ListOverviewAccumulateReportRequest struct {
	*requests.RpcRequest
	Date          string    `position:"Query" name:"Date"`
	InstanceId    string    `position:"Query" name:"InstanceId"`
	Dimension     string    `position:"Query" name:"Dimension"`
	IndicatorName *[]string `position:"Query" name:"IndicatorName"  type:"Repeated"`
}

// ListOverviewAccumulateReportResponse is the response struct for api ListOverviewAccumulateReport
type ListOverviewAccumulateReportResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           string `json:"Data" xml:"Data"`
}

// CreateListOverviewAccumulateReportRequest creates a request to invoke ListOverviewAccumulateReport API
func CreateListOverviewAccumulateReportRequest() (request *ListOverviewAccumulateReportRequest) {
	request = &ListOverviewAccumulateReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListOverviewAccumulateReport", "", "")
	request.Method = requests.POST
	return
}

// CreateListOverviewAccumulateReportResponse creates a response to parse from ListOverviewAccumulateReport response
func CreateListOverviewAccumulateReportResponse() (response *ListOverviewAccumulateReportResponse) {
	response = &ListOverviewAccumulateReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
