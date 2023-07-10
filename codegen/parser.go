package codegen

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yoheimuta/go-protoparser/v4"
)

type AutoGenerated struct {
	Syntax struct {
		ProtobufVersion      string `json:"ProtobufVersion"`
		ProtobufVersionQuote string `json:"ProtobufVersionQuote"`
		Comments             any    `json:"Comments"`
		InlineComment        any    `json:"InlineComment"`
		Meta                 struct {
			Pos struct {
				Filename string `json:"Filename"`
				Offset   int    `json:"Offset"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			} `json:"Pos"`
			LastPos struct {
				Filename string `json:"Filename"`
				Offset   int    `json:"Offset"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			} `json:"LastPos"`
		} `json:"Meta"`
	} `json:"Syntax"`
	ProtoBody []struct {
		Name          string `json:"Name,omitempty"`
		Comments      any    `json:"Comments"`
		InlineComment any    `json:"InlineComment"`
		Meta          struct {
			Pos struct {
				Filename string `json:"Filename"`
				Offset   int    `json:"Offset"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			} `json:"Pos"`
			LastPos struct {
				Filename string `json:"Filename"`
				Offset   int    `json:"Offset"`
				Line     int    `json:"Line"`
				Column   int    `json:"Column"`
			} `json:"LastPos"`
		} `json:"Meta"`
		ServiceName string `json:"ServiceName,omitempty"`
		ServiceBody []struct {
			RPCName    string `json:"RPCName"`
			RPCRequest struct {
				IsStream    bool   `json:"IsStream"`
				MessageType string `json:"MessageType"`
				Meta        struct {
					Pos struct {
						Filename string `json:"Filename"`
						Offset   int    `json:"Offset"`
						Line     int    `json:"Line"`
						Column   int    `json:"Column"`
					} `json:"Pos"`
					LastPos struct {
						Filename string `json:"Filename"`
						Offset   int    `json:"Offset"`
						Line     int    `json:"Line"`
						Column   int    `json:"Column"`
					} `json:"LastPos"`
				} `json:"Meta"`
			} `json:"RPCRequest"`
			RPCResponse struct {
				IsStream    bool   `json:"IsStream"`
				MessageType string `json:"MessageType"`
				Meta        struct {
					Pos struct {
						Filename string `json:"Filename"`
						Offset   int    `json:"Offset"`
						Line     int    `json:"Line"`
						Column   int    `json:"Column"`
					} `json:"Pos"`
					LastPos struct {
						Filename string `json:"Filename"`
						Offset   int    `json:"Offset"`
						Line     int    `json:"Line"`
						Column   int    `json:"Column"`
					} `json:"LastPos"`
				} `json:"Meta"`
			} `json:"RPCResponse"`
			Options                      any `json:"Options"`
			Comments                     any `json:"Comments"`
			InlineComment                any `json:"InlineComment"`
			InlineCommentBehindLeftCurly any `json:"InlineCommentBehindLeftCurly"`
			Meta                         struct {
				Pos struct {
					Filename string `json:"Filename"`
					Offset   int    `json:"Offset"`
					Line     int    `json:"Line"`
					Column   int    `json:"Column"`
				} `json:"Pos"`
				LastPos struct {
					Filename string `json:"Filename"`
					Offset   int    `json:"Offset"`
					Line     int    `json:"Line"`
					Column   int    `json:"Column"`
				} `json:"LastPos"`
			} `json:"Meta"`
		} `json:"ServiceBody,omitempty"`
		InlineCommentBehindLeftCurly any    `json:"InlineCommentBehindLeftCurly,omitempty"`
		MessageName                  string `json:"MessageName,omitempty"`
		MessageBody                  []struct {
			IsRepeated    bool   `json:"IsRepeated"`
			IsRequired    bool   `json:"IsRequired"`
			IsOptional    bool   `json:"IsOptional"`
			Type          string `json:"Type"`
			FieldName     string `json:"FieldName"`
			FieldNumber   string `json:"FieldNumber"`
			FieldOptions  any    `json:"FieldOptions"`
			Comments      any    `json:"Comments"`
			InlineComment any    `json:"InlineComment"`
			Meta          struct {
				Pos struct {
					Filename string `json:"Filename"`
					Offset   int    `json:"Offset"`
					Line     int    `json:"Line"`
					Column   int    `json:"Column"`
				} `json:"Pos"`
				LastPos struct {
					Filename string `json:"Filename"`
					Offset   int    `json:"Offset"`
					Line     int    `json:"Line"`
					Column   int    `json:"Column"`
				} `json:"LastPos"`
			} `json:"Meta"`
		} `json:"MessageBody,omitempty"`
	} `json:"ProtoBody"`
	Meta struct {
		Filename string `json:"Filename"`
	} `json:"Meta"`
}

func Parse(path string) (Service, error) {
	reader, err := os.Open(path)
	if err != nil {
		return Service{}, fmt.Errorf("error: Unable to open the file: %v", err)
	}
	defer func() error {
		if err := reader.Close(); err != nil {
			return fmt.Errorf("error: Unable to close the file: %v", err)
		}
		return nil
	}()

	got, err := protoparser.Parse(
		reader,
		protoparser.WithFilename(filepath.Base(path)),
	)
	if err != nil {
		return Service{}, fmt.Errorf("error: Unable to parse the file: %v", err)
	}

	gotJSON, err := json.MarshalIndent(got, "", "  ")
	if err != nil {
		return Service{}, fmt.Errorf("error: Unable to marshal the file: %v", err)
	}

	//fmt.Println(string(gotJSON))

	//unmarshal the json into a struct
	var testProto AutoGenerated
	err = json.Unmarshal(gotJSON, &testProto)
	if err != nil {
		return Service{}, fmt.Errorf("error: Unable to unmarshal the file: %v", err)
	}

	//Now we can use the testProto struct to extract only the data we need
	var service Service
	service.ServiceName = testProto.ProtoBody[1].ServiceName
	service.FileName = strings.Split(testProto.Meta.Filename, ".")[0]

	// We need to loop over the field servicebody to determine the amount of RPCs
	for _, v := range testProto.ProtoBody[1].ServiceBody {
		service.Rpc = append(service.Rpc, Rpc{
			RpcName:         v.RPCName,
			Request:         v.RPCRequest.MessageType,
			Response:        v.RPCResponse.MessageType,
			RequestMessage:  []Message{},
			ResponseMessage: []Message{},
		})
	}

	//Loop over protoBody to map the request and response messages to the RPCs
	for _, v := range testProto.ProtoBody {
		for i, rpc := range service.Rpc {
			if rpc.Request == v.MessageName {
				for _, message := range v.MessageBody {
					service.Rpc[i].RequestMessage = append(service.Rpc[i].RequestMessage, Message{
						IsRepeated: message.IsRepeated,
						FieldName:  message.FieldName,
						Type:       message.Type,
					})
				}
			}
			if rpc.Response == v.MessageName {
				for _, message := range v.MessageBody {
					service.Rpc[i].ResponseMessage = append(service.Rpc[i].ResponseMessage, Message{
						IsRepeated: message.IsRepeated,
						FieldName:  message.FieldName,
						Type:       message.Type,
					})
				}
			}
		}
	}

	//marshal the service struct to json
	serviceJSON, err := json.MarshalIndent(service, "", "  ")
	if err != nil {
		return Service{}, fmt.Errorf("error: Unable to marshal the file: %v", err)
	}

	//print the json
	fmt.Println(string(serviceJSON))

	return service, nil
}

type Service struct {
	ServiceName string `json:"ServiceName"`
	FileName    string `json:"FileName"`
	Rpc         []Rpc  `json:"Rpc"`
}

type Rpc struct {
	RpcName         string    `json:"RpcName"`
	RequestMessage  []Message `json:"RequestMessage"`
	ResponseMessage []Message `json:"ResponseMessage"`
	Request         string    `json:"Request"`
	Response        string    `json:"Response"`
}

type Message struct {
	IsRepeated bool   `json:"IsRepeated"`
	FieldName  string `json:"FieldName"`
	Type       string `json:"Type"`
}
