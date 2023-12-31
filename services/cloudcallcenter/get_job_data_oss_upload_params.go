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

// GetJobDataOssUploadParams invokes the cloudcallcenter.GetJobDataOssUploadParams API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getjobdataossuploadparams.html
func (client *Client) GetJobDataOssUploadParams(request *GetJobDataOssUploadParamsRequest) (response *GetJobDataOssUploadParamsResponse, err error) {
	response = CreateGetJobDataOssUploadParamsResponse()
	err = client.DoAction(request, response)
	return
}

// GetJobDataOssUploadParamsWithChan invokes the cloudcallcenter.GetJobDataOssUploadParams API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getjobdataossuploadparams.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetJobDataOssUploadParamsWithChan(request *GetJobDataOssUploadParamsRequest) (<-chan *GetJobDataOssUploadParamsResponse, <-chan error) {
	responseChan := make(chan *GetJobDataOssUploadParamsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJobDataOssUploadParams(request)
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

// GetJobDataOssUploadParamsWithCallback invokes the cloudcallcenter.GetJobDataOssUploadParams API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getjobdataossuploadparams.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetJobDataOssUploadParamsWithCallback(request *GetJobDataOssUploadParamsRequest, callback func(response *GetJobDataOssUploadParamsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJobDataOssUploadParamsResponse
		var err error
		defer close(result)
		response, err = client.GetJobDataOssUploadParams(request)
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

// GetJobDataOssUploadParamsRequest is the request struct for api GetJobDataOssUploadParams
type GetJobDataOssUploadParamsRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	FileName   string `position:"Query" name:"FileName"`
}

// GetJobDataOssUploadParamsResponse is the response struct for api GetJobDataOssUploadParams
type GetJobDataOssUploadParamsResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Success        bool         `json:"Success" xml:"Success"`
	Code           string       `json:"Code" xml:"Code"`
	Message        string       `json:"Message" xml:"Message"`
	HttpStatusCode int          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	UploadParams   UploadParams `json:"UploadParams" xml:"UploadParams"`
}

// CreateGetJobDataOssUploadParamsRequest creates a request to invoke GetJobDataOssUploadParams API
func CreateGetJobDataOssUploadParamsRequest() (request *GetJobDataOssUploadParamsRequest) {
	request = &GetJobDataOssUploadParamsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetJobDataOssUploadParams", "", "")
	request.Method = requests.POST
	return
}

// CreateGetJobDataOssUploadParamsResponse creates a response to parse from GetJobDataOssUploadParams response
func CreateGetJobDataOssUploadParamsResponse() (response *GetJobDataOssUploadParamsResponse) {
	response = &GetJobDataOssUploadParamsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
