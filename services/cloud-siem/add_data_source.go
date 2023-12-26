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

// AddDataSource invokes the cloud_siem.AddDataSource API synchronously
func (client *Client) AddDataSource(request *AddDataSourceRequest) (response *AddDataSourceResponse, err error) {
	response = CreateAddDataSourceResponse()
	err = client.DoAction(request, response)
	return
}

// AddDataSourceWithChan invokes the cloud_siem.AddDataSource API asynchronously
func (client *Client) AddDataSourceWithChan(request *AddDataSourceRequest) (<-chan *AddDataSourceResponse, <-chan error) {
	responseChan := make(chan *AddDataSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDataSource(request)
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

// AddDataSourceWithCallback invokes the cloud_siem.AddDataSource API asynchronously
func (client *Client) AddDataSourceWithCallback(request *AddDataSourceRequest, callback func(response *AddDataSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDataSourceResponse
		var err error
		defer close(result)
		response, err = client.AddDataSource(request)
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

// AddDataSourceRequest is the request struct for api AddDataSource
type AddDataSourceRequest struct {
	*requests.RpcRequest
	DataSourceType           string `position:"Body" name:"DataSourceType"`
	CloudCode                string `position:"Body" name:"CloudCode"`
	DataSourceInstanceName   string `position:"Body" name:"DataSourceInstanceName"`
	AccountId                string `position:"Body" name:"AccountId"`
	DataSourceInstanceRemark string `position:"Body" name:"DataSourceInstanceRemark"`
	DataSourceInstanceParams string `position:"Body" name:"DataSourceInstanceParams"`
}

// AddDataSourceResponse is the response struct for api AddDataSource
type AddDataSourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAddDataSourceRequest creates a request to invoke AddDataSource API
func CreateAddDataSourceRequest() (request *AddDataSourceRequest) {
	request = &AddDataSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "AddDataSource", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddDataSourceResponse creates a response to parse from AddDataSource response
func CreateAddDataSourceResponse() (response *AddDataSourceResponse) {
	response = &AddDataSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}