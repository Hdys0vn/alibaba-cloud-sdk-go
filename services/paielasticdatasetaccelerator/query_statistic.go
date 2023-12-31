package paielasticdatasetaccelerator

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

// QueryStatistic invokes the paielasticdatasetaccelerator.QueryStatistic API synchronously
func (client *Client) QueryStatistic(request *QueryStatisticRequest) (response *QueryStatisticResponse, err error) {
	response = CreateQueryStatisticResponse()
	err = client.DoAction(request, response)
	return
}

// QueryStatisticWithChan invokes the paielasticdatasetaccelerator.QueryStatistic API asynchronously
func (client *Client) QueryStatisticWithChan(request *QueryStatisticRequest) (<-chan *QueryStatisticResponse, <-chan error) {
	responseChan := make(chan *QueryStatisticResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryStatistic(request)
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

// QueryStatisticWithCallback invokes the paielasticdatasetaccelerator.QueryStatistic API asynchronously
func (client *Client) QueryStatisticWithCallback(request *QueryStatisticRequest, callback func(response *QueryStatisticResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryStatisticResponse
		var err error
		defer close(result)
		response, err = client.QueryStatistic(request)
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

// QueryStatisticRequest is the request struct for api QueryStatistic
type QueryStatisticRequest struct {
	*requests.RoaRequest
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
	Fields    string `position:"Query" name:"Fields"`
}

// QueryStatisticResponse is the response struct for api QueryStatistic
type QueryStatisticResponse struct {
	*responses.BaseResponse
	RequestId                string                 `json:"RequestId" xml:"RequestId"`
	InstanceCapacityEachType map[string]interface{} `json:"InstanceCapacityEachType" xml:"InstanceCapacityEachType"`
	InstanceNumEachType      map[string]interface{} `json:"InstanceNumEachType" xml:"InstanceNumEachType"`
	SlotNumEachType          map[string]interface{} `json:"SlotNumEachType" xml:"SlotNumEachType"`
}

// CreateQueryStatisticRequest creates a request to invoke QueryStatistic API
func CreateQueryStatisticRequest() (request *QueryStatisticRequest) {
	request = &QueryStatisticRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PAIElasticDatasetAccelerator", "2022-08-01", "QueryStatistic", "/api/v1/statistics/action/query", "datasetacc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryStatisticResponse creates a response to parse from QueryStatistic response
func CreateQueryStatisticResponse() (response *QueryStatisticResponse) {
	response = &QueryStatisticResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
