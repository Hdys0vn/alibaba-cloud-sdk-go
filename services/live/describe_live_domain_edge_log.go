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

// DescribeLiveDomainEdgeLog invokes the live.DescribeLiveDomainEdgeLog API synchronously
func (client *Client) DescribeLiveDomainEdgeLog(request *DescribeLiveDomainEdgeLogRequest) (response *DescribeLiveDomainEdgeLogResponse, err error) {
	response = CreateDescribeLiveDomainEdgeLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainEdgeLogWithChan invokes the live.DescribeLiveDomainEdgeLog API asynchronously
func (client *Client) DescribeLiveDomainEdgeLogWithChan(request *DescribeLiveDomainEdgeLogRequest) (<-chan *DescribeLiveDomainEdgeLogResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainEdgeLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainEdgeLog(request)
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

// DescribeLiveDomainEdgeLogWithCallback invokes the live.DescribeLiveDomainEdgeLog API asynchronously
func (client *Client) DescribeLiveDomainEdgeLogWithCallback(request *DescribeLiveDomainEdgeLogRequest, callback func(response *DescribeLiveDomainEdgeLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainEdgeLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainEdgeLog(request)
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

// DescribeLiveDomainEdgeLogRequest is the request struct for api DescribeLiveDomainEdgeLog
type DescribeLiveDomainEdgeLogRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	LogType    string           `position:"Query" name:"LogType"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveDomainEdgeLogResponse is the response struct for api DescribeLiveDomainEdgeLog
type DescribeLiveDomainEdgeLogResponse struct {
	*responses.BaseResponse
	DomainName       string                                      `json:"DomainName" xml:"DomainName"`
	RequestId        string                                      `json:"RequestId" xml:"RequestId"`
	DomainLogDetails DomainLogDetailsInDescribeLiveDomainEdgeLog `json:"DomainLogDetails" xml:"DomainLogDetails"`
}

// CreateDescribeLiveDomainEdgeLogRequest creates a request to invoke DescribeLiveDomainEdgeLog API
func CreateDescribeLiveDomainEdgeLogRequest() (request *DescribeLiveDomainEdgeLogRequest) {
	request = &DescribeLiveDomainEdgeLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainEdgeLog", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainEdgeLogResponse creates a response to parse from DescribeLiveDomainEdgeLog response
func CreateDescribeLiveDomainEdgeLogResponse() (response *DescribeLiveDomainEdgeLogResponse) {
	response = &DescribeLiveDomainEdgeLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
