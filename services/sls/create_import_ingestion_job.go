package sls

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

// CreateImportIngestionJob invokes the sls.CreateImportIngestionJob API synchronously
func (client *Client) CreateImportIngestionJob(request *CreateImportIngestionJobRequest) (response *CreateImportIngestionJobResponse, err error) {
	response = CreateCreateImportIngestionJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateImportIngestionJobWithChan invokes the sls.CreateImportIngestionJob API asynchronously
func (client *Client) CreateImportIngestionJobWithChan(request *CreateImportIngestionJobRequest) (<-chan *CreateImportIngestionJobResponse, <-chan error) {
	responseChan := make(chan *CreateImportIngestionJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateImportIngestionJob(request)
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

// CreateImportIngestionJobWithCallback invokes the sls.CreateImportIngestionJob API asynchronously
func (client *Client) CreateImportIngestionJobWithCallback(request *CreateImportIngestionJobRequest, callback func(response *CreateImportIngestionJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateImportIngestionJobResponse
		var err error
		defer close(result)
		response, err = client.CreateImportIngestionJob(request)
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

// CreateImportIngestionJobRequest is the request struct for api CreateImportIngestionJob
type CreateImportIngestionJobRequest struct {
	*requests.RpcRequest
	Args            string           `position:"Body" name:"Args"`
	Image           string           `position:"Body" name:"Image"`
	EnvFromSecret   string           `position:"Body" name:"EnvFromSecret"`
	RestartPolicy   string           `position:"Body" name:"RestartPolicy"`
	Parallelism     requests.Integer `position:"Body" name:"Parallelism"`
	Namespace       string           `position:"Body" name:"Namespace"`
	Completions     requests.Integer `position:"Body" name:"Completions"`
	EnvVars         string           `position:"Body" name:"EnvVars"`
	ImagePullSecret string           `position:"Body" name:"ImagePullSecret"`
	CallerId        string           `position:"Body" name:"CallerId"`
	Region          string           `position:"Body" name:"Region"`
	JobName         string           `position:"Body" name:"JobName"`
}

// CreateImportIngestionJobResponse is the response struct for api CreateImportIngestionJob
type CreateImportIngestionJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCreateImportIngestionJobRequest creates a request to invoke CreateImportIngestionJob API
func CreateCreateImportIngestionJobRequest() (request *CreateImportIngestionJobRequest) {
	request = &CreateImportIngestionJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2019-10-23", "CreateImportIngestionJob", "sls", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateImportIngestionJobResponse creates a response to parse from CreateImportIngestionJob response
func CreateCreateImportIngestionJobResponse() (response *CreateImportIngestionJobResponse) {
	response = &CreateImportIngestionJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
