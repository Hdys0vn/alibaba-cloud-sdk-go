package alidns

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

// DescribePdnsRequestStatistic invokes the alidns.DescribePdnsRequestStatistic API synchronously
func (client *Client) DescribePdnsRequestStatistic(request *DescribePdnsRequestStatisticRequest) (response *DescribePdnsRequestStatisticResponse, err error) {
	response = CreateDescribePdnsRequestStatisticResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePdnsRequestStatisticWithChan invokes the alidns.DescribePdnsRequestStatistic API asynchronously
func (client *Client) DescribePdnsRequestStatisticWithChan(request *DescribePdnsRequestStatisticRequest) (<-chan *DescribePdnsRequestStatisticResponse, <-chan error) {
	responseChan := make(chan *DescribePdnsRequestStatisticResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePdnsRequestStatistic(request)
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

// DescribePdnsRequestStatisticWithCallback invokes the alidns.DescribePdnsRequestStatistic API asynchronously
func (client *Client) DescribePdnsRequestStatisticWithCallback(request *DescribePdnsRequestStatisticRequest, callback func(response *DescribePdnsRequestStatisticResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePdnsRequestStatisticResponse
		var err error
		defer close(result)
		response, err = client.DescribePdnsRequestStatistic(request)
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

// DescribePdnsRequestStatisticRequest is the request struct for api DescribePdnsRequestStatistic
type DescribePdnsRequestStatisticRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	Type       string `position:"Query" name:"Type"`
	StartDate  string `position:"Query" name:"StartDate"`
	EndDate    string `position:"Query" name:"EndDate"`
	SubDomain  string `position:"Query" name:"SubDomain"`
	Lang       string `position:"Query" name:"Lang"`
}

// DescribePdnsRequestStatisticResponse is the response struct for api DescribePdnsRequestStatistic
type DescribePdnsRequestStatisticResponse struct {
	*responses.BaseResponse
	RequestId string          `json:"RequestId" xml:"RequestId"`
	Data      []StatisticItem `json:"Data" xml:"Data"`
}

// CreateDescribePdnsRequestStatisticRequest creates a request to invoke DescribePdnsRequestStatistic API
func CreateDescribePdnsRequestStatisticRequest() (request *DescribePdnsRequestStatisticRequest) {
	request = &DescribePdnsRequestStatisticRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribePdnsRequestStatistic", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePdnsRequestStatisticResponse creates a response to parse from DescribePdnsRequestStatistic response
func CreateDescribePdnsRequestStatisticResponse() (response *DescribePdnsRequestStatisticResponse) {
	response = &DescribePdnsRequestStatisticResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
