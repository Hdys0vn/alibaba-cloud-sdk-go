package cloudphoto

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

// SetQuota invokes the cloudphoto.SetQuota API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/setquota.html
func (client *Client) SetQuota(request *SetQuotaRequest) (response *SetQuotaResponse, err error) {
	response = CreateSetQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// SetQuotaWithChan invokes the cloudphoto.SetQuota API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/setquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetQuotaWithChan(request *SetQuotaRequest) (<-chan *SetQuotaResponse, <-chan error) {
	responseChan := make(chan *SetQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetQuota(request)
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

// SetQuotaWithCallback invokes the cloudphoto.SetQuota API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/setquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetQuotaWithCallback(request *SetQuotaRequest, callback func(response *SetQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetQuotaResponse
		var err error
		defer close(result)
		response, err = client.SetQuota(request)
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

// SetQuotaRequest is the request struct for api SetQuota
type SetQuotaRequest struct {
	*requests.RpcRequest
	TotalQuota requests.Integer `position:"Query" name:"TotalQuota"`
	LibraryId  string           `position:"Query" name:"LibraryId"`
	StoreName  string           `position:"Query" name:"StoreName"`
}

// SetQuotaResponse is the response struct for api SetQuota
type SetQuotaResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
}

// CreateSetQuotaRequest creates a request to invoke SetQuota API
func CreateSetQuotaRequest() (request *SetQuotaRequest) {
	request = &SetQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "SetQuota", "cloudphoto", "openAPI")
	return
}

// CreateSetQuotaResponse creates a response to parse from SetQuota response
func CreateSetQuotaResponse() (response *SetQuotaResponse) {
	response = &SetQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
