package clickhouse

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

// DescribeLorneTasksMCount invokes the clickhouse.DescribeLorneTasksMCount API synchronously
func (client *Client) DescribeLorneTasksMCount(request *DescribeLorneTasksMCountRequest) (response *DescribeLorneTasksMCountResponse, err error) {
	response = CreateDescribeLorneTasksMCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLorneTasksMCountWithChan invokes the clickhouse.DescribeLorneTasksMCount API asynchronously
func (client *Client) DescribeLorneTasksMCountWithChan(request *DescribeLorneTasksMCountRequest) (<-chan *DescribeLorneTasksMCountResponse, <-chan error) {
	responseChan := make(chan *DescribeLorneTasksMCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLorneTasksMCount(request)
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

// DescribeLorneTasksMCountWithCallback invokes the clickhouse.DescribeLorneTasksMCount API asynchronously
func (client *Client) DescribeLorneTasksMCountWithCallback(request *DescribeLorneTasksMCountRequest, callback func(response *DescribeLorneTasksMCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLorneTasksMCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeLorneTasksMCount(request)
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

// DescribeLorneTasksMCountRequest is the request struct for api DescribeLorneTasksMCount
type DescribeLorneTasksMCountRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	MetricName           string           `position:"Query" name:"MetricName"`
	TaskId               string           `position:"Query" name:"TaskId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLorneTasksMCountResponse is the response struct for api DescribeLorneTasksMCount
type DescribeLorneTasksMCountResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Data      float64 `json:"Data" xml:"Data"`
}

// CreateDescribeLorneTasksMCountRequest creates a request to invoke DescribeLorneTasksMCount API
func CreateDescribeLorneTasksMCountRequest() (request *DescribeLorneTasksMCountRequest) {
	request = &DescribeLorneTasksMCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeLorneTasksMCount", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLorneTasksMCountResponse creates a response to parse from DescribeLorneTasksMCount response
func CreateDescribeLorneTasksMCountResponse() (response *DescribeLorneTasksMCountResponse) {
	response = &DescribeLorneTasksMCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}