package cloud_siem

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

// DescribeAlertsWithEvent invokes the cloud_siem.DescribeAlertsWithEvent API synchronously
func (client *Client) DescribeAlertsWithEvent(request *DescribeAlertsWithEventRequest) (response *DescribeAlertsWithEventResponse, err error) {
	response = CreateDescribeAlertsWithEventResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAlertsWithEventWithChan invokes the cloud_siem.DescribeAlertsWithEvent API asynchronously
func (client *Client) DescribeAlertsWithEventWithChan(request *DescribeAlertsWithEventRequest) (<-chan *DescribeAlertsWithEventResponse, <-chan error) {
	responseChan := make(chan *DescribeAlertsWithEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAlertsWithEvent(request)
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

// DescribeAlertsWithEventWithCallback invokes the cloud_siem.DescribeAlertsWithEvent API asynchronously
func (client *Client) DescribeAlertsWithEventWithCallback(request *DescribeAlertsWithEventRequest, callback func(response *DescribeAlertsWithEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAlertsWithEventResponse
		var err error
		defer close(result)
		response, err = client.DescribeAlertsWithEvent(request)
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

// DescribeAlertsWithEventRequest is the request struct for api DescribeAlertsWithEvent
type DescribeAlertsWithEventRequest struct {
	*requests.RpcRequest
	Source       string           `position:"Body" name:"Source"`
	IsDefend     string           `position:"Body" name:"IsDefend"`
	SubUserId    requests.Integer `position:"Body" name:"SubUserId"`
	PageSize     requests.Integer `position:"Body" name:"PageSize"`
	Level        *[]string        `position:"Body" name:"Level"  type:"Repeated"`
	AlertTitle   string           `position:"Body" name:"AlertTitle"`
	CurrentPage  requests.Integer `position:"Body" name:"CurrentPage"`
	IncidentUuid string           `position:"Body" name:"IncidentUuid"`
}

// DescribeAlertsWithEventResponse is the response struct for api DescribeAlertsWithEvent
type DescribeAlertsWithEventResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeAlertsWithEventRequest creates a request to invoke DescribeAlertsWithEvent API
func CreateDescribeAlertsWithEventRequest() (request *DescribeAlertsWithEventRequest) {
	request = &DescribeAlertsWithEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeAlertsWithEvent", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAlertsWithEventResponse creates a response to parse from DescribeAlertsWithEvent response
func CreateDescribeAlertsWithEventResponse() (response *DescribeAlertsWithEventResponse) {
	response = &DescribeAlertsWithEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
