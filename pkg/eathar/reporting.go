package eathar

/*
Copyright © 2022 Rory McCune <rorym@mccune.org.uk>

*/

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
)

func reportPSS(f []Finding, options *pflag.FlagSet, check string) {
	jsonrep, _ := options.GetBool("jsonrep")
	file, _ := options.GetString("file")

	if !jsonrep {
		var rep *os.File
		if file != "" {
			rep, _ = os.OpenFile(file+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		} else {
			rep = os.Stdout
		}

		fmt.Fprintf(rep, "Findings for the %s check\n", check)
		if f != nil {
			for _, i := range f {
				switch i.Check {
				case "hostpid", "hostnet", "hostipc", "privileged", "allowprivesc", "HostProcess", "Seccomp Disabled", "Unmasked Procmount", "Apparmor Disabled":
					if i.Container != "" {
						fmt.Fprintf(rep, "namespace %s : pod %s : container %s\n", i.Namespace, i.Pod, i.Container)
					} else {
						fmt.Fprintf(rep, "namespace %s : pod %s\n", i.Namespace, i.Pod)
					}
				case "Added Capabilities":
					fmt.Fprintf(rep, "namespace %s : pod %s : container %s added %s capabilities\n", i.Namespace, i.Pod, i.Container, strings.Join(i.Capabilities[:], ","))
				case "Dropped Capabilities":
					fmt.Fprintf(rep, "namespace %s : pod %s : container %s dropped %s capabilities\n", i.Namespace, i.Pod, i.Container, strings.Join(i.Capabilities[:], ","))
				case "Host Ports":
					fmt.Fprintf(rep, "namespace %s : pod %s : container %s : port %d\n", i.Namespace, i.Pod, i.Container, i.Hostport)
				case "Host Path":
					fmt.Fprintf(rep, "namespace %s : pod %s : volume %s : path %s\n", i.Namespace, i.Pod, i.Volume, i.Path)
				case "Unsafe Sysctl":
					fmt.Fprintf(rep, "namespace %s : pod %s : unsafe sysctl %s", i.Namespace, i.Pod, i.Sysctl)
				case "Image List":
					//fmt.Fprintf(rep, "namespace %s : pod %s : container %s : image %s\n", i.Namespace, i.Pod, i.Container, i.Image)
					//Let's just print the unique image name
					fmt.Fprintf(rep, "%s\n", i.Image)

				}
			}
		} else {
			fmt.Fprintln(rep, "No findings!")
		}
		fmt.Fprintln(rep, "")
	} else {
		var rep *os.File
		if file != "" {
			rep, _ = os.OpenFile(file+".json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		} else {
			rep = os.Stdout
		}
		if f != nil {
			js, err := json.MarshalIndent(f, "", "  ")
			if err != nil {
				log.Print(err)
			}
			fmt.Fprintln(rep, string(js))
		}
	}

}

func reportImage(f []string, options *pflag.FlagSet, check string) {
	jsonrep, _ := options.GetBool("jsonrep")
	file, _ := options.GetString("file")

	if !jsonrep {
		var rep *os.File
		if file != "" {
			rep, _ = os.OpenFile(file+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		} else {
			rep = os.Stdout
		}

		fmt.Fprintf(rep, "Findings for the %s check\n", check)
		if f != nil {
			for _, i := range f {
				fmt.Fprintf(rep, "%s\n", i)
			}
		} else {
			fmt.Fprintln(rep, "No findings!")
		}
		fmt.Fprintln(rep, "")
	} else {
		var rep *os.File
		if file != "" {
			rep, _ = os.OpenFile(file+".json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		} else {
			rep = os.Stdout
		}
		if f != nil {
			js, err := json.MarshalIndent(f, "", "  ")
			if err != nil {
				log.Print(err)
			}
			fmt.Fprintln(rep, string(js))
		}
	}
}
