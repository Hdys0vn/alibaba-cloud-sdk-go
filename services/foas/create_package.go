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

// CreatePackage invokes the foas.CreatePackage API synchronously
func (client *Client) CreatePackage(request *CreatePackageRequest) (response *CreatePackageResponse, err error) {
	response = CreateCreatePackageResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePackageWithChan invokes the foas.CreatePackage API asynchronously
func (client *Client) CreatePackageWithChan(request *CreatePackageRequest) (<-chan *CreatePackageResponse, <-chan error) {
	responseChan := make(chan *CreatePackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePackage(request)
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

// CreatePackageWithCallback invokes the foas.CreatePackage API asynchronously
func (client *Client) CreatePackageWithCallback(request *CreatePackageRequest, callback func(response *CreatePackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePackageResponse
		var err error
		defer close(result)
		response, err = client.CreatePackage(request)
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

// CreatePackageRequest is the request struct for api CreatePackage
type CreatePackageRequest struct {
	*requests.RoaRequest
	ProjectName string `position:"Path" name:"projectName"`
	OssBucket   string `position:"Body" name:"ossBucket"`
	OssOwner    string `position:"Body" name:"ossOwner"`
	PackageName string `position:"Body" name:"packageName"`
	OssEndpoint string `position:"Body" name:"ossEndpoint"`
	Description string `position:"Body" name:"description"`
	Tag         string `position:"Body" name:"tag"`
	OriginName  string `position:"Body" name:"originName"`
	Type        string `position:"Body" name:"type"`
	OssPath     string `position:"Body" name:"ossPath"`
	Md5         string `position:"Body" name:"md5"`
}

// CreatePackageResponse is the response struct for api CreatePackage
type CreatePackageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreatePackageRequest creates a request to invoke CreatePackage API
func CreateCreatePackageRequest() (request *CreatePackageRequest) {
	request = &CreatePackageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "CreatePackage", "/api/v2/projects/[projectName]/packages", "foas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreatePackageResponse creates a response to parse from CreatePackage response
func CreateCreatePackageResponse() (response *CreatePackageResponse) {
	response = &CreatePackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
