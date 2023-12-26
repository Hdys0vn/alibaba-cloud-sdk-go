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

// DescribeDcdnUserSecDropByMinute invokes the dcdn.DescribeDcdnUserSecDropByMinute API synchronously
func (client *Client) DescribeDcdnUserSecDropByMinute(request *DescribeDcdnUserSecDropByMinuteRequest) (response *DescribeDcdnUserSecDropByMinuteResponse, err error) {
	response = CreateDescribeDcdnUserSecDropByMinuteResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnUserSecDropByMinuteWithChan invokes the dcdn.DescribeDcdnUserSecDropByMinute API asynchronously
func (client *Client) DescribeDcdnUserSecDropByMinuteWithChan(request *DescribeDcdnUserSecDropByMinuteRequest) (<-chan *DescribeDcdnUserSecDropByMinuteResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnUserSecDropByMinuteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnUserSecDropByMinute(request)
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

// DescribeDcdnUserSecDropByMinuteWithCallback invokes the dcdn.DescribeDcdnUserSecDropByMinute API asynchronously
func (client *Client) DescribeDcdnUserSecDropByMinuteWithCallback(request *DescribeDcdnUserSecDropByMinuteRequest, callback func(response *DescribeDcdnUserSecDropByMinuteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnUserSecDropByMinuteResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnUserSecDropByMinute(request)
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

// DescribeDcdnUserSecDropByMinuteRequest is the request struct for api DescribeDcdnUserSecDropByMinute
type DescribeDcdnUserSecDropByMinuteRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	EndTime    string           `position:"Query" name:"EndTime"`
	RuleName   string           `position:"Query" name:"RuleName"`
	StartTime  string           `position:"Query" name:"StartTime"`
	SecFunc    string           `position:"Query" name:"SecFunc"`
	Lang       string           `position:"Query" name:"Lang"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Object     string           `position:"Query" name:"Object"`
}

// DescribeDcdnUserSecDropByMinuteResponse is the response struct for api DescribeDcdnUserSecDropByMinute
type DescribeDcdnUserSecDropByMinuteResponse struct {
	*responses.BaseResponse
	RequestId   string     `json:"RequestId" xml:"RequestId"`
	Description string     `json:"Description" xml:"Description"`
	Len         int        `json:"Len" xml:"Len"`
	PageNumber  int        `json:"PageNumber" xml:"PageNumber"`
	PageSize    int        `json:"PageSize" xml:"PageSize"`
	TotalCount  int        `json:"TotalCount" xml:"TotalCount"`
	Rows        []RowsItem `json:"Rows" xml:"Rows"`
}

// CreateDescribeDcdnUserSecDropByMinuteRequest creates a request to invoke DescribeDcdnUserSecDropByMinute API
func CreateDescribeDcdnUserSecDropByMinuteRequest() (request *DescribeDcdnUserSecDropByMinuteRequest) {
	request = &DescribeDcdnUserSecDropByMinuteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnUserSecDropByMinute", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnUserSecDropByMinuteResponse creates a response to parse from DescribeDcdnUserSecDropByMinute response
func CreateDescribeDcdnUserSecDropByMinuteResponse() (response *DescribeDcdnUserSecDropByMinuteResponse) {
	response = &DescribeDcdnUserSecDropByMinuteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
