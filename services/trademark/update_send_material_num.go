package trademark

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

// UpdateSendMaterialNum invokes the trademark.UpdateSendMaterialNum API synchronously
// api document: https://help.aliyun.com/api/trademark/updatesendmaterialnum.html
func (client *Client) UpdateSendMaterialNum(request *UpdateSendMaterialNumRequest) (response *UpdateSendMaterialNumResponse, err error) {
	response = CreateUpdateSendMaterialNumResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSendMaterialNumWithChan invokes the trademark.UpdateSendMaterialNum API asynchronously
// api document: https://help.aliyun.com/api/trademark/updatesendmaterialnum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSendMaterialNumWithChan(request *UpdateSendMaterialNumRequest) (<-chan *UpdateSendMaterialNumResponse, <-chan error) {
	responseChan := make(chan *UpdateSendMaterialNumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSendMaterialNum(request)
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

// UpdateSendMaterialNumWithCallback invokes the trademark.UpdateSendMaterialNum API asynchronously
// api document: https://help.aliyun.com/api/trademark/updatesendmaterialnum.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSendMaterialNumWithCallback(request *UpdateSendMaterialNumRequest, callback func(response *UpdateSendMaterialNumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSendMaterialNumResponse
		var err error
		defer close(result)
		response, err = client.UpdateSendMaterialNum(request)
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

// UpdateSendMaterialNumRequest is the request struct for api UpdateSendMaterialNum
type UpdateSendMaterialNumRequest struct {
	*requests.RpcRequest
	Num         string           `position:"Query" name:"Num"`
	BizId       string           `position:"Query" name:"BizId"`
	OperateType requests.Integer `position:"Query" name:"OperateType"`
}

// UpdateSendMaterialNumResponse is the response struct for api UpdateSendMaterialNum
type UpdateSendMaterialNumResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateUpdateSendMaterialNumRequest creates a request to invoke UpdateSendMaterialNum API
func CreateUpdateSendMaterialNumRequest() (request *UpdateSendMaterialNumRequest) {
	request = &UpdateSendMaterialNumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "UpdateSendMaterialNum", "trademark", "openAPI")
	return
}

// CreateUpdateSendMaterialNumResponse creates a response to parse from UpdateSendMaterialNum response
func CreateUpdateSendMaterialNumResponse() (response *UpdateSendMaterialNumResponse) {
	response = &UpdateSendMaterialNumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
