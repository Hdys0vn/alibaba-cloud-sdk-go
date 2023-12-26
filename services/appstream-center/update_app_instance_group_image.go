package appstream_center

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

// UpdateAppInstanceGroupImage invokes the appstream_center.UpdateAppInstanceGroupImage API synchronously
func (client *Client) UpdateAppInstanceGroupImage(request *UpdateAppInstanceGroupImageRequest) (response *UpdateAppInstanceGroupImageResponse, err error) {
	response = CreateUpdateAppInstanceGroupImageResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAppInstanceGroupImageWithChan invokes the appstream_center.UpdateAppInstanceGroupImage API asynchronously
func (client *Client) UpdateAppInstanceGroupImageWithChan(request *UpdateAppInstanceGroupImageRequest) (<-chan *UpdateAppInstanceGroupImageResponse, <-chan error) {
	responseChan := make(chan *UpdateAppInstanceGroupImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAppInstanceGroupImage(request)
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

// UpdateAppInstanceGroupImageWithCallback invokes the appstream_center.UpdateAppInstanceGroupImage API asynchronously
func (client *Client) UpdateAppInstanceGroupImageWithCallback(request *UpdateAppInstanceGroupImageRequest, callback func(response *UpdateAppInstanceGroupImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAppInstanceGroupImageResponse
		var err error
		defer close(result)
		response, err = client.UpdateAppInstanceGroupImage(request)
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

// UpdateAppInstanceGroupImageRequest is the request struct for api UpdateAppInstanceGroupImage
type UpdateAppInstanceGroupImageRequest struct {
	*requests.RpcRequest
	BizRegionId        string `position:"Query" name:"BizRegionId"`
	UpdateMode         string `position:"Query" name:"UpdateMode"`
	ProductType        string `position:"Query" name:"ProductType"`
	AppInstanceGroupId string `position:"Query" name:"AppInstanceGroupId"`
	AppCenterImageId   string `position:"Query" name:"AppCenterImageId"`
}

// UpdateAppInstanceGroupImageResponse is the response struct for api UpdateAppInstanceGroupImage
type UpdateAppInstanceGroupImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdateAppInstanceGroupImageRequest creates a request to invoke UpdateAppInstanceGroupImage API
func CreateUpdateAppInstanceGroupImageRequest() (request *UpdateAppInstanceGroupImageRequest) {
	request = &UpdateAppInstanceGroupImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("appstream-center", "2021-09-01", "UpdateAppInstanceGroupImage", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateAppInstanceGroupImageResponse creates a response to parse from UpdateAppInstanceGroupImage response
func CreateUpdateAppInstanceGroupImageResponse() (response *UpdateAppInstanceGroupImageResponse) {
	response = &UpdateAppInstanceGroupImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}