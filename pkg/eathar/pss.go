package eathar

/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/
import (
	"strings"

	"github.com/spf13/pflag"
)

//This needs to be exported to work with the JSON marshalling
// omitempty thing is there as container won't always be relevant (e.g. hostPID)
type Finding struct {
	Check        string
	Namespace    string
	Pod          string
	Container    string   `json:",omitempty"`
	Capabilities []string `json:",omitempty"`
	Hostport     int      `json:",omitempty"`
	Volume       string   `json:",omitempty"`
	Path         string   `json:",omitempty"`
	Sysctl       string   `json:",omitempty"`
	Image        string   `json:",omitempty"`
}

func Hostnet(options *pflag.FlagSet) []Finding {
	var hostnetcont []Finding
	pods := connectWithPods(options)
	for _, pod := range pods.Items {

		if pod.Spec.HostNetwork {
			p := Finding{Check: "hostnet", Namespace: pod.Namespace, Pod: pod.Name}
			hostnetcont = append(hostnetcont, p)
		}
	}
	return hostnetcont
}

func Hostpid(options *pflag.FlagSet) []Finding {
	var hostpidcont []Finding
	pods := connectWithPods(options)

	for _, pod := range pods.Items {

		if pod.Spec.HostPID {
			p := Finding{Check: "hostpid", Namespace: pod.Namespace, Pod: pod.Name, Container: ""}
			hostpidcont = append(hostpidcont, p)
		}
	}
	return hostpidcont

}

func Hostipc(options *pflag.FlagSet) []Finding {
	var hostipccont []Finding
	pods := connectWithPods(options)

	for _, pod := range pods.Items {

		if pod.Spec.HostIPC {
			p := Finding{Check: "hostipc", Namespace: pod.Namespace, Pod: pod.Name, Container: ""}
			hostipccont = append(hostipccont, p)
		}
	}
	return hostipccont
}

func HostProcess(options *pflag.FlagSet) []Finding {
	var hostprocesscont []Finding
	pods := connectWithPods(options)
	for _, pod := range pods.Items {
		hostProcessPod := pod.Spec.SecurityContext.WindowsOptions != nil && *pod.Spec.SecurityContext.WindowsOptions.HostProcess
		if hostProcessPod {
			p := Finding{Check: "HostProcess", Namespace: pod.Namespace, Pod: pod.Name}
			hostprocesscont = append(hostprocesscont, p)
		}
		for _, container := range pod.Spec.Containers {
			hostProcessCont := container.SecurityContext != nil && container.SecurityContext.WindowsOptions != nil && *container.SecurityContext.WindowsOptions.HostProcess
			if hostProcessCont {
				p := Finding{Check: "HostProcess", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name}
				hostprocesscont = append(hostprocesscont, p)
			}
		}
		for _, init_container := range pod.Spec.InitContainers {
			hostProcessCont := init_container.SecurityContext != nil && init_container.SecurityContext.WindowsOptions != nil && *init_container.SecurityContext.WindowsOptions.HostProcess
			if hostProcessCont {
				p := Finding{Check: "HostProcess", Namespace: pod.Namespace, Pod: pod.Name, Container: init_container.Name}
				hostprocesscont = append(hostprocesscont, p)
			}
		}
		for _, eph_container := range pod.Spec.EphemeralContainers {
			hostProcessCont := eph_container.SecurityContext != nil && eph_container.SecurityContext.WindowsOptions != nil && *eph_container.SecurityContext.WindowsOptions.HostProcess
			if hostProcessCont {
				p := Finding{Check: "HostProcess", Namespace: pod.Namespace, Pod: pod.Name, Container: eph_container.Name}
				hostprocesscont = append(hostprocesscont, p)
			}
		}
	}
	return hostprocesscont

}

