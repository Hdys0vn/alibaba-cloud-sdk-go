package ledgerdb

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

// GetJournal invokes the ledgerdb.GetJournal API synchronously
// api document: https://help.aliyun.com/api/ledgerdb/getjournal.html
func (client *Client) GetJournal(request *GetJournalRequest) (response *GetJournalResponse, err error) {
	response = CreateGetJournalResponse()
	err = client.DoAction(request, response)
	return
}

// GetJournalWithChan invokes the ledgerdb.GetJournal API asynchronously
// api document: https://help.aliyun.com/api/ledgerdb/getjournal.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetJournalWithChan(request *GetJournalRequest) (<-chan *GetJournalResponse, <-chan error) {
	responseChan := make(chan *GetJournalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJournal(request)
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

// GetJournalWithCallback invokes the ledgerdb.GetJournal API asynchronously
// api document: https://help.aliyun.com/api/ledgerdb/getjournal.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetJournalWithCallback(request *GetJournalRequest, callback func(response *GetJournalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJournalResponse
		var err error
		defer close(result)
		response, err = client.GetJournal(request)
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

// GetJournalRequest is the request struct for api GetJournal
type GetJournalRequest struct {
	*requests.RpcRequest
	JournalId requests.Integer `position:"Query" name:"JournalId"`
	LedgerId  string           `position:"Query" name:"LedgerId"`
}

// GetJournalResponse is the response struct for api GetJournal
type GetJournalResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Journal   Journal `json:"Journal" xml:"Journal"`
}

// CreateGetJournalRequest creates a request to invoke GetJournal API
func CreateGetJournalRequest() (request *GetJournalRequest) {
	request = &GetJournalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ledgerdb", "2019-11-22", "GetJournal", "ledgerdb", "openAPI")
	return
}

// CreateGetJournalResponse creates a response to parse from GetJournal response
func CreateGetJournalResponse() (response *GetJournalResponse) {
	response = &GetJournalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
