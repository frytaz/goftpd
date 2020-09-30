package script

import (
	"context"

	"github.com/goftpd/goftpd/ftp/cmd"
	"github.com/pkg/errors"
)

var (
	ErrNotExist = errors.New("does not exist")
)

type Engine interface {
	Do(context.Context, []string, ScriptHook, cmd.Session) error
}

type ScriptHook string

const (
	ScriptHookPre     ScriptHook = "pre"
	ScriptHookPost               = "post"
	ScriptHookCommand            = "command"
)

var stringToScriptHook = map[string]ScriptHook{
	string(ScriptHookPre):     ScriptHookPre,
	string(ScriptHookPost):    ScriptHookPost,
	string(ScriptHookCommand): ScriptHookCommand,
}

type ScriptType string

const (
	ScriptTypeTrigger ScriptType = "trigger"
	ScriptTypeEvent              = "event"
)

type Command struct {
	FTPCommand string
	Hook       ScriptHook
	ScriptType ScriptType
	Path       string
}
