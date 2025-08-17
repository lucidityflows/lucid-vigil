package block_ip

import (
	"context"
	"fmt"
	"net"
	os/exec"

	"github.com/rs/zerolog/log"
)

// BlockIPAction implements the actions.Action interface for blocking an IP address.
type BlockIPAction struct{}

// Name returns the name of the action.
func (bia *BlockIPAction) Name() string {
	return "block_ip"
}

// Execute attempts to block an IP address using iptables.
// The `data` map is expected to contain an "ip" key with the IP address as a string.
func (bia *BlockIPAction) Execute(ctx context.Context, data map[string]interface{}) error {
	ip, ok := data["ip"].(string)
	if !ok || ip == "" {
		return fmt.Errorf("missing or invalid 'ip' in action data for block_ip action")
	}

	// Validate IP address format
	if net.ParseIP(ip) == nil {
		return fmt.Errorf("invalid IP address format: %s", ip)
	}

	log.Info().Str("ip", ip).Msg("Attempting to block IP using iptables...")

	// Command to add a rule to block the IP in the INPUT chain
	cmd := exec.CommandContext(ctx, "sudo", "iptables", "-A", "INPUT", "-s", ip, "-j", "DROP")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to block IP %s: %w\nOutput: %s", ip, err, string(out))
	}

	log.Info().Str("ip", ip).Msg("Successfully blocked IP using iptables.")
	return nil
}
