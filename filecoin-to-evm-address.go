package filecoin_to_evm_address

import (
	"encoding/hex"
	"fmt"
	"github.com/filecoin-project/go-address"
)

func FilecoinToEVM(filecoinAddress string) (string, error) {
	// Parse the Filecoin address
	addr, err := address.NewFromString(filecoinAddress)
	if err != nil {
		return "", fmt.Errorf("failed to parse Filecoin address: %w", err)
	}

	// Check if the address is a Delegated Address (Protocol 4)
	if addr.Protocol() != address.Delegated {
		return "", fmt.Errorf("unsupported address protocol: %d", addr.Protocol())
	}

	// Extract the payload
	payload := addr.Payload()

	// Depending on how the EVM address is embedded, you might need to process the payload.
	// For this example, we'll assume the payload directly contains the EVM address as the last 20 bytes.
	if len(payload) < 20 {
		return "", fmt.Errorf("payload too short to contain an EVM address")
	}

	// Extract the last 20 bytes for the EVM address
	evmAddressBytes := payload[len(payload)-20:]

	// Convert to hexadecimal string and prepend '0x'
	evmAddress := "0x" + hex.EncodeToString(evmAddressBytes)

	return evmAddress, nil
}
