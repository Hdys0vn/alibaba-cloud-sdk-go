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

// Get400CoppdfOSSUrl invokes the cloudcallcenter.Get400CoppdfOSSUrl API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/get400coppdfossurl.html
func (client *Client) Get400CoppdfOSSUrl(request *Get400CoppdfOSSUrlRequest) (response *Get400CoppdfOSSUrlResponse, err error) {
	response = CreateGet400CoppdfOSSUrlResponse()
	err = client.DoAction(request, response)
	return
}

// Get400CoppdfOSSUrlWithChan invokes the cloudcallcenter.Get400CoppdfOSSUrl API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/get400coppdfossurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) Get400CoppdfOSSUrlWithChan(request *Get400CoppdfOSSUrlRequest) (<-chan *Get400CoppdfOSSUrlResponse, <-chan error) {
	responseChan := make(chan *Get400CoppdfOSSUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Get400CoppdfOSSUrl(request)
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

// Get400CoppdfOSSUrlWithCallback invokes the cloudcallcenter.Get400CoppdfOSSUrl API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/get400coppdfossurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) Get400CoppdfOSSUrlWithCallback(request *Get400CoppdfOSSUrlRequest, callback func(response *Get400CoppdfOSSUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *Get400CoppdfOSSUrlResponse
		var err error
		defer close(result)
		response, err = client.Get400CoppdfOSSUrl(request)
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

// Get400CoppdfOSSUrlRequest is the request struct for api Get400CoppdfOSSUrl
type Get400CoppdfOSSUrlRequest struct {
	*requests.RpcRequest
	ManagerCertExpDate string `position:"Query" name:"ManagerCertExpDate"`
	ManagerMobilePhone string `position:"Query" name:"ManagerMobilePhone"`
	Number             string `position:"Query" name:"Number"`
	ManagerCertEffDate string `position:"Query" name:"ManagerCertEffDate"`
	ManagerName        string `position:"Query" name:"ManagerName"`
	ManagerCertPic     string `position:"Query" name:"ManagerCertPic"`
	ManagerCertNumber  string `position:"Query" name:"ManagerCertNumber"`
	ManagerCertType    string `position:"Query" name:"ManagerCertType"`
}

// Get400CoppdfOSSUrlResponse is the response struct for api Get400CoppdfOSSUrl
type Get400CoppdfOSSUrlResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           string `json:"Data" xml:"Data"`
}

// CreateGet400CoppdfOSSUrlRequest creates a request to invoke Get400CoppdfOSSUrl API
func CreateGet400CoppdfOSSUrlRequest() (request *Get400CoppdfOSSUrlRequest) {
	request = &Get400CoppdfOSSUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "Get400CoppdfOSSUrl", "", "")
	request.Method = requests.POST
	return
}

// CreateGet400CoppdfOSSUrlResponse creates a response to parse from Get400CoppdfOSSUrl response
func CreateGet400CoppdfOSSUrlResponse() (response *Get400CoppdfOSSUrlResponse) {
	response = &Get400CoppdfOSSUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
