package dataworks_public

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

// CreateDagTest invokes the dataworks_public.CreateDagTest API synchronously
func (client *Client) CreateDagTest(request *CreateDagTestRequest) (response *CreateDagTestResponse, err error) {
	response = CreateCreateDagTestResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDagTestWithChan invokes the dataworks_public.CreateDagTest API asynchronously
func (client *Client) CreateDagTestWithChan(request *CreateDagTestRequest) (<-chan *CreateDagTestResponse, <-chan error) {
	responseChan := make(chan *CreateDagTestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDagTest(request)
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

// CreateDagTestWithCallback invokes the dataworks_public.CreateDagTest API asynchronously
func (client *Client) CreateDagTestWithCallback(request *CreateDagTestRequest, callback func(response *CreateDagTestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDagTestResponse
		var err error
		defer close(result)
		response, err = client.CreateDagTest(request)
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

// CreateDagTestRequest is the request struct for api CreateDagTest
type CreateDagTestRequest struct {
	*requests.RpcRequest
	ProjectEnv string           `position:"Body" name:"ProjectEnv"`
	Bizdate    string           `position:"Body" name:"Bizdate"`
	Name       string           `position:"Body" name:"Name"`
	NodeParams string           `position:"Body" name:"NodeParams"`
	NodeId     requests.Integer `position:"Body" name:"NodeId"`
}

// CreateDagTestResponse is the response struct for api CreateDagTest
type CreateDagTestResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           int64  `json:"Data" xml:"Data"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateCreateDagTestRequest creates a request to invoke CreateDagTest API
func CreateCreateDagTestRequest() (request *CreateDagTestRequest) {
	request = &CreateDagTestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateDagTest", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDagTestResponse creates a response to parse from CreateDagTest response
func CreateCreateDagTestResponse() (response *CreateDagTestResponse) {
	response = &CreateDagTestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
