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

// GetAllApModel invokes the cloudwf.GetAllApModel API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getallapmodel.html
func (client *Client) GetAllApModel(request *GetAllApModelRequest) (response *GetAllApModelResponse, err error) {
	response = CreateGetAllApModelResponse()
	err = client.DoAction(request, response)
	return
}

// GetAllApModelWithChan invokes the cloudwf.GetAllApModel API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getallapmodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAllApModelWithChan(request *GetAllApModelRequest) (<-chan *GetAllApModelResponse, <-chan error) {
	responseChan := make(chan *GetAllApModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAllApModel(request)
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

// GetAllApModelWithCallback invokes the cloudwf.GetAllApModel API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getallapmodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAllApModelWithCallback(request *GetAllApModelRequest, callback func(response *GetAllApModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAllApModelResponse
		var err error
		defer close(result)
		response, err = client.GetAllApModel(request)
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

// GetAllApModelRequest is the request struct for api GetAllApModel
type GetAllApModelRequest struct {
	*requests.RpcRequest
}

// GetAllApModelResponse is the response struct for api GetAllApModel
type GetAllApModelResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetAllApModelRequest creates a request to invoke GetAllApModel API
func CreateGetAllApModelRequest() (request *GetAllApModelRequest) {
	request = &GetAllApModelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetAllApModel", "cloudwf", "openAPI")
	return
}

// CreateGetAllApModelResponse creates a response to parse from GetAllApModel response
func CreateGetAllApModelResponse() (response *GetAllApModelResponse) {
	response = &GetAllApModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
