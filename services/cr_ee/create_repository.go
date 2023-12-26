package cr_ee

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

// CreateRepository invokes the cr.CreateRepository API synchronously
// api document: https://help.aliyun.com/api/cr/createrepository.html
func (client *Client) CreateRepository(request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
	response = CreateCreateRepositoryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRepositoryWithChan invokes the cr.CreateRepository API asynchronously
// api document: https://help.aliyun.com/api/cr/createrepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRepositoryWithChan(request *CreateRepositoryRequest) (<-chan *CreateRepositoryResponse, <-chan error) {
	responseChan := make(chan *CreateRepositoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRepository(request)
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

// CreateRepositoryWithCallback invokes the cr.CreateRepository API asynchronously
// api document: https://help.aliyun.com/api/cr/createrepository.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRepositoryWithCallback(request *CreateRepositoryRequest, callback func(response *CreateRepositoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRepositoryResponse
		var err error
		defer close(result)
		response, err = client.CreateRepository(request)
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

// CreateRepositoryRequest is the request struct for api CreateRepository
type CreateRepositoryRequest struct {
	*requests.RpcRequest
	RepoType          string `position:"Query" name:"RepoType"`
	Summary           string `position:"Query" name:"Summary"`
	InstanceId        string `position:"Query" name:"InstanceId"`
	RepoName          string `position:"Query" name:"RepoName"`
	RepoNamespaceName string `position:"Query" name:"RepoNamespaceName"`
	Detail            string `position:"Query" name:"Detail"`
}

// CreateRepositoryResponse is the response struct for api CreateRepository
type CreateRepositoryResponse struct {
	*responses.BaseResponse
	CreateRepositoryIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	Code                      string `json:"Code" xml:"Code"`
	RequestId                 string `json:"RequestId" xml:"RequestId"`
	RepoId                    string `json:"RepoId" xml:"RepoId"`
}

// CreateCreateRepositoryRequest creates a request to invoke CreateRepository API
func CreateCreateRepositoryRequest() (request *CreateRepositoryRequest) {
	request = &CreateRepositoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "CreateRepository", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateRepositoryResponse creates a response to parse from CreateRepository response
func CreateCreateRepositoryResponse() (response *CreateRepositoryResponse) {
	response = &CreateRepositoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
