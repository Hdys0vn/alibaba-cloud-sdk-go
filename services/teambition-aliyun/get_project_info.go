package teambition_aliyun

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

// GetProjectInfo invokes the teambition_aliyun.GetProjectInfo API synchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/getprojectinfo.html
func (client *Client) GetProjectInfo(request *GetProjectInfoRequest) (response *GetProjectInfoResponse, err error) {
	response = CreateGetProjectInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetProjectInfoWithChan invokes the teambition_aliyun.GetProjectInfo API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/getprojectinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProjectInfoWithChan(request *GetProjectInfoRequest) (<-chan *GetProjectInfoResponse, <-chan error) {
	responseChan := make(chan *GetProjectInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProjectInfo(request)
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

// GetProjectInfoWithCallback invokes the teambition_aliyun.GetProjectInfo API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/getprojectinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProjectInfoWithCallback(request *GetProjectInfoRequest, callback func(response *GetProjectInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProjectInfoResponse
		var err error
		defer close(result)
		response, err = client.GetProjectInfo(request)
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

// GetProjectInfoRequest is the request struct for api GetProjectInfo
type GetProjectInfoRequest struct {
	*requests.RpcRequest
	ProjectId string `position:"Body" name:"ProjectId"`
	OrgId     string `position:"Body" name:"OrgId"`
}

// GetProjectInfoResponse is the response struct for api GetProjectInfo
type GetProjectInfoResponse struct {
	*responses.BaseResponse
	Successful bool   `json:"Successful" xml:"Successful"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Object     Object `json:"Object" xml:"Object"`
}

// CreateGetProjectInfoRequest creates a request to invoke GetProjectInfo API
func CreateGetProjectInfoRequest() (request *GetProjectInfoRequest) {
	request = &GetProjectInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("teambition-aliyun", "2020-02-26", "GetProjectInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateGetProjectInfoResponse creates a response to parse from GetProjectInfo response
func CreateGetProjectInfoResponse() (response *GetProjectInfoResponse) {
	response = &GetProjectInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
