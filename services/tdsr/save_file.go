package tdsr

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

// SaveFile invokes the tdsr.SaveFile API synchronously
// api document: https://help.aliyun.com/api/tdsr/savefile.html
func (client *Client) SaveFile(request *SaveFileRequest) (response *SaveFileResponse, err error) {
	response = CreateSaveFileResponse()
	err = client.DoAction(request, response)
	return
}

// SaveFileWithChan invokes the tdsr.SaveFile API asynchronously
// api document: https://help.aliyun.com/api/tdsr/savefile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveFileWithChan(request *SaveFileRequest) (<-chan *SaveFileResponse, <-chan error) {
	responseChan := make(chan *SaveFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveFile(request)
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

// SaveFileWithCallback invokes the tdsr.SaveFile API asynchronously
// api document: https://help.aliyun.com/api/tdsr/savefile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveFileWithCallback(request *SaveFileRequest, callback func(response *SaveFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveFileResponse
		var err error
		defer close(result)
		response, err = client.SaveFile(request)
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

// SaveFileRequest is the request struct for api SaveFile
type SaveFileRequest struct {
	*requests.RpcRequest
	SubSceneUuid string `position:"Query" name:"SubSceneUuid"`
	ParamFile    string `position:"Query" name:"ParamFile"`
}

// SaveFileResponse is the response struct for api SaveFile
type SaveFileResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrMessage   string `json:"ErrMessage" xml:"ErrMessage"`
	Data         string `json:"Data" xml:"Data"`
	ObjectString string `json:"ObjectString" xml:"ObjectString"`
}

// CreateSaveFileRequest creates a request to invoke SaveFile API
func CreateSaveFileRequest() (request *SaveFileRequest) {
	request = &SaveFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("tdsr", "2020-01-01", "SaveFile", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveFileResponse creates a response to parse from SaveFile response
func CreateSaveFileResponse() (response *SaveFileResponse) {
	response = &SaveFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}