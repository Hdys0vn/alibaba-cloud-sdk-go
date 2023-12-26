package ivpd

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

// GetJobResult invokes the ivpd.GetJobResult API synchronously
// api document: https://help.aliyun.com/api/ivpd/getjobresult.html
func (client *Client) GetJobResult(request *GetJobResultRequest) (response *GetJobResultResponse, err error) {
	response = CreateGetJobResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetJobResultWithChan invokes the ivpd.GetJobResult API asynchronously
// api document: https://help.aliyun.com/api/ivpd/getjobresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetJobResultWithChan(request *GetJobResultRequest) (<-chan *GetJobResultResponse, <-chan error) {
	responseChan := make(chan *GetJobResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJobResult(request)
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

// GetJobResultWithCallback invokes the ivpd.GetJobResult API asynchronously
// api document: https://help.aliyun.com/api/ivpd/getjobresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetJobResultWithCallback(request *GetJobResultRequest, callback func(response *GetJobResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJobResultResponse
		var err error
		defer close(result)
		response, err = client.GetJobResult(request)
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

// GetJobResultRequest is the request struct for api GetJobResult
type GetJobResultRequest struct {
	*requests.RpcRequest
	JobId string `position:"Body" name:"JobId"`
}

// GetJobResultResponse is the response struct for api GetJobResult
type GetJobResultResponse struct {
	*responses.BaseResponse
	RequestId string             `json:"RequestId" xml:"RequestId"`
	Code      string             `json:"Code" xml:"Code"`
	Message   string             `json:"Message" xml:"Message"`
	Data      DataInGetJobResult `json:"Data" xml:"Data"`
}

// CreateGetJobResultRequest creates a request to invoke GetJobResult API
func CreateGetJobResultRequest() (request *GetJobResultRequest) {
	request = &GetJobResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivpd", "2019-06-25", "GetJobResult", "ivpd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetJobResultResponse creates a response to parse from GetJobResult response
func CreateGetJobResultResponse() (response *GetJobResultResponse) {
	response = &GetJobResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
