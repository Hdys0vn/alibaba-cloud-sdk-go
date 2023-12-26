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

// CreateDIJob invokes the dataworks_public.CreateDIJob API synchronously
func (client *Client) CreateDIJob(request *CreateDIJobRequest) (response *CreateDIJobResponse, err error) {
	response = CreateCreateDIJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDIJobWithChan invokes the dataworks_public.CreateDIJob API asynchronously
func (client *Client) CreateDIJobWithChan(request *CreateDIJobRequest) (<-chan *CreateDIJobResponse, <-chan error) {
	responseChan := make(chan *CreateDIJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDIJob(request)
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

// CreateDIJobWithCallback invokes the dataworks_public.CreateDIJob API asynchronously
func (client *Client) CreateDIJobWithCallback(request *CreateDIJobRequest, callback func(response *CreateDIJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDIJobResponse
		var err error
		defer close(result)
		response, err = client.CreateDIJob(request)
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

// CreateDIJobRequest is the request struct for api CreateDIJob
type CreateDIJobRequest struct {
	*requests.RpcRequest
	SourceDataSourceType          string           `position:"Body" name:"SourceDataSourceType"`
	Description                   string           `position:"Body" name:"Description"`
	TransformationRules           string           `position:"Body" name:"TransformationRules"`
	DestinationDataSourceType     string           `position:"Body" name:"DestinationDataSourceType"`
	DestinationDataSourceSettings string           `position:"Body" name:"DestinationDataSourceSettings"`
	SourceDataSourceSettings      string           `position:"Body" name:"SourceDataSourceSettings"`
	ResourceSettings              string           `position:"Body" name:"ResourceSettings"`
	MigrationType                 string           `position:"Body" name:"MigrationType"`
	SystemDebug                   string           `position:"Query" name:"SystemDebug"`
	ProjectId                     requests.Integer `position:"Body" name:"ProjectId"`
	JobName                       string           `position:"Body" name:"JobName"`
	TableMappings                 string           `position:"Body" name:"TableMappings"`
	JobSettings                   string           `position:"Body" name:"JobSettings"`
}

// CreateDIJobResponse is the response struct for api CreateDIJob
type CreateDIJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	DIJobId   int64  `json:"DIJobId" xml:"DIJobId"`
}

// CreateCreateDIJobRequest creates a request to invoke CreateDIJob API
func CreateCreateDIJobRequest() (request *CreateDIJobRequest) {
	request = &CreateDIJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateDIJob", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDIJobResponse creates a response to parse from CreateDIJob response
func CreateCreateDIJobResponse() (response *CreateDIJobResponse) {
	response = &CreateDIJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
