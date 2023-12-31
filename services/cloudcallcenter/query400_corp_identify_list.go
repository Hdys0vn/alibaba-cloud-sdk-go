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

// Query400CorpIdentifyList invokes the cloudcallcenter.Query400CorpIdentifyList API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/query400corpidentifylist.html
func (client *Client) Query400CorpIdentifyList(request *Query400CorpIdentifyListRequest) (response *Query400CorpIdentifyListResponse, err error) {
	response = CreateQuery400CorpIdentifyListResponse()
	err = client.DoAction(request, response)
	return
}

// Query400CorpIdentifyListWithChan invokes the cloudcallcenter.Query400CorpIdentifyList API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/query400corpidentifylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) Query400CorpIdentifyListWithChan(request *Query400CorpIdentifyListRequest) (<-chan *Query400CorpIdentifyListResponse, <-chan error) {
	responseChan := make(chan *Query400CorpIdentifyListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Query400CorpIdentifyList(request)
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

// Query400CorpIdentifyListWithCallback invokes the cloudcallcenter.Query400CorpIdentifyList API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/query400corpidentifylist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) Query400CorpIdentifyListWithCallback(request *Query400CorpIdentifyListRequest, callback func(response *Query400CorpIdentifyListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *Query400CorpIdentifyListResponse
		var err error
		defer close(result)
		response, err = client.Query400CorpIdentifyList(request)
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

// Query400CorpIdentifyListRequest is the request struct for api Query400CorpIdentifyList
type Query400CorpIdentifyListRequest struct {
	*requests.RpcRequest
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Status      string           `position:"Query" name:"Status"`
}

// Query400CorpIdentifyListResponse is the response struct for api Query400CorpIdentifyList
type Query400CorpIdentifyListResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateQuery400CorpIdentifyListRequest creates a request to invoke Query400CorpIdentifyList API
func CreateQuery400CorpIdentifyListRequest() (request *Query400CorpIdentifyListRequest) {
	request = &Query400CorpIdentifyListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "Query400CorpIdentifyList", "", "")
	request.Method = requests.POST
	return
}

// CreateQuery400CorpIdentifyListResponse creates a response to parse from Query400CorpIdentifyList response
func CreateQuery400CorpIdentifyListResponse() (response *Query400CorpIdentifyListResponse) {
	response = &Query400CorpIdentifyListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
