// Credit to mattermost for this container code

package storetest

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os/exec"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/yogyrahmawan/grpc_logger_service/src/utils"
)

// Container represent container struct
type Container struct {
	ID              string
	NetworkSettings struct {
		Ports map[string][]struct {
			HostPort string
		}
	}
}

// RunningContainer represent running docker container
type RunningContainer struct {
	Container
}

// Stop stopping running container
func (c *RunningContainer) Stop() error {
	log.Info("Removing container: %v", c.ID)
	return exec.Command("docker", "rm", "-f", c.ID).Run()
}

// RunCustomCommand run docker command
func (c *RunningContainer) RunCustomCommand(args []string) error {
	log.Info("docker execute command")
	return exec.Command("docker", args...).Run()
}

// NewMongoDBContainer instantiate mongodb container
func NewMongoDBContainer() (*RunningContainer, string, error) {
	container, err := runContainer([]string{
		"-p", "30000:27017",
		"mongo",
	})
	if err != nil {
		return nil, "", err
	}
	log.Info("Waiting for mongodb connectivity")
	port := container.NetworkSettings.Ports["27017/tcp"][0].HostPort
	if err := waitForPort(port); err != nil {
		container.Stop()
		return nil, "", err
	}
	return container, "mongodb://localhost:" + port + "/integration_test", nil
}

func runContainer(args []string) (*RunningContainer, error) {
	name := "logger-storetest-" + utils.GenerateUUID()
	dockerArgs := append([]string{"run", "-d", "-P", "--name", name}, args...)
	out, err := exec.Command("docker", dockerArgs...).Output()
	if err != nil {
		return nil, err
	}
	id := strings.TrimSpace(string(out))
	out, err = exec.Command("docker", "inspect", id).Output()
	if err != nil {
		exec.Command("docker", "rm", "-f", id).Run()
		return nil, err
	}
	var containers []Container
	if err := json.Unmarshal(out, &containers); err != nil {
		exec.Command("docker", "rm", "-f", id).Run()
		return nil, err
	}
	log.Info("Running container: %v", id)
	return &RunningContainer{containers[0]}, nil
}

func waitForPort(port string) error {
	deadline := time.Now().Add(time.Minute * 10)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", "127.0.0.1:"+port, time.Minute)
		if err != nil {
			return err
		}
		if err = conn.SetReadDeadline(time.Now().Add(time.Millisecond * 500)); err != nil {
			return err
		}
		_, err = conn.Read(make([]byte, 1))
		conn.Close()
		if err == nil {
			return nil
		}
		if e, ok := err.(net.Error); ok && e.Timeout() {
			return nil
		}
		if err != io.EOF {
			return err
		}
		time.Sleep(time.Millisecond * 200)
	}
	return fmt.Errorf("timeout waiting for port %v", port)
}
