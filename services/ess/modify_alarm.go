package ess

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

// ModifyAlarm invokes the ess.ModifyAlarm API synchronously
func (client *Client) ModifyAlarm(request *ModifyAlarmRequest) (response *ModifyAlarmResponse, err error) {
	response = CreateModifyAlarmResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAlarmWithChan invokes the ess.ModifyAlarm API asynchronously
func (client *Client) ModifyAlarmWithChan(request *ModifyAlarmRequest) (<-chan *ModifyAlarmResponse, <-chan error) {
	responseChan := make(chan *ModifyAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAlarm(request)
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

// ModifyAlarmWithCallback invokes the ess.ModifyAlarm API asynchronously
func (client *Client) ModifyAlarmWithCallback(request *ModifyAlarmRequest, callback func(response *ModifyAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAlarmResponse
		var err error
		defer close(result)
		response, err = client.ModifyAlarm(request)
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

// ModifyAlarmRequest is the request struct for api ModifyAlarm
type ModifyAlarmRequest struct {
	*requests.RpcRequest
	MetricType               string                   `position:"Query" name:"MetricType"`
	Description              string                   `position:"Query" name:"Description"`
	ExpressionsLogicOperator string                   `position:"Query" name:"ExpressionsLogicOperator"`
	AlarmAction              *[]string                `position:"Query" name:"AlarmAction"  type:"Repeated"`
	Threshold                requests.Float           `position:"Query" name:"Threshold"`
	Effective                string                   `position:"Query" name:"Effective"`
	EvaluationCount          requests.Integer         `position:"Query" name:"EvaluationCount"`
	MetricName               string                   `position:"Query" name:"MetricName"`
	Dimension                *[]ModifyAlarmDimension  `position:"Query" name:"Dimension"  type:"Repeated"`
	Period                   requests.Integer         `position:"Query" name:"Period"`
	Expression               *[]ModifyAlarmExpression `position:"Query" name:"Expression"  type:"Repeated"`
	ResourceOwnerAccount     string                   `position:"Query" name:"ResourceOwnerAccount"`
	GroupId                  requests.Integer         `position:"Query" name:"GroupId"`
	OwnerId                  requests.Integer         `position:"Query" name:"OwnerId"`
	AlarmTaskId              string                   `position:"Query" name:"AlarmTaskId"`
	Name                     string                   `position:"Query" name:"Name"`
	ComparisonOperator       string                   `position:"Query" name:"ComparisonOperator"`
	Statistics               string                   `position:"Query" name:"Statistics"`
}

// ModifyAlarmDimension is a repeated param struct in ModifyAlarmRequest
type ModifyAlarmDimension struct {
	DimensionValue string `name:"DimensionValue"`
	DimensionKey   string `name:"DimensionKey"`
}

// ModifyAlarmExpression is a repeated param struct in ModifyAlarmRequest
type ModifyAlarmExpression struct {
	Period             string `name:"Period"`
	Threshold          string `name:"Threshold"`
	MetricName         string `name:"MetricName"`
	ComparisonOperator string `name:"ComparisonOperator"`
	Statistics         string `name:"Statistics"`
}

// ModifyAlarmResponse is the response struct for api ModifyAlarm
type ModifyAlarmResponse struct {
	*responses.BaseResponse
	AlarmTaskId string `json:"AlarmTaskId" xml:"AlarmTaskId"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAlarmRequest creates a request to invoke ModifyAlarm API
func CreateModifyAlarmRequest() (request *ModifyAlarmRequest) {
	request = &ModifyAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ModifyAlarm", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyAlarmResponse creates a response to parse from ModifyAlarm response
func CreateModifyAlarmResponse() (response *ModifyAlarmResponse) {
	response = &ModifyAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
