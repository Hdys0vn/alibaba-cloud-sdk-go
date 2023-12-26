package mse

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

// ExportZookeeperData invokes the mse.ExportZookeeperData API synchronously
func (client *Client) ExportZookeeperData(request *ExportZookeeperDataRequest) (response *ExportZookeeperDataResponse, err error) {
	response = CreateExportZookeeperDataResponse()
	err = client.DoAction(request, response)
	return
}

// ExportZookeeperDataWithChan invokes the mse.ExportZookeeperData API asynchronously
func (client *Client) ExportZookeeperDataWithChan(request *ExportZookeeperDataRequest) (<-chan *ExportZookeeperDataResponse, <-chan error) {
	responseChan := make(chan *ExportZookeeperDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportZookeeperData(request)
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

// ExportZookeeperDataWithCallback invokes the mse.ExportZookeeperData API asynchronously
func (client *Client) ExportZookeeperDataWithCallback(request *ExportZookeeperDataRequest, callback func(response *ExportZookeeperDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportZookeeperDataResponse
		var err error
		defer close(result)
		response, err = client.ExportZookeeperData(request)
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

// ExportZookeeperDataRequest is the request struct for api ExportZookeeperData
type ExportZookeeperDataRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	RequestPars    string `position:"Query" name:"RequestPars"`
	ExportType     string `position:"Query" name:"ExportType"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// ExportZookeeperDataResponse is the response struct for api ExportZookeeperData
type ExportZookeeperDataResponse struct {
	*responses.BaseResponse
	Success        bool                      `json:"Success" xml:"Success"`
	Message        string                    `json:"Message" xml:"Message"`
	DynamicMessage string                    `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId      string                    `json:"RequestId" xml:"RequestId"`
	ErrorCode      string                    `json:"ErrorCode" xml:"ErrorCode"`
	HttpStatusCode string                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInExportZookeeperData `json:"Data" xml:"Data"`
}

// CreateExportZookeeperDataRequest creates a request to invoke ExportZookeeperData API
func CreateExportZookeeperDataRequest() (request *ExportZookeeperDataRequest) {
	request = &ExportZookeeperDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ExportZookeeperData", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExportZookeeperDataResponse creates a response to parse from ExportZookeeperData response
func CreateExportZookeeperDataResponse() (response *ExportZookeeperDataResponse) {
	response = &ExportZookeeperDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