func AllowPrivEsc(options *pflag.FlagSet) []Finding {
	var allowprivesccont []Finding
	pods := connectWithPods(options)
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			// Logic here is if there's no security context, or there is a security context and no mention of allow privilege escalation then the default is true
			// We don't catch the case of someone explicitly setting it to true, but that seems unlikely
			allowPrivilegeEscalation := (container.SecurityContext == nil) || (container.SecurityContext != nil && container.SecurityContext.AllowPrivilegeEscalation == nil)
			if allowPrivilegeEscalation {
				p := Finding{Check: "allowprivesc", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name}
				allowprivesccont = append(allowprivesccont, p)
			}
		}
		for _, init_container := range pod.Spec.InitContainers {
			allowPrivilegeEscalation := (init_container.SecurityContext == nil) || (init_container.SecurityContext != nil && init_container.SecurityContext.AllowPrivilegeEscalation == nil)
			if allowPrivilegeEscalation {
				p := Finding{Check: "allowprivesc", Namespace: pod.Namespace, Pod: pod.Name, Container: init_container.Name}
				allowprivesccont = append(allowprivesccont, p)
			}
		}
		for _, eph_container := range pod.Spec.EphemeralContainers {
			allowPrivilegeEscalation := (eph_container.SecurityContext == nil) || (eph_container.SecurityContext != nil && eph_container.SecurityContext.AllowPrivilegeEscalation == nil)
			if allowPrivilegeEscalation {
				p := Finding{Check: "allowprivesc", Namespace: pod.Namespace, Pod: pod.Name, Container: eph_container.Name}
				allowprivesccont = append(allowprivesccont, p)
			}
		}
	}
	return allowprivesccont

}

func Privileged(options *pflag.FlagSet) []Finding {
	var privcont []Finding
	pods := connectWithPods(options)
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			privileged_container := container.SecurityContext != nil && container.SecurityContext.Privileged != nil && *container.SecurityContext.Privileged
			if privileged_container {
				p := Finding{Check: "privileged", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name}
				privcont = append(privcont, p)
			}
		}
		for _, init_container := range pod.Spec.InitContainers {
			privileged_container := init_container.SecurityContext != nil && init_container.SecurityContext.Privileged != nil && *init_container.SecurityContext.Privileged
			if privileged_container {
				p := Finding{Check: "privileged", Namespace: pod.Namespace, Pod: pod.Name, Container: init_container.Name}
				privcont = append(privcont, p)
			}
		}
		for _, eph_container := range pod.Spec.EphemeralContainers {
			privileged_container := eph_container.SecurityContext != nil && eph_container.SecurityContext.Privileged != nil && *eph_container.SecurityContext.Privileged
			if privileged_container {
				p := Finding{Check: "privileged", Namespace: pod.Namespace, Pod: pod.Name, Container: eph_container.Name}
				privcont = append(privcont, p)
			}
		}
	}
	return privcont

}

func AddedCapabilities(options *pflag.FlagSet) []Finding {
	var capadded []Finding
	pods := connectWithPods(options)
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			cap_added := container.SecurityContext != nil && container.SecurityContext.Capabilities != nil && container.SecurityContext.Capabilities.Add != nil
			if cap_added {
				//Need to convert the capabilities struct to strings.
				var added_caps []string
				for _, cap := range container.SecurityContext.Capabilities.Add {
					added_caps = append(added_caps, string(cap))
				}
				p := Finding{Check: "Added Capabilities", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name, Capabilities: added_caps}
				capadded = append(capadded, p)
			}
		}

		for _, init_container := range pod.Spec.InitContainers {
			cap_added := init_container.SecurityContext != nil && init_container.SecurityContext.Capabilities != nil && init_container.SecurityContext.Capabilities.Add != nil
			if cap_added {
				var added_caps []string
				for _, cap := range init_container.SecurityContext.Capabilities.Add {
					added_caps = append(added_caps, string(cap))
				}
				p := Finding{Check: "Added Capabilities", Namespace: pod.Namespace, Pod: pod.Name, Container: init_container.Name, Capabilities: added_caps}
				capadded = append(capadded, p)
			}
		}

		for _, eph_container := range pod.Spec.EphemeralContainers {
			cap_added := eph_container.SecurityContext != nil && eph_container.SecurityContext.Capabilities != nil && eph_container.SecurityContext.Capabilities.Add != nil
			if cap_added {
				var added_caps []string
				for _, cap := range eph_container.SecurityContext.Capabilities.Add {
					added_caps = append(added_caps, string(cap))
				}
				p := Finding{Check: "Added Capabilities", Namespace: pod.Namespace, Pod: pod.Name, Container: eph_container.Name, Capabilities: added_caps}
				capadded = append(capadded, p)
			}
		}
	}
	return capadded

}

