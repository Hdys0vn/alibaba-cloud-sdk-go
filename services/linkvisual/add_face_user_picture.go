package linkvisual

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

// AddFaceUserPicture invokes the linkvisual.AddFaceUserPicture API synchronously
func (client *Client) AddFaceUserPicture(request *AddFaceUserPictureRequest) (response *AddFaceUserPictureResponse, err error) {
	response = CreateAddFaceUserPictureResponse()
	err = client.DoAction(request, response)
	return
}

// AddFaceUserPictureWithChan invokes the linkvisual.AddFaceUserPicture API asynchronously
func (client *Client) AddFaceUserPictureWithChan(request *AddFaceUserPictureRequest) (<-chan *AddFaceUserPictureResponse, <-chan error) {
	responseChan := make(chan *AddFaceUserPictureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddFaceUserPicture(request)
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

// AddFaceUserPictureWithCallback invokes the linkvisual.AddFaceUserPicture API asynchronously
func (client *Client) AddFaceUserPictureWithCallback(request *AddFaceUserPictureRequest, callback func(response *AddFaceUserPictureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddFaceUserPictureResponse
		var err error
		defer close(result)
		response, err = client.AddFaceUserPicture(request)
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

// AddFaceUserPictureRequest is the request struct for api AddFaceUserPicture
type AddFaceUserPictureRequest struct {
	*requests.RpcRequest
	IsolationId string `position:"Query" name:"IsolationId"`
	UserId      string `position:"Query" name:"UserId"`
	FacePicUrl  string `position:"Query" name:"FacePicUrl"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// AddFaceUserPictureResponse is the response struct for api AddFaceUserPicture
type AddFaceUserPictureResponse struct {
	*responses.BaseResponse
	Code         string                 `json:"Code" xml:"Code"`
	Data         map[string]interface{} `json:"Data" xml:"Data"`
	ErrorMessage string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string                 `json:"RequestId" xml:"RequestId"`
	Success      bool                   `json:"Success" xml:"Success"`
}

// CreateAddFaceUserPictureRequest creates a request to invoke AddFaceUserPicture API
func CreateAddFaceUserPictureRequest() (request *AddFaceUserPictureRequest) {
	request = &AddFaceUserPictureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "AddFaceUserPicture", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddFaceUserPictureResponse creates a response to parse from AddFaceUserPicture response
func CreateAddFaceUserPictureResponse() (response *AddFaceUserPictureResponse) {
	response = &AddFaceUserPictureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}