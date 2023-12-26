package gpdb

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

// DeleteCollectionData invokes the gpdb.DeleteCollectionData API synchronously
func (client *Client) DeleteCollectionData(request *DeleteCollectionDataRequest) (response *DeleteCollectionDataResponse, err error) {
	response = CreateDeleteCollectionDataResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCollectionDataWithChan invokes the gpdb.DeleteCollectionData API asynchronously
func (client *Client) DeleteCollectionDataWithChan(request *DeleteCollectionDataRequest) (<-chan *DeleteCollectionDataResponse, <-chan error) {
	responseChan := make(chan *DeleteCollectionDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCollectionData(request)
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

// DeleteCollectionDataWithCallback invokes the gpdb.DeleteCollectionData API asynchronously
func (client *Client) DeleteCollectionDataWithCallback(request *DeleteCollectionDataRequest, callback func(response *DeleteCollectionDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCollectionDataResponse
		var err error
		defer close(result)
		response, err = client.DeleteCollectionData(request)
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

// DeleteCollectionDataRequest is the request struct for api DeleteCollectionData
type DeleteCollectionDataRequest struct {
	*requests.RpcRequest
	CollectionDataFilter string           `position:"Query" name:"CollectionDataFilter"`
	CollectionData       string           `position:"Query" name:"CollectionData"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	Collection           string           `position:"Query" name:"Collection"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NamespacePassword    string           `position:"Query" name:"NamespacePassword"`
	Namespace            string           `position:"Query" name:"Namespace"`
}

// DeleteCollectionDataResponse is the response struct for api DeleteCollectionData
type DeleteCollectionDataResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Message     string `json:"Message" xml:"Message"`
	Status      string `json:"Status" xml:"Status"`
	AppliedRows int64  `json:"AppliedRows" xml:"AppliedRows"`
}

// CreateDeleteCollectionDataRequest creates a request to invoke DeleteCollectionData API
func CreateDeleteCollectionDataRequest() (request *DeleteCollectionDataRequest) {
	request = &DeleteCollectionDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DeleteCollectionData", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteCollectionDataResponse creates a response to parse from DeleteCollectionData response
func CreateDeleteCollectionDataResponse() (response *DeleteCollectionDataResponse) {
	response = &DeleteCollectionDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
