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

// DeleteABMetric invokes the pairecservice.DeleteABMetric API synchronously
func (client *Client) DeleteABMetric(request *DeleteABMetricRequest) (response *DeleteABMetricResponse, err error) {
	response = CreateDeleteABMetricResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteABMetricWithChan invokes the pairecservice.DeleteABMetric API asynchronously
func (client *Client) DeleteABMetricWithChan(request *DeleteABMetricRequest) (<-chan *DeleteABMetricResponse, <-chan error) {
	responseChan := make(chan *DeleteABMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteABMetric(request)
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

// DeleteABMetricWithCallback invokes the pairecservice.DeleteABMetric API asynchronously
func (client *Client) DeleteABMetricWithCallback(request *DeleteABMetricRequest, callback func(response *DeleteABMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteABMetricResponse
		var err error
		defer close(result)
		response, err = client.DeleteABMetric(request)
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

// DeleteABMetricRequest is the request struct for api DeleteABMetric
type DeleteABMetricRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	ABMetricId string `position:"Path" name:"ABMetricId"`
}

// DeleteABMetricResponse is the response struct for api DeleteABMetric
type DeleteABMetricResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteABMetricRequest creates a request to invoke DeleteABMetric API
func CreateDeleteABMetricRequest() (request *DeleteABMetricRequest) {
	request = &DeleteABMetricRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "DeleteABMetric", "/api/v1/abmetrics/[ABMetricId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteABMetricResponse creates a response to parse from DeleteABMetric response
func CreateDeleteABMetricResponse() (response *DeleteABMetricResponse) {
	response = &DeleteABMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}