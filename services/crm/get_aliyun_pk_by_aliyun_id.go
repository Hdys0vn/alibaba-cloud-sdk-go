package crm

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

// GetAliyunPkByAliyunId invokes the crm.GetAliyunPkByAliyunId API synchronously
// api document: https://help.aliyun.com/api/crm/getaliyunpkbyaliyunid.html
func (client *Client) GetAliyunPkByAliyunId(request *GetAliyunPkByAliyunIdRequest) (response *GetAliyunPkByAliyunIdResponse, err error) {
	response = CreateGetAliyunPkByAliyunIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetAliyunPkByAliyunIdWithChan invokes the crm.GetAliyunPkByAliyunId API asynchronously
// api document: https://help.aliyun.com/api/crm/getaliyunpkbyaliyunid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAliyunPkByAliyunIdWithChan(request *GetAliyunPkByAliyunIdRequest) (<-chan *GetAliyunPkByAliyunIdResponse, <-chan error) {
	responseChan := make(chan *GetAliyunPkByAliyunIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAliyunPkByAliyunId(request)
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

// GetAliyunPkByAliyunIdWithCallback invokes the crm.GetAliyunPkByAliyunId API asynchronously
// api document: https://help.aliyun.com/api/crm/getaliyunpkbyaliyunid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAliyunPkByAliyunIdWithCallback(request *GetAliyunPkByAliyunIdRequest, callback func(response *GetAliyunPkByAliyunIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAliyunPkByAliyunIdResponse
		var err error
		defer close(result)
		response, err = client.GetAliyunPkByAliyunId(request)
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

// GetAliyunPkByAliyunIdRequest is the request struct for api GetAliyunPkByAliyunId
type GetAliyunPkByAliyunIdRequest struct {
	*requests.RpcRequest
	AliyunId string `position:"Query" name:"AliyunId"`
}

// GetAliyunPkByAliyunIdResponse is the response struct for api GetAliyunPkByAliyunId
type GetAliyunPkByAliyunIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	AliyunPk  string `json:"AliyunPk" xml:"AliyunPk"`
}

// CreateGetAliyunPkByAliyunIdRequest creates a request to invoke GetAliyunPkByAliyunId API
func CreateGetAliyunPkByAliyunIdRequest() (request *GetAliyunPkByAliyunIdRequest) {
	request = &GetAliyunPkByAliyunIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Crm", "2015-04-08", "GetAliyunPkByAliyunId", "crm", "openAPI")
	return
}

// CreateGetAliyunPkByAliyunIdResponse creates a response to parse from GetAliyunPkByAliyunId response
func CreateGetAliyunPkByAliyunIdResponse() (response *GetAliyunPkByAliyunIdResponse) {
	response = &GetAliyunPkByAliyunIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