func DroppedCapabilities(options *pflag.FlagSet) []Finding {
	var capdropped []Finding
	pods := connectWithPods(options)
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			cap_dropped := container.SecurityContext != nil && container.SecurityContext.Capabilities != nil && container.SecurityContext.Capabilities.Drop != nil
			if cap_dropped {
				var dropped_caps []string
				for _, cap := range container.SecurityContext.Capabilities.Drop {
					dropped_caps = append(dropped_caps, string(cap))
				}
				p := Finding{Check: "Dropped Capabilities", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name, Capabilities: dropped_caps}
				capdropped = append(capdropped, p)
			}
		}

		for _, init_container := range pod.Spec.InitContainers {
			cap_dropped := init_container.SecurityContext != nil && init_container.SecurityContext.Capabilities != nil && init_container.SecurityContext.Capabilities.Drop != nil
			if cap_dropped {
				var dropped_caps []string
				for _, cap := range init_container.SecurityContext.Capabilities.Drop {
					dropped_caps = append(dropped_caps, string(cap))
				}
				p := Finding{Check: "Dropped Capabilities", Namespace: pod.Namespace, Pod: pod.Name, Container: init_container.Name, Capabilities: dropped_caps}
				capdropped = append(capdropped, p)
			}
		}

		for _, eph_container := range pod.Spec.EphemeralContainers {
			cap_dropped := eph_container.SecurityContext != nil && eph_container.SecurityContext.Capabilities != nil && eph_container.SecurityContext.Capabilities.Drop != nil
			if cap_dropped {
				var dropped_caps []string
				for _, cap := range eph_container.SecurityContext.Capabilities.Drop {
					dropped_caps = append(dropped_caps, string(cap))
				}
				p := Finding{Check: "Dropped Capabilities", Namespace: pod.Namespace, Pod: pod.Name, Container: eph_container.Name, Capabilities: dropped_caps}
				capdropped = append(capdropped, p)
			}
		}
	}
	return capdropped

}

func HostPorts(options *pflag.FlagSet) []Finding {
	var hostports []Finding
	pods := connectWithPods(options)
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			//Does the container have ports specified
			cports := container.Ports != nil
			if cports {
				for _, port := range container.Ports {
					// Is the port a host port
					if port.HostPort != 0 {
						p := Finding{Check: "Host Ports", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name, Hostport: int(port.HostPort)}
						hostports = append(hostports, p)
					}
				}
			}
		}
		for _, init_container := range pod.Spec.InitContainers {
			cports := init_container.Ports != nil
			if cports {
				for _, port := range init_container.Ports {
					if port.HostPort != 0 {
						p := Finding{Check: "Host Ports", Namespace: pod.Namespace, Pod: pod.Name, Container: init_container.Name, Hostport: int(port.HostPort)}
						hostports = append(hostports, p)
					}
				}
			}
		}
		for _, eph_container := range pod.Spec.EphemeralContainers {
			cports := eph_container.Ports != nil
			if cports {
				for _, port := range eph_container.Ports {
					if port.HostPort != 0 {
						p := Finding{Check: "Host Ports", Namespace: pod.Namespace, Pod: pod.Name, Container: eph_container.Name, Hostport: int(port.HostPort)}
						hostports = append(hostports, p)
					}
				}
			}
		}
	}
	return hostports

}

func Seccomp(options *pflag.FlagSet) []Finding {
	var seccomp []Finding
	pods := connectWithPods(options)
	// The logic here is that if the pod is unconfined & the container is unconfined, it's unconfined.
	// In theory if all the containers in the pod are unconfined we could just mark it at pod level, but that's more complex :P
	for _, pod := range pods.Items {
		unconfined_pod := (pod.Spec.SecurityContext == nil) || (pod.Spec.SecurityContext != nil && pod.Spec.SecurityContext.SeccompProfile == nil) || (pod.Spec.SecurityContext != nil && pod.Spec.SecurityContext.SeccompProfile != nil && pod.Spec.SecurityContext.SeccompProfile.Type == "Unconfined")
		if unconfined_pod {
			//log.Printf("Pod name %s was unconfined at pod level", pod.Name)
			for _, container := range pod.Spec.Containers {
				unconfined_container := (container.SecurityContext == nil) || (container.SecurityContext != nil && container.SecurityContext.SeccompProfile == nil) || (container.SecurityContext != nil && container.SecurityContext.SeccompProfile != nil && container.SecurityContext.SeccompProfile.Type == "Unconfined")
				if unconfined_container {
					p := Finding{Check: "Seccomp Disabled", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name}
					seccomp = append(seccomp, p)
				}
			}
			for _, init_container := range pod.Spec.InitContainers {
				unconfined_init_container := (init_container.SecurityContext == nil) || (init_container.SecurityContext != nil && init_container.SecurityContext.SeccompProfile == nil) || (init_container.SecurityContext != nil && init_container.SecurityContext.SeccompProfile != nil && init_container.SecurityContext.SeccompProfile.Type == "Unconfined")
				if unconfined_init_container {
					p := Finding{Check: "Seccomp Disabled", Namespace: pod.Namespace, Pod: pod.Name, Container: init_container.Name}
					seccomp = append(seccomp, p)
				}
			}
			for _, eph_container := range pod.Spec.EphemeralContainers {
				unconfined_eph_container := (eph_container.SecurityContext == nil) || (eph_container.SecurityContext != nil && eph_container.SecurityContext.SeccompProfile == nil) || (eph_container.SecurityContext != nil && eph_container.SecurityContext.SeccompProfile != nil && eph_container.SecurityContext.SeccompProfile.Type == "Unconfined")
				if unconfined_eph_container {
					p := Finding{Check: "Seccomp Disabled", Namespace: pod.Namespace, Pod: pod.Name, Container: eph_container.Name}
					seccomp = append(seccomp, p)
				}
			}
		}
	}
	return seccomp

}

