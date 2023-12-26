package faas

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

// DeleteFpgaImage invokes the faas.DeleteFpgaImage API synchronously
// api document: https://help.aliyun.com/api/faas/deletefpgaimage.html
func (client *Client) DeleteFpgaImage(request *DeleteFpgaImageRequest) (response *DeleteFpgaImageResponse, err error) {
	response = CreateDeleteFpgaImageResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFpgaImageWithChan invokes the faas.DeleteFpgaImage API asynchronously
// api document: https://help.aliyun.com/api/faas/deletefpgaimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFpgaImageWithChan(request *DeleteFpgaImageRequest) (<-chan *DeleteFpgaImageResponse, <-chan error) {
	responseChan := make(chan *DeleteFpgaImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFpgaImage(request)
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

// DeleteFpgaImageWithCallback invokes the faas.DeleteFpgaImage API asynchronously
// api document: https://help.aliyun.com/api/faas/deletefpgaimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFpgaImageWithCallback(request *DeleteFpgaImageRequest, callback func(response *DeleteFpgaImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFpgaImageResponse
		var err error
		defer close(result)
		response, err = client.DeleteFpgaImage(request)
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

// DeleteFpgaImageRequest is the request struct for api DeleteFpgaImage
type DeleteFpgaImageRequest struct {
	*requests.RpcRequest
	FpgaImageUniqueId string `position:"Query" name:"FpgaImageUniqueId"`
	ECSImageId        string `position:"Query" name:"ECSImageId"`
	OwnerAlias        string `position:"Query" name:"OwnerAlias"`
}

// DeleteFpgaImageResponse is the response struct for api DeleteFpgaImage
type DeleteFpgaImageResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	FpgaImageUniqueId string `json:"FpgaImageUniqueId" xml:"FpgaImageUniqueId"`
}

// CreateDeleteFpgaImageRequest creates a request to invoke DeleteFpgaImage API
func CreateDeleteFpgaImageRequest() (request *DeleteFpgaImageRequest) {
	request = &DeleteFpgaImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2020-02-17", "DeleteFpgaImage", "faas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFpgaImageResponse creates a response to parse from DeleteFpgaImage response
func CreateDeleteFpgaImageResponse() (response *DeleteFpgaImageResponse) {
	response = &DeleteFpgaImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
