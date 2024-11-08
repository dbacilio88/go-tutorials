# gRPC

<!-- TOC -->

* [gRPC](#grpc)
* [Overview](#overview)
* [grpc-producer-app](#grpc-producer-app)
* [Proto Help](#proto-help)
    * [Generate Proto](#generate-proto)

<!-- TOC -->

# Overview

**HTTP 1.1**

- Transfiere texto plano, aumenta tamaño
- Se tiene que romper la

**HTTP 2**

**Stream**

**Frame**

**Message**

# Patrones de comunicación

- Unary
- Server Streaming
- Client Streaming
- Bidirectional Streaming (EOF)

# Comparaciones

- gRPC
- REST
- GraphQL

# grpc-producer-app

grpc-producer-app

# Proto Help

```shell
protoc
```

```textmate
Usage: C:\proto3\bin\protoc.exe [OPTION] PROTO_FILES
Parse PROTO_FILES and generate output based on the options given:
  -IPATH, --proto_path=PATH   Specify the directory in which to search for
                              imports.  May be specified multiple times;
                              directories will be searched in order.  If not
                              given, the current working directory is used.
                              If not found in any of the these directories,
                              the --descriptor_set_in descriptors will be
                              checked for required proto file.
  --version                   Show version info and exit.
  -h, --help                  Show this text and exit.
  --encode=MESSAGE_TYPE       Read a text-format message of the given type
                              from standard input and write it in binary
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --deterministic_output      When using --encode, ensure map fields are
                              deterministically ordered. Note that this order
                              is not canonical, and changes across builds or
                              releases of protoc.
  --decode=MESSAGE_TYPE       Read a binary message of the given type from
                              standard input and write it in text format
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --decode_raw                Read an arbitrary protocol message from
                              standard input and write the raw tag/value
                              pairs in text format to standard output.  No
                              PROTO_FILES should be given when using this
                              flag.
  --descriptor_set_in=FILES   Specifies a delimited list of FILES
                              each containing a FileDescriptorSet (a
                              protocol buffer defined in descriptor.proto).
                              The FileDescriptor for each of the PROTO_FILES
                              provided will be loaded from these
                              FileDescriptorSets. If a FileDescriptor
                              appears multiple times, the first occurrence
                              will be used.
  -oFILE,                     Writes a FileDescriptorSet (a protocol buffer,
    --descriptor_set_out=FILE defined in descriptor.proto) containing all of
                              the input files to FILE.
  --include_imports           When using --descriptor_set_out, also include
                              all dependencies of the input files in the
                              set, so that the set is self-contained.
  --include_source_info       When using --descriptor_set_out, do not strip
                              SourceCodeInfo from the FileDescriptorProto.
                              This results in vastly larger descriptors that
                              include information about the original
                              location of each decl in the source file as
                              well as surrounding comments.
  --retain_options            When using --descriptor_set_out, do not strip
                              any options from the FileDescriptorProto.
                              This results in potentially larger descriptors
                              that include information about options that were
                              only meant to be useful during compilation.
  --dependency_out=FILE       Write a dependency output file in the format
                              expected by make. This writes the transitive
                              set of input file paths to FILE
  --error_format=FORMAT       Set the format in which to print errors.
                              FORMAT may be 'gcc' (the default) or 'msvs'
                              (Microsoft Visual Studio format).
  --fatal_warnings            Make warnings be fatal (similar to -Werr in
                              gcc). This flag will make protoc return
                              with a non-zero exit code if any warnings
                              are generated.
  --print_free_field_numbers  Print the free field numbers of the messages
                              defined in the given proto files. Extension ranges
                              are counted as occupied fields numbers.
  --enable_codegen_trace      Enables tracing which parts of protoc are
                              responsible for what codegen output. Not supported
                              by all backends or on all platforms.
  --plugin=EXECUTABLE         Specifies a plugin executable to use.
                              Normally, protoc searches the PATH for
                              plugins, but you may specify additional
                              executables not in the path using this flag.
                              Additionally, EXECUTABLE may be of the form
                              NAME=PATH, in which case the given plugin name
                              is mapped to the given executable even if
                              the executable's own name differs.
  --cpp_out=OUT_DIR           Generate C++ header and source.
  --csharp_out=OUT_DIR        Generate C# source file.
  --java_out=OUT_DIR          Generate Java source file.
  --kotlin_out=OUT_DIR        Generate Kotlin file.
  --objc_out=OUT_DIR          Generate Objective-C header and source.
  --php_out=OUT_DIR           Generate PHP source file.
  --pyi_out=OUT_DIR           Generate python pyi stub.
  --python_out=OUT_DIR        Generate Python source file.
  --ruby_out=OUT_DIR          Generate Ruby source file.
  --rust_out=OUT_DIR          Generate Rust sources.
  @<filename>                 Read options and filenames from file. If a
                              relative file path is specified, the file
                              will be searched in the working directory.
                              The --proto_path option will not affect how
                              this argument file is searched. Content of
                              the file will be expanded in the position of
                              @<filename> as in the argument list. Note
                              that shell expansion is not applied to the
                              content of the file (i.e., you cannot use
                              quotes, wildcards, escapes, commands, etc.).
                              Each line corresponds to a single argument,
                              even if it contains spaces.
```

## Generate Proto

Code python Example

```shell
protoc -I=proto --python_out=python proto/name.proto
```

Code Go Example

Prerequisites

[grpc/grpc-go](https://github.com/grpc/grpc-go)

[go/generate](https://protobuf.dev/reference/go/go-generated/)

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get -u google.golang.org/grpc
``` 

Deprecated:

```shell
go get -u github.com/golang/protobuf/protoc-gen-go
```

Use:

```shell
go get -u google.golang.org/protobuf
```

Example: Unary

```shell
 protoc --go_out=. --go-grpc_out=. .\proto\hello.proto
 protoc --go_out=. --go-grpc_out=. proto/product/*.proto
```

# gRPC Error Handling

https://grpc.io/docs/guides/error/

# gRPC Error Core

https://grpc.github.io/grpc/core/md_doc_statuscodes.html

# SSL Commands

SERVER_CN=localhost

Generar el certificado de autoridad, Generate Certificate Authority

```shell
openssl genrsa --passout pass:1111 -des3 -out ca.key 4069
```

Genera del certificado de confianza: Trust Certificate (ca.crt)

```shell
openssl req --passin pass:1111 -new -x509 -days 100 -key ca.key -out ca.crt -subj "//CN=${SERVER_CN}"
```

Genera la llave privada del servidor

```shell
openssl genrsa -passout pass:1111 -des3 -out server.key 4096
```

Genera certificado para firmar peticiones

```shell
openssl req --passin pass:1111 -new -key server.key -out server.csr -subj "//CN=${SERVER_CN}"
```

Firma el certificado con el CA

```shell
openssl x509 -req -passin pass:1111 -days 100 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt
```

Convertir el certificado a un formato pem que podemos usar en grpc

```shell
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem
```

```
package main

import (
	"context"
	"fmt"
	"github.com/dbacilio88/grpc-producer-app/proto/hello"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"net"
	"time"
)

/**
*
* main
* <p>
* main file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author cbaciliod
* @author dbacilio88@outlook.es
* @since 16/07/2024
*
 */

type servers struct {
	hello.UnimplementedHelloServiceServer
}

func (s *servers) Hello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Printf("call hello from client %v", req)
	request := req.GetHello()
	prefix := request.GetPrefix()
	firstName := request.GetFirstName()
	message := "Hello: " + prefix + ", " + firstName + " welcome. "
	response := &hello.HelloResponse{
		CustomHello: message,
	}

	return response, nil
}

func (s *servers) HelloManyLanguages(req *hello.HelloManyLanguageRequest, stream hello.HelloService_HelloManyLanguagesServer) error {
	fmt.Printf("call hello from client request %v", req)
	request := req.GetHello()
	prefix := request.GetPrefix()
	firstName := request.GetFirstName()

	languages := []string{
		"Hello!",    // Inglés
		"Bonjour !", // Francés
		"¡Hola!",    // Español
		"Hallo!",    // Alemán
		"Ciao!",     // Italiano
		"Olá!",      // Portugués
	}

	size := len(languages)
	fmt.Println(size)
	for _, language := range languages {
		helloLanguage := language + prefix + ", " + firstName + " "
		response := &hello.HelloManyLanguageResponse{
			HelloLanguage: helloLanguage,
		}
		if err := stream.Send(response); err != nil {
			fmt.Printf("error in sending response to client %v", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (s *servers) HelloGoodBye(stream hello.HelloService_HelloGoodByeServer) error {

	fmt.Println("call hello from client")

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			//once is finished the streams we're going to send the response
			return stream.SendAndClose(&hello.HellosGoodByeResponse{
				Goodbye: "good bye",
			})
		}
		if err != nil {
			fmt.Printf("error in receiving from client %v", err)
			return err
		}

		request := message.GetHello()
		prefix := request.GetPrefix()
		firstName := request.GetFirstName()

		messages := "Hello: " + prefix + ", " + firstName + " welcome. "
		fmt.Println(messages)
	}
}

func (s servers) GoodBye(stream hello.HelloService_GoodByeServer) error {
	fmt.Println("call goodbye from client")
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Printf("error in receiving from client %v", err)
			return err
		}

		request := message.GetHello()
		prefix := request.GetPrefix()
		firstName := request.GetFirstName()
		messages := "Goodbye: " + prefix + ", " + firstName
		err = stream.Send(&hello.GoodByeResponse{
			Goodbye: messages,
		})
		if err != nil {
			fmt.Printf("error in sending response to client %v", err)
			return err
		}

	}
}

func (s *servers) Transfer(ctx context.Context, req *hello.TransferRequest) (*hello.TransferResponse, error) {
	fmt.Printf("call transfer from client %v", req)

	fromAccount := req.GetFromAccount()

	amount := req.GetAmount()

	if amount <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("amount must be > 0"))
	}

	if fromAccount != "1234567890" {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("fromAccount must begin with '1234567890'"))
	}

	year, month, day := time.Now().Date()

	response := &hello.TransferResponse{
		CreateAt: &date.Date{
			Day:   int32(day),
			Month: int32(month),
			Year:  int32(year),
		},
	}

	return response, nil
}

func main() {
	fmt.Println("hello, Go server is running...")
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		_ = fmt.Errorf("failed to listen: %v", err)
	}

	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"

	cred, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		fmt.Println("Failed to load TLS credentials")
		return
	}

	//grpcServer := grpc.NewServer()
	grpcServer := grpc.NewServer(grpc.Creds(cred))
	//hello.RegisterHelloServiceServer(grpcServer, &servers{})
	if err := grpcServer.Serve(listen); err != nil {
		_ = fmt.Errorf("failed to serve: %v", err)
	}
	stop := make(chan bool)
	go func() {

		fmt.Println("grpc server started")

	}()

	<-stop

}
```