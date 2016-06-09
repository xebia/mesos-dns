package main // import "github.com/xebia/mesos-dns"

import "syscall"
import "encoding/json"
import "fmt"
import "os"
import "log"
import "strconv"
import "regexp"
import "strings"
import "io/ioutil"

type MesosDNSConfig struct {
  Zk string `json:"zk,omitempty"`
  Masters []string `json:"masters,omitempty"`
  RefreshSeconds int `json:"refreshSeconds,omitempty"`
  TTL int `json:"ttl,omitempty"`
  Domain string `json:"domain,omitempty"`
  Port int `json:"port,omitempty"`
  Resolvers []string `json:"resolvers,omitempty"`
  Timeout int `json:"timeout,omitempty" `
  Httpon bool `json:"httpon,omitempty"`
  Dnson bool `json:"dnson,omitempty"`
  Httpport int `json:"httpport,omitempty"`
  Externalon bool `json:"externalon,omitempty"`
  Listener string `json:"listener,omitempty"`
  SOAMname string `json:"SOAMname,omitempty"`
  SOARname string `json:"SOARname,omitempty"`
  SOARefresh int `json:"SOARefresh,omitempty"`
  SOARetry int `json:"SOARetry,omitempty"`
  SOAExpire int `json:"SOAExpire,omitempty"`
  SOAMinTTL int `json:"SOAMinttl,omitempty"`
  IPSources []string `json:"IPSources,omitempty"`
  Recurseon bool `json:"recurseon,omitempty"`
  EnforceRFC952 bool `json:"enforceRFC952,omitempty"`
}

func set_bool_from_env(value *bool, name string) {
	if v := os.Getenv(name) ; v != "" {
		if i, err := strconv.ParseBool(v); err == nil {
			*value = i
		} else {
			log.Fatal(fmt.Sprintf("%s is not a valid boolean value for %s", v, name))
		}
	}
}

func set_string_from_env(value *string, name string) {
	if v := os.Getenv(name) ; v != "" {
		*value = v
	}
}

func set_string_array_from_env(value *[]string, name string) {
	if v := os.Getenv(name) ; v != "" {
		*value = regexp.MustCompile("[[:space:]]+").Split(v, 32767)
	}
}

func set_int_from_env(value *int, name string) {
	if v := os.Getenv(name) ; v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			*value = i
		} else {
			log.Fatal(fmt.Sprintf("%s is not a valid integer value for %s", v, name))
		}
	}
}

func main() {
	config := MesosDNSConfig{}
	set_string_from_env(&config.Zk, "MESOS_DNS_ZK")
	set_string_array_from_env(&config.Masters, "MESOS_DNS_MASTERS")
	set_string_from_env(&config.Domain, "MESOS_DNS_DOMAIN")
	set_int_from_env(&config.RefreshSeconds, "MESOS_DNS_REFRESH_SECONDS")
	set_int_from_env(&config.TTL, "MESOS_DNS_TTL")
	set_int_from_env(&config.Port, "MESOS_DNS_PORT")
	set_string_array_from_env(&config.Resolvers, "MESOS_DNS_RESOLVERS")
	set_int_from_env(&config.Timeout, "MESOS_DNS_TIMEOUT")
	set_bool_from_env(&config.Httpon, "MESOS_DNS_HTTP_ON")
	set_bool_from_env(&config.Dnson, "MESOS_DNS_DNS_ON")
	set_int_from_env(&config.Httpport, "MESOS_DNS_HTTP_PORT")
	set_bool_from_env(&config.Externalon, "MESOS_DNS_EXTERNAL_ON")
	set_string_from_env(&config.Listener, "MESOS_DNS_LISTENER")
	set_string_from_env(&config.SOAMname, "MESOS_DNS_SOA_MNAME")
	set_string_from_env(&config.SOARname, "MESOS_DNS_SOA_RNAME")
	set_int_from_env(&config.SOARefresh, "MESOS_DNS_SOA_REFRESH")
	set_int_from_env(&config.SOARetry, "MESOS_DNS_SOA_RETRY")
	set_int_from_env(&config.SOAExpire, "MESOS_DNS_SOA_EXPIRE")
	set_int_from_env(&config.SOAMinTTL, "MESOS_DNS_MIN_TTL")
	set_string_array_from_env(&config.IPSources, "MESOS_DNS_IP_SOURCES")
	set_bool_from_env(&config.Recurseon, "MESOS_DNS_RECURSE_ON")
	set_bool_from_env(&config.EnforceRFC952, "MESOS_DNS_ENFORCE_RFC952")

	args := os.Args[1:]
	filename := ""
	for i, arg := range args {
		if strings.HasPrefix(arg, "-config") {
			if arg == "-config" &&  i + 1 < len(args) {
				filename = args[i+1]
		        } else if strings.HasPrefix(arg, "-config=") { 
				filename = arg[8:]
				fmt.Println(filename)
			}
		}
	}

	if filename == "" {
		log.Fatal("-f option is missing")
	}

	str, err := json.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}
		
	log.Println(string(str))
	err = ioutil.WriteFile(filename, str, 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = syscall.Exec(args[0], args, os.Environ())
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to start %s, %s", args[0], err))
	}
}
