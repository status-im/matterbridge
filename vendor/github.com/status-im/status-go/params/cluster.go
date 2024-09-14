package params

// Define available fleets.
const (
	FleetUndefined     = ""
	FleetProd          = "eth.prod"
	FleetStatusStaging = "status.staging"
	FleetShardsTest    = "shards.test"
	FleetStatusProd    = "status.prod"
	FleetStatusTest    = "status.test"
	FleetWakuSandbox   = "waku.sandbox"
	FleetWakuTest      = "waku.test"
)

// Cluster defines a list of Ethereum nodes.
type Cluster struct {
	StaticNodes     []string `json:"staticnodes"`
	BootNodes       []string `json:"bootnodes"`
	MailServers     []string `json:"mailservers"` // list of trusted mail servers
	RendezvousNodes []string `json:"rendezvousnodes"`
}

type FleetName = string
type NodeType = string

const (
	WakuNodes            NodeType = "WakuNodes"
	DiscV5BootstrapNodes NodeType = "DiscV5BootstrapNodes"
)

// DefaultWakuNodes is a list of "supported" fleets. This list is populated to clients UI settings.
var supportedFleets = map[FleetName]map[NodeType][]string{
	FleetStatusStaging: {
		WakuNodes: {
			"enrtree://AI4W5N5IFEUIHF5LESUAOSMV6TKWF2MB6GU2YK7PU4TYUGUNOCEPW@boot.staging.status.nodes.status.im",
		},
		DiscV5BootstrapNodes: {
			"enrtree://AI4W5N5IFEUIHF5LESUAOSMV6TKWF2MB6GU2YK7PU4TYUGUNOCEPW@boot.staging.status.nodes.status.im",
			"enr:-QEQuEDsh6FgAb_36cReaX7W4gWx_7_GNpsUki7bXMoMrrrWij5pDEyV3guR-urDW_6GJTAzpQiJV61F-CfNn_NxPbY-AYJpZIJ2NIJpcISPxvrpim11bHRpYWRkcnO4YAAtNihib290LTAxLmRvLWFtczMuc2hhcmRzLnN0YWdpbmcuc3RhdHVzLmltBnZfAC82KGJvb3QtMDEuZG8tYW1zMy5zaGFyZHMuc3RhZ2luZy5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaEDIH8BcuEzgnmwPQTu7BPYyg4u4om7K9qekKA2gT_H2wSDdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
			"enr:-QEiuED2UusuHo1d6WN2-tHjtj0T0gdnsOh7aRZnFF6OEYLDbyxOtQo2_4dFUHhc9xm5SHNrWJJq8X7FRsxc4VCMGjjbAYJpZIJ2NIJpcIRoxQVgim11bHRpYWRkcnO4cgA2NjFib290LTAxLmdjLXVzLWNlbnRyYWwxLWEuc3RhdHVzLnN0YWdpbmcuc3RhdHVzLmltBnZfADg2MWJvb3QtMDEuZ2MtdXMtY2VudHJhbDEtYS5zdGF0dXMuc3RhZ2luZy5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaEDNAvlGjekD1YV4WpmjwArGAH2g9kHFJnMRfgUhcIkoA2DdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
			"enr:-QEiuECJPv2vL00Jp5sTEMAFyW7qXkK2cFgphlU_G8-FJuJqoW_D5aWIy3ylGdv2K8DkiG7PWgng4Ql_VI7Qc2RhBdwfAYJpZIJ2NIJpcIQvTKi6im11bHRpYWRkcnO4cgA2NjFib290LTAxLmFjLWNuLWhvbmdrb25nLWMuc3RhdHVzLnN0YWdpbmcuc3RhdHVzLmltBnZfADg2MWJvb3QtMDEuYWMtY24taG9uZ2tvbmctYy5zdGF0dXMuc3RhZ2luZy5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaEDkbgV7oqPNmFtX5FzSPi9WH8kkmrPB1R3n9xRXge91M-DdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
		},
	},
	FleetShardsTest: {
		WakuNodes: {
			"enrtree://AMOJVZX4V6EXP7NTJPMAYJYST2QP6AJXYW76IU6VGJS7UVSNDYZG4@boot.test.shards.nodes.status.im",
		},
		DiscV5BootstrapNodes: {
			"enrtree://AMOJVZX4V6EXP7NTJPMAYJYST2QP6AJXYW76IU6VGJS7UVSNDYZG4@boot.test.shards.nodes.status.im",
			"enr:-QEKuECA0zhRJej2eaOoOPddNcYr7-5NdRwuoLCe2EE4wfEYkAZhFotg6Kkr8K15pMAGyUyt0smHkZCjLeld0BUzogNtAYJpZIJ2NIJpcISnYxMvim11bHRpYWRkcnO4WgAqNiVib290LTAxLmRvLWFtczMuc2hhcmRzLnRlc3Quc3RhdHVzLmltBnZfACw2JWJvb3QtMDEuZG8tYW1zMy5zaGFyZHMudGVzdC5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaEC3rRtFQSgc24uWewzXaxTY8hDAHB8sgnxr9k8Rjb5GeSDdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
			"enr:-QEcuEAgXDqrYd_TrpUWtn3zmxZ9XPm7O3GS6lV7aMJJOTsbOAAeQwSd_eoHcCXqVzTUtwTyB4855qtbd8DARnExyqHPAYJpZIJ2NIJpcIQihw1Xim11bHRpYWRkcnO4bAAzNi5ib290LTAxLmdjLXVzLWNlbnRyYWwxLWEuc2hhcmRzLnRlc3Quc3RhdHVzLmltBnZfADU2LmJvb3QtMDEuZ2MtdXMtY2VudHJhbDEtYS5zaGFyZHMudGVzdC5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaECxjqgDQ0WyRSOilYU32DA5k_XNlDis3m1VdXkK9xM6kODdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
			"enr:-QEcuEAX6Qk-vVAoJLxR4A_4UVogGhvQrqKW4DFKlf8MA1PmCjgowL-LBtSC9BLjXbb8gf42FdDHGtSjEvvWKD10erxqAYJpZIJ2NIJpcIQI2hdMim11bHRpYWRkcnO4bAAzNi5ib290LTAxLmFjLWNuLWhvbmdrb25nLWMuc2hhcmRzLnRlc3Quc3RhdHVzLmltBnZfADU2LmJvb3QtMDEuYWMtY24taG9uZ2tvbmctYy5zaGFyZHMudGVzdC5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaEDP7CbRk-YKJwOFFM4Z9ney0GPc7WPJaCwGkpNRyla7mCDdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
		},
	},
	FleetWakuSandbox: {
		WakuNodes: {
			"enrtree://AIRVQ5DDA4FFWLRBCHJWUWOO6X6S4ZTZ5B667LQ6AJU6PEYDLRD5O@sandbox.waku.nodes.status.im",
		},
		DiscV5BootstrapNodes: {
			"enrtree://AIRVQ5DDA4FFWLRBCHJWUWOO6X6S4ZTZ5B667LQ6AJU6PEYDLRD5O@sandbox.waku.nodes.status.im",
		},
	},
	FleetWakuTest: {
		WakuNodes: {
			"enrtree://AOGYWMBYOUIMOENHXCHILPKY3ZRFEULMFI4DOM442QSZ73TT2A7VI@test.waku.nodes.status.im",
		},
		DiscV5BootstrapNodes: {
			"enrtree://AOGYWMBYOUIMOENHXCHILPKY3ZRFEULMFI4DOM442QSZ73TT2A7VI@test.waku.nodes.status.im",
		},
	},
}

func DefaultWakuNodes(fleet string) []string {
	return supportedFleets[fleet][WakuNodes]
}

func DefaultDiscV5Nodes(fleet string) []string {
	return supportedFleets[fleet][DiscV5BootstrapNodes]

}

func IsFleetSupported(fleet string) bool {
	_, ok := supportedFleets[fleet]
	return ok
}

func GetSupportedFleets() map[FleetName]map[NodeType][]string {
	return supportedFleets
}