//
//  Copyright (c) 2021 Datastax, Inc.
//
//  Licensed to the Apache Software Foundation (ASF) under one
//  or more contributor license agreements.  See the NOTICE file
//  distributed with this work for additional information
//  regarding copyright ownership.  The ASF licenses this file
//  to you under the Apache License, Version 2.0 (the
//  "License"); you may not use this file except in compliance
//  with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an
//  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
//  KIND, either express or implied.  See the License for the
//  specific language governing permissions and limitations
//  under the License.
//

package cfg

import (
	pd "github.com/PagerDuty/go-pagerduty"
	log "github.com/apex/log"
)

const (
	trigger     = "trigger"
	acknowledge = "acknowledge"
	resolve     = "resolve"
)

// CreatePDIncident creates PagerDuty incident
func CreatePDIncident(component, alias, msg, pdIntegrationKey string) {
	payload := pd.V2Payload{
		Summary:   component + ":" + msg,
		Source:    "pulsar-heartbeat",
		Severity:  "critical",
		Component: component,
	}
	PdV2Event(trigger, alias, pdIntegrationKey, &payload)
}

// ResolvePDIncident resolves PagerDuty incident
func ResolvePDIncident(component, alias, pdIntegrationKey string) {
	payload := pd.V2Payload{
		Summary:   component,
		Source:    "pulsar-heartbeat",
		Severity:  "critical",
		Component: component,
	}
	PdV2Event(resolve, alias, pdIntegrationKey, &payload)
}

// PdV2Event is pd client
func PdV2Event(action, dedupKey, routingKey string, payload *pd.V2Payload) {
	if routingKey == "" {
		return
	}
	v2Event := pd.V2Event{
		RoutingKey: routingKey,
		DedupKey:   dedupKey,
		Action:     action,
		Payload:    payload,
	}
	resp, err := pd.ManageEvent(v2Event)
	if err != nil {
		log.Errorf("failed V2Event to PagerDuty error - %v", err)
	} else {
		log.Infof("PagerDuty V2Event sent with response - %v", resp)
	}
}
