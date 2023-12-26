package cloudcallcenter

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

// DownloadUnreachableContacts invokes the cloudcallcenter.DownloadUnreachableContacts API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadunreachablecontacts.html
func (client *Client) DownloadUnreachableContacts(request *DownloadUnreachableContactsRequest) (response *DownloadUnreachableContactsResponse, err error) {
	response = CreateDownloadUnreachableContactsResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadUnreachableContactsWithChan invokes the cloudcallcenter.DownloadUnreachableContacts API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadunreachablecontacts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadUnreachableContactsWithChan(request *DownloadUnreachableContactsRequest) (<-chan *DownloadUnreachableContactsResponse, <-chan error) {
	responseChan := make(chan *DownloadUnreachableContactsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadUnreachableContacts(request)
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

// DownloadUnreachableContactsWithCallback invokes the cloudcallcenter.DownloadUnreachableContacts API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadunreachablecontacts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadUnreachableContactsWithCallback(request *DownloadUnreachableContactsRequest, callback func(response *DownloadUnreachableContactsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadUnreachableContactsResponse
		var err error
		defer close(result)
		response, err = client.DownloadUnreachableContacts(request)
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

// DownloadUnreachableContactsRequest is the request struct for api DownloadUnreachableContacts
type DownloadUnreachableContactsRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	JobGroupId string `position:"Query" name:"JobGroupId"`
}

// DownloadUnreachableContactsResponse is the response struct for api DownloadUnreachableContacts
type DownloadUnreachableContactsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DownloadParams DownloadParams `json:"DownloadParams" xml:"DownloadParams"`
}

// CreateDownloadUnreachableContactsRequest creates a request to invoke DownloadUnreachableContacts API
func CreateDownloadUnreachableContactsRequest() (request *DownloadUnreachableContactsRequest) {
	request = &DownloadUnreachableContactsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DownloadUnreachableContacts", "", "")
	request.Method = requests.POST
	return
}

// CreateDownloadUnreachableContactsResponse creates a response to parse from DownloadUnreachableContacts response
func CreateDownloadUnreachableContactsResponse() (response *DownloadUnreachableContactsResponse) {
	response = &DownloadUnreachableContactsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
