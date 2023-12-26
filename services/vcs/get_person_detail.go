package vcs

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

// GetPersonDetail invokes the vcs.GetPersonDetail API synchronously
func (client *Client) GetPersonDetail(request *GetPersonDetailRequest) (response *GetPersonDetailResponse, err error) {
	response = CreateGetPersonDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetPersonDetailWithChan invokes the vcs.GetPersonDetail API asynchronously
func (client *Client) GetPersonDetailWithChan(request *GetPersonDetailRequest) (<-chan *GetPersonDetailResponse, <-chan error) {
	responseChan := make(chan *GetPersonDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPersonDetail(request)
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

// GetPersonDetailWithCallback invokes the vcs.GetPersonDetail API asynchronously
func (client *Client) GetPersonDetailWithCallback(request *GetPersonDetailRequest, callback func(response *GetPersonDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPersonDetailResponse
		var err error
		defer close(result)
		response, err = client.GetPersonDetail(request)
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

// GetPersonDetailRequest is the request struct for api GetPersonDetail
type GetPersonDetailRequest struct {
	*requests.RpcRequest
	AlgorithmType string `position:"Body" name:"AlgorithmType"`
	CorpId        string `position:"Body" name:"CorpId"`
	PersonID      string `position:"Body" name:"PersonID"`
}

// GetPersonDetailResponse is the response struct for api GetPersonDetail
type GetPersonDetailResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetPersonDetailRequest creates a request to invoke GetPersonDetail API
func CreateGetPersonDetailRequest() (request *GetPersonDetailRequest) {
	request = &GetPersonDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetPersonDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPersonDetailResponse creates a response to parse from GetPersonDetail response
func CreateGetPersonDetailResponse() (response *GetPersonDetailResponse) {
	response = &GetPersonDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
