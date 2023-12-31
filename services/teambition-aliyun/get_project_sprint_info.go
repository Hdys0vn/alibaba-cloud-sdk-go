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

// GetProjectSprintInfo invokes the teambition_aliyun.GetProjectSprintInfo API synchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/getprojectsprintinfo.html
func (client *Client) GetProjectSprintInfo(request *GetProjectSprintInfoRequest) (response *GetProjectSprintInfoResponse, err error) {
	response = CreateGetProjectSprintInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetProjectSprintInfoWithChan invokes the teambition_aliyun.GetProjectSprintInfo API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/getprojectsprintinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProjectSprintInfoWithChan(request *GetProjectSprintInfoRequest) (<-chan *GetProjectSprintInfoResponse, <-chan error) {
	responseChan := make(chan *GetProjectSprintInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProjectSprintInfo(request)
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

// GetProjectSprintInfoWithCallback invokes the teambition_aliyun.GetProjectSprintInfo API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/getprojectsprintinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetProjectSprintInfoWithCallback(request *GetProjectSprintInfoRequest, callback func(response *GetProjectSprintInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProjectSprintInfoResponse
		var err error
		defer close(result)
		response, err = client.GetProjectSprintInfo(request)
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

// GetProjectSprintInfoRequest is the request struct for api GetProjectSprintInfo
type GetProjectSprintInfoRequest struct {
	*requests.RpcRequest
	SprintId string `position:"Body" name:"SprintId"`
	OrgId    string `position:"Body" name:"OrgId"`
}

// GetProjectSprintInfoResponse is the response struct for api GetProjectSprintInfo
type GetProjectSprintInfoResponse struct {
	*responses.BaseResponse
	Successful bool   `json:"Successful" xml:"Successful"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Object     Object `json:"Object" xml:"Object"`
}

// CreateGetProjectSprintInfoRequest creates a request to invoke GetProjectSprintInfo API
func CreateGetProjectSprintInfoRequest() (request *GetProjectSprintInfoRequest) {
	request = &GetProjectSprintInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("teambition-aliyun", "2020-02-26", "GetProjectSprintInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateGetProjectSprintInfoResponse creates a response to parse from GetProjectSprintInfo response
func CreateGetProjectSprintInfoResponse() (response *GetProjectSprintInfoResponse) {
	response = &GetProjectSprintInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
