package cloudwf

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

// SetUpgradeImgByModel invokes the cloudwf.SetUpgradeImgByModel API synchronously
// api document: https://help.aliyun.com/api/cloudwf/setupgradeimgbymodel.html
func (client *Client) SetUpgradeImgByModel(request *SetUpgradeImgByModelRequest) (response *SetUpgradeImgByModelResponse, err error) {
	response = CreateSetUpgradeImgByModelResponse()
	err = client.DoAction(request, response)
	return
}

// SetUpgradeImgByModelWithChan invokes the cloudwf.SetUpgradeImgByModel API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/setupgradeimgbymodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetUpgradeImgByModelWithChan(request *SetUpgradeImgByModelRequest) (<-chan *SetUpgradeImgByModelResponse, <-chan error) {
	responseChan := make(chan *SetUpgradeImgByModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetUpgradeImgByModel(request)
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

// SetUpgradeImgByModelWithCallback invokes the cloudwf.SetUpgradeImgByModel API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/setupgradeimgbymodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetUpgradeImgByModelWithCallback(request *SetUpgradeImgByModelRequest, callback func(response *SetUpgradeImgByModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetUpgradeImgByModelResponse
		var err error
		defer close(result)
		response, err = client.SetUpgradeImgByModel(request)
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

// SetUpgradeImgByModelRequest is the request struct for api SetUpgradeImgByModel
type SetUpgradeImgByModelRequest struct {
	*requests.RpcRequest
	ImgAddr    string           `position:"Query" name:"ImgAddr"`
	ImgVersion string           `position:"Query" name:"ImgVersion"`
	ApModelId  requests.Integer `position:"Query" name:"ApModelId"`
	Comment    string           `position:"Query" name:"Comment"`
}

// SetUpgradeImgByModelResponse is the response struct for api SetUpgradeImgByModel
type SetUpgradeImgByModelResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateSetUpgradeImgByModelRequest creates a request to invoke SetUpgradeImgByModel API
func CreateSetUpgradeImgByModelRequest() (request *SetUpgradeImgByModelRequest) {
	request = &SetUpgradeImgByModelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "SetUpgradeImgByModel", "cloudwf", "openAPI")
	return
}

// CreateSetUpgradeImgByModelResponse creates a response to parse from SetUpgradeImgByModel response
func CreateSetUpgradeImgByModelResponse() (response *SetUpgradeImgByModelResponse) {
	response = &SetUpgradeImgByModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
