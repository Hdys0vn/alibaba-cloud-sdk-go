package dyvmsapi

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

// QueryCallInTransferRecord invokes the dyvmsapi.QueryCallInTransferRecord API synchronously
func (client *Client) QueryCallInTransferRecord(request *QueryCallInTransferRecordRequest) (response *QueryCallInTransferRecordResponse, err error) {
	response = CreateQueryCallInTransferRecordResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCallInTransferRecordWithChan invokes the dyvmsapi.QueryCallInTransferRecord API asynchronously
func (client *Client) QueryCallInTransferRecordWithChan(request *QueryCallInTransferRecordRequest) (<-chan *QueryCallInTransferRecordResponse, <-chan error) {
	responseChan := make(chan *QueryCallInTransferRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCallInTransferRecord(request)
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

// QueryCallInTransferRecordWithCallback invokes the dyvmsapi.QueryCallInTransferRecord API asynchronously
func (client *Client) QueryCallInTransferRecordWithCallback(request *QueryCallInTransferRecordRequest, callback func(response *QueryCallInTransferRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCallInTransferRecordResponse
		var err error
		defer close(result)
		response, err = client.QueryCallInTransferRecord(request)
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

// QueryCallInTransferRecordRequest is the request struct for api QueryCallInTransferRecord
type QueryCallInTransferRecordRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PhoneNumber          string           `position:"Query" name:"PhoneNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	QueryDate            string           `position:"Query" name:"QueryDate"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
	CallInCaller         string           `position:"Query" name:"CallInCaller"`
}

// QueryCallInTransferRecordResponse is the response struct for api QueryCallInTransferRecord
type QueryCallInTransferRecordResponse struct {
	*responses.BaseResponse
	RequestId string                          `json:"RequestId" xml:"RequestId"`
	Code      string                          `json:"Code" xml:"Code"`
	Message   string                          `json:"Message" xml:"Message"`
	Data      DataInQueryCallInTransferRecord `json:"Data" xml:"Data"`
}

// CreateQueryCallInTransferRecordRequest creates a request to invoke QueryCallInTransferRecord API
func CreateQueryCallInTransferRecordRequest() (request *QueryCallInTransferRecordRequest) {
	request = &QueryCallInTransferRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "QueryCallInTransferRecord", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryCallInTransferRecordResponse creates a response to parse from QueryCallInTransferRecord response
func CreateQueryCallInTransferRecordResponse() (response *QueryCallInTransferRecordResponse) {
	response = &QueryCallInTransferRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
