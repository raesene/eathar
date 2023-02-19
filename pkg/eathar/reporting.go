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
	v1 "k8s.io/api/rbac/v1"
)

var style string = ` <style>
body {
	font: normal 14px;
	font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
	color: #C41230;
	background: #FFFFFF;
}
#kubernetes-analyzer {
	font-weight: bold;
	font-size: 48px;
	color: #C41230;
}
.master-node, .worker-node, .vuln-node {
	background: #F5F5F5;
	border: 1px solid black;
	padding-left: 6px;
}
#api-server-results {
	font-weight: italic;
	font-size: 36px;
	color: #C41230;
}
table, th, td {
	border-collapse: collapse;
	border: 1px solid black;
}
th {
 font: bold 11px;
 color: #C41230;
 background: #999999;
 letter-spacing: 2px;
 text-transform: uppercase;
 text-align: left;
 padding: 6px 6px 6px 12px;
}
td {
background: #FFFFFF;
padding: 6px 6px 6px 12px;
color: #333333;
}
.container{
	display: flex;
} 
.fixed{
	width: 300px;
}
.flex-item{
	flex-grow: 1;
}
</style>`

func ReportPSS(f []Finding, options *pflag.FlagSet, check string) {
	jsonrep, _ := options.GetBool("jsonrep")
	htmlrep, _ := options.GetBool("htmlrep")
	file, _ := options.GetString("file")
	switch {
	case jsonrep:
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
	case htmlrep:
		var rep *os.File
		if file != "" {
			rep, _ = os.OpenFile(file+".html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		} else {
			rep = os.Stdout
		}
		fmt.Fprintf(rep, "<html><head>%s<title>Findings for the %s check</title></head><body>", style, check)
		fmt.Fprintf(rep, "<h1>Findings for the %s check</h1>", check)
		if f != nil {
			fmt.Fprintln(rep, "<table>")
			switch f[0].Check {
			case "hostpid", "hostnet", "hostipc", "privileged", "allowprivesc", "HostProcess", "Seccomp Disabled", "Unmasked Procmount", "Apparmor Disabled":
				if f[0].Container != "" {
					fmt.Fprintf(rep, "<tr><th>Namespace</th><th>Pod</th><th>Container</th></tr>")
				} else {
					fmt.Fprintf(rep, "<tr><th>Namespace</th><th>Pod</th></tr>")
				}
			case "Added Capabilities":
				fmt.Fprintf(rep, "<tr><th>namespace</th><th>pod</th><th>container</th><th>added capabilities</th></tr>")
			case "Dropped Capabilities":
				fmt.Fprintf(rep, "<tr><th>namespace</th><th>pod</th><th>container</th><th>dropped capabilities</th></tr>")
			case "Host Ports":
				fmt.Fprintf(rep, "<tr><th>namespace</th><th>pod</th><th>container</th><th>port</th</tr>")
			case "Host Path":
				fmt.Fprintf(rep, "<tr><th>namespace</th><th>pod</th><th>volume</th><th>path</th></tr>")
			case "Unsafe Sysctl":
				fmt.Fprintf(rep, "<tr><th>namespace</th><th>pod</th><th>unsafe sysctl</th></tr>")
			}
			for _, i := range f {
				switch i.Check {
				case "hostpid", "hostnet", "hostipc", "privileged", "allowprivesc", "HostProcess", "Seccomp Disabled", "Unmasked Procmount", "Apparmor Disabled":
					if i.Container != "" {
						fmt.Fprintf(rep, "<tr><td>%s</td><td>%s</td><td>%s</td></tr>", i.Namespace, i.Pod, i.Container)
					} else {
						fmt.Fprintf(rep, "<tr><td>%s</td><td>%s</td></tr>", i.Namespace, i.Pod)
					}
				case "Added Capabilities":
					fmt.Fprintf(rep, "<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", i.Namespace, i.Pod, i.Container, strings.Join(i.Capabilities[:], ","))
				case "Dropped Capabilities":
					fmt.Fprintf(rep, "<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", i.Namespace, i.Pod, i.Container, strings.Join(i.Capabilities[:], ","))
				case "Host Ports":
					fmt.Fprintf(rep, "<tr><td>%s</td><td>%s</td><td>%s</td><td>%d</td></tr>", i.Namespace, i.Pod, i.Container, i.Hostport)
				case "Host Path":
					fmt.Fprintf(rep, "<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", i.Namespace, i.Pod, i.Volume, i.Path)
				case "Unsafe Sysctl":
					fmt.Fprintf(rep, "<tr><td>%s</td><td>%s</td><td>%s</td></tr>", i.Namespace, i.Pod, i.Sysctl)
				}
			}
			fmt.Fprintln(rep, "</table></body></html>")
		} else {
			fmt.Fprintln(rep, "<p>No findings</p>")
		}

	default:
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
					fmt.Fprintf(rep, "namespace %s : pod %s : container %s added capabilities %s \n", i.Namespace, i.Pod, i.Container, strings.Join(i.Capabilities[:], ","))
				case "Dropped Capabilities":
					fmt.Fprintf(rep, "namespace %s : pod %s : container %s dropped capabilities %s \n", i.Namespace, i.Pod, i.Container, strings.Join(i.Capabilities[:], ","))
				case "Host Ports":
					fmt.Fprintf(rep, "namespace %s : pod %s : container %s : port %d\n", i.Namespace, i.Pod, i.Container, i.Hostport)
				case "Host Path":
					fmt.Fprintf(rep, "namespace %s : pod %s : volume %s : path %s\n", i.Namespace, i.Pod, i.Volume, i.Path)
				case "Unsafe Sysctl":
					fmt.Fprintf(rep, "namespace %s : pod %s : unsafe sysctl %s", i.Namespace, i.Pod, i.Sysctl)

				}
			}
		} else {
			fmt.Fprintln(rep, "No findings!")
		}
		fmt.Fprintln(rep, "")
	}
}

