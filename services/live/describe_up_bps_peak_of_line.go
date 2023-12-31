package live

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

// DescribeUpBpsPeakOfLine invokes the live.DescribeUpBpsPeakOfLine API synchronously
func (client *Client) DescribeUpBpsPeakOfLine(request *DescribeUpBpsPeakOfLineRequest) (response *DescribeUpBpsPeakOfLineResponse, err error) {
	response = CreateDescribeUpBpsPeakOfLineResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUpBpsPeakOfLineWithChan invokes the live.DescribeUpBpsPeakOfLine API asynchronously
func (client *Client) DescribeUpBpsPeakOfLineWithChan(request *DescribeUpBpsPeakOfLineRequest) (<-chan *DescribeUpBpsPeakOfLineResponse, <-chan error) {
	responseChan := make(chan *DescribeUpBpsPeakOfLineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUpBpsPeakOfLine(request)
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

// DescribeUpBpsPeakOfLineWithCallback invokes the live.DescribeUpBpsPeakOfLine API asynchronously
func (client *Client) DescribeUpBpsPeakOfLineWithCallback(request *DescribeUpBpsPeakOfLineRequest, callback func(response *DescribeUpBpsPeakOfLineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUpBpsPeakOfLineResponse
		var err error
		defer close(result)
		response, err = client.DescribeUpBpsPeakOfLine(request)
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

// DescribeUpBpsPeakOfLineRequest is the request struct for api DescribeUpBpsPeakOfLine
type DescribeUpBpsPeakOfLineRequest struct {
	*requests.RpcRequest
	Line         string           `position:"Query" name:"Line"`
	StartTime    string           `position:"Query" name:"StartTime"`
	DomainName   string           `position:"Query" name:"DomainName"`
	EndTime      string           `position:"Query" name:"EndTime"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	DomainSwitch string           `position:"Query" name:"DomainSwitch"`
}

// DescribeUpBpsPeakOfLineResponse is the response struct for api DescribeUpBpsPeakOfLine
type DescribeUpBpsPeakOfLineResponse struct {
	*responses.BaseResponse
	RequestId                string                   `json:"RequestId" xml:"RequestId"`
	DescribeUpBpsPeakOfLines DescribeUpBpsPeakOfLines `json:"DescribeUpBpsPeakOfLines" xml:"DescribeUpBpsPeakOfLines"`
}

// CreateDescribeUpBpsPeakOfLineRequest creates a request to invoke DescribeUpBpsPeakOfLine API
func CreateDescribeUpBpsPeakOfLineRequest() (request *DescribeUpBpsPeakOfLineRequest) {
	request = &DescribeUpBpsPeakOfLineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeUpBpsPeakOfLine", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUpBpsPeakOfLineResponse creates a response to parse from DescribeUpBpsPeakOfLine response
func CreateDescribeUpBpsPeakOfLineResponse() (response *DescribeUpBpsPeakOfLineResponse) {
	response = &DescribeUpBpsPeakOfLineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
