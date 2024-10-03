## Filecoin Address to Ethereum (EVM) Address Converter
This Go program decodes Filecoin addresses (specifically Delegated Addresses in Protocol 4) into their corresponding Ethereum EVM addresses. The program extracts the last 20 bytes of the Filecoin address payload to generate an Ethereum-compatible address.

## Table of Contents

- Introduction
- Installation
- Usage
- Code Structure
- Address Protocols
- Troubleshooting

**Introduction** <br>
- Filecoin and Ethereum use different address formats, and this project helps convert Filecoin Delegated Addresses (Protocol 4) into Ethereum EVM addresses. The goal is to take a Filecoin address, decode its payload, and extract the last 20 bytes as the Ethereum address.

**Example** <br>
Given a Filecoin address: <br>
```
f410f2oekwcmo2pueydmaq53eic2i62crtbeyuzx2gmy
```

The program will output the corresponding Ethereum address: <br>
```
0xd388ab098ed3e84c0d808776440b48f685198498
```
**Installation** <br>

Prerequisites <br>
Go 1.16+ installed on your machine.
Internet connection to install the required Go packages.

**Steps:**

Clone the repository:
```
git clone https://github.com/yourusername/filecoin-to-evm-converter.git
cd filecoin-to-evm-converter
```

***Install the dependencies:***
```
go mod tidy
```
This will fetch the go-address package from the Filecoin project.

If not already installed, you can install the go-address library manually:
```
go get github.com/filecoin-project/go-address
```

**Usage:**
To run the program and convert a Filecoin address to an Ethereum EVM address, follow these steps:

- Modify the filecoinAddress variable in main.go with your desired Filecoin address.

- Run the program:
```
go run main.go
```

- The output will display the corresponding Ethereum EVM address.

**Example:**
```
go run main.go
```

**Output:**
```
EVM Address: 0xd388ab098ed3e84c0d808776440b48f685198498
```

**Code Structure:**<br>
- main.go: Contains the core logic to parse the Filecoin address, decode it, and extract the EVM address.
- go.mod: Defines the dependencies for the project, including the Filecoin go-address package.
  
***Core Function*** <br>
The function filecoinToEVM(filecoinAddress string) performs the conversion:

```
Input: A Filecoin address (e.g., t410...).
Output: An Ethereum EVM address (e.g., 0x...).
```

**Address Protocols** <br>
This program is designed to handle Filecoin Delegated Addresses (Protocol 4). The address type is identified by the prefix (t4 for testnet, f4 for mainnet).

**Protocol 4 (Delegated Address):** <br>
The address wraps an EVM-compatible account, and the last 20 bytes of the address payload represent the Ethereum address.
If you encounter an address with a different protocol, the program will return an error message.

## Troubleshooting

**Common Issues** <br>
- "Failed to parse Filecoin address": Ensure that the provided address is valid and belongs to Protocol 4.
- "Payload too short to contain an EVM address": The payload must contain at least 20 bytes to represent a valid Ethereum address. Verify the structure of your Filecoin address.
- "Unsupported Protocols": Currently, this program supports only Delegated (Protocol 4) addresses. You may extend the program to support other protocols.
