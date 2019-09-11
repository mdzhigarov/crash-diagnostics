package exec

import (
	"fmt"

	"gitlab.eng.vmware.com/vivienv/flare/script"
)

func exeSSH(src *script.Script) (*script.SSHConfigCommand, error) {
	sshCmds, ok := src.Preambles[script.CmdSSHConfig]
	if !ok {
		return nil, fmt.Errorf("Script missing valid %s", script.CmdSSHConfig)
	}
	sshCmd := sshCmds[0].(*script.SSHConfigCommand)
	return sshCmd, nil
}
