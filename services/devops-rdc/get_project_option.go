package devops_rdc

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

// GetProjectOption invokes the devops_rdc.GetProjectOption API synchronously
func (client *Client) GetProjectOption(request *GetProjectOptionRequest) (response *GetProjectOptionResponse, err error) {
	response = CreateGetProjectOptionResponse()
	err = client.DoAction(request, response)
	return
}

// GetProjectOptionWithChan invokes the devops_rdc.GetProjectOption API asynchronously
func (client *Client) GetProjectOptionWithChan(request *GetProjectOptionRequest) (<-chan *GetProjectOptionResponse, <-chan error) {
	responseChan := make(chan *GetProjectOptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProjectOption(request)
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

// GetProjectOptionWithCallback invokes the devops_rdc.GetProjectOption API asynchronously
func (client *Client) GetProjectOptionWithCallback(request *GetProjectOptionRequest, callback func(response *GetProjectOptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProjectOptionResponse
		var err error
		defer close(result)
		response, err = client.GetProjectOption(request)
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

// GetProjectOptionRequest is the request struct for api GetProjectOption
type GetProjectOptionRequest struct {
	*requests.RpcRequest
	Query     string `position:"Body" name:"Query"`
	Type      string `position:"Body" name:"Type"`
	ProjectId string `position:"Body" name:"ProjectId"`
	OrgId     string `position:"Body" name:"OrgId"`
}

// GetProjectOptionResponse is the response struct for api GetProjectOption
type GetProjectOptionResponse struct {
	*responses.BaseResponse
	Successful bool     `json:"Successful" xml:"Successful"`
	ErrorCode  string   `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string   `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	Object     []Option `json:"Object" xml:"Object"`
}

// CreateGetProjectOptionRequest creates a request to invoke GetProjectOption API
func CreateGetProjectOptionRequest() (request *GetProjectOptionRequest) {
	request = &GetProjectOptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "GetProjectOption", "", "")
	request.Method = requests.POST
	return
}

// CreateGetProjectOptionResponse creates a response to parse from GetProjectOption response
func CreateGetProjectOptionResponse() (response *GetProjectOptionResponse) {
	response = &GetProjectOptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
