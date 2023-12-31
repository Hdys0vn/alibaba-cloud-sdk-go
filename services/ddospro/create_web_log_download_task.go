package ddospro

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

// CreateWebLogDownloadTask invokes the ddospro.CreateWebLogDownloadTask API synchronously
// api document: https://help.aliyun.com/api/ddospro/createweblogdownloadtask.html
func (client *Client) CreateWebLogDownloadTask(request *CreateWebLogDownloadTaskRequest) (response *CreateWebLogDownloadTaskResponse, err error) {
	response = CreateCreateWebLogDownloadTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateWebLogDownloadTaskWithChan invokes the ddospro.CreateWebLogDownloadTask API asynchronously
// api document: https://help.aliyun.com/api/ddospro/createweblogdownloadtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateWebLogDownloadTaskWithChan(request *CreateWebLogDownloadTaskRequest) (<-chan *CreateWebLogDownloadTaskResponse, <-chan error) {
	responseChan := make(chan *CreateWebLogDownloadTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateWebLogDownloadTask(request)
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

// CreateWebLogDownloadTaskWithCallback invokes the ddospro.CreateWebLogDownloadTask API asynchronously
// api document: https://help.aliyun.com/api/ddospro/createweblogdownloadtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateWebLogDownloadTaskWithCallback(request *CreateWebLogDownloadTaskRequest, callback func(response *CreateWebLogDownloadTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateWebLogDownloadTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateWebLogDownloadTask(request)
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

// CreateWebLogDownloadTaskRequest is the request struct for api CreateWebLogDownloadTask
type CreateWebLogDownloadTaskRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Domain          string           `position:"Query" name:"Domain"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
}

// CreateWebLogDownloadTaskResponse is the response struct for api CreateWebLogDownloadTask
type CreateWebLogDownloadTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateWebLogDownloadTaskRequest creates a request to invoke CreateWebLogDownloadTask API
func CreateCreateWebLogDownloadTaskRequest() (request *CreateWebLogDownloadTaskRequest) {
	request = &CreateWebLogDownloadTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "CreateWebLogDownloadTask", "", "")
	return
}

// CreateCreateWebLogDownloadTaskResponse creates a response to parse from CreateWebLogDownloadTask response
func CreateCreateWebLogDownloadTaskResponse() (response *CreateWebLogDownloadTaskResponse) {
	response = &CreateWebLogDownloadTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
