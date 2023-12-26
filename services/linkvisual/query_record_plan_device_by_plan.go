package linkvisual

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

// QueryRecordPlanDeviceByPlan invokes the linkvisual.QueryRecordPlanDeviceByPlan API synchronously
func (client *Client) QueryRecordPlanDeviceByPlan(request *QueryRecordPlanDeviceByPlanRequest) (response *QueryRecordPlanDeviceByPlanResponse, err error) {
	response = CreateQueryRecordPlanDeviceByPlanResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRecordPlanDeviceByPlanWithChan invokes the linkvisual.QueryRecordPlanDeviceByPlan API asynchronously
func (client *Client) QueryRecordPlanDeviceByPlanWithChan(request *QueryRecordPlanDeviceByPlanRequest) (<-chan *QueryRecordPlanDeviceByPlanResponse, <-chan error) {
	responseChan := make(chan *QueryRecordPlanDeviceByPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRecordPlanDeviceByPlan(request)
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

// QueryRecordPlanDeviceByPlanWithCallback invokes the linkvisual.QueryRecordPlanDeviceByPlan API asynchronously
func (client *Client) QueryRecordPlanDeviceByPlanWithCallback(request *QueryRecordPlanDeviceByPlanRequest, callback func(response *QueryRecordPlanDeviceByPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRecordPlanDeviceByPlanResponse
		var err error
		defer close(result)
		response, err = client.QueryRecordPlanDeviceByPlan(request)
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

// QueryRecordPlanDeviceByPlanRequest is the request struct for api QueryRecordPlanDeviceByPlan
type QueryRecordPlanDeviceByPlanRequest struct {
	*requests.RpcRequest
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	ApiProduct  string           `position:"Body" name:"ApiProduct"`
	ApiRevision string           `position:"Body" name:"ApiRevision"`
	PlanId      string           `position:"Query" name:"PlanId"`
}

// QueryRecordPlanDeviceByPlanResponse is the response struct for api QueryRecordPlanDeviceByPlan
type QueryRecordPlanDeviceByPlanResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateQueryRecordPlanDeviceByPlanRequest creates a request to invoke QueryRecordPlanDeviceByPlan API
func CreateQueryRecordPlanDeviceByPlanRequest() (request *QueryRecordPlanDeviceByPlanRequest) {
	request = &QueryRecordPlanDeviceByPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "QueryRecordPlanDeviceByPlan", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryRecordPlanDeviceByPlanResponse creates a response to parse from QueryRecordPlanDeviceByPlan response
func CreateQueryRecordPlanDeviceByPlanResponse() (response *QueryRecordPlanDeviceByPlanResponse) {
	response = &QueryRecordPlanDeviceByPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
