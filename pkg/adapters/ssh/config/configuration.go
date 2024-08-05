package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

/**
*
* configuration
* <p>
* configuration file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author bxcode
* @author dbacilio88@outlook.es
* @since 3/08/2024
*
 */

var Config Configurations

func Load(path string) {
	viper.SetConfigName("application")
	viper.AddConfigPath(path)
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		var file viper.ConfigFileNotFoundError
		if errors.As(err, &file) {
			log.Fatalf("Error reading config file, %s", file)
			return
		}
		return
	}
	err := viper.WriteConfig()
	if err != nil {
		log.Fatalf("Error writing config file, %s", err)
		return
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Error unmarshalling config, %s", err)
		return
	}

	log.Println("Successfully loaded config")

}

type Configurations struct {
	Server    Server    `mapstructure:"server" yaml:"server"`
	Ssh       Ssh       `mapstructure:"ssh" yaml:"ssh"`
	Scheduler Scheduler `mapstructure:"scheduler" yaml:"scheduler"`
}

// Server use mapstructure in document github.com/go-viper/mapstructure/v2
type Server struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port string `mapstructure:"port" yaml:"port"`
}
type Ssh struct {
	Host       string `mapstructure:"host" yaml:"host"`
	Port       string `mapstructure:"port" yaml:"port"`
	Protocol   string `mapstructure:"protocol" yaml:"protocol"`
	Username   string `mapstructure:"username" yaml:"username"`
	Password   string `mapstructure:"password" yaml:"password"`
	PrivateKey string `mapstructure:"private_key" yaml:"private_key"`
	PublicKey  string `mapstructure:"public_key" yaml:"public_key"`
	KnownHosts string `mapstructure:"known_hosts" yaml:"known_hosts"`
	SftpPath   string `mapstructure:"sftp_path" yaml:"sftp_path"`
	Enable     bool   `mapstructure:"enable" yaml:"enable"`
}
type Scheduler struct {
	Enable bool `mapstructure:"enable" yaml:"enable"`
}
