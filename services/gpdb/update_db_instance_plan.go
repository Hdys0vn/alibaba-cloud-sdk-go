package gpdb

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

// UpdateDBInstancePlan invokes the gpdb.UpdateDBInstancePlan API synchronously
func (client *Client) UpdateDBInstancePlan(request *UpdateDBInstancePlanRequest) (response *UpdateDBInstancePlanResponse, err error) {
	response = CreateUpdateDBInstancePlanResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDBInstancePlanWithChan invokes the gpdb.UpdateDBInstancePlan API asynchronously
func (client *Client) UpdateDBInstancePlanWithChan(request *UpdateDBInstancePlanRequest) (<-chan *UpdateDBInstancePlanResponse, <-chan error) {
	responseChan := make(chan *UpdateDBInstancePlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDBInstancePlan(request)
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

// UpdateDBInstancePlanWithCallback invokes the gpdb.UpdateDBInstancePlan API asynchronously
func (client *Client) UpdateDBInstancePlanWithCallback(request *UpdateDBInstancePlanRequest, callback func(response *UpdateDBInstancePlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDBInstancePlanResponse
		var err error
		defer close(result)
		response, err = client.UpdateDBInstancePlan(request)
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

// UpdateDBInstancePlanRequest is the request struct for api UpdateDBInstancePlan
type UpdateDBInstancePlanRequest struct {
	*requests.RpcRequest
	PlanStartDate string           `position:"Query" name:"PlanStartDate"`
	PlanConfig    string           `position:"Query" name:"PlanConfig"`
	PlanName      string           `position:"Query" name:"PlanName"`
	DBInstanceId  string           `position:"Query" name:"DBInstanceId"`
	PlanDesc      string           `position:"Query" name:"PlanDesc"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	PlanEndDate   string           `position:"Query" name:"PlanEndDate"`
	PlanId        string           `position:"Query" name:"PlanId"`
}

// UpdateDBInstancePlanResponse is the response struct for api UpdateDBInstancePlan
type UpdateDBInstancePlanResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Status       string `json:"Status" xml:"Status"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	PlanId       string `json:"PlanId" xml:"PlanId"`
}

// CreateUpdateDBInstancePlanRequest creates a request to invoke UpdateDBInstancePlan API
func CreateUpdateDBInstancePlanRequest() (request *UpdateDBInstancePlanRequest) {
	request = &UpdateDBInstancePlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "UpdateDBInstancePlan", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateDBInstancePlanResponse creates a response to parse from UpdateDBInstancePlan response
func CreateUpdateDBInstancePlanResponse() (response *UpdateDBInstancePlanResponse) {
	response = &UpdateDBInstancePlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
