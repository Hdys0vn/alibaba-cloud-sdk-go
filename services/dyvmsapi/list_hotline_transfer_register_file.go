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

// ListHotlineTransferRegisterFile invokes the dyvmsapi.ListHotlineTransferRegisterFile API synchronously
func (client *Client) ListHotlineTransferRegisterFile(request *ListHotlineTransferRegisterFileRequest) (response *ListHotlineTransferRegisterFileResponse, err error) {
	response = CreateListHotlineTransferRegisterFileResponse()
	err = client.DoAction(request, response)
	return
}

// ListHotlineTransferRegisterFileWithChan invokes the dyvmsapi.ListHotlineTransferRegisterFile API asynchronously
func (client *Client) ListHotlineTransferRegisterFileWithChan(request *ListHotlineTransferRegisterFileRequest) (<-chan *ListHotlineTransferRegisterFileResponse, <-chan error) {
	responseChan := make(chan *ListHotlineTransferRegisterFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListHotlineTransferRegisterFile(request)
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

// ListHotlineTransferRegisterFileWithCallback invokes the dyvmsapi.ListHotlineTransferRegisterFile API asynchronously
func (client *Client) ListHotlineTransferRegisterFileWithCallback(request *ListHotlineTransferRegisterFileRequest, callback func(response *ListHotlineTransferRegisterFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListHotlineTransferRegisterFileResponse
		var err error
		defer close(result)
		response, err = client.ListHotlineTransferRegisterFile(request)
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

// ListHotlineTransferRegisterFileRequest is the request struct for api ListHotlineTransferRegisterFile
type ListHotlineTransferRegisterFileRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HotlineNumber        string           `position:"Query" name:"HotlineNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	QualificationId      string           `position:"Query" name:"QualificationId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
}

// ListHotlineTransferRegisterFileResponse is the response struct for api ListHotlineTransferRegisterFile
type ListHotlineTransferRegisterFileResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListHotlineTransferRegisterFileRequest creates a request to invoke ListHotlineTransferRegisterFile API
func CreateListHotlineTransferRegisterFileRequest() (request *ListHotlineTransferRegisterFileRequest) {
	request = &ListHotlineTransferRegisterFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "ListHotlineTransferRegisterFile", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListHotlineTransferRegisterFileResponse creates a response to parse from ListHotlineTransferRegisterFile response
func CreateListHotlineTransferRegisterFileResponse() (response *ListHotlineTransferRegisterFileResponse) {
	response = &ListHotlineTransferRegisterFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}