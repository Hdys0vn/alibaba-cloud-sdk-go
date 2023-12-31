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

// UploadCorpIdentifyFile invokes the cloudcallcenter.UploadCorpIdentifyFile API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/uploadcorpidentifyfile.html
func (client *Client) UploadCorpIdentifyFile(request *UploadCorpIdentifyFileRequest) (response *UploadCorpIdentifyFileResponse, err error) {
	response = CreateUploadCorpIdentifyFileResponse()
	err = client.DoAction(request, response)
	return
}

// UploadCorpIdentifyFileWithChan invokes the cloudcallcenter.UploadCorpIdentifyFile API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/uploadcorpidentifyfile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadCorpIdentifyFileWithChan(request *UploadCorpIdentifyFileRequest) (<-chan *UploadCorpIdentifyFileResponse, <-chan error) {
	responseChan := make(chan *UploadCorpIdentifyFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadCorpIdentifyFile(request)
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

// UploadCorpIdentifyFileWithCallback invokes the cloudcallcenter.UploadCorpIdentifyFile API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/uploadcorpidentifyfile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UploadCorpIdentifyFileWithCallback(request *UploadCorpIdentifyFileRequest, callback func(response *UploadCorpIdentifyFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadCorpIdentifyFileResponse
		var err error
		defer close(result)
		response, err = client.UploadCorpIdentifyFile(request)
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

// UploadCorpIdentifyFileRequest is the request struct for api UploadCorpIdentifyFile
type UploadCorpIdentifyFileRequest struct {
	*requests.RpcRequest
	DayuKey string `position:"Query" name:"DayuKey"`
	CccKey  string `position:"Query" name:"CccKey"`
}

// UploadCorpIdentifyFileResponse is the response struct for api UploadCorpIdentifyFile
type UploadCorpIdentifyFileResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateUploadCorpIdentifyFileRequest creates a request to invoke UploadCorpIdentifyFile API
func CreateUploadCorpIdentifyFileRequest() (request *UploadCorpIdentifyFileRequest) {
	request = &UploadCorpIdentifyFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "UploadCorpIdentifyFile", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadCorpIdentifyFileResponse creates a response to parse from UploadCorpIdentifyFile response
func CreateUploadCorpIdentifyFileResponse() (response *UploadCorpIdentifyFileResponse) {
	response = &UploadCorpIdentifyFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
