// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * AUSF Configuration Factory
 */

package factory

import (
	"fmt"
	"net/url"
	"os"

	"github.com/omec-project/ausf/logger"
	"gopkg.in/yaml.v2"
)

var AusfConfig Config

// TODO: Support configuration update from REST api
func InitConfigFactory(f string) error {
	content, err := os.ReadFile(f)
	if err != nil {
		return err
	}
	AusfConfig = Config{}

	if err = yaml.Unmarshal(content, &AusfConfig); err != nil {
		return err
	}
	if AusfConfig.Configuration.WebuiUri == "" {
		AusfConfig.Configuration.WebuiUri = "http://webui:5001"
		logger.CfgLog.Infof("webuiUri not set in configuration file. Using %v", AusfConfig.Configuration.WebuiUri)
		return nil
	}

	if AusfConfig.Configuration.ManualConfigs != nil {
		logger.CfgLog.Infof("Manual Configuration provided for network functions")
		for nfType, nfs := range AusfConfig.Configuration.ManualConfigs.NFs {
			for _, nf := range nfs {
				logger.CfgLog.Debugf("Manual Configuration - NF Type: %s, Name: %s, URL: %s, Port: %d", nfType, nf.NfInstanceName, nf.NfServices)
			}
		}
	} else {
		logger.CfgLog.Infof("No manual configuration provided for network functions")
	}
	err = validateWebuiUri(AusfConfig.Configuration.WebuiUri)
	return err
}

func CheckConfigVersion() error {
	currentVersion := AusfConfig.GetVersion()

	if currentVersion != AUSF_EXPECTED_CONFIG_VERSION {
		return fmt.Errorf("config version is [%s], but expected is [%s]",
			currentVersion, AUSF_EXPECTED_CONFIG_VERSION)
	}
	logger.CfgLog.Infof("config version [%s]", currentVersion)

	return nil
}

func validateWebuiUri(uri string) error {
	parsedUrl, err := url.ParseRequestURI(uri)
	if err != nil {
		return err
	}
	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return fmt.Errorf("unsupported scheme for webuiUri: %s", parsedUrl.Scheme)
	}
	if parsedUrl.Hostname() == "" {
		return fmt.Errorf("missing host in webuiUri")
	}
	return nil
}
