package dcdn

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

// DescribeDcdnDomainCcActivityLog invokes the dcdn.DescribeDcdnDomainCcActivityLog API synchronously
func (client *Client) DescribeDcdnDomainCcActivityLog(request *DescribeDcdnDomainCcActivityLogRequest) (response *DescribeDcdnDomainCcActivityLogResponse, err error) {
	response = CreateDescribeDcdnDomainCcActivityLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainCcActivityLogWithChan invokes the dcdn.DescribeDcdnDomainCcActivityLog API asynchronously
func (client *Client) DescribeDcdnDomainCcActivityLogWithChan(request *DescribeDcdnDomainCcActivityLogRequest) (<-chan *DescribeDcdnDomainCcActivityLogResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainCcActivityLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainCcActivityLog(request)
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

// DescribeDcdnDomainCcActivityLogWithCallback invokes the dcdn.DescribeDcdnDomainCcActivityLog API asynchronously
func (client *Client) DescribeDcdnDomainCcActivityLogWithCallback(request *DescribeDcdnDomainCcActivityLogRequest, callback func(response *DescribeDcdnDomainCcActivityLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainCcActivityLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainCcActivityLog(request)
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

// DescribeDcdnDomainCcActivityLogRequest is the request struct for api DescribeDcdnDomainCcActivityLog
type DescribeDcdnDomainCcActivityLogRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	EndTime       string           `position:"Query" name:"EndTime"`
	RuleName      string           `position:"Query" name:"RuleName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	TriggerObject string           `position:"Query" name:"TriggerObject"`
	Value         string           `position:"Query" name:"Value"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeDcdnDomainCcActivityLogResponse is the response struct for api DescribeDcdnDomainCcActivityLog
type DescribeDcdnDomainCcActivityLogResponse struct {
	*responses.BaseResponse
	PageIndex   int64     `json:"PageIndex" xml:"PageIndex"`
	RequestId   string    `json:"RequestId" xml:"RequestId"`
	PageSize    int64     `json:"PageSize" xml:"PageSize"`
	Total       int64     `json:"Total" xml:"Total"`
	ActivityLog []LogInfo `json:"ActivityLog" xml:"ActivityLog"`
}

// CreateDescribeDcdnDomainCcActivityLogRequest creates a request to invoke DescribeDcdnDomainCcActivityLog API
func CreateDescribeDcdnDomainCcActivityLogRequest() (request *DescribeDcdnDomainCcActivityLogRequest) {
	request = &DescribeDcdnDomainCcActivityLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainCcActivityLog", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnDomainCcActivityLogResponse creates a response to parse from DescribeDcdnDomainCcActivityLog response
func CreateDescribeDcdnDomainCcActivityLogResponse() (response *DescribeDcdnDomainCcActivityLogResponse) {
	response = &DescribeDcdnDomainCcActivityLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
