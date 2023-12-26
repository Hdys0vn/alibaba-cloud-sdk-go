package cloud_siem

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

// ModifyDataSourceLog invokes the cloud_siem.ModifyDataSourceLog API synchronously
func (client *Client) ModifyDataSourceLog(request *ModifyDataSourceLogRequest) (response *ModifyDataSourceLogResponse, err error) {
	response = CreateModifyDataSourceLogResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDataSourceLogWithChan invokes the cloud_siem.ModifyDataSourceLog API asynchronously
func (client *Client) ModifyDataSourceLogWithChan(request *ModifyDataSourceLogRequest) (<-chan *ModifyDataSourceLogResponse, <-chan error) {
	responseChan := make(chan *ModifyDataSourceLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDataSourceLog(request)
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

// ModifyDataSourceLogWithCallback invokes the cloud_siem.ModifyDataSourceLog API asynchronously
func (client *Client) ModifyDataSourceLogWithCallback(request *ModifyDataSourceLogRequest, callback func(response *ModifyDataSourceLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDataSourceLogResponse
		var err error
		defer close(result)
		response, err = client.ModifyDataSourceLog(request)
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

// ModifyDataSourceLogRequest is the request struct for api ModifyDataSourceLog
type ModifyDataSourceLogRequest struct {
	*requests.RpcRequest
	DataSourceType         string `position:"Body" name:"DataSourceType"`
	CloudCode              string `position:"Body" name:"CloudCode"`
	AccountId              string `position:"Body" name:"AccountId"`
	LogCode                string `position:"Body" name:"LogCode"`
	LogInstanceId          string `position:"Body" name:"LogInstanceId"`
	DataSourceInstanceLogs string `position:"Body" name:"DataSourceInstanceLogs"`
	DataSourceInstanceId   string `position:"Body" name:"DataSourceInstanceId"`
}

// ModifyDataSourceLogResponse is the response struct for api ModifyDataSourceLog
type ModifyDataSourceLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateModifyDataSourceLogRequest creates a request to invoke ModifyDataSourceLog API
func CreateModifyDataSourceLogRequest() (request *ModifyDataSourceLogRequest) {
	request = &ModifyDataSourceLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "ModifyDataSourceLog", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDataSourceLogResponse creates a response to parse from ModifyDataSourceLog response
func CreateModifyDataSourceLogResponse() (response *ModifyDataSourceLogResponse) {
	response = &ModifyDataSourceLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
