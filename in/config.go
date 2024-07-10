/*
Copyright 2024 The nsqcc Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package in

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	ntls "github.com/bytemind-io/nsqcc/tls"
)

// Config is the configuration for the reader.
type Config struct {
	Addresses       []string `json:",optional,env=NSQ_ADDRESSES,default=127.0.0.1:4150"                 envconfig:"NSQ_ADDRESSES"                   default:"127.0.0.1:4150"`        // Nsqd 地址列表
	LookupAddresses []string `json:",optional,env=NSQ_LOOKUP_ADDRESSES,default=127.0.0.1:4161"          envconfig:"NSQ_LOOKUP_ADDRESSES"            default:"127.0.0.1:4161"`        // NSQLookupd 地址列表
	Topic           string   `json:",optional,env=NSQ_TOPIC"                                            envconfig:"NSQ_TOPIC"`                                                       // 消费的主题名
	Channel         string   `json:",optional,env=NSQ_CHANNEL,default=default"                          envconfig:"NSQ_CHANNEL"                     default:"default"`               // 消费的频道名
	UserAgent       string   `json:",optional,env=NSQ_USER_AGENT,default=DeepAutoConsumer/1.0"          envconfig:"NSQ_USER_AGENT"                  default:"DeepAuto Consumer/1.0"` // 连接时使用的用户UA
	MaxInFlight     int      `json:",optional,env=NSQ_MAX_IN_FLIGHT,default=64"                         envconfig:"NSQ_MAX_IN_FLIGHT"               default:"64"`                    // 同时处理的最大消息数量.
	MaxAttempts     uint16   `json:",optional,env=NSQ_MAX_ATTEMPTS,default=3"                           envconfig:"NSQ_MAX_ATTEMPTS"                default:"3"`                     // 消息最大重试次数
	TLS             ntls.Config
}

// NewConfig creates a new Config with default values.
func NewConfig() Config {
	return Config{
		Channel:         "default",
		Addresses:       []string{"127.0.0.1:4150"},
		LookupAddresses: []string{"127.0.0.1:4161"},
		UserAgent:       "DeepAuto NSQ/1.0",
		MaxInFlight:     64,
		MaxAttempts:     5,
		TLS:             ntls.NewConfig(),
	}
}

// Validate validates the configuration.
func (c Config) Validate() error {
	if len(c.Addresses) == 0 {
		return fmt.Errorf("nsq address is required")
	}

	if len(c.LookupAddresses) == 0 {
		return fmt.Errorf("nsq lookupd addresses is required")
	}

	if govalidator.IsNull(c.Channel) {
		return fmt.Errorf("nsq channel is required")
	}
	return nil
}
