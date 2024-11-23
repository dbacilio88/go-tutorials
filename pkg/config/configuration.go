package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
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
* @author christian
* @author dbacilio88@outlook.es
* @since 3/08/2024
*
 */

var Config Configurations

func Load() {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		log.Fatal("La variable de entorno CONFIG_PATH no está definida")
	}

	viper.SetConfigName("application")
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	// Si se desea pasar el archivo por variable de entorno:
	//viper.SetConfigFile(path)
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
	Database  Database  `mapstructure:"database" yaml:"database"`
	Rabbitmq  Rabbitmq  `mapstructure:"rabbitmq" yaml:"rabbitmq"`
	Queue     Queue     `mapstructure:"queue" yaml:"queue"`
	Grpc      Grpc      `mapstructure:"grpc" yaml:"grpc"`
}

// Server use mapstructure in document github.com/go-viper/mapstructure/v2
type Server struct {
	Host        string `mapstructure:"host" yaml:"host"`
	Port        string `mapstructure:"port" yaml:"port"`
	Name        string `mapstructure:"name" yaml:"name"`
	Timeout     int    `mapstructure:"timeout" yaml:"timeout"`
	Logging     string `mapstructure:"logging" yaml:"logging"`
	Environment string `mapstructure:"environment" yaml:"environment"`
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

type Database struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     string `mapstructure:"port" yaml:"port"`
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	Dbname   string `mapstructure:"dbname" yaml:"dbname"`
}

type Rabbitmq struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     string `mapstructure:"port" yaml:"port"`
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	Vhost    string `mapstructure:"vhost" yaml:"vhost"`
	Protocol string `mapstructure:"protocol" yaml:"protocol"`
}

type Queue struct {
	Consumer string `mapstructure:"consumer" yaml:"consumer"`
	Producer string `mapstructure:"producer" yaml:"producer"`
}

type Grpc struct {
	Server   string `mapstructure:"server" yaml:"server"`
	Client   string `mapstructure:"client" yaml:"client"`
	Protocol string `mapstructure:"protocol" yaml:"protocol"`
	Cert     string `mapstructure:"cert" yaml:"cert"`
	Key      string `mapstructure:"key" yaml:"key"`
}

func GetDomainRabbitConnection() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/",
		Config.Rabbitmq.Protocol,
		Config.Rabbitmq.User,
		Config.Rabbitmq.Password,
		Config.Rabbitmq.Host,
		Config.Rabbitmq.Port)
}
