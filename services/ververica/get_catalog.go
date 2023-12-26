package ververica

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

// GetCatalog invokes the ververica.GetCatalog API synchronously
func (client *Client) GetCatalog(request *GetCatalogRequest) (response *GetCatalogResponse, err error) {
	response = CreateGetCatalogResponse()
	err = client.DoAction(request, response)
	return
}

// GetCatalogWithChan invokes the ververica.GetCatalog API asynchronously
func (client *Client) GetCatalogWithChan(request *GetCatalogRequest) (<-chan *GetCatalogResponse, <-chan error) {
	responseChan := make(chan *GetCatalogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCatalog(request)
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

// GetCatalogWithCallback invokes the ververica.GetCatalog API asynchronously
func (client *Client) GetCatalogWithCallback(request *GetCatalogRequest, callback func(response *GetCatalogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCatalogResponse
		var err error
		defer close(result)
		response, err = client.GetCatalog(request)
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

// GetCatalogRequest is the request struct for api GetCatalog
type GetCatalogRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
	Cat       string `position:"Path" name:"cat"`
	Namespace string `position:"Path" name:"namespace"`
}

// GetCatalogResponse is the response struct for api GetCatalog
type GetCatalogResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateGetCatalogRequest creates a request to invoke GetCatalog API
func CreateGetCatalogRequest() (request *GetCatalogRequest) {
	request = &GetCatalogRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "GetCatalog", "/pop/workspaces/[workspace]/catalog/v1beta2/namespaces/[namespace]/catalogs/[cat]:getCatalog", "", "")
	request.Method = requests.GET
	return
}

// CreateGetCatalogResponse creates a response to parse from GetCatalog response
func CreateGetCatalogResponse() (response *GetCatalogResponse) {
	response = &GetCatalogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}