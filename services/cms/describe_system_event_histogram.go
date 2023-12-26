package cms

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

// DescribeSystemEventHistogram invokes the cms.DescribeSystemEventHistogram API synchronously
func (client *Client) DescribeSystemEventHistogram(request *DescribeSystemEventHistogramRequest) (response *DescribeSystemEventHistogramResponse, err error) {
	response = CreateDescribeSystemEventHistogramResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSystemEventHistogramWithChan invokes the cms.DescribeSystemEventHistogram API asynchronously
func (client *Client) DescribeSystemEventHistogramWithChan(request *DescribeSystemEventHistogramRequest) (<-chan *DescribeSystemEventHistogramResponse, <-chan error) {
	responseChan := make(chan *DescribeSystemEventHistogramResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSystemEventHistogram(request)
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

// DescribeSystemEventHistogramWithCallback invokes the cms.DescribeSystemEventHistogram API asynchronously
func (client *Client) DescribeSystemEventHistogramWithCallback(request *DescribeSystemEventHistogramRequest, callback func(response *DescribeSystemEventHistogramResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSystemEventHistogramResponse
		var err error
		defer close(result)
		response, err = client.DescribeSystemEventHistogram(request)
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

// DescribeSystemEventHistogramRequest is the request struct for api DescribeSystemEventHistogram
type DescribeSystemEventHistogramRequest struct {
	*requests.RpcRequest
	StartTime      string `position:"Query" name:"StartTime"`
	SearchKeywords string `position:"Query" name:"SearchKeywords"`
	Product        string `position:"Query" name:"Product"`
	Level          string `position:"Query" name:"Level"`
	GroupId        string `position:"Query" name:"GroupId"`
	EndTime        string `position:"Query" name:"EndTime"`
	Name           string `position:"Query" name:"Name"`
	EventType      string `position:"Query" name:"EventType"`
	Status         string `position:"Query" name:"Status"`
}

// DescribeSystemEventHistogramResponse is the response struct for api DescribeSystemEventHistogram
type DescribeSystemEventHistogramResponse struct {
	*responses.BaseResponse
	Code                  string                `json:"Code" xml:"Code"`
	Message               string                `json:"Message" xml:"Message"`
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	Success               string                `json:"Success" xml:"Success"`
	SystemEventHistograms SystemEventHistograms `json:"SystemEventHistograms" xml:"SystemEventHistograms"`
}

// CreateDescribeSystemEventHistogramRequest creates a request to invoke DescribeSystemEventHistogram API
func CreateDescribeSystemEventHistogramRequest() (request *DescribeSystemEventHistogramRequest) {
	request = &DescribeSystemEventHistogramRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeSystemEventHistogram", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSystemEventHistogramResponse creates a response to parse from DescribeSystemEventHistogram response
func CreateDescribeSystemEventHistogramResponse() (response *DescribeSystemEventHistogramResponse) {
	response = &DescribeSystemEventHistogramResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
