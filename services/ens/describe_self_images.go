package ens

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

// DescribeSelfImages invokes the ens.DescribeSelfImages API synchronously
func (client *Client) DescribeSelfImages(request *DescribeSelfImagesRequest) (response *DescribeSelfImagesResponse, err error) {
	response = CreateDescribeSelfImagesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSelfImagesWithChan invokes the ens.DescribeSelfImages API asynchronously
func (client *Client) DescribeSelfImagesWithChan(request *DescribeSelfImagesRequest) (<-chan *DescribeSelfImagesResponse, <-chan error) {
	responseChan := make(chan *DescribeSelfImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSelfImages(request)
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

// DescribeSelfImagesWithCallback invokes the ens.DescribeSelfImages API asynchronously
func (client *Client) DescribeSelfImagesWithCallback(request *DescribeSelfImagesRequest, callback func(response *DescribeSelfImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSelfImagesResponse
		var err error
		defer close(result)
		response, err = client.DescribeSelfImages(request)
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

// DescribeSelfImagesRequest is the request struct for api DescribeSelfImages
type DescribeSelfImagesRequest struct {
	*requests.RpcRequest
	ImageId    string           `position:"Query" name:"ImageId"`
	SnapshotId string           `position:"Query" name:"SnapshotId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	ImageName  string           `position:"Query" name:"ImageName"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeSelfImagesResponse is the response struct for api DescribeSelfImages
type DescribeSelfImagesResponse struct {
	*responses.BaseResponse
	Code      int                        `json:"Code" xml:"Code"`
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Images    ImagesInDescribeSelfImages `json:"Images" xml:"Images"`
}

// CreateDescribeSelfImagesRequest creates a request to invoke DescribeSelfImages API
func CreateDescribeSelfImagesRequest() (request *DescribeSelfImagesRequest) {
	request = &DescribeSelfImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeSelfImages", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSelfImagesResponse creates a response to parse from DescribeSelfImages response
func CreateDescribeSelfImagesResponse() (response *DescribeSelfImagesResponse) {
	response = &DescribeSelfImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
