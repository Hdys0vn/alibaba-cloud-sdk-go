package iot

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

// QuerySummarySceneRuleLog invokes the iot.QuerySummarySceneRuleLog API synchronously
func (client *Client) QuerySummarySceneRuleLog(request *QuerySummarySceneRuleLogRequest) (response *QuerySummarySceneRuleLogResponse, err error) {
	response = CreateQuerySummarySceneRuleLogResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySummarySceneRuleLogWithChan invokes the iot.QuerySummarySceneRuleLog API asynchronously
func (client *Client) QuerySummarySceneRuleLogWithChan(request *QuerySummarySceneRuleLogRequest) (<-chan *QuerySummarySceneRuleLogResponse, <-chan error) {
	responseChan := make(chan *QuerySummarySceneRuleLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySummarySceneRuleLog(request)
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

// QuerySummarySceneRuleLogWithCallback invokes the iot.QuerySummarySceneRuleLog API asynchronously
func (client *Client) QuerySummarySceneRuleLogWithCallback(request *QuerySummarySceneRuleLogRequest, callback func(response *QuerySummarySceneRuleLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySummarySceneRuleLogResponse
		var err error
		defer close(result)
		response, err = client.QuerySummarySceneRuleLog(request)
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

// QuerySummarySceneRuleLogRequest is the request struct for api QuerySummarySceneRuleLog
type QuerySummarySceneRuleLogRequest struct {
	*requests.RpcRequest
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	RuleId        string           `position:"Query" name:"RuleId"`
	Status        string           `position:"Query" name:"Status"`
}

// QuerySummarySceneRuleLogResponse is the response struct for api QuerySummarySceneRuleLog
type QuerySummarySceneRuleLogResponse struct {
	*responses.BaseResponse
	RequestId    string                         `json:"RequestId" xml:"RequestId"`
	Success      bool                           `json:"Success" xml:"Success"`
	ErrorMessage string                         `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string                         `json:"Code" xml:"Code"`
	Data         DataInQuerySummarySceneRuleLog `json:"Data" xml:"Data"`
}

// CreateQuerySummarySceneRuleLogRequest creates a request to invoke QuerySummarySceneRuleLog API
func CreateQuerySummarySceneRuleLogRequest() (request *QuerySummarySceneRuleLogRequest) {
	request = &QuerySummarySceneRuleLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QuerySummarySceneRuleLog", "", "")
	request.Method = requests.POST
	return
}

// CreateQuerySummarySceneRuleLogResponse creates a response to parse from QuerySummarySceneRuleLog response
func CreateQuerySummarySceneRuleLogResponse() (response *QuerySummarySceneRuleLogResponse) {
	response = &QuerySummarySceneRuleLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
