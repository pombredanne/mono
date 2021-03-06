package onepassword

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/peco/peco"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"

	"go.f110.dev/mono/go/pkg/logger"
	"go.f110.dev/mono/go/pkg/opvault"
)

const (
	ConfigDirName = ".1p"
)

var subcommands = []func(command *cobra.Command){
	Daemon,
	Shutdown,
	UseVault,
	Info,
	Unlock,
	List,
	Get,
}

func Main() error {
	client, err := dial()
	if err != nil && err == ErrDaemonNotExist {
		cmd := exec.Command(os.Args[0], "daemon")
		if err := cmd.Run(); err != nil {
			return xerrors.Errorf(": %w", err)
		}
		time.Sleep(100 * time.Millisecond)
		client, err = dial()
		if err != nil {
			return xerrors.Errorf(": %w", err)
		}
	} else if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	info, err := client.Info(ctx, &RequestInfo{})
	cancel()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if info.Locked {
		cmd := exec.Command(os.Args[0], "unlock")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			return xerrors.Errorf(": %w", err)
		}
	}

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	list, err := client.List(ctx, &RequestList{})
	cancel()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	sort.Slice(list.Items, func(i, j int) bool {
		return list.Items[i].Title < list.Items[j].Title
	})
	input := new(bytes.Buffer)
	for _, v := range list.Items {
		if opvault.Category(v.Category) != opvault.CategoryLogin {
			continue
		}
		fmt.Fprintf(input, "%s %s\n", v.Uuid, v.Title)
	}
	selector := peco.New()
	selector.Stdin = input
	selector.Run(context.Background())

	selected, err := selector.CurrentLineBuffer().LineAt(selector.Location().LineNumber())
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	s := strings.SplitN(selected.DisplayString(), " ", 2)
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	_, err = client.SetClipboard(ctx, &RequestSetClipboard{Uuid: s[0]})
	cancel()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return nil
}

func Daemon(rootCmd *cobra.Command) {
	daemonize := false
	foreground := false
	daemonCmd := &cobra.Command{
		Use: "daemon",
		RunE: func(_ *cobra.Command, _ []string) error {
			if !foreground {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					return xerrors.Errorf(": %w", err)
				}
				socketFile := filepath.Join(homeDir, ConfigDirName, socketFilename)
				if _, err := os.Stat(socketFile); !os.IsNotExist(err) {
					_, err := dial()
					if err != nil {
						os.Remove(socketFile)
					} else {
						// Already running
						logger.Log.Debug("The daemon is already running")
						return nil
					}
				}

				if !daemonize {
					cmd := exec.Command(os.Args[0], "daemon", "--daemonize")
					if err := cmd.Start(); err != nil {
						return xerrors.Errorf(": %w", err)
					}
					pid := cmd.Process.Pid
					if err := ioutil.WriteFile(filepath.Join(homeDir, ConfigDirName, "1p.pid"), []byte(strconv.Itoa(pid)), 0644); err != nil {
						return xerrors.Errorf(": %w", err)
					}
					return nil
				}

				defer func() {
					os.Remove(filepath.Join(homeDir, ConfigDirName, "1p.pid"))
				}()
			}

			d := NewDaemon()
			return d.Loop()
		},
	}
	daemonCmd.Flags().BoolVar(&daemonize, "daemonize", false, "Daemonize")
	daemonCmd.Flags().BoolVar(&foreground, "foreground", false, "")

	rootCmd.AddCommand(daemonCmd)
}

func Shutdown(rootCmd *cobra.Command) {
	shutdownCmd := &cobra.Command{
		Use: "shutdown",
		RunE: func(_ *cobra.Command, _ []string) error {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			pidFile := filepath.Join(homeDir, ConfigDirName, "1p.pid")
			if _, err := os.Stat(pidFile); os.IsNotExist(err) {
				return nil
			}

			buf, err := ioutil.ReadFile(pidFile)
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			pid, err := strconv.Atoi(string(buf))
			if err != nil {
				os.Remove(pidFile)
				return xerrors.Errorf(": %w", err)
			}
			if exists, err := process.PidExists(int32(pid)); err != nil {
				return xerrors.Errorf(": %w", err)
			} else if exists {
				proc, err := os.FindProcess(pid)
				if err != nil {
					return xerrors.Errorf(": %w", err)
				}
				proc.Signal(syscall.SIGTERM)
			} else {
				os.Remove(pidFile)
			}

			return nil
		},
	}
	rootCmd.AddCommand(shutdownCmd)
}

