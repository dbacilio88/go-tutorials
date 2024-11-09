package validation

import (
	"bytes"
	"encoding/json"
	"github.com/dbacilio88/go/models/request"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
	"io"
)

/**
*
* validationService
* <p>
* validationService file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author bxcode
* @author dbacilio88@outlook.es
* @since 9/11/2024
*
 */

type Executor interface {
	ValidateRequest(data interface{}) (request.MessageRequest, error)
}

type ValidatorService struct {
	console *zap.Logger
}

func NewValidatorService(console *zap.Logger) *ValidatorService {
	return &ValidatorService{
		console: console,
	}
}
func (svc *ValidatorService) ValidateRequest(data interface{}) (request.MessageRequest, error) {
	var req request.MessageRequest
	var err error
	var decoder *json.Decoder

	validate := validator.New()

	if parse, ok := data.(io.Reader); ok {
		decoder = json.NewDecoder(parse)
	} else {
		marshal, _ := json.Marshal(data)
		decoder = json.NewDecoder(bytes.NewReader(marshal))
	}
	if err = decoder.Decode(&req); err != nil {
		return request.MessageRequest{}, err
	}

	err = validate.Struct(req)
	if err != nil {
		return request.MessageRequest{}, err
	}
	return req, nil
}
