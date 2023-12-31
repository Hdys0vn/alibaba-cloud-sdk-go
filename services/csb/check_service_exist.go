package csb

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

// CheckServiceExist invokes the csb.CheckServiceExist API synchronously
// api document: https://help.aliyun.com/api/csb/checkserviceexist.html
func (client *Client) CheckServiceExist(request *CheckServiceExistRequest) (response *CheckServiceExistResponse, err error) {
	response = CreateCheckServiceExistResponse()
	err = client.DoAction(request, response)
	return
}

// CheckServiceExistWithChan invokes the csb.CheckServiceExist API asynchronously
// api document: https://help.aliyun.com/api/csb/checkserviceexist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckServiceExistWithChan(request *CheckServiceExistRequest) (<-chan *CheckServiceExistResponse, <-chan error) {
	responseChan := make(chan *CheckServiceExistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckServiceExist(request)
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

// CheckServiceExistWithCallback invokes the csb.CheckServiceExist API asynchronously
// api document: https://help.aliyun.com/api/csb/checkserviceexist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckServiceExistWithCallback(request *CheckServiceExistRequest, callback func(response *CheckServiceExistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckServiceExistResponse
		var err error
		defer close(result)
		response, err = client.CheckServiceExist(request)
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

// CheckServiceExistRequest is the request struct for api CheckServiceExist
type CheckServiceExistRequest struct {
	*requests.RpcRequest
	CsbId       requests.Integer `position:"Query" name:"CsbId"`
	ServiceName string           `position:"Query" name:"ServiceName"`
}

// CheckServiceExistResponse is the response struct for api CheckServiceExist
type CheckServiceExistResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCheckServiceExistRequest creates a request to invoke CheckServiceExist API
func CreateCheckServiceExistRequest() (request *CheckServiceExistRequest) {
	request = &CheckServiceExistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "CheckServiceExist", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckServiceExistResponse creates a response to parse from CheckServiceExist response
func CreateCheckServiceExistResponse() (response *CheckServiceExistResponse) {
	response = &CheckServiceExistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
