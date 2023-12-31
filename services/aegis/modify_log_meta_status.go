package aegis

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

// ModifyLogMetaStatus invokes the aegis.ModifyLogMetaStatus API synchronously
// api document: https://help.aliyun.com/api/aegis/modifylogmetastatus.html
func (client *Client) ModifyLogMetaStatus(request *ModifyLogMetaStatusRequest) (response *ModifyLogMetaStatusResponse, err error) {
	response = CreateModifyLogMetaStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLogMetaStatusWithChan invokes the aegis.ModifyLogMetaStatus API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifylogmetastatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLogMetaStatusWithChan(request *ModifyLogMetaStatusRequest) (<-chan *ModifyLogMetaStatusResponse, <-chan error) {
	responseChan := make(chan *ModifyLogMetaStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLogMetaStatus(request)
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

// ModifyLogMetaStatusWithCallback invokes the aegis.ModifyLogMetaStatus API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifylogmetastatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyLogMetaStatusWithCallback(request *ModifyLogMetaStatusRequest, callback func(response *ModifyLogMetaStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLogMetaStatusResponse
		var err error
		defer close(result)
		response, err = client.ModifyLogMetaStatus(request)
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

// ModifyLogMetaStatusRequest is the request struct for api ModifyLogMetaStatus
type ModifyLogMetaStatusRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Project  string `position:"Query" name:"Project"`
	From     string `position:"Query" name:"From"`
	Lang     string `position:"Query" name:"Lang"`
	LogStore string `position:"Query" name:"LogStore"`
	Status   string `position:"Query" name:"Status"`
}

// ModifyLogMetaStatusResponse is the response struct for api ModifyLogMetaStatus
type ModifyLogMetaStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyLogMetaStatusRequest creates a request to invoke ModifyLogMetaStatus API
func CreateModifyLogMetaStatusRequest() (request *ModifyLogMetaStatusRequest) {
	request = &ModifyLogMetaStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifyLogMetaStatus", "vipaegis", "openAPI")
	return
}

// CreateModifyLogMetaStatusResponse creates a response to parse from ModifyLogMetaStatus response
func CreateModifyLogMetaStatusResponse() (response *ModifyLogMetaStatusResponse) {
	response = &ModifyLogMetaStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
