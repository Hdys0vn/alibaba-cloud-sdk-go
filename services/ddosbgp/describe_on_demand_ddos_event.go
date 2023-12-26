package ddosbgp

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

// DescribeOnDemandDdosEvent invokes the ddosbgp.DescribeOnDemandDdosEvent API synchronously
func (client *Client) DescribeOnDemandDdosEvent(request *DescribeOnDemandDdosEventRequest) (response *DescribeOnDemandDdosEventResponse, err error) {
	response = CreateDescribeOnDemandDdosEventResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOnDemandDdosEventWithChan invokes the ddosbgp.DescribeOnDemandDdosEvent API asynchronously
func (client *Client) DescribeOnDemandDdosEventWithChan(request *DescribeOnDemandDdosEventRequest) (<-chan *DescribeOnDemandDdosEventResponse, <-chan error) {
	responseChan := make(chan *DescribeOnDemandDdosEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOnDemandDdosEvent(request)
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

// DescribeOnDemandDdosEventWithCallback invokes the ddosbgp.DescribeOnDemandDdosEvent API asynchronously
func (client *Client) DescribeOnDemandDdosEventWithCallback(request *DescribeOnDemandDdosEventRequest, callback func(response *DescribeOnDemandDdosEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOnDemandDdosEventResponse
		var err error
		defer close(result)
		response, err = client.DescribeOnDemandDdosEvent(request)
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

// DescribeOnDemandDdosEventRequest is the request struct for api DescribeOnDemandDdosEvent
type DescribeOnDemandDdosEventRequest struct {
	*requests.RpcRequest
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Ip              string           `position:"Query" name:"Ip"`
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	PageNo          requests.Integer `position:"Query" name:"PageNo"`
}

// DescribeOnDemandDdosEventResponse is the response struct for api DescribeOnDemandDdosEvent
type DescribeOnDemandDdosEventResponse struct {
	*responses.BaseResponse
	Total     int64   `json:"Total" xml:"Total"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Events    []Event `json:"Events" xml:"Events"`
}

// CreateDescribeOnDemandDdosEventRequest creates a request to invoke DescribeOnDemandDdosEvent API
func CreateDescribeOnDemandDdosEventRequest() (request *DescribeOnDemandDdosEventRequest) {
	request = &DescribeOnDemandDdosEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "DescribeOnDemandDdosEvent", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeOnDemandDdosEventResponse creates a response to parse from DescribeOnDemandDdosEvent response
func CreateDescribeOnDemandDdosEventResponse() (response *DescribeOnDemandDdosEventResponse) {
	response = &DescribeOnDemandDdosEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