func HostPath(options *pflag.FlagSet) []Finding {
	var hostpath []Finding
	pods := connectWithPods(options)
	for _, pod := range pods.Items {
		host_path := pod.Spec.Volumes != nil
		if host_path {
			for _, vol := range pod.Spec.Volumes {
				if vol.HostPath != nil {
					p := Finding{Check: "Host Path", Namespace: pod.Namespace, Pod: pod.Name, Volume: vol.Name, Path: vol.HostPath.Path}
					hostpath = append(hostpath, p)
				}
			}
		}
	}
	return hostpath

}

func Apparmor(options *pflag.FlagSet) []Finding {
	var apparmor []Finding
	pods := connectWithPods(options)
	for _, pod := range pods.Items {
		// Default should be apparmor is set (well it is for docker anyway), so we only care if it's explicitly set to unconfined
		if pod.Annotations != nil {
			for key, val := range pod.Annotations {
				if val == "unconfined" && strings.Split(key, "/")[0] == "container.apparmor.security.beta.kubernetes.io" {
					p := Finding{Check: "Apparmor Disabled", Namespace: pod.Namespace, Pod: pod.Name}
					apparmor = append(apparmor, p)
				}
			}
		}
	}
	return apparmor

}

func Procmount(options *pflag.FlagSet) []Finding {
	var unmaskedproc []Finding
	pods := connectWithPods(options)
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			unmask := container.SecurityContext != nil && container.SecurityContext.ProcMount != nil && *container.SecurityContext.ProcMount == "Unmasked"
			if unmask {
				p := Finding{Check: "Unmasked procmount", Namespace: pod.Namespace, Pod: pod.Name, Container: container.Name}
				unmaskedproc = append(unmaskedproc, p)
			}
		}
		for _, init_container := range pod.Spec.InitContainers {
			unmask := init_container.SecurityContext != nil && init_container.SecurityContext.ProcMount != nil && *init_container.SecurityContext.ProcMount == "Unmasked"
			if unmask {
				p := Finding{Check: "Unmasked procmount", Namespace: pod.Namespace, Pod: pod.Name, Container: init_container.Name}
				unmaskedproc = append(unmaskedproc, p)
			}
		}
		for _, eph_container := range pod.Spec.EphemeralContainers {
			unmask := eph_container.SecurityContext != nil && eph_container.SecurityContext.ProcMount != nil && *eph_container.SecurityContext.ProcMount == "Unmasked"
			if unmask {
				p := Finding{Check: "Unmasked procmount", Namespace: pod.Namespace, Pod: pod.Name, Container: eph_container.Name}
				unmaskedproc = append(unmaskedproc, p)
			}
		}
	}
	return unmaskedproc

}

func Sysctl(options *pflag.FlagSet) []Finding {
	var sysctls []Finding
	pods := connectWithPods(options)
	for _, pod := range pods.Items {
		sysctl := pod.Spec.SecurityContext != nil && pod.Spec.SecurityContext.Sysctls != nil
		if sysctl {
			for _, sys := range pod.Spec.SecurityContext.Sysctls {
				safe := []string{"kernel.shm_rmid_forced", "net.ipv4.ip_local_port_range", "net.ipv4.ip_unprivileged_port_start", "net.ipv4.tcp_syncookies", "net.ipv4.ping_group_range"}
				safe_sys := false
				for _, s := range safe {
					if sys.Name == s {
						safe_sys = true
					}
				}
				if !safe_sys {
					p := Finding{Check: "Unsafe Sysctl", Namespace: pod.Namespace, Sysctl: sys.Name}
					sysctls = append(sysctls, p)
				}
			}
		}
	}
	return sysctls

}
