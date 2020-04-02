package install

import (
	"github.com/cuisongliu/sshcmd/pkg/sshutil"
	"github.com/fanux/lvscare/care"
	"regexp"
)

var (
	MasterIPs   []string
	NodeIPs     []string
	ApiServerCertSANs []string
	VIP         string
	PkgUrl      string
	KubeadmFile string
	Version     string
	SSHConfig   sshutil.SSH
	Ipvs        care.LvsCare
	Kustomize   bool
	ApiServer   string
	Repo        string
	PodCIDR     string
	SvcCIDR     string
	// network type, calico or flannel etc..
	Network string
	// if true don't install cni plugin
	WithoutCNI bool
	//network interface name, like "eth.*|en.*"
	Interface string
	// the ipip mode of the calico
	IPIP bool
	// mtu size
	MTU string

	YesRx = regexp.MustCompile("^(?i:y(?:es)?)$")

	CleanForce bool
)