func UseVault(rootCmd *cobra.Command) {
	path := ""
	useVaultCmd := &cobra.Command{
		Use: "use-vault",
		RunE: func(_ *cobra.Command, _ []string) error {
			if path == "" {
				return xerrors.New("--path is mandatory")
			}

			client, err := dial()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			_, err = client.UseVault(ctx, &RequestUseVault{Path: path})
			cancel()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}

			return nil
		},
	}
	useVaultCmd.Flags().StringVar(&path, "path", "", "The path to opvault")

	rootCmd.AddCommand(useVaultCmd)
}

func Info(rootCmd *cobra.Command) {
	infoCmd := &cobra.Command{
		Use: "info",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := dial()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			res, err := client.Info(ctx, &RequestInfo{})
			cancel()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			fmt.Fprintf(os.Stdout, "Current opvault is %s\n", res.Path)
			if res.Locked {
				fmt.Fprintln(os.Stdout, "Vault is Locked")
			} else {
				fmt.Fprintln(os.Stdout, "Vault is Unlocked")
			}

			return nil
		},
	}

	rootCmd.AddCommand(infoCmd)
}

func Unlock(rootCmd *cobra.Command) {
	unlockCmd := &cobra.Command{
		Use: "unlock",
		RunE: func(_ *cobra.Command, _ []string) error {
			client, err := dial()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			info, err := client.Info(ctx, &RequestInfo{})
			cancel()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			if !info.Locked {
				fmt.Fprintln(os.Stdout, "Already unlocked")
				return nil
			}

			fmt.Printf("Master passowrd: ")
			masterPassword, err := term.ReadPassword(syscall.Stdin)
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			fmt.Println()
			ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
			res, err := client.Unlock(ctx, &RequestUnlock{MasterPassword: masterPassword})
			cancel()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			if !res.Success {
				return xerrors.New("unlock failed.")
			}
			fmt.Println("Unlock succeeded")

			return nil
		},
	}

	rootCmd.AddCommand(unlockCmd)
}

func List(rootCmd *cobra.Command) {
	listCmd := &cobra.Command{
		Use: "list",
		RunE: func(_ *cobra.Command, _ []string) error {
			client, err := dial()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			info, err := client.Info(ctx, &RequestInfo{})
			cancel()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			if info.Locked {
				return xerrors.New("Vault is locked")
			}

			ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
			list, err := client.List(ctx, &RequestList{})
			cancel()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			sort.Slice(list.Items, func(i, j int) bool {
				return list.Items[i].Title < list.Items[j].Title
			})
			for _, v := range list.Items {
				fmt.Printf("%s %s\n", v.Uuid, v.Title)
			}

			return nil
		},
	}

	rootCmd.AddCommand(listCmd)
}

func Get(rootCmd *cobra.Command) {
	getCmd := &cobra.Command{
		Use: "get UUID",
		RunE: func(_ *cobra.Command, args []string) error {
			log.Print(args)
			if len(args) != 1 {
				return xerrors.New("UUID is required")
			}
			client, err := dial()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			info, err := client.Info(ctx, &RequestInfo{})
			cancel()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			if info.Locked {
				return xerrors.New("Vault is locked")
			}

			ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
			res, err := client.Get(ctx, &RequestGet{Uuid: args[0]})
			cancel()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			fmt.Printf("UUID: %s\n", res.Item.Uuid)
			fmt.Printf("Title: %s\n", res.Item.Title)
			fmt.Printf("Password: %s\n", res.Item.Password)

			return nil
		},
	}

	rootCmd.AddCommand(getCmd)
}

var ErrDaemonNotExist = xerrors.New("daemon not exist")

func dial() (OnePasswordClient, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	confDir := filepath.Join(homeDir, ConfigDirName)
	if _, err := os.Stat(confDir); os.IsNotExist(err) {
		if err := os.Mkdir(confDir, 0700); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
	}
	pidFile := filepath.Join(homeDir, ConfigDirName, "1p.pid")
	if _, err := os.Stat(pidFile); os.IsNotExist(err) {
		return nil, ErrDaemonNotExist
	}

	buf, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	pid, err := strconv.Atoi(string(buf))
	if err != nil {
		os.Remove(pidFile)
		return nil, ErrDaemonNotExist
	}
	if exists, err := process.PidExists(int32(pid)); err != nil || !exists {
		return nil, ErrDaemonNotExist
	}

	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		var d net.Dialer
		return d.DialContext(ctx, "unix", addr)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	conn, err := grpc.DialContext(
		ctx,
		filepath.Join(confDir, socketFilename),
		grpc.WithInsecure(),
		grpc.WithContextDialer(dialer),
		grpc.WithBlock(),
	)
	cancel()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return NewOnePasswordClient(conn), nil
}

func AddCommand(rootCmd *cobra.Command) {
	for _, v := range subcommands {
		v(rootCmd)
	}
}