func ReportPrincipal(f []string, options *pflag.FlagSet, check string) {
	jsonrep, _ := options.GetBool("jsonrep")
	htmlrep, _ := options.GetBool("htmlrep")
	file, _ := options.GetString("file")
	var rep *os.File
	switch {
	case jsonrep:
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
	case htmlrep:
		if file != "" {
			rep, _ = os.OpenFile(file+".html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		} else {
			rep = os.Stdout
		}
		fmt.Fprintf(rep, "<html><head>%s<title>%s</title></head><body><table><tr><th>Name</th></tr>", style, check)
		if f != nil {
			for _, i := range f {
				fmt.Fprintf(rep, "<tr><td>%s</td></tr>", i)
			}
		} else {
			fmt.Fprintln(rep, "<tr><td>No findings</td></tr>")
		}
		fmt.Fprintln(rep, "</table></body></html>")
	default:
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
	}
}

func ReportImage(f []string, options *pflag.FlagSet, check string) {
	jsonrep, _ := options.GetBool("jsonrep")
	htmlrep, _ := options.GetBool("htmlrep")
	file, _ := options.GetString("file")

	var rep *os.File
	switch {
	case jsonrep:
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
	case htmlrep:
		if file != "" {
			rep, _ = os.OpenFile(file+".html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		} else {
			rep = os.Stdout
		}
		fmt.Fprintf(rep, "<html><head>%s<title>Image List</title></head><body><table><tr><th>Image</th></tr>", style)
		if f != nil {
			for _, i := range f {
				fmt.Fprintf(rep, "<tr><td>%s</td></tr>", i)
			}
		} else {
			fmt.Fprintln(rep, "<tr><td>No findings</td></tr>")
		}
		fmt.Fprintln(rep, "</table></body></html>")
	default:
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
	}
}

func ReportRBAC(f v1.ClusterRoleBindingList, options *pflag.FlagSet, check string) {
	jsonrep, _ := options.GetBool("jsonrep")
	htmlrep, _ := options.GetBool("htmlrep")
	file, _ := options.GetString("file")

	var rep *os.File
	switch {
	case jsonrep:
		if file != "" {
			rep, _ = os.OpenFile(file+".json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		} else {
			rep = os.Stdout
		}
		if f.Items != nil {
			js, err := json.MarshalIndent(f, "", "  ")
			if err != nil {
				log.Print(err)
			}
			fmt.Fprintln(rep, string(js))
		}
	case htmlrep:
		if file != "" {
			rep, _ = os.OpenFile(file+".html", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		} else {
			rep = os.Stdout
		}
		fmt.Fprintf(rep, "<html><head>%s<title>RBAC Report</title></head><body>", style)
		if f.Items != nil {
			fmt.Fprintf(rep, "<table><tr><th>ClusterRoleBinding</th><th>Subjects</th><th>Role Ref</th></tr>")
			for _, i := range f.Items {
				fmt.Fprintf(rep, "<tr><td>%s</td>", i.Name)
				for _, s := range i.Subjects {
					if s.Kind == "ServiceAccount" {
						fmt.Fprintf(rep, "<td>Kind: %s, Name: %s, Namespace: %s</td>", s.Kind, s.Name, s.Namespace)
					} else {
						fmt.Fprintf(rep, "<td>Kind: %s, Name: %s</td>", s.Kind, s.Name)
					}
				}
				fmt.Fprintf(rep, "<td>Kind: %s, Name: %s</td></tr>", i.RoleRef.Kind, i.RoleRef.Name)
			}
			fmt.Fprintln(rep, "</table></body></html>")
		} else {
			fmt.Fprintln(rep, "No findings!")
		}
	default:
		if file != "" {
			rep, _ = os.OpenFile(file+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		} else {
			rep = os.Stdout
		}
		fmt.Fprintf(rep, "Findings for the %s check\n", check)
		if f.Items != nil {
			for _, i := range f.Items {
				fmt.Fprintf(rep, "ClusterRoleBinding %s\n", i.Name)
				fmt.Fprintf(rep, "Subjects:\n")
				for _, s := range i.Subjects {
					if s.Kind == "ServiceAccount" {
						fmt.Fprintf(rep, "  Kind: %s, Name: %s, Namespace: %s\n", s.Kind, s.Name, s.Namespace)
					} else {
						fmt.Fprintf(rep, "  Kind: %s, Name: %s\n", s.Kind, s.Name)
					}
				}
				fmt.Fprintf(rep, "RoleRef:\n")
				fmt.Fprintf(rep, "  Kind: %s, Name: %s, APIGroup: %s\n", i.RoleRef.Kind, i.RoleRef.Name, i.RoleRef.APIGroup)
				fmt.Fprintln(rep, "------------------------")
			}
		}
	}
}
