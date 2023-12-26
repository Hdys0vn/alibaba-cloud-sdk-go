package sas

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

// GetCheckDetail invokes the sas.GetCheckDetail API synchronously
func (client *Client) GetCheckDetail(request *GetCheckDetailRequest) (response *GetCheckDetailResponse, err error) {
	response = CreateGetCheckDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetCheckDetailWithChan invokes the sas.GetCheckDetail API asynchronously
func (client *Client) GetCheckDetailWithChan(request *GetCheckDetailRequest) (<-chan *GetCheckDetailResponse, <-chan error) {
	responseChan := make(chan *GetCheckDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCheckDetail(request)
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

// GetCheckDetailWithCallback invokes the sas.GetCheckDetail API asynchronously
func (client *Client) GetCheckDetailWithCallback(request *GetCheckDetailRequest, callback func(response *GetCheckDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCheckDetailResponse
		var err error
		defer close(result)
		response, err = client.GetCheckDetail(request)
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

// GetCheckDetailRequest is the request struct for api GetCheckDetail
type GetCheckDetailRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	Lang     string           `position:"Query" name:"Lang"`
	CheckId  requests.Integer `position:"Query" name:"CheckId"`
}

// GetCheckDetailResponse is the response struct for api GetCheckDetail
type GetCheckDetailResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Description Description `json:"Description" xml:"Description"`
	Solution    Solution    `json:"Solution" xml:"Solution"`
	AssistInfo  AssistInfo  `json:"AssistInfo" xml:"AssistInfo"`
}

// CreateGetCheckDetailRequest creates a request to invoke GetCheckDetail API
func CreateGetCheckDetailRequest() (request *GetCheckDetailRequest) {
	request = &GetCheckDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "GetCheckDetail", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetCheckDetailResponse creates a response to parse from GetCheckDetail response
func CreateGetCheckDetailResponse() (response *GetCheckDetailResponse) {
	response = &GetCheckDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}