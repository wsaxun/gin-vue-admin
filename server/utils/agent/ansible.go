package agent

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
)

const playbookTemplate = `
---
- hosts: all
  vars:
    - src_script: SRC_SCRIPT_PATH
    - dest_script: DEST_SCRIPT_PATH	
  gather_facts: false
  tasks:
    - name: copy
      copy:
        src: "{{ src_script }}"
        dest: "{{ dest_script }}"
        mode: '0744'
    - name: exec
      shell: "{{ dest_script }}"
    - name: rm
      file:
        path: "{{ dest_script }}"
        state: absent
`

func createFile(filename, context string) (string, error) {
	TmpFile, err := os.CreateTemp("", filename)
	if err != nil {
		return "", err
	}

	if _, err := TmpFile.WriteString(context); err != nil {
		return "", err
	}
	return TmpFile.Name(), nil
}

func GetPlayBookTemplate() string {
	content := playbookTemplate
	return content
}

func RunPlaybook(scriptContent string, hosts []string, timeout int) error {
	hostStr := ""
	if len(hosts) == 1 {
		hostStr = hosts[0] + ","
	} else {
		hostStr = strings.Join(hosts, ",")
	}
	sum := fmt.Sprintf("%x", md5.Sum([]byte(scriptContent)))
	fmt.Println(sum)
	fileShell, err := createFile(sum, scriptContent)
	if err != nil {
		return err
	}
	ymlContent := GetPlayBookTemplate()
	tmp := strings.ReplaceAll(ymlContent, "SRC_SCRIPT_PATH", fileShell)
	yaml := strings.ReplaceAll(tmp, "DEST_SCRIPT_PATH", fmt.Sprintf("/tmp/%s", sum))
	fileYaml, err := createFile("ansible-playbook.*.yml", yaml)
	if err != nil {
		return err
	}
	defer os.RemoveAll(fileShell)
	defer os.RemoveAll(fileYaml)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	buff := new(bytes.Buffer)

	playbooksList := []string{fileYaml}
	playbookCmd := &playbook.AnsiblePlaybookCmd{
		Playbooks: playbooksList,
		Exec: execute.NewDefaultExecute(
			execute.WithWrite(io.Writer(buff)),
		),
		Options: &playbook.AnsiblePlaybookOptions{
			Inventory: hostStr,
		},
		StdoutCallback: "json",
	}

	err = playbookCmd.Run(ctx)
	if err != nil {
		return err
	}
	res, err := results.ParseJSONResultsStream(io.Reader(buff))
	if err != nil {
		return err
	}
	for _, play := range res.Plays {
		for _, task := range play.Tasks {
			if task.Task.Name == "exec" {
				for _, content := range task.Hosts {
					fmt.Println(content.Stdout)
				}
			}
		}
	}

	return nil
}
