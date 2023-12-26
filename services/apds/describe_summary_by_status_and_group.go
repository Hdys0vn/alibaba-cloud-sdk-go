package apds

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

// DescribeSummaryByStatusAndGroup invokes the apds.DescribeSummaryByStatusAndGroup API synchronously
func (client *Client) DescribeSummaryByStatusAndGroup(request *DescribeSummaryByStatusAndGroupRequest) (response *DescribeSummaryByStatusAndGroupResponse, err error) {
	response = CreateDescribeSummaryByStatusAndGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSummaryByStatusAndGroupWithChan invokes the apds.DescribeSummaryByStatusAndGroup API asynchronously
func (client *Client) DescribeSummaryByStatusAndGroupWithChan(request *DescribeSummaryByStatusAndGroupRequest) (<-chan *DescribeSummaryByStatusAndGroupResponse, <-chan error) {
	responseChan := make(chan *DescribeSummaryByStatusAndGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSummaryByStatusAndGroup(request)
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

// DescribeSummaryByStatusAndGroupWithCallback invokes the apds.DescribeSummaryByStatusAndGroup API asynchronously
func (client *Client) DescribeSummaryByStatusAndGroupWithCallback(request *DescribeSummaryByStatusAndGroupRequest, callback func(response *DescribeSummaryByStatusAndGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSummaryByStatusAndGroupResponse
		var err error
		defer close(result)
		response, err = client.DescribeSummaryByStatusAndGroup(request)
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

// DescribeSummaryByStatusAndGroupRequest is the request struct for api DescribeSummaryByStatusAndGroup
type DescribeSummaryByStatusAndGroupRequest struct {
	*requests.RoaRequest
}

// DescribeSummaryByStatusAndGroupResponse is the response struct for api DescribeSummaryByStatusAndGroup
type DescribeSummaryByStatusAndGroupResponse struct {
	*responses.BaseResponse
	Code    string `json:"Code" xml:"Code"`
	Error   string `json:"Error" xml:"Error"`
	Success bool   `json:"Success" xml:"Success"`
	Data    string `json:"Data" xml:"Data"`
}

// CreateDescribeSummaryByStatusAndGroupRequest creates a request to invoke DescribeSummaryByStatusAndGroup API
func CreateDescribeSummaryByStatusAndGroupRequest() (request *DescribeSummaryByStatusAndGroupRequest) {
	request = &DescribeSummaryByStatusAndGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("apds", "2022-03-31", "DescribeSummaryByStatusAndGroup", "/okss-services/summary/summary-by-status-and-region", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeSummaryByStatusAndGroupResponse creates a response to parse from DescribeSummaryByStatusAndGroup response
func CreateDescribeSummaryByStatusAndGroupResponse() (response *DescribeSummaryByStatusAndGroupResponse) {
	response = &DescribeSummaryByStatusAndGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
