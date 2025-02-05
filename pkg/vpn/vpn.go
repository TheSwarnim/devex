package vpn

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/pquerna/otp/totp"
	"gopkg.in/yaml.v2"
)

// Config is used to define the structure of config file
type Config struct {
	VPNs         map[string]VPNConfig `yaml:"vpns"`
	ClientCLIPath string              `yaml:"client_cli_path"`
}

// VPNConfig holds the configuration for a VPN
type VPNConfig struct {
	Type     string `yaml:"type"`
	VPNID    string `yaml:"vpn_id"`
	PIN      string `yaml:"pin"`
	AuthKey  string `yaml:"auth_key"`
}

// loadConfig reads the config from the yaml file
func loadConfig() (*Config, error) {
	configPath := filepath.Join(os.Getenv("HOME"), ".devex", "config.yaml")
	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{VPNs: make(map[string]VPNConfig)}, nil
		}
		return nil, err
	}
	defer file.Close()
	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// saveConfig writes the config to the yaml file
func saveConfig(config *Config) error {
	configPath := filepath.Join(os.Getenv("HOME"), ".devex", "config.yaml")
	os.MkdirAll(filepath.Dir(configPath), os.ModePerm)
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := yaml.NewEncoder(file)
	return encoder.Encode(config)
}

// generateTOTP generates a TOTP using the secret for a given VPN
func generateTOTP(vpnName string) (string, error) {
	config, err := loadConfig()
	if err != nil {
		return "", err
	}
	vpnConfig, exists := config.VPNs[vpnName]
	if !exists {
		return "", errors.New("VPN not found")
	}
	if vpnConfig.Type != "pritunl" {
		return "", errors.New("VPN type not supported")
	}
	return totp.GenerateCode(vpnConfig.AuthKey, time.Now())
}

// Connect connects to a given VPN by name
func Connect(vpnName string) error {
	if vpnName == "" {
		return errors.New("VPN name is required")
	}
	config, err := loadConfig()
	if err != nil {
		return err
	}
	vpnConfig, exists := config.VPNs[vpnName]
	if !exists {
		return errors.New("VPN not found")
	}
	if vpnConfig.Type != "pritunl" {
		return errors.New("VPN type not supported")
	}
	otp, err := generateTOTP(vpnName)
	if err != nil {
		return err
	}
	clientCLIPath := config.ClientCLIPath
	if clientCLIPath == "" {
		clientCLIPath = "/Applications/Pritunl.app/Contents/Resources/pritunl-client"
	}
	password := vpnConfig.PIN + otp
	cmd := exec.Command(clientCLIPath, "start", vpnConfig.VPNID, "--password", password)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	return err
}

// Disconnect disconnects from a given VPN by name, or all if name is empty
func Disconnect(vpnName string) error {
	config, err := loadConfig()
	if err != nil {
		return err
	}
	clientCLIPath := config.ClientCLIPath
	if clientCLIPath == "" {
		clientCLIPath = "/Applications/Pritunl.app/Contents/Resources/pritunl-client"
	}
	if vpnName == "" {
		return errors.New("VPN name is required")
	} else {
		vpnConfig, exists := config.VPNs[vpnName]
		if !exists {
			return errors.New("VPN not found")
		}
		if vpnConfig.Type != "pritunl" {
			return errors.New("VPN type not supported")
		}
		cmd := exec.Command(clientCLIPath, "stop", vpnConfig.VPNID)
		output, err := cmd.CombinedOutput()
		fmt.Println(string(output))
		return err
	}
}

// AddVPN adds a new VPN configuration with the provided name and details
func AddVPN(vpnName, vpnType, vpnID, pin, authKey string) error {
	config, err := loadConfig()
	if err != nil {
		return err
	}
	if _, exists := config.VPNs[vpnName]; exists {
		return errors.New("VPN already exists")
	}
	config.VPNs[vpnName] = VPNConfig{
		Type:    vpnType,
		VPNID:   vpnID,
		PIN:     pin,
		AuthKey: authKey,
	}
	return saveConfig(config)
}

// SetClientCLIPath sets the client CLI path in the config file
func SetClientCLIPath(path string) error {
	config, err := loadConfig()
	if err != nil {
		return err
	}
	config.ClientCLIPath = path
	return saveConfig(config)
}
