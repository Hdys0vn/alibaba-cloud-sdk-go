package reid

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

// ListDevicesImages invokes the reid.ListDevicesImages API synchronously
// api document: https://help.aliyun.com/api/reid/listdevicesimages.html
func (client *Client) ListDevicesImages(request *ListDevicesImagesRequest) (response *ListDevicesImagesResponse, err error) {
	response = CreateListDevicesImagesResponse()
	err = client.DoAction(request, response)
	return
}

// ListDevicesImagesWithChan invokes the reid.ListDevicesImages API asynchronously
// api document: https://help.aliyun.com/api/reid/listdevicesimages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDevicesImagesWithChan(request *ListDevicesImagesRequest) (<-chan *ListDevicesImagesResponse, <-chan error) {
	responseChan := make(chan *ListDevicesImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDevicesImages(request)
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

// ListDevicesImagesWithCallback invokes the reid.ListDevicesImages API asynchronously
// api document: https://help.aliyun.com/api/reid/listdevicesimages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDevicesImagesWithCallback(request *ListDevicesImagesRequest, callback func(response *ListDevicesImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDevicesImagesResponse
		var err error
		defer close(result)
		response, err = client.ListDevicesImages(request)
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

// ListDevicesImagesRequest is the request struct for api ListDevicesImages
type ListDevicesImagesRequest struct {
	*requests.RpcRequest
	IpcIdList string           `position:"Body" name:"IpcIdList"`
	StoreId   requests.Integer `position:"Body" name:"StoreId"`
}

// ListDevicesImagesResponse is the response struct for api ListDevicesImages
type ListDevicesImagesResponse struct {
	*responses.BaseResponse
	ErrorCode      string        `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string        `json:"ErrorMessage" xml:"ErrorMessage"`
	Message        string        `json:"Message" xml:"Message"`
	Code           string        `json:"Code" xml:"Code"`
	DynamicCode    string        `json:"DynamicCode" xml:"DynamicCode"`
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	Success        bool          `json:"Success" xml:"Success"`
	DynamicMessage string        `json:"DynamicMessage" xml:"DynamicMessage"`
	DeviceImages   []DeviceImage `json:"DeviceImages" xml:"DeviceImages"`
}

// CreateListDevicesImagesRequest creates a request to invoke ListDevicesImages API
func CreateListDevicesImagesRequest() (request *ListDevicesImagesRequest) {
	request = &ListDevicesImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "ListDevicesImages", "1.1.8.2", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDevicesImagesResponse creates a response to parse from ListDevicesImages response
func CreateListDevicesImagesResponse() (response *ListDevicesImagesResponse) {
	response = &ListDevicesImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
