package cloudwf

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

// GetApTop invokes the cloudwf.GetApTop API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getaptop.html
func (client *Client) GetApTop(request *GetApTopRequest) (response *GetApTopResponse, err error) {
	response = CreateGetApTopResponse()
	err = client.DoAction(request, response)
	return
}

// GetApTopWithChan invokes the cloudwf.GetApTop API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getaptop.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApTopWithChan(request *GetApTopRequest) (<-chan *GetApTopResponse, <-chan error) {
	responseChan := make(chan *GetApTopResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetApTop(request)
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

// GetApTopWithCallback invokes the cloudwf.GetApTop API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getaptop.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApTopWithCallback(request *GetApTopRequest, callback func(response *GetApTopResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetApTopResponse
		var err error
		defer close(result)
		response, err = client.GetApTop(request)
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

// GetApTopRequest is the request struct for api GetApTop
type GetApTopRequest struct {
	*requests.RpcRequest
	ApgroupId requests.Integer `position:"Query" name:"ApgroupId"`
}

// GetApTopResponse is the response struct for api GetApTop
type GetApTopResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetApTopRequest creates a request to invoke GetApTop API
func CreateGetApTopRequest() (request *GetApTopRequest) {
	request = &GetApTopRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetApTop", "cloudwf", "openAPI")
	return
}

// CreateGetApTopResponse creates a response to parse from GetApTop response
func CreateGetApTopResponse() (response *GetApTopResponse) {
	response = &GetApTopResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
