package das

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

// EnableSqlConcurrencyControl invokes the das.EnableSqlConcurrencyControl API synchronously
func (client *Client) EnableSqlConcurrencyControl(request *EnableSqlConcurrencyControlRequest) (response *EnableSqlConcurrencyControlResponse, err error) {
	response = CreateEnableSqlConcurrencyControlResponse()
	err = client.DoAction(request, response)
	return
}

// EnableSqlConcurrencyControlWithChan invokes the das.EnableSqlConcurrencyControl API asynchronously
func (client *Client) EnableSqlConcurrencyControlWithChan(request *EnableSqlConcurrencyControlRequest) (<-chan *EnableSqlConcurrencyControlResponse, <-chan error) {
	responseChan := make(chan *EnableSqlConcurrencyControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableSqlConcurrencyControl(request)
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

// EnableSqlConcurrencyControlWithCallback invokes the das.EnableSqlConcurrencyControl API asynchronously
func (client *Client) EnableSqlConcurrencyControlWithCallback(request *EnableSqlConcurrencyControlRequest, callback func(response *EnableSqlConcurrencyControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableSqlConcurrencyControlResponse
		var err error
		defer close(result)
		response, err = client.EnableSqlConcurrencyControl(request)
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

// EnableSqlConcurrencyControlRequest is the request struct for api EnableSqlConcurrencyControl
type EnableSqlConcurrencyControlRequest struct {
	*requests.RpcRequest
	SqlType                string           `position:"Query" name:"SqlType"`
	SqlKeywords            string           `position:"Query" name:"SqlKeywords"`
	ConsoleContext         string           `position:"Query" name:"ConsoleContext"`
	InstanceId             string           `position:"Query" name:"InstanceId"`
	ConcurrencyControlTime requests.Integer `position:"Query" name:"ConcurrencyControlTime"`
	MaxConcurrency         requests.Integer `position:"Query" name:"MaxConcurrency"`
}

// EnableSqlConcurrencyControlResponse is the response struct for api EnableSqlConcurrencyControl
type EnableSqlConcurrencyControlResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateEnableSqlConcurrencyControlRequest creates a request to invoke EnableSqlConcurrencyControl API
func CreateEnableSqlConcurrencyControlRequest() (request *EnableSqlConcurrencyControlRequest) {
	request = &EnableSqlConcurrencyControlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "EnableSqlConcurrencyControl", "", "")
	request.Method = requests.POST
	return
}

// CreateEnableSqlConcurrencyControlResponse creates a response to parse from EnableSqlConcurrencyControl response
func CreateEnableSqlConcurrencyControlResponse() (response *EnableSqlConcurrencyControlResponse) {
	response = &EnableSqlConcurrencyControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
