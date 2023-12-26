package codeup

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

// GetRepositoryInfo invokes the codeup.GetRepositoryInfo API synchronously
func (client *Client) GetRepositoryInfo(request *GetRepositoryInfoRequest) (response *GetRepositoryInfoResponse, err error) {
	response = CreateGetRepositoryInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetRepositoryInfoWithChan invokes the codeup.GetRepositoryInfo API asynchronously
func (client *Client) GetRepositoryInfoWithChan(request *GetRepositoryInfoRequest) (<-chan *GetRepositoryInfoResponse, <-chan error) {
	responseChan := make(chan *GetRepositoryInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRepositoryInfo(request)
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

// GetRepositoryInfoWithCallback invokes the codeup.GetRepositoryInfo API asynchronously
func (client *Client) GetRepositoryInfoWithCallback(request *GetRepositoryInfoRequest, callback func(response *GetRepositoryInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRepositoryInfoResponse
		var err error
		defer close(result)
		response, err = client.GetRepositoryInfo(request)
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

// GetRepositoryInfoRequest is the request struct for api GetRepositoryInfo
type GetRepositoryInfoRequest struct {
	*requests.RoaRequest
	OrganizationId string `position:"Query" name:"OrganizationId"`
	Identity       string `position:"Query" name:"Identity"`
	AccessToken    string `position:"Query" name:"AccessToken"`
}

// GetRepositoryInfoResponse is the response struct for api GetRepositoryInfo
type GetRepositoryInfoResponse struct {
	*responses.BaseResponse
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Result       Result `json:"Result" xml:"Result"`
}

// CreateGetRepositoryInfoRequest creates a request to invoke GetRepositoryInfo API
func CreateGetRepositoryInfoRequest() (request *GetRepositoryInfoRequest) {
	request = &GetRepositoryInfoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("codeup", "2020-04-14", "GetRepositoryInfo", "/api/v3/projects/info", "", "")
	request.Method = requests.GET
	return
}

// CreateGetRepositoryInfoResponse creates a response to parse from GetRepositoryInfo response
func CreateGetRepositoryInfoResponse() (response *GetRepositoryInfoResponse) {
	response = &GetRepositoryInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
