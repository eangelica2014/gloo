package endpointdiscovery

import "github.com/solo-io/glue/pkg/api/types/v1"

// groups endpoints by their respective upstream name
type EndpointGroups map[string][]Endpoint

type Endpoint struct {
	Address string
	Port    int32
}

type Interface interface {
	// starts the discovery service
	Run(stop <-chan struct{})

	// tells the discovery to track endpoints for the given upstreams
	TrackUpstreams(upstreams []*v1.Upstream)

	// endpoint groups are pushed here whenever they are updated
	Endpoints() <-chan EndpointGroups

	// should show valid if the most recent update passed, otherwise a useful error
	Error() <-chan error
}
