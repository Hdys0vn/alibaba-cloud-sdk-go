package crm

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

// BatchGetAliyunIdByAliyunPk invokes the crm.BatchGetAliyunIdByAliyunPk API synchronously
// api document: https://help.aliyun.com/api/crm/batchgetaliyunidbyaliyunpk.html
func (client *Client) BatchGetAliyunIdByAliyunPk(request *BatchGetAliyunIdByAliyunPkRequest) (response *BatchGetAliyunIdByAliyunPkResponse, err error) {
	response = CreateBatchGetAliyunIdByAliyunPkResponse()
	err = client.DoAction(request, response)
	return
}

// BatchGetAliyunIdByAliyunPkWithChan invokes the crm.BatchGetAliyunIdByAliyunPk API asynchronously
// api document: https://help.aliyun.com/api/crm/batchgetaliyunidbyaliyunpk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetAliyunIdByAliyunPkWithChan(request *BatchGetAliyunIdByAliyunPkRequest) (<-chan *BatchGetAliyunIdByAliyunPkResponse, <-chan error) {
	responseChan := make(chan *BatchGetAliyunIdByAliyunPkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchGetAliyunIdByAliyunPk(request)
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

// BatchGetAliyunIdByAliyunPkWithCallback invokes the crm.BatchGetAliyunIdByAliyunPk API asynchronously
// api document: https://help.aliyun.com/api/crm/batchgetaliyunidbyaliyunpk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetAliyunIdByAliyunPkWithCallback(request *BatchGetAliyunIdByAliyunPkRequest, callback func(response *BatchGetAliyunIdByAliyunPkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchGetAliyunIdByAliyunPkResponse
		var err error
		defer close(result)
		response, err = client.BatchGetAliyunIdByAliyunPk(request)
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

// BatchGetAliyunIdByAliyunPkRequest is the request struct for api BatchGetAliyunIdByAliyunPk
type BatchGetAliyunIdByAliyunPkRequest struct {
	*requests.RpcRequest
	PkList *[]string `position:"Query" name:"PkList"  type:"Repeated"`
}

// BatchGetAliyunIdByAliyunPkResponse is the response struct for api BatchGetAliyunIdByAliyunPk
type BatchGetAliyunIdByAliyunPkResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	ProfileList ProfileList `json:"ProfileList" xml:"ProfileList"`
}

// CreateBatchGetAliyunIdByAliyunPkRequest creates a request to invoke BatchGetAliyunIdByAliyunPk API
func CreateBatchGetAliyunIdByAliyunPkRequest() (request *BatchGetAliyunIdByAliyunPkRequest) {
	request = &BatchGetAliyunIdByAliyunPkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Crm", "2015-04-08", "BatchGetAliyunIdByAliyunPk", "crm", "openAPI")
	return
}

// CreateBatchGetAliyunIdByAliyunPkResponse creates a response to parse from BatchGetAliyunIdByAliyunPk response
func CreateBatchGetAliyunIdByAliyunPkResponse() (response *BatchGetAliyunIdByAliyunPkResponse) {
	response = &BatchGetAliyunIdByAliyunPkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
