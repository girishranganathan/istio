// Copyright 2018 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v2

import (
	xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/gogo/protobuf/types"
	"github.com/prometheus/client_golang/prometheus"

	"istio.io/istio/pilot/pkg/model"
)

// clusters aggregate a DiscoveryResponse for pushing.
func (con *XdsConnection) clusters(response []*xdsapi.Cluster) *xdsapi.DiscoveryResponse {
	out := &xdsapi.DiscoveryResponse{
		// All resources for CDS ought to be of the type ClusterLoadAssignment
		TypeUrl: ClusterType,

		// Pilot does not really care for versioning. It always supplies what's currently
		// available to it, irrespective of whether Envoy chooses to accept or reject CDS
		// responses. Pilot believes in eventual consistency and that at some point, Envoy
		// will begin seeing results it deems to be good.
		VersionInfo: versionInfo(),
		Nonce:       nonce(),
	}

	for _, c := range response {
		cc, _ := types.MarshalAny(c)
		out.Resources = append(out.Resources, *cc)
	}

	return out
}

func (s *DiscoveryServer) pushCds(node model.Proxy, con *XdsConnection) error {
	rawClusters, err := s.ConfigGenerator.BuildClusters(s.env, *con.modelNode)
	if err != nil {
		adsLog.Warnf("CDS: Failed to generate clusters %v", err)
		pushes.With(prometheus.Labels{"type": "cds_builderr"}).Add(1)
		return err
	}

	con.HTTPClusters = rawClusters
	response := con.clusters(rawClusters)
	err = con.send(response)
	if err != nil {
		adsLog.Warnf("CDS: Send failure, closing grpc %s %v", node.ID, err)
		pushes.With(prometheus.Labels{"type": "cds_senderr"}).Add(1)
		return err
	}
	pushes.With(prometheus.Labels{"type": "cds"}).Add(1)

	// The response can't be easily read due to 'any' marshalling.
	adsLog.Infof("CDS: PUSH for %s %q, Response: %d",
		node, con.PeerAddr, len(rawClusters))
	return nil
}
