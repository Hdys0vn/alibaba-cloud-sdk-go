package domain

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

// SaveSingleTaskForDeletingDSRecord invokes the domain.SaveSingleTaskForDeletingDSRecord API synchronously
func (client *Client) SaveSingleTaskForDeletingDSRecord(request *SaveSingleTaskForDeletingDSRecordRequest) (response *SaveSingleTaskForDeletingDSRecordResponse, err error) {
	response = CreateSaveSingleTaskForDeletingDSRecordResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForDeletingDSRecordWithChan invokes the domain.SaveSingleTaskForDeletingDSRecord API asynchronously
func (client *Client) SaveSingleTaskForDeletingDSRecordWithChan(request *SaveSingleTaskForDeletingDSRecordRequest) (<-chan *SaveSingleTaskForDeletingDSRecordResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForDeletingDSRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForDeletingDSRecord(request)
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

// SaveSingleTaskForDeletingDSRecordWithCallback invokes the domain.SaveSingleTaskForDeletingDSRecord API asynchronously
func (client *Client) SaveSingleTaskForDeletingDSRecordWithCallback(request *SaveSingleTaskForDeletingDSRecordRequest, callback func(response *SaveSingleTaskForDeletingDSRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForDeletingDSRecordResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForDeletingDSRecord(request)
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

// SaveSingleTaskForDeletingDSRecordRequest is the request struct for api SaveSingleTaskForDeletingDSRecord
type SaveSingleTaskForDeletingDSRecordRequest struct {
	*requests.RpcRequest
	KeyTag       requests.Integer `position:"Query" name:"KeyTag"`
	DomainName   string           `position:"Query" name:"DomainName"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
}

// SaveSingleTaskForDeletingDSRecordResponse is the response struct for api SaveSingleTaskForDeletingDSRecord
type SaveSingleTaskForDeletingDSRecordResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForDeletingDSRecordRequest creates a request to invoke SaveSingleTaskForDeletingDSRecord API
func CreateSaveSingleTaskForDeletingDSRecordRequest() (request *SaveSingleTaskForDeletingDSRecordRequest) {
	request = &SaveSingleTaskForDeletingDSRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForDeletingDSRecord", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveSingleTaskForDeletingDSRecordResponse creates a response to parse from SaveSingleTaskForDeletingDSRecord response
func CreateSaveSingleTaskForDeletingDSRecordResponse() (response *SaveSingleTaskForDeletingDSRecordResponse) {
	response = &SaveSingleTaskForDeletingDSRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
