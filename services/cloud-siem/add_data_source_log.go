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

// AddDataSourceLog invokes the cloud_siem.AddDataSourceLog API synchronously
func (client *Client) AddDataSourceLog(request *AddDataSourceLogRequest) (response *AddDataSourceLogResponse, err error) {
	response = CreateAddDataSourceLogResponse()
	err = client.DoAction(request, response)
	return
}

// AddDataSourceLogWithChan invokes the cloud_siem.AddDataSourceLog API asynchronously
func (client *Client) AddDataSourceLogWithChan(request *AddDataSourceLogRequest) (<-chan *AddDataSourceLogResponse, <-chan error) {
	responseChan := make(chan *AddDataSourceLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDataSourceLog(request)
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

// AddDataSourceLogWithCallback invokes the cloud_siem.AddDataSourceLog API asynchronously
func (client *Client) AddDataSourceLogWithCallback(request *AddDataSourceLogRequest, callback func(response *AddDataSourceLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDataSourceLogResponse
		var err error
		defer close(result)
		response, err = client.AddDataSourceLog(request)
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

// AddDataSourceLogRequest is the request struct for api AddDataSourceLog
type AddDataSourceLogRequest struct {
	*requests.RpcRequest
	CloudCode              string `position:"Body" name:"CloudCode"`
	AccountId              string `position:"Body" name:"AccountId"`
	LogCode                string `position:"Body" name:"LogCode"`
	DataSourceInstanceLogs string `position:"Body" name:"DataSourceInstanceLogs"`
	DataSourceInstanceId   string `position:"Body" name:"DataSourceInstanceId"`
}

// AddDataSourceLogResponse is the response struct for api AddDataSourceLog
type AddDataSourceLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAddDataSourceLogRequest creates a request to invoke AddDataSourceLog API
func CreateAddDataSourceLogRequest() (request *AddDataSourceLogRequest) {
	request = &AddDataSourceLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "AddDataSourceLog", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddDataSourceLogResponse creates a response to parse from AddDataSourceLog response
func CreateAddDataSourceLogResponse() (response *AddDataSourceLogResponse) {
	response = &AddDataSourceLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}