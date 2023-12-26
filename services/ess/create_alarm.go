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

// CreateAlarm invokes the ess.CreateAlarm API synchronously
func (client *Client) CreateAlarm(request *CreateAlarmRequest) (response *CreateAlarmResponse, err error) {
	response = CreateCreateAlarmResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAlarmWithChan invokes the ess.CreateAlarm API asynchronously
func (client *Client) CreateAlarmWithChan(request *CreateAlarmRequest) (<-chan *CreateAlarmResponse, <-chan error) {
	responseChan := make(chan *CreateAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAlarm(request)
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

// CreateAlarmWithCallback invokes the ess.CreateAlarm API asynchronously
func (client *Client) CreateAlarmWithCallback(request *CreateAlarmRequest, callback func(response *CreateAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAlarmResponse
		var err error
		defer close(result)
		response, err = client.CreateAlarm(request)
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

// CreateAlarmRequest is the request struct for api CreateAlarm
type CreateAlarmRequest struct {
	*requests.RpcRequest
	MetricType               string                   `position:"Query" name:"MetricType"`
	ScalingGroupId           string                   `position:"Query" name:"ScalingGroupId"`
	Description              string                   `position:"Query" name:"Description"`
	ExpressionsLogicOperator string                   `position:"Query" name:"ExpressionsLogicOperator"`
	AlarmAction              *[]string                `position:"Query" name:"AlarmAction"  type:"Repeated"`
	Threshold                requests.Float           `position:"Query" name:"Threshold"`
	Effective                string                   `position:"Query" name:"Effective"`
	EvaluationCount          requests.Integer         `position:"Query" name:"EvaluationCount"`
	MetricName               string                   `position:"Query" name:"MetricName"`
	Dimension                *[]CreateAlarmDimension  `position:"Query" name:"Dimension"  type:"Repeated"`
	Period                   requests.Integer         `position:"Query" name:"Period"`
	Expression               *[]CreateAlarmExpression `position:"Query" name:"Expression"  type:"Repeated"`
	ResourceOwnerAccount     string                   `position:"Query" name:"ResourceOwnerAccount"`
	GroupId                  requests.Integer         `position:"Query" name:"GroupId"`
	OwnerId                  requests.Integer         `position:"Query" name:"OwnerId"`
	Name                     string                   `position:"Query" name:"Name"`
	ComparisonOperator       string                   `position:"Query" name:"ComparisonOperator"`
	Statistics               string                   `position:"Query" name:"Statistics"`
}

// CreateAlarmDimension is a repeated param struct in CreateAlarmRequest
type CreateAlarmDimension struct {
	DimensionValue string `name:"DimensionValue"`
	DimensionKey   string `name:"DimensionKey"`
}

// CreateAlarmExpression is a repeated param struct in CreateAlarmRequest
type CreateAlarmExpression struct {
	Period             string `name:"Period"`
	Threshold          string `name:"Threshold"`
	MetricName         string `name:"MetricName"`
	ComparisonOperator string `name:"ComparisonOperator"`
	Statistics         string `name:"Statistics"`
}

// CreateAlarmResponse is the response struct for api CreateAlarm
type CreateAlarmResponse struct {
	*responses.BaseResponse
	AlarmTaskId string `json:"AlarmTaskId" xml:"AlarmTaskId"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateAlarmRequest creates a request to invoke CreateAlarm API
func CreateCreateAlarmRequest() (request *CreateAlarmRequest) {
	request = &CreateAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "CreateAlarm", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAlarmResponse creates a response to parse from CreateAlarm response
func CreateCreateAlarmResponse() (response *CreateAlarmResponse) {
	response = &CreateAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
