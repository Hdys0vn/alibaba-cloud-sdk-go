package hbase

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

// CheckComponentsVersion invokes the hbase.CheckComponentsVersion API synchronously
func (client *Client) CheckComponentsVersion(request *CheckComponentsVersionRequest) (response *CheckComponentsVersionResponse, err error) {
	response = CreateCheckComponentsVersionResponse()
	err = client.DoAction(request, response)
	return
}

// CheckComponentsVersionWithChan invokes the hbase.CheckComponentsVersion API asynchronously
func (client *Client) CheckComponentsVersionWithChan(request *CheckComponentsVersionRequest) (<-chan *CheckComponentsVersionResponse, <-chan error) {
	responseChan := make(chan *CheckComponentsVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckComponentsVersion(request)
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

// CheckComponentsVersionWithCallback invokes the hbase.CheckComponentsVersion API asynchronously
func (client *Client) CheckComponentsVersionWithCallback(request *CheckComponentsVersionRequest, callback func(response *CheckComponentsVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckComponentsVersionResponse
		var err error
		defer close(result)
		response, err = client.CheckComponentsVersion(request)
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

// CheckComponentsVersionRequest is the request struct for api CheckComponentsVersion
type CheckComponentsVersionRequest struct {
	*requests.RpcRequest
	Components string `position:"Query" name:"Components"`
	ClusterId  string `position:"Query" name:"ClusterId"`
}

// CheckComponentsVersionResponse is the response struct for api CheckComponentsVersion
type CheckComponentsVersionResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Components Components `json:"Components" xml:"Components"`
}

// CreateCheckComponentsVersionRequest creates a request to invoke CheckComponentsVersion API
func CreateCheckComponentsVersionRequest() (request *CheckComponentsVersionRequest) {
	request = &CheckComponentsVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "CheckComponentsVersion", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckComponentsVersionResponse creates a response to parse from CheckComponentsVersion response
func CreateCheckComponentsVersionResponse() (response *CheckComponentsVersionResponse) {
	response = &CheckComponentsVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
