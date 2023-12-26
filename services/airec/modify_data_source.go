package airec

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

// ModifyDataSource invokes the airec.ModifyDataSource API synchronously
func (client *Client) ModifyDataSource(request *ModifyDataSourceRequest) (response *ModifyDataSourceResponse, err error) {
	response = CreateModifyDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDataSourceWithChan invokes the airec.ModifyDataSource API asynchronously
func (client *Client) ModifyDataSourceWithChan(request *ModifyDataSourceRequest) (<-chan *ModifyDataSourceResponse, <-chan error) {
	responseChan := make(chan *ModifyDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDataSource(request)
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

// ModifyDataSourceWithCallback invokes the airec.ModifyDataSource API asynchronously
func (client *Client) ModifyDataSourceWithCallback(request *ModifyDataSourceRequest, callback func(response *ModifyDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDataSourceResponse
		var err error
		defer close(result)
		response, err = client.ModifyDataSource(request)
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

// ModifyDataSourceRequest is the request struct for api ModifyDataSource
type ModifyDataSourceRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"instanceId"`
	TableName  string `position:"Path" name:"tableName"`
}

// ModifyDataSourceResponse is the response struct for api ModifyDataSource
type ModifyDataSourceResponse struct {
	*responses.BaseResponse
	Code      string                   `json:"code" xml:"code"`
	Message   string                   `json:"message" xml:"message"`
	RequestId string                   `json:"requestId" xml:"requestId"`
	Result    ResultInModifyDataSource `json:"result" xml:"result"`
}

// CreateModifyDataSourceRequest creates a request to invoke ModifyDataSource API
func CreateModifyDataSourceRequest() (request *ModifyDataSourceRequest) {
	request = &ModifyDataSourceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "ModifyDataSource", "/v2/openapi/instances/[instanceId]/dataSources/[tableName]", "airec", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateModifyDataSourceResponse creates a response to parse from ModifyDataSource response
func CreateModifyDataSourceResponse() (response *ModifyDataSourceResponse) {
	response = &ModifyDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
