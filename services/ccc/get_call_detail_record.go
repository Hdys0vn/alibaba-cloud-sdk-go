package ccc

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

// GetCallDetailRecord invokes the ccc.GetCallDetailRecord API synchronously
func (client *Client) GetCallDetailRecord(request *GetCallDetailRecordRequest) (response *GetCallDetailRecordResponse, err error) {
	response = CreateGetCallDetailRecordResponse()
	err = client.DoAction(request, response)
	return
}

// GetCallDetailRecordWithChan invokes the ccc.GetCallDetailRecord API asynchronously
func (client *Client) GetCallDetailRecordWithChan(request *GetCallDetailRecordRequest) (<-chan *GetCallDetailRecordResponse, <-chan error) {
	responseChan := make(chan *GetCallDetailRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCallDetailRecord(request)
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

// GetCallDetailRecordWithCallback invokes the ccc.GetCallDetailRecord API asynchronously
func (client *Client) GetCallDetailRecordWithCallback(request *GetCallDetailRecordRequest, callback func(response *GetCallDetailRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCallDetailRecordResponse
		var err error
		defer close(result)
		response, err = client.GetCallDetailRecord(request)
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

// GetCallDetailRecordRequest is the request struct for api GetCallDetailRecord
type GetCallDetailRecordRequest struct {
	*requests.RpcRequest
	ContactId  string `position:"Query" name:"ContactId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetCallDetailRecordResponse is the response struct for api GetCallDetailRecord
type GetCallDetailRecordResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetCallDetailRecordRequest creates a request to invoke GetCallDetailRecord API
func CreateGetCallDetailRecordRequest() (request *GetCallDetailRecordRequest) {
	request = &GetCallDetailRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "GetCallDetailRecord", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetCallDetailRecordResponse creates a response to parse from GetCallDetailRecord response
func CreateGetCallDetailRecordResponse() (response *GetCallDetailRecordResponse) {
	response = &GetCallDetailRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
