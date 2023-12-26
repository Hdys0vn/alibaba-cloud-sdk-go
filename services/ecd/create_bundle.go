package ecd

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

// CreateBundle invokes the ecd.CreateBundle API synchronously
func (client *Client) CreateBundle(request *CreateBundleRequest) (response *CreateBundleResponse, err error) {
	response = CreateCreateBundleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBundleWithChan invokes the ecd.CreateBundle API asynchronously
func (client *Client) CreateBundleWithChan(request *CreateBundleRequest) (<-chan *CreateBundleResponse, <-chan error) {
	responseChan := make(chan *CreateBundleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBundle(request)
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

// CreateBundleWithCallback invokes the ecd.CreateBundle API asynchronously
func (client *Client) CreateBundleWithCallback(request *CreateBundleRequest, callback func(response *CreateBundleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBundleResponse
		var err error
		defer close(result)
		response, err = client.CreateBundle(request)
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

// CreateBundleRequest is the request struct for api CreateBundle
type CreateBundleRequest struct {
	*requests.RpcRequest
	RootDiskPerformanceLevel string           `position:"Query" name:"RootDiskPerformanceLevel"`
	ImageId                  string           `position:"Query" name:"ImageId"`
	Description              string           `position:"Query" name:"Description"`
	Language                 string           `position:"Query" name:"Language"`
	UserDiskPerformanceLevel string           `position:"Query" name:"UserDiskPerformanceLevel"`
	DesktopType              string           `position:"Query" name:"DesktopType"`
	BundleName               string           `position:"Query" name:"BundleName"`
	UserDiskSizeGib          *[]string        `position:"Query" name:"UserDiskSizeGib"  type:"Repeated"`
	RootDiskSizeGib          requests.Integer `position:"Query" name:"RootDiskSizeGib"`
}

// CreateBundleResponse is the response struct for api CreateBundle
type CreateBundleResponse struct {
	*responses.BaseResponse
	BundleId  string `json:"BundleId" xml:"BundleId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateBundleRequest creates a request to invoke CreateBundle API
func CreateCreateBundleRequest() (request *CreateBundleRequest) {
	request = &CreateBundleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CreateBundle", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateBundleResponse creates a response to parse from CreateBundle response
func CreateCreateBundleResponse() (response *CreateBundleResponse) {
	response = &CreateBundleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
