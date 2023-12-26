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

// GetInstanceSummaryReportByInterval invokes the cloudcallcenter.GetInstanceSummaryReportByInterval API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinstancesummaryreportbyinterval.html
func (client *Client) GetInstanceSummaryReportByInterval(request *GetInstanceSummaryReportByIntervalRequest) (response *GetInstanceSummaryReportByIntervalResponse, err error) {
	response = CreateGetInstanceSummaryReportByIntervalResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceSummaryReportByIntervalWithChan invokes the cloudcallcenter.GetInstanceSummaryReportByInterval API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinstancesummaryreportbyinterval.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceSummaryReportByIntervalWithChan(request *GetInstanceSummaryReportByIntervalRequest) (<-chan *GetInstanceSummaryReportByIntervalResponse, <-chan error) {
	responseChan := make(chan *GetInstanceSummaryReportByIntervalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceSummaryReportByInterval(request)
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

// GetInstanceSummaryReportByIntervalWithCallback invokes the cloudcallcenter.GetInstanceSummaryReportByInterval API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinstancesummaryreportbyinterval.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceSummaryReportByIntervalWithCallback(request *GetInstanceSummaryReportByIntervalRequest, callback func(response *GetInstanceSummaryReportByIntervalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceSummaryReportByIntervalResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceSummaryReportByInterval(request)
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

// GetInstanceSummaryReportByIntervalRequest is the request struct for api GetInstanceSummaryReportByInterval
type GetInstanceSummaryReportByIntervalRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	IntervalType string `position:"Query" name:"IntervalType"`
	EndTime      string `position:"Query" name:"EndTime"`
	StartTime    string `position:"Query" name:"StartTime"`
}

// GetInstanceSummaryReportByIntervalResponse is the response struct for api GetInstanceSummaryReportByInterval
type GetInstanceSummaryReportByIntervalResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	Success                    bool                       `json:"Success" xml:"Success"`
	Code                       string                     `json:"Code" xml:"Code"`
	Message                    string                     `json:"Message" xml:"Message"`
	HttpStatusCode             int                        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	InstanceTimeIntervalReport InstanceTimeIntervalReport `json:"InstanceTimeIntervalReport" xml:"InstanceTimeIntervalReport"`
}

// CreateGetInstanceSummaryReportByIntervalRequest creates a request to invoke GetInstanceSummaryReportByInterval API
func CreateGetInstanceSummaryReportByIntervalRequest() (request *GetInstanceSummaryReportByIntervalRequest) {
	request = &GetInstanceSummaryReportByIntervalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetInstanceSummaryReportByInterval", "", "")
	request.Method = requests.POST
	return
}

// CreateGetInstanceSummaryReportByIntervalResponse creates a response to parse from GetInstanceSummaryReportByInterval response
func CreateGetInstanceSummaryReportByIntervalResponse() (response *GetInstanceSummaryReportByIntervalResponse) {
	response = &GetInstanceSummaryReportByIntervalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
