package driver

import (
	"context"

	csipb "github.com/shravaniphadolibm/csi-driver/api"
)

type IdentityServer struct {
	csipb.UnimplementedIdentityServer
}

func (s *IdentityServer) GetPluginInfo(ctx context.Context, req *csipb.GetPluginInfoRequest) (*csipb.GetPluginInfoResponse, error) {
	// Return dummy plugin info for now
	return &csipb.GetPluginInfoResponse{
		Name:          "csi-hostpath-driver",
		VendorVersion: "1.0.0",
	}, nil
}
func (s *IdentityServer) GetPluginCapabilities(ctx context.Context, req *csipb.GetPluginCapabilitiesRequest) (*csipb.GetPluginCapabilitiesResponse, error) {
	// Return dummy capabilities for now
	return &csipb.GetPluginCapabilitiesResponse{
		Capabilities: []*csipb.PluginCapability{
			&csipb.PluginCapability{
				Type: &csipb.PluginCapability_Service{
					Service: csipb.PluginCapability_Service_CONTROLLER_SERVICE,
				},
			},
			&csipb.PluginCapability{
				Type: &csipb.PluginCapability_Service{
					Service: csipb.PluginCapability_Service_NODE_SERVICE,
				},
			},
		},
	}, nil
}
