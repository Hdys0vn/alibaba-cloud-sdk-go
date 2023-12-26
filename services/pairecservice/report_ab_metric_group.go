package pairecservice

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

// ReportABMetricGroup invokes the pairecservice.ReportABMetricGroup API synchronously
func (client *Client) ReportABMetricGroup(request *ReportABMetricGroupRequest) (response *ReportABMetricGroupResponse, err error) {
	response = CreateReportABMetricGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ReportABMetricGroupWithChan invokes the pairecservice.ReportABMetricGroup API asynchronously
func (client *Client) ReportABMetricGroupWithChan(request *ReportABMetricGroupRequest) (<-chan *ReportABMetricGroupResponse, <-chan error) {
	responseChan := make(chan *ReportABMetricGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportABMetricGroup(request)
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

// ReportABMetricGroupWithCallback invokes the pairecservice.ReportABMetricGroup API asynchronously
func (client *Client) ReportABMetricGroupWithCallback(request *ReportABMetricGroupRequest, callback func(response *ReportABMetricGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportABMetricGroupResponse
		var err error
		defer close(result)
		response, err = client.ReportABMetricGroup(request)
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

// ReportABMetricGroupRequest is the request struct for api ReportABMetricGroup
type ReportABMetricGroupRequest struct {
	*requests.RoaRequest
	ABMetricGroupId string `position:"Path" name:"ABMetricGroupId"`
	Body            string `position:"Body" name:"body"`
}

// ReportABMetricGroupResponse is the response struct for api ReportABMetricGroup
type ReportABMetricGroupResponse struct {
	*responses.BaseResponse
	RequestId        string                 `json:"RequestId" xml:"RequestId"`
	ExperimentReport map[string]interface{} `json:"ExperimentReport" xml:"ExperimentReport"`
	GroupDimension   []string               `json:"GroupDimension" xml:"GroupDimension"`
}

// CreateReportABMetricGroupRequest creates a request to invoke ReportABMetricGroup API
func CreateReportABMetricGroupRequest() (request *ReportABMetricGroupRequest) {
	request = &ReportABMetricGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "ReportABMetricGroup", "/api/v1/abmetricgroups/[ABMetricGroupId]/action/report", "", "")
	request.Method = requests.POST
	return
}

// CreateReportABMetricGroupResponse creates a response to parse from ReportABMetricGroup response
func CreateReportABMetricGroupResponse() (response *ReportABMetricGroupResponse) {
	response = &ReportABMetricGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}