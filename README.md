<!--
SPDX-FileCopyrightText: 2025 Canonical Ltd
SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
Copyright 2019 free5GC.org

SPDX-License-Identifier: Apache-2.0
-->
[![Go Report Card](https://goreportcard.com/badge/github.com/omec-project/ausf)](https://goreportcard.com/report/github.com/omec-project/ausf)

# ausf

AUSF is an important service function in the 5G Core network. It provides the
basic authentication for both 3gpp and non-3gpp access. Crucial for secured
network access, the AUSF is responsible for the security procedure for SIM
authentication using the 5G-AKA authentication method. AUSF connects and also
provides services with UDM (Unified Data Management) and AMF (Access and
Mobility Management Function) through SBI.

AMF requests the authentication of the UE by providing UE related information
and the serving network name and the 5G AKA is selected. The NF Service Consumer
(AMF) shall then return to the AUSF the result received from the UE.

## Repository Structure

Below is a high-level view of the repository and its main components:

```
.
├── ausf.go                         # Main entry point for the AUSF service
├── ausf_test.go
├── callback                        # Notification and subscription callback handling
│   ├── api_nf_subscribe_notify.go
│   └── router.go
├── consumer                        # Communication with external NFs 
│   ├── nf_discovery.go
│   └── nf_management.go
├── context                         # Runtime context and AUSF state management
│   ├── ausf_context_init.go
│   └── context.go
├── dev-container.ps1
├── dev-container.sh
├── DEV_README.md
├── Dockerfile                      # Container build definitions
├── Dockerfile_dev
├── Dockerfile.fast
├── factory                         # Configuration management and initialization
│   ├── ausfcfg_with_custom_webui_url.yaml
│   ├── ausfcfg.yaml
│   ├── ausf_config_test.go
│   ├── config.go
│   ├── factory.go
│   └── manual_config.go
├── go.mod
├── go.mod.license
├── go.sum
├── go.sum.license
├── LICENSES
│   └── Apache-2.0.txt
├── logger                          # Centralized logging utilities
│   └── logger.go
├── Makefile                        # Build and test automation
├── metrics                         # Telemetry and metrics collection
│   └── telemetry.go
├── nfregistration                  # NF registration and service discovery logic
│   ├── nf_registration.go
│   └── nf_registration_test.go
├── NOTICE.txt
├── polling                         # NF configuration polling and updates
│   ├── nf_configuration.go
│   └── nf_configuration_test.go
├── producer                        # API handlers and callbacks for incoming requests
│   ├── callback.go
│   ├── functions.go
│   ├── ue_authentication.go
│   └── ue_authentication_test.go
├── README.md
├── service                         # AUSF service initialization and startup
│   └── init.go
├── sorprotection                   # Service of Roaming (SOR) protection API implementation
│   ├── api_default.go
│   └── routers.go
├── Taskfile.yml                    # Build and test automation
├── test-mirror.txt
├── ueauthentication                # UE authentication APIs and routing
│   ├── api_default.go
│   └── routers.go
├── upuprotection                   # User Plane Integrity Protection (UPU) API implementation
│   ├── api_default.go
│   └── routers.go
├── util                            # Utility functions and test configurations
│   ├── manual_configs_functions.go
│   └── test
│       ├── testAmfcfg2.yaml
│       └── testAmfcfg.yaml
├── VERSION
└── VERSION.license

17 directories, 52 files
```

## Configuration and Deployment


**Docker**

To build the container image:

```
task mod-start
task build
task docker-build-fast
```

**Kubernetes**

The standard deployment uses Helm charts from the Aether project. The version of the Chart can be found in the OnRamp repository in the `vars/main.yml` file.


## Quick Navigation

| Directory           | Description                                                        |
| ------------------- | ------------------------------------------------------------------ |
| `producer/`         | Implements core authentication flows and exposes REST APIs to AMF. |
| `consumer/`         | Handles NF discovery and communication with UDM and NRF.           |
| `context/`          | Maintains runtime data for active authentication sessions.         |
| `factory/`          | Loads, validates, and manages AUSF configuration files.            |
| `nfregistration/`   | Manages AUSF registration to the NRF.                              |
| `polling/`          | Handles NF configuration synchronization and polling.              |
| `sorprotection/`    | Implements Service of Roaming (SOR) protection APIs.               |
| `upuprotection/`    | Provides User Plane Integrity Protection services.                 |
| `ueauthentication/` | Handles UE authentication API endpoints and routing.               |
| `metrics/`          | Collects telemetry and exports metrics for monitoring.             |
| `logger/`           | Provides structured and centralized logging for the component.     |

## Supported Features

1. Supports Nudm_UEAuthentication Services Procedure
2. Nausf_UEAuthentication (Authentication and Key Agreement)
3. Provides service on SBI interface Nausf

Compliance of the 5G Network functions can be found at [5G Compliance](https://docs.sd-core.opennetworking.org/main/overview/3gpp-compliance-5g.html)

## Dynamic Network configuration (via webconsole)

AUSF polls the webconsole every 5 seconds to fetch the latest PLMN configuration.

### Setting Up Polling

Include the `webuiUri` of the webconsole in the configuration file
```
configuration:
  ...
  webuiUri: https://webui:5001 # or http://webui:5001
  ...
```
The scheme (http:// or https://) must be explicitly specified. If no parameter is specified,
AUSF will use `http://webui:5001` by default.

### HTTPS Support

If the webconsole is served over HTTPS and uses a custom or self-signed certificate,
you must install the root CA certificate into the trust store of the AUSF environment.

Check the official guide for installing root CA certificates on Ubuntu:
[Install a Root CA Certificate in the Trust Store](https://documentation.ubuntu.com/server/how-to/security/install-a-root-ca-certificate-in-the-trust-store/index.html)

## Reach out to us through

1. #sdcore-dev channel in [ONF Community Slack](https://aether5g-project.slack.com)
2. Raise Github [issues](https://github.com/omec-project/ausf/issues/new)
