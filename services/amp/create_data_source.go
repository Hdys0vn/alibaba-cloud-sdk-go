package amp

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

// CreateDataSource invokes the amp.CreateDataSource API synchronously
func (client *Client) CreateDataSource(request *CreateDataSourceRequest) (response *CreateDataSourceResponse, err error) {
	response = CreateCreateDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataSourceWithChan invokes the amp.CreateDataSource API asynchronously
func (client *Client) CreateDataSourceWithChan(request *CreateDataSourceRequest) (<-chan *CreateDataSourceResponse, <-chan error) {
	responseChan := make(chan *CreateDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataSource(request)
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

// CreateDataSourceWithCallback invokes the amp.CreateDataSource API asynchronously
func (client *Client) CreateDataSourceWithCallback(request *CreateDataSourceRequest, callback func(response *CreateDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataSourceResponse
		var err error
		defer close(result)
		response, err = client.CreateDataSource(request)
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

// CreateDataSourceRequest is the request struct for api CreateDataSource
type CreateDataSourceRequest struct {
	*requests.RoaRequest
	Body string `position:"Body" name:"body"`
}

// CreateDataSourceResponse is the response struct for api CreateDataSource
type CreateDataSourceResponse struct {
	*responses.BaseResponse
	DataSourceId string `json:"DataSourceId" xml:"DataSourceId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDataSourceRequest creates a request to invoke CreateDataSource API
func CreateCreateDataSourceRequest() (request *CreateDataSourceRequest) {
	request = &CreateDataSourceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("amp", "2020-07-08", "CreateDataSource", "/api/v1/datasources", "ServiceCode", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDataSourceResponse creates a response to parse from CreateDataSource response
func CreateCreateDataSourceResponse() (response *CreateDataSourceResponse) {
	response = &CreateDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}