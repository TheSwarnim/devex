package cmd

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/theswarnim/devex/pkg/vpn"
	"golang.org/x/crypto/ssh/terminal"
)

var clientCLIPath string

func init() {
	vpnCmd := &cobra.Command{
		Use:   "vpn",
		Short: "Manage VPN connections.",
	}

	connectCmd := &cobra.Command{
		Use:   "connect <vpn_name>",
		Short: "Connect to a VPN",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vpnName := args[0]
			err := vpn.Connect(vpnName)
			if err != nil {
				fmt.Println("Error connecting to VPN:", err)
			} else {
				fmt.Println("Connected to VPN:", vpnName)
			}
		},
	}

	disconnectCmd := &cobra.Command{
		Use:   "disconnect [vpn_name]",
		Short: "Disconnect from VPN",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var vpnName string
			if len(args) > 0 {
				vpnName = args[0]
			}
			err := vpn.Disconnect(vpnName)
			if err != nil {
				fmt.Println("Error disconnecting from VPN:", err)
			} else if vpnName == "" {
				fmt.Println("Disconnected from all VPNs.")
			} else {
				fmt.Println("Disconnected from VPN:", vpnName)
			}
		},
	}

	addCmd := &cobra.Command{
		Use:   "add <vpn_name>",
		Short: "Add a new VPN configuration.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			vpnName := args[0]
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Print("Enter the VPN type (default: pritunl): ")
			scanner.Scan()
			vpnType := scanner.Text()
			if vpnType == "" {
				vpnType = "pritunl"
			}

			if vpnType != "pritunl" {
				fmt.Println("VPN type not supported")
				return
			}

			fmt.Print("Enter the VPN ID: ")
			scanner.Scan()
			vpnID := scanner.Text()

			fmt.Print("Enter the PIN: ")
			pinBytes, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				fmt.Println("Error reading PIN:", err)
				return
			}
			pin := string(pinBytes)
			fmt.Println()

			fmt.Print("Enter the two-step authentication key: ")
			authKeyBytes, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				fmt.Println("Error reading authentication key:", err)
				return
			}
			authKey := string(authKeyBytes)
			fmt.Println()

			err = vpn.AddVPN(vpnName, vpnType, vpnID, pin, authKey)
			if err != nil {
				fmt.Println()
				fmt.Println("Error adding VPN:", err)
			} else {
				fmt.Println("VPN added successfully:", vpnName)
			}
		},
	}

	setClientCLICmd := &cobra.Command{
		Use:   "set-client-cli <path_to_pritunl_client_cli>",
		Short: "Set the Pritunl client CLI location.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			clientCLIPath = args[0]
			err := vpn.SetClientCLIPath(clientCLIPath)
			if err != nil {
				fmt.Println("Error setting client CLI path:", err)
			} else {
				fmt.Println("Client CLI path set successfully:", clientCLIPath)
			}
		},
	}

	vpnCmd.AddCommand(connectCmd)
	vpnCmd.AddCommand(disconnectCmd)
	vpnCmd.AddCommand(addCmd)
	vpnCmd.AddCommand(setClientCLICmd)
	rootCmd.AddCommand(vpnCmd)
}
