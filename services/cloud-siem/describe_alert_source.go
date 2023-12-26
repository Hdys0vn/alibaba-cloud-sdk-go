package cloud_siem

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

// DescribeAlertSource invokes the cloud_siem.DescribeAlertSource API synchronously
func (client *Client) DescribeAlertSource(request *DescribeAlertSourceRequest) (response *DescribeAlertSourceResponse, err error) {
	response = CreateDescribeAlertSourceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAlertSourceWithChan invokes the cloud_siem.DescribeAlertSource API asynchronously
func (client *Client) DescribeAlertSourceWithChan(request *DescribeAlertSourceRequest) (<-chan *DescribeAlertSourceResponse, <-chan error) {
	responseChan := make(chan *DescribeAlertSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAlertSource(request)
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

// DescribeAlertSourceWithCallback invokes the cloud_siem.DescribeAlertSource API asynchronously
func (client *Client) DescribeAlertSourceWithCallback(request *DescribeAlertSourceRequest, callback func(response *DescribeAlertSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAlertSourceResponse
		var err error
		defer close(result)
		response, err = client.DescribeAlertSource(request)
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

// DescribeAlertSourceRequest is the request struct for api DescribeAlertSource
type DescribeAlertSourceRequest struct {
	*requests.RpcRequest
	StartTime requests.Integer `position:"Body" name:"StartTime"`
	Level     *[]string        `position:"Body" name:"Level"  type:"Repeated"`
	EndTime   requests.Integer `position:"Body" name:"EndTime"`
}

// DescribeAlertSourceResponse is the response struct for api DescribeAlertSource
type DescribeAlertSourceResponse struct {
	*responses.BaseResponse
	Success   bool       `json:"Success" xml:"Success"`
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeAlertSourceRequest creates a request to invoke DescribeAlertSource API
func CreateDescribeAlertSourceRequest() (request *DescribeAlertSourceRequest) {
	request = &DescribeAlertSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeAlertSource", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAlertSourceResponse creates a response to parse from DescribeAlertSource response
func CreateDescribeAlertSourceResponse() (response *DescribeAlertSourceResponse) {
	response = &DescribeAlertSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}