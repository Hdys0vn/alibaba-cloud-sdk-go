package dms_enterprise

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

// GetDatabase invokes the dms_enterprise.GetDatabase API synchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdatabase.html
func (client *Client) GetDatabase(request *GetDatabaseRequest) (response *GetDatabaseResponse, err error) {
	response = CreateGetDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// GetDatabaseWithChan invokes the dms_enterprise.GetDatabase API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDatabaseWithChan(request *GetDatabaseRequest) (<-chan *GetDatabaseResponse, <-chan error) {
	responseChan := make(chan *GetDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDatabase(request)
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

// GetDatabaseWithCallback invokes the dms_enterprise.GetDatabase API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDatabaseWithCallback(request *GetDatabaseRequest, callback func(response *GetDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDatabaseResponse
		var err error
		defer close(result)
		response, err = client.GetDatabase(request)
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

// GetDatabaseRequest is the request struct for api GetDatabase
type GetDatabaseRequest struct {
	*requests.RpcRequest
	SchemaName string           `position:"Query" name:"SchemaName"`
	Port       requests.Integer `position:"Query" name:"Port"`
	Host       string           `position:"Query" name:"Host"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	Sid        string           `position:"Query" name:"Sid"`
}

// GetDatabaseResponse is the response struct for api GetDatabase
type GetDatabaseResponse struct {
	*responses.BaseResponse
	RequestId    string   `json:"RequestId" xml:"RequestId"`
	Success      bool     `json:"Success" xml:"Success"`
	ErrorMessage string   `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string   `json:"ErrorCode" xml:"ErrorCode"`
	Database     Database `json:"Database" xml:"Database"`
}

// CreateGetDatabaseRequest creates a request to invoke GetDatabase API
func CreateGetDatabaseRequest() (request *GetDatabaseRequest) {
	request = &GetDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetDatabase", "dmsenterprise", "openAPI")
	return
}

// CreateGetDatabaseResponse creates a response to parse from GetDatabase response
func CreateGetDatabaseResponse() (response *GetDatabaseResponse) {
	response = &GetDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
