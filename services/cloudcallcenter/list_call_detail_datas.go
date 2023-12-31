package cloudcallcenter

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

// ListCallDetailDatas invokes the cloudcallcenter.ListCallDetailDatas API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listcalldetaildatas.html
func (client *Client) ListCallDetailDatas(request *ListCallDetailDatasRequest) (response *ListCallDetailDatasResponse, err error) {
	response = CreateListCallDetailDatasResponse()
	err = client.DoAction(request, response)
	return
}

// ListCallDetailDatasWithChan invokes the cloudcallcenter.ListCallDetailDatas API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listcalldetaildatas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCallDetailDatasWithChan(request *ListCallDetailDatasRequest) (<-chan *ListCallDetailDatasResponse, <-chan error) {
	responseChan := make(chan *ListCallDetailDatasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCallDetailDatas(request)
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

// ListCallDetailDatasWithCallback invokes the cloudcallcenter.ListCallDetailDatas API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listcalldetaildatas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCallDetailDatasWithCallback(request *ListCallDetailDatasRequest, callback func(response *ListCallDetailDatasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCallDetailDatasResponse
		var err error
		defer close(result)
		response, err = client.ListCallDetailDatas(request)
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

// ListCallDetailDatasRequest is the request struct for api ListCallDetailDatas
type ListCallDetailDatasRequest struct {
	*requests.RpcRequest
	InstanceId    string           `position:"Query" name:"InstanceId"`
	Criteria      string           `position:"Query" name:"Criteria"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	StopTime      requests.Integer `position:"Query" name:"StopTime"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	WithRecording requests.Boolean `position:"Query" name:"WithRecording"`
}

// ListCallDetailDatasResponse is the response struct for api ListCallDetailDatas
type ListCallDetailDatasResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Success           bool              `json:"Success" xml:"Success"`
	Code              string            `json:"Code" xml:"Code"`
	Message           string            `json:"Message" xml:"Message"`
	HttpStatusCode    int               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	CallDetailRecords CallDetailRecords `json:"CallDetailRecords" xml:"CallDetailRecords"`
}

// CreateListCallDetailDatasRequest creates a request to invoke ListCallDetailDatas API
func CreateListCallDetailDatasRequest() (request *ListCallDetailDatasRequest) {
	request = &ListCallDetailDatasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListCallDetailDatas", "", "")
	request.Method = requests.POST
	return
}

// CreateListCallDetailDatasResponse creates a response to parse from ListCallDetailDatas response
func CreateListCallDetailDatasResponse() (response *ListCallDetailDatasResponse) {
	response = &ListCallDetailDatasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
