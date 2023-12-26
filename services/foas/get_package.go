package foas

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

// GetPackage invokes the foas.GetPackage API synchronously
func (client *Client) GetPackage(request *GetPackageRequest) (response *GetPackageResponse, err error) {
	response = CreateGetPackageResponse()
	err = client.DoAction(request, response)
	return
}

// GetPackageWithChan invokes the foas.GetPackage API asynchronously
func (client *Client) GetPackageWithChan(request *GetPackageRequest) (<-chan *GetPackageResponse, <-chan error) {
	responseChan := make(chan *GetPackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPackage(request)
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

// GetPackageWithCallback invokes the foas.GetPackage API asynchronously
func (client *Client) GetPackageWithCallback(request *GetPackageRequest, callback func(response *GetPackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPackageResponse
		var err error
		defer close(result)
		response, err = client.GetPackage(request)
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

// GetPackageRequest is the request struct for api GetPackage
type GetPackageRequest struct {
	*requests.RoaRequest
	ProjectName string `position:"Path" name:"projectName"`
	PackageName string `position:"Path" name:"packageName"`
}

// GetPackageResponse is the response struct for api GetPackage
type GetPackageResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Package   Package `json:"Package" xml:"Package"`
}

// CreateGetPackageRequest creates a request to invoke GetPackage API
func CreateGetPackageRequest() (request *GetPackageRequest) {
	request = &GetPackageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "GetPackage", "/api/v2/projects/[projectName]/packages/[packageName]", "foas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetPackageResponse creates a response to parse from GetPackage response
func CreateGetPackageResponse() (response *GetPackageResponse) {
	response = &GetPackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
