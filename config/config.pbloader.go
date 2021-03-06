// Code generated by hypertrace/goagent/config. DO NOT EDIT.

package config

import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *AgentConfig) loadFromEnv(prefix string, defaultValues *AgentConfig) {
	if val, ok := getStringEnv(prefix + "SERVICE_NAME"); ok {
		x.ServiceName = &wrappers.StringValue{Value: val}
	} else if x.ServiceName == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.ServiceName = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.ServiceName != nil {
			x.ServiceName = &wrappers.StringValue{Value: defaultValues.ServiceName.Value}
		}
	}
	if x.Reporting == nil {
		x.Reporting = new(Reporting)
	}
	x.Reporting.loadFromEnv(prefix+"REPORTING_", defaultValues.Reporting)
	if x.DataCapture == nil {
		x.DataCapture = new(DataCapture)
	}
	x.DataCapture.loadFromEnv(prefix+"DATA_CAPTURE_", defaultValues.DataCapture)
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *Reporting) loadFromEnv(prefix string, defaultValues *Reporting) {
	if val, ok := getStringEnv(prefix + "ENDPOINT"); ok {
		x.Endpoint = &wrappers.StringValue{Value: val}
	} else if x.Endpoint == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Endpoint = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.Endpoint != nil {
			x.Endpoint = &wrappers.StringValue{Value: defaultValues.Endpoint.Value}
		}
	}
	if val, ok := getBoolEnv(prefix + "SECURE"); ok {
		x.Secure = &wrappers.BoolValue{Value: val}
	} else if x.Secure == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Secure = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Secure != nil {
			x.Secure = &wrappers.BoolValue{Value: defaultValues.Secure.Value}
		}
	}
	if val, ok := getStringEnv(prefix + "TOKEN"); ok {
		x.Token = &wrappers.StringValue{Value: val}
	} else if x.Token == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Token = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.Token != nil {
			x.Token = &wrappers.StringValue{Value: defaultValues.Token.Value}
		}
	}
	if x.Opa == nil {
		x.Opa = new(Opa)
	}
	x.Opa.loadFromEnv(prefix+"OPA_", defaultValues.Opa)
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *Opa) loadFromEnv(prefix string, defaultValues *Opa) {
	if val, ok := getStringEnv(prefix + "ENDPOINT"); ok {
		x.Endpoint = &wrappers.StringValue{Value: val}
	} else if x.Endpoint == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Endpoint = new(wrappers.StringValue)
		if defaultValues != nil && defaultValues.Endpoint != nil {
			x.Endpoint = &wrappers.StringValue{Value: defaultValues.Endpoint.Value}
		}
	}
	if val, ok := getInt32Env(prefix + "POLL_PERIOD_SECONDS"); ok {
		x.PollPeriodSeconds = &wrappers.Int32Value{Value: val}
	} else if x.PollPeriodSeconds == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.PollPeriodSeconds = new(wrappers.Int32Value)
		if defaultValues != nil && defaultValues.PollPeriodSeconds != nil {
			x.PollPeriodSeconds = &wrappers.Int32Value{Value: defaultValues.PollPeriodSeconds.Value}
		}
	}
	if val, ok := getBoolEnv(prefix + "ENABLED"); ok {
		x.Enabled = &wrappers.BoolValue{Value: val}
	} else if x.Enabled == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Enabled = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Enabled != nil {
			x.Enabled = &wrappers.BoolValue{Value: defaultValues.Enabled.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *Message) loadFromEnv(prefix string, defaultValues *Message) {
	if val, ok := getBoolEnv(prefix + "REQUEST"); ok {
		x.Request = &wrappers.BoolValue{Value: val}
	} else if x.Request == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Request = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Request != nil {
			x.Request = &wrappers.BoolValue{Value: defaultValues.Request.Value}
		}
	}
	if val, ok := getBoolEnv(prefix + "RESPONSE"); ok {
		x.Response = &wrappers.BoolValue{Value: val}
	} else if x.Response == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.Response = new(wrappers.BoolValue)
		if defaultValues != nil && defaultValues.Response != nil {
			x.Response = &wrappers.BoolValue{Value: defaultValues.Response.Value}
		}
	}
}

// loadFromEnv loads the data from env vars, defaults and makes sure all values are initialized.
func (x *DataCapture) loadFromEnv(prefix string, defaultValues *DataCapture) {
	if x.HttpHeaders == nil {
		x.HttpHeaders = new(Message)
	}
	x.HttpHeaders.loadFromEnv(prefix+"HTTP_HEADERS_", defaultValues.HttpHeaders)
	if x.HttpBody == nil {
		x.HttpBody = new(Message)
	}
	x.HttpBody.loadFromEnv(prefix+"HTTP_BODY_", defaultValues.HttpBody)
	if x.RpcMetadata == nil {
		x.RpcMetadata = new(Message)
	}
	x.RpcMetadata.loadFromEnv(prefix+"RPC_METADATA_", defaultValues.RpcMetadata)
	if x.RpcBody == nil {
		x.RpcBody = new(Message)
	}
	x.RpcBody.loadFromEnv(prefix+"RPC_BODY_", defaultValues.RpcBody)
	if val, ok := getInt32Env(prefix + "BODY_MAX_SIZE_BYTES"); ok {
		x.BodyMaxSizeBytes = &wrappers.Int32Value{Value: val}
	} else if x.BodyMaxSizeBytes == nil {
		// when there is no value to set we still prefer to initialize the variable to avoid
		// `nil` checks in the consumers.
		x.BodyMaxSizeBytes = new(wrappers.Int32Value)
		if defaultValues != nil && defaultValues.BodyMaxSizeBytes != nil {
			x.BodyMaxSizeBytes = &wrappers.Int32Value{Value: defaultValues.BodyMaxSizeBytes.Value}
		}
	}
}
