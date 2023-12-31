package edas

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

// ImportK8sCluster invokes the edas.ImportK8sCluster API synchronously
func (client *Client) ImportK8sCluster(request *ImportK8sClusterRequest) (response *ImportK8sClusterResponse, err error) {
	response = CreateImportK8sClusterResponse()
	err = client.DoAction(request, response)
	return
}

// ImportK8sClusterWithChan invokes the edas.ImportK8sCluster API asynchronously
func (client *Client) ImportK8sClusterWithChan(request *ImportK8sClusterRequest) (<-chan *ImportK8sClusterResponse, <-chan error) {
	responseChan := make(chan *ImportK8sClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportK8sCluster(request)
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

// ImportK8sClusterWithCallback invokes the edas.ImportK8sCluster API asynchronously
func (client *Client) ImportK8sClusterWithCallback(request *ImportK8sClusterRequest, callback func(response *ImportK8sClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportK8sClusterResponse
		var err error
		defer close(result)
		response, err = client.ImportK8sCluster(request)
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

// ImportK8sClusterRequest is the request struct for api ImportK8sCluster
type ImportK8sClusterRequest struct {
	*requests.RoaRequest
	Mode        requests.Integer `position:"Query" name:"Mode"`
	EnableAsm   requests.Boolean `position:"Query" name:"EnableAsm"`
	NamespaceId string           `position:"Query" name:"NamespaceId"`
	ClusterId   string           `position:"Query" name:"ClusterId"`
}

// ImportK8sClusterResponse is the response struct for api ImportK8sCluster
type ImportK8sClusterResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateImportK8sClusterRequest creates a request to invoke ImportK8sCluster API
func CreateImportK8sClusterRequest() (request *ImportK8sClusterRequest) {
	request = &ImportK8sClusterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ImportK8sCluster", "/pop/v5/import_k8s_cluster", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateImportK8sClusterResponse creates a response to parse from ImportK8sCluster response
func CreateImportK8sClusterResponse() (response *ImportK8sClusterResponse) {
	response = &ImportK8sClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
